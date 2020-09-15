package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/stretchr/testify/require"
)

var testdataPath = filepath.Join("testdata", "azure-quickstart-templates", "*", "azuredeploy.json")

func TestTemplateCoverage(t *testing.T) {
	var metadata provider.AzureApiMetadata
	f, err := os.Open("../../../../cmd/pulumi-resource-azure-nextgen/schema/metadata.json")
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&metadata))
	f.Close()
	matches, err := filepath.Glob(testdataPath)
	require.NoError(t, err)
	matches = []string{"testdata/azure-quickstart-templates/101-acs-swarm/azuredeploy.json"}
	for _, match := range matches {
		t.Run(match, func(t *testing.T) {
			body, err := gen.RenderFile(match, &metadata)
			require.NoError(t, err)
			fmt.Println(body)

			programBody := fmt.Sprintf("%v", body)

			parser := syntax.NewParser()
			require.NoError(t, parser.ParseFile(strings.NewReader(programBody), "program.pp"))

			if parser.Diagnostics.HasErrors() {
				fmt.Print(programBody)
				parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
				t.Fail()
			}

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.Cache(hcl2.NewPackageCache()))
			if err != nil {
				log.Fatalf("failed to bind program: %v", err)
			}
			if diags.HasErrors() {
				log.Printf(programBody)
				program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
				t.Fail()
			}

			languages := []string{"nodejs"}
			languageExample := map[string]string{}
			for _, lang := range languages {
				var files map[string][]byte

				switch lang {
				case "dotnet":
					files, diags, err = dotnet.GenerateProgram(program)
				case "go":
					files, diags, err = gogen.GenerateProgram(program)
				case "nodejs":
					files, diags, err = nodejs.GenerateProgram(program)
				case "python":
					files, diags, err = python.GenerateProgram(program)
				case "schema":
					continue
				}
				if err != nil {
					log.Fatalf("failed to generate program: %v", err)
				}
				if diags.HasErrors() {
					log.Printf(programBody)
					program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
					os.Exit(-1)
				}

				buf := strings.Builder{}
				for _, f := range files {
					_, err := buf.Write(f)
					require.NoError(t, err)
				}
				languageExample[lang] = buf.String()
				fmt.Printf("%s\n", buf.String())
			}
			time.Sleep(5 * time.Second)
		})

	}
}
