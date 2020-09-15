package gen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/schollz/progressbar/v3"
)

type description string
type programText string
type language string

type languageToExampleProgram map[language]programText
type exampleRenderData struct {
	ExampleDescription       description
	LanguageToExampleProgram languageToExampleProgram
}
type resourceExamplesRenderData struct {
	Data []exampleRenderData
}

// Examples renders Azure API examples to the pkgSpec for the specified list of languages.
func Examples(pkgSpec *schema.PackageSpec, metadata *provider.AzureAPIMetadata,
	resExamples map[string][]provider.AzureAPIExample, languages []string) error {
	sortedKeys := codegen.SortedKeys(pkgSpec.Resources) // To generate in deterministic order

	// Use a progress bar to show progress since this can be a long running process
	bar := progressbar.Default(int64(len(sortedKeys)), "Resources processed:")

	// cache to speed up code generation
	hcl2Cache := hcl2.Cache(hcl2.NewPackageCache())
	for _, pulumiToken := range sortedKeys {
		err := bar.Add(1)
		if err != nil {
			log.Printf("error in the progress bar %v", err)
		}

		if shouldExclude(pulumiToken) {
			log.Printf("Skipping '%s' since it matches exclusion pattern", pulumiToken)
			continue
		}

		debug.Log("Processing '%s'", pulumiToken)
		resource := metadata.Resources[pulumiToken]
		seen := codegen.NewStringSet()
		split := strings.Split(pulumiToken, ":")
		if len(split) == 0 {
			return fmt.Errorf("invalid resourcename: %s", pulumiToken)
		}
		resourceName := pcl.Camel(split[len(split)-1])

		resourceExamples, ok := resExamples[pulumiToken]
		if !ok {
			continue
		}
		examplesRenderData := resourceExamplesRenderData{}
		for _, example := range resourceExamples {
			var items []model.BodyItem
			if seen.Has(example.Location) {
				continue
			}
			seen.Add(example.Location)
			f, err := os.Open(example.Location)
			if err != nil {
				return err
			}
			var exampleJSON map[string]interface{}
			if err = json.NewDecoder(f).Decode(&exampleJSON); err != nil {
				return err
			}
			if err = f.Close(); err != nil {
				return err
			}
			if _, ok := exampleJSON["parameters"]; !ok {
				return fmt.Errorf("example %s missing expected key: 'parameters'", example.Location)
			}

			resourceParams := map[string]provider.AzureAPIParameter{}
			for _, param := range resource.PutParameters {
				resourceParams[param.Name] = param
			}

			if _, ok := exampleJSON["parameters"].(map[string]interface{}); !ok {
				fmt.Printf("Expect parameters to be a map, received: %T for resource: %s, skipping.",
					exampleJSON["parameters"], pulumiToken)
				continue
			}
			exampleParams := exampleJSON["parameters"].(map[string]interface{})

			flattened, err := flattenInput(exampleParams, resourceParams, metadata.Types)
			if err != nil {
				fmt.Printf("tranforming input for example %s for resource %s: %v", example.Description, pulumiToken, err)
				continue
			}

			keys := codegen.SortedKeys(flattened)
			for _, k := range keys {
				v := flattened[k]
				val, err := pcl.RenderValue(v)
				if err != nil {
					return err
				}
				items = append(items, &model.Attribute{
					Name:  k,
					Value: val,
				})
			}

			block := model.Block{
				Type:   "resource",
				Body:   &model.Body{Items: items},
				Labels: []string{resourceName, pulumiToken},
			}
			body := &model.Body{Items: []model.BodyItem{&block}}
			pcl.FormatBody(body)
			languageExample, err := generateExamplePrograms(example, body, languages, hcl2Cache)
			if err != nil {
				return err
			}

			examplesRenderData.Data = append(examplesRenderData.Data,
				exampleRenderData{
					ExampleDescription:       description(example.Description),
					LanguageToExampleProgram: languageExample,
				})
		}
		if len(examplesRenderData.Data) > 0 {
			err := renderExampleToSchema(pkgSpec, pulumiToken, &examplesRenderData)
			if err != nil {
				return err
			}
		}

		metadata.Resources[pulumiToken] = resource
	}

	return nil
}

type programGenFn func(*hcl2.Program) (map[string][]byte, hcl.Diagnostics, error)

func generateExamplePrograms(example provider.AzureAPIExample, body *model.Body, languages []string,
	bindOptions ...hcl2.BindOption) (languageToExampleProgram, error) {
	programBody := fmt.Sprintf("%v", body)
	debug.Log(programBody)
	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programBody), "program.pp"); err != nil {
		return nil, fmt.Errorf("failed to parse IR - file: %s: %v", example.Location, err)
	}
	if parser.Diagnostics.HasErrors() {
		fmt.Print(programBody)
		err := parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	program, diags, err := hcl2.BindProgram(parser.Files, bindOptions...)
	if err != nil {
		log.Fatalf("failed to bind program: %v", err)
	}
	if diags.HasErrors() {
		log.Print(programBody)
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	languageExample := languageToExampleProgram{}

	for _, lang := range languages {
		var files map[string][]byte

		switch lang {
		case "dotnet":
			files, err = recoverableProgramGen(program, dotnet.GenerateProgram)
		case "go":
			files, err = recoverableProgramGen(program, gogen.GenerateProgram)
		case "nodejs":
			files, err = recoverableProgramGen(program, nodejs.GenerateProgram)
		case "python":
			files, err = recoverableProgramGen(program, python.GenerateProgram)
		default:
			continue
		}
		if err != nil {
			log.Printf("Program generation failed for language: %s for example %s, continuing", lang, example.Location)
			continue
		}

		buf := strings.Builder{}
		for _, f := range files {
			_, err := buf.Write(f)
			if err != nil {
				return nil, err
			}
		}
		languageExample[language(lang)] = programText(buf.String())
		debug.Log("Generated %s equivalent for %s", lang, example.Location)
		debug.Log("%s", buf.String())
	}

	return languageExample, nil
}

func recoverableProgramGen(program *hcl2.Program, fn programGenFn) (files map[string][]byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered during generation: %v", r)
		}
	}()

	var d hcl.Diagnostics
	files, d, err = fn(program)

	if err != nil {
		return nil, err
	}
	if d.HasErrors() {
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(d)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}
	return
}

func renderExampleToSchema(pkgSpec *schema.PackageSpec, resourceName string,
	examplesRenderData *resourceExamplesRenderData) error {
	const tmpl = `

{{"{{% examples %}}"}}
## Example Usage
{{- range .Data }}
{{ "{{% example %}}" }}
### {{ .ExampleDescription }}

{{- range $lang, $example := .LanguageToExampleProgram }}
{{ beginLanguage $lang }}
{{ $example }}
{{ endLanguage }}
{{ end }}
{{"{{% /example %}}"}}
{{- end }}
{{"{{% /examples %}}"}}
`
	res, ok := pkgSpec.Resources[resourceName]
	if !ok {
		return fmt.Errorf("missing resource from schema: %s", resourceName)
	}

	t, err := template.New("examples").Funcs(template.FuncMap{
		"beginLanguage": func(lang interface{}) string {
			l := fmt.Sprintf("%s", lang)
			switch l {
			case "nodejs":
				l = "typescript"
			case "dotnet":
				l = "csharp"
			}
			return fmt.Sprintf("```%s", l)
		},
		"endLanguage": func() string {
			return "```"
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}
	b := strings.Builder{}
	if err = t.Execute(&b, examplesRenderData); err != nil {
		return err
	}
	res.Description += b.String()
	pkgSpec.Resources[resourceName] = res
	return nil
}
