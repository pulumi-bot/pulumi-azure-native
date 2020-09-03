// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-azurerm/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// PulumiSchema will generate a Pulumi schema for the given Azure providers and resources map.
func PulumiSchema(providerMap openapi.AzureProviders) (*pschema.PackageSpec, *provider.AzureApiMetadata, error) {
	pkg := pschema.PackageSpec{
		Name:        "azurerm",
		Version:     "0.1.0", // TODO
		Description: "A Pulumi package for creating and managing Azure resources.",
		License:     "TODO",
		Keywords:    []string{"pulumi", "azure"},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-azurerm",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{},
		},
		Types:     map[string]pschema.ObjectTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]json.RawMessage{},
	}
	metadata := provider.AzureApiMetadata{
		Types:     map[string]provider.AzureApiType{},
		Resources: map[string]provider.AzureApiResource{},
		Invokes:   map[string]provider.AzureApiInvoke{},
	}

	csharpNamespaces := map[string]string{
		"azurerm": "AzureRM",
	}
	pythonModuleNames := map[string]string{}

	var providers []string
	for prov := range providerMap {
		providers = append(providers, prov)
	}
	sort.Strings(providers)

	for _, providerName := range providers {
		versionMap := providerMap[providerName]
		var versions []string
		for version := range versionMap {
			versions = append(versions, version)
		}
		sort.Strings(versions)

		for _, version := range versions {
			gen := packageGenerator{pkg: &pkg, metadata: &metadata, apiVersion: version}

			// Populate C# and Python module mapping.
			module := gen.providerToModule(providerName)
			csVersion := strings.Title(strings.Replace(version, "preview", "Preview", 1))
			csharpNamespaces[module] = fmt.Sprintf("%s.%s", providerName, csVersion)
			pythonModuleNames[module] = module

			// Populate resources and get invokes.
			items := versionMap[version]
			var resources []string
			for resource := range items.Resources {
				resources = append(resources, resource)
			}
			sort.Strings(resources)

			for _, typeName := range resources {
				resource := items.Resources[typeName]
				gen.genResources(providerName, typeName, resource.Path, resource.PathItem, resource.Swagger, resource.CompatibleVersions)
			}

			// Populate list invokes.
			var invokes []string
			for invoke := range items.Invokes {
				invokes = append(invokes, invoke)
			}
			sort.Strings(invokes)

			for _, typeName := range invokes {
				invoke := items.Invokes[typeName]
				gen.genListFunctions(providerName, typeName, invoke.Path, invoke.PathItem, invoke.Swagger)
			}
		}
	}

	const goBasePath = "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm"
	golangImportAliases := map[string]string{}

	re := regexp.MustCompile(`(?P<pkgName>.*):(?P<module>.*)/(?P<version>.*):(?P<type>.*)`)
	for resName := range pkg.Resources {
		subMatches := re.FindStringSubmatch(resName)
		subMatchMap := map[string]string{}

		if len(subMatches) != len(re.SubexpNames()) {
			return nil, nil, fmt.Errorf("unexpected resource format: %s", resName)
		}
		for i, name := range re.SubexpNames() {
			if i != 0 {
				subMatchMap[name] = subMatches[i]
			}
		}
		packageImport := filepath.Join(goBasePath, subMatchMap["module"], subMatchMap["version"])
		golangImportAliases[packageImport] = subMatchMap["module"]
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":       goBasePath,
		"packageImportAliases": golangImportAliases,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^2.0.0",
		},
		"readme": pkg.Description, // TODO: add a proper readme.
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=2.0.0,<3.0.0",
		},
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "2.*",
			"System.Collections.Immutable": "1.6.0",
		},
		"namespaces": csharpNamespaces,
	})

	return &pkg, &metadata, nil
}

// Microsoft-specific API extension constants.
const (
	extensionClientFlatten     = "x-ms-client-flatten"
	extensionClientName        = "x-ms-client-name"
	extensionEnum              = "x-ms-enum"
	extensionMutability        = "x-ms-mutability"
	extensionMutabilityCreate  = "create"
	extensionMutabilityUpdate  = "update"
	extensionParameterLocation = "x-ms-parameter-location"
)

type packageGenerator struct {
	pkg        *pschema.PackageSpec
	metadata   *provider.AzureApiMetadata
	apiVersion string
}

func (g *packageGenerator) genResources(prov, typeName, key string, path *spec.PathItem, swagger *openapi.Spec, otherVersions []string) {
	module := g.providerToModule(prov)
	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName)

	// Generate the resource.
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		resourceToken: resourceTok,
		visitedTypes:  make(map[string]bool),
	}

	parameters := swagger.MergeParameters(path.Put.Parameters, path.Parameters)
	resourceRequest, err := gen.genMethodParameters(parameters, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", resourceTok, err.Error())
		return
	}

	resourceResponse, err := gen.genResponse(path.Put.Responses.StatusCodeResponses, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", resourceTok, err.Error())
		return
	}

	if len(resourceResponse.specs) == 0 {
		// Response is specified empty, do not generate a resource for it.
		return
	}

	gen.escapeCSharpNames(typeName, resourceResponse)

	// Add an alias for each API version that has the same path in it.
	var aliases []pschema.AliasSpec
	for _, version := range otherVersions {
		alias := fmt.Sprintf("%s:%s/%s:%s", g.pkg.Name, strings.ToLower(prov), version, typeName)
		aliases = append(aliases, pschema.AliasSpec{Type: &alias})
	}

	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: resourceResponse.description,
			Type:        "object",
			Properties:  resourceResponse.specs,
			Required:    resourceResponse.requiredSpecs.SortedValues(),
		},
		InputProperties: resourceRequest.specs,
		RequiredInputs:  resourceRequest.requiredSpecs.SortedValues(),
		Aliases:         aliases,
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, typeName)

	parameters = swagger.MergeParameters(path.Get.Parameters, path.Parameters)
	requestFunction, err := gen.genMethodParameters(parameters, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	responseFunction, err := gen.genResponse(path.Get.Responses.StatusCodeResponses, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	functionSpec := pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Description: requestFunction.description,
			Type:        "object",
			Properties:  requestFunction.specs,
			Required:    requestFunction.requiredSpecs.SortedValues(),
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: responseFunction.description,
			Type:        "object",
			Properties:  responseFunction.specs,
			Required:    responseFunction.requiredSpecs.SortedValues(),
		},
	}
	g.pkg.Functions[functionTok] = functionSpec

	r := provider.AzureApiResource{
		ApiVersion:    swagger.Info.Version,
		Path:          key,
		GetParameters: requestFunction.parameters,
		PutParameters: resourceRequest.parameters,
		Examples:      g.generateExampleReferences(prov, key, path, swagger),
		Response:      resourceResponse.properties,
	}
	g.metadata.Resources[resourceTok] = r

	f := provider.AzureApiInvoke{
		ApiVersion:    swagger.Info.Version,
		Path:          key,
		GetParameters: requestFunction.parameters,
		Response:      responseFunction.properties,
	}
	g.metadata.Invokes[functionTok] = f
}

func (g *packageGenerator) generateExampleReferences(providerName, key string, path *spec.PathItem, swagger *openapi.Spec) []provider.AzureApiExample {
	raw, ok := path.Put.Extensions["x-ms-examples"]
	if !ok {
		return nil
	}

	examples := raw.(map[string]interface{})

	result := make([]provider.AzureApiExample, 0, len(examples))
	for k, v := range examples {
		resolved := v.(map[string]interface{})
		if _, ok := resolved["$ref"]; !ok {
			continue
		}

		relativePath := resolved["$ref"].(string)
		relativeURL, err := url.Parse(relativePath)
		if err != nil {
			panic(err)
		}
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		url, err := swagger.ResolveReference(relativeURL.String())
		if err != nil {
			panic(err)
		}

		url, err = filepath.Rel(cwd, url)
		if err != nil {
			panic(err)
		}
		if _, err := os.Stat(url); err != nil {
			panic(err)
		}
		result = append(result, provider.AzureApiExample{
			Description: k,
			Location:    url,
		})
	}

	// Deterministic ordering of the examples.
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Description < result[j].Description
	})

	return result
}

// genListFunctions defines functions for list* (listKeys, listSecrets, etc.) POST endpoints.
func (g *packageGenerator) genListFunctions(prov, typeName, path string, pathItem *spec.PathItem, swagger *openapi.Spec) {
	module := g.providerToModule(prov)
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		resourceToken: fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName),
		visitedTypes:  make(map[string]bool),
	}

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:list%s`, g.pkg.Name, module, typeName)

	parameters := swagger.MergeParameters(pathItem.Post.Parameters, pathItem.Parameters)
	request, err := gen.genMethodParameters(parameters, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	response, err := gen.genResponse(pathItem.Post.Responses.StatusCodeResponses, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	if len(response.specs) == 0 {
		// Response is specified empty, do not generate an invoke for it.
		return
	}

	functionSpec := pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Description: request.description,
			Type:        "object",
			Properties:  request.specs,
			Required:    request.requiredSpecs.SortedValues(),
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: response.description,
			Type:        "object",
			Properties:  response.specs,
			Required:    response.requiredSpecs.SortedValues(),
		},
	}
	g.pkg.Functions[functionTok] = functionSpec

	f := provider.AzureApiInvoke{
		ApiVersion:     swagger.Info.Version,
		Path:           path,
		PostParameters: request.parameters,
		Response:       response.properties,
	}
	g.metadata.Invokes[functionTok] = f
}

// providerToModule produces the module name from the provider name and the API version (e.g. (`Compute`, `2020-07-01` => `compute/v20200701`).
func (g *packageGenerator) providerToModule(prov string) string {
	return fmt.Sprintf("%s/%s", strings.ToLower(prov), g.apiVersion)
}

type moduleGenerator struct {
	pkg           *pschema.PackageSpec
	metadata      *provider.AzureApiMetadata
	module        string
	resourceToken string
	visitedTypes  map[string]bool
}

func (m *moduleGenerator) escapeCSharpNames(typeName string, resourceResponse *propertyBag) {
	for name, swagger := range resourceResponse.specs {
		// C# doesn't allow properties to have the same name as its containing type.
		if strings.Title(name) == typeName {
			swagger.Language = map[string]json.RawMessage{
				"csharp": rawMessage(map[string]interface{}{
					"name": fmt.Sprintf("%sValue", typeName),
				}),
			}
			resourceResponse.specs[name] = swagger
		}
	}
}

func (m *moduleGenerator) genMethodParameters(parameters []spec.Parameter, ctx *openapi.ReferenceContext) (*parameterBag, error) {
	result := newParameterBag()

	for _, param := range parameters {
		param, err := ctx.ResolveParameter(param)
		if err != nil {
			return nil, err
		}

		location, _ := param.Extensions.GetString(extensionParameterLocation)
		apiParameter := provider.AzureApiParameter{
			Name:       param.Name,
			Location:   param.In,
			Source:     location,
			IsRequired: param.Required,
			Value: &provider.AzureApiProperty{
				Type:      param.Type,
				MinLength: param.MinLength,
				MaxLength: param.MaxLength,
				Pattern:   param.Pattern,
			},
		}

		switch {
		case param.In == "header":
			continue // Header parameters aren't mapped to the SDK.
		case param.Name == "subscriptionId":
		case param.Name == "api-version":
			continue // No need to include these in the schema, they are added automatically by the provider.
		case param.In == "body":
			// The body parameter is flattened, so that all its properties become the properties of the type.
			if param.Schema == nil {
				return nil, errors.Errorf("no schema for body parameter '%s'", param.Name)
			}

			resolvedSchema, err := param.ResolveSchema(param.Schema)
			if err != nil {
				return nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}

			props, err := m.genProperties(resolvedSchema, false /* isOutput */, false /* isType */)
			if err != nil {
				return nil, err
			}

			result.merge(props)
			apiParameter.Body = &provider.AzureApiType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}

		default:
			name := param.Name
			if clientName, ok := param.Extensions.GetString(extensionClientName); ok {
				name = clientName
			}

			// Change the name to lowerCamelCase.
			name = toLowerCamel(name)
			if name != param.Name {
				apiParameter.Value.SdkName = name
			}

			propertySpec := pschema.PropertySpec{
				Description: param.Description,
				TypeSpec: pschema.TypeSpec{
					Type: param.Type,
				},
			}
			result.specs[name] = propertySpec
			if param.Required {
				result.requiredSpecs.Add(name)
			}
		}

		result.parameters = append(result.parameters, apiParameter)
	}

	return result, nil
}

func (m *moduleGenerator) genResponse(statusCodeResponses map[int]spec.Response, ctx *openapi.ReferenceContext) (*propertyBag, error) {
	var responseSchema *openapi.Schema

	// Find all 2xx codes and sort them to make codegen deterministic.
	var codes []int
	for code := range statusCodeResponses {
		if code >= 300 {
			continue
		}

		codes = append(codes, code)
	}
	sort.Ints(codes)

	if len(codes) == 0 {
		return nil, errors.New("no 2xx response found")
	}

	// Find the lowest 2xx response with a schema definition and derive response properties from it.
	for _, code := range codes {
		resp := statusCodeResponses[code]
		response, err := ctx.ResolveResponse(resp)
		if err != nil {
			return nil, err
		}

		if response.Schema == nil {
			continue
		}

		responseSchema, err = response.ResolveSchema(response.Schema)
		if err != nil {
			return nil, err
		}

		if len(responseSchema.Type) > 0 && responseSchema.Type[0] == "array" {
			return nil, errors.New("array responses are not implemented yet (see issue #120)")
		}

		result, err := m.genProperties(responseSchema, true /* isOutput */, false /* isType */)
		if err != nil {
			return nil, err
		}

		// Skip empty result objects.
		if len(result.specs) == 0 {
			continue
		}

		// Id is a property of the base Custom Resource, we don't want to introduce it on derived resources.
		delete(result.specs, "id")
		result.requiredSpecs.Delete("id")

		// Special case the 'properties' output property as required.
		// This should be gone when we apply flattening.
		if _, ok := result.specs["properties"]; ok {
			result.requiredSpecs.Add("properties")
		}

		result.description = responseSchema.Description
		return result, nil
	}

	// There was at least one 2xx response defined, but it has no schema. This is not a valid resource for us,
	// skip its processing.
	return &propertyBag{}, nil
}

func (m *moduleGenerator) genProperties(resolvedSchema *openapi.Schema, isOutput, isType bool) (*propertyBag, error) {
	result := newPropertyBag()

	// Sort properties to make codegen deterministic.
	var names []string
	for name := range resolvedSchema.Properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		property := resolvedSchema.Properties[name]
		sdkName := name
		if clientName, ok := property.Extensions.GetString(extensionClientName); ok {
			sdkName = firstToLower(clientName)
		}
		// Change the name to lowerCamelCase.
		sdkName = toLowerCamel(sdkName)

		// Flattened properties aren't modelled in the SDK explicitly: their sub-properties are merged directly to the parent.
		if flatten, ok := property.Extensions.GetBool(extensionClientFlatten); ok && flatten {
			resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
			if err != nil {
				return nil, err
			}

			bag, err := m.genProperties(resolvedProperty, isOutput, isType)
			if err != nil {
				return nil, err
			}

			// Adjust every property to mark them as flattened.
			newProperties := map[string]provider.AzureApiProperty{}
			for n, value := range bag.properties {
				value.Container = name
				newProperties[n] = value
			}
			bag.properties = newProperties

			result.merge(bag)
			continue
		}

		// Skip read-only properties for input types.
		if property.ReadOnly && !isOutput {
			continue
		}

		propertySpec, err := m.genProperty(name, &property, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, err
		}

		// Skip properties that yield degenerate types (e.g., when an input type has only read-only properties).
		if propertySpec == nil {
			continue
		}

		var apiProperty provider.AzureApiProperty
		if isOutput {
			if property.ReadOnly {
				result.requiredSpecs.Add(sdkName)
			}
			apiProperty = provider.AzureApiProperty{
				Ref: propertySpec.Ref,
			}
		} else {
			resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
			if err != nil {
				return nil, err
			}
			apiProperty = provider.AzureApiProperty{
				Type:      propertySpec.Type,
				Enum:      m.getEnumValues(&property),
				Ref:       propertySpec.Ref,
				Minimum:   resolvedProperty.Minimum,
				Maximum:   resolvedProperty.Maximum,
				MinLength: resolvedProperty.MinLength,
				MaxLength: resolvedProperty.MaxLength,
				Pattern:   resolvedProperty.Pattern,
			}

			// Mutability extension signals whether a property can be updated in-place. Lack of the extension means
			// updatable by default.
			// Note: a non-updatable property at a subtype level (a property of a property of a resource) does not
			// mandate the replacement of the whole resource. Anyway, it's used very seldom (2 places at the time of writing).
			// Example: `StorageAccount.encryption.services.blob.keyType` is non-updatable, but a user can remove `blob`
			// and then re-add it with the new `keyType` without replacing the whole storage account (which would be
			// very disruptive).
			if mutability, ok := property.Extensions.GetStringSlice(extensionMutability); ok && !isType {
				operations := codegen.NewStringSet(mutability...)
				apiProperty.ForceNew = operations.Has(extensionMutabilityCreate) && !operations.Has(extensionMutabilityUpdate)
			}
		}

		if sdkName != name {
			apiProperty.SdkName = sdkName
		}
		if propertySpec.Items != nil {
			apiProperty.Items = &provider.AzureApiProperty{
				Type: propertySpec.Items.Type,
				Ref:  propertySpec.Items.Ref,
			}
		}
		result.properties[name] = apiProperty
		result.specs[sdkName] = *propertySpec
	}

	for _, s := range resolvedSchema.AllOf {
		allOfSchema, err := resolvedSchema.ResolveSchema(&s)
		if err != nil {
			return nil, err
		}

		allOfProperties, err := m.genProperties(allOfSchema, isOutput, isType)
		if err != nil {
			return nil, err
		}

		result.merge(allOfProperties)
	}

	for _, name := range resolvedSchema.Required {
		if prop, ok := result.properties[name]; ok {
			if prop.SdkName != "" {
				result.requiredSpecs.Add(prop.SdkName)
			} else {
				result.requiredSpecs.Add(name)
			}
			result.requiredProperties.Add(name)
		}
	}
	return result, nil
}

func (m *moduleGenerator) getEnumValues(property *spec.Schema) (enum []string) {
	if property.Enum == nil {
		return
	}

	restrictive := true
	// If x-ms-enum is present and modelAsString is set to false, the enum is not strict, so we don't want to enforce it.
	if extension, ok := property.Extensions[extensionEnum]; ok {
		if modelAsString, ok := extension.(map[string]interface{})["modelAsString"]; ok {
			if v, ok := modelAsString.(bool); ok {
				restrictive = !v
			}
		}
	}
	if restrictive {
		for _, value := range property.Enum {
			if s, ok := value.(string); ok {
				enum = append(enum, s)
			}
		}
	}
	return
}

func (m *moduleGenerator) genProperty(name string, schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.PropertySpec, error) {
	description := schema.Description
	if description == "" {
		resolvedSchema, err := context.ResolveSchema(schema)
		if err != nil {
			return nil, err
		}

		description = resolvedSchema.Description
	}

	typeSpec, err := m.genTypeSpec(name, schema, context, isOutput)
	if err != nil {
		return nil, err
	}

	// Nil type means empty: e.g., when an input type has only read-only properties. Bubble the nil up.
	if typeSpec == nil {
		return nil, nil
	}

	propertySpec := pschema.PropertySpec{
		Description: description,
		TypeSpec:    *typeSpec,
	}

	return &propertySpec, nil
}

func (m *moduleGenerator) genTypeSpec(propertyName string, schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.TypeSpec, error) {
	resolvedSchema, err := context.ResolveSchema(schema)
	if err != nil {
		return nil, err
	}

	// Type specification specifies either a primitive type (e.g. 'string') or a reference to a separately defined
	// strongly-typed object. Either `primitiveTypeName` or `referencedTypeName` should be filled, but not both.
	var tok, primitiveTypeName, referencedTypeName string
	if len(resolvedSchema.Properties) > 0 || len(resolvedSchema.AllOf) > 0 {
		ptr := schema.Ref.GetPointer()
		if ptr != nil && !ptr.IsEmpty() {
			tok = m.typeName(resolvedSchema.ReferenceContext, isOutput)
		} else {
			// Inline properties have no type in the Open API schema, so we use parent type's name + property name.
			tok = m.typeName(context, isOutput) + strings.Title(propertyName)
		}
	} else if len(resolvedSchema.Type) > 0 {
		primitiveTypeName = resolvedSchema.Type[0]
	} else {
		primitiveTypeName = "object"
	}

	// If an object type is referenced, add its definition to the type map.
	if tok != "" {
		referencedTypeName = fmt.Sprintf("#/types/%s", tok)

		if _, ok := m.visitedTypes[tok]; !ok {
			m.visitedTypes[tok] = true

			props, err := m.genProperties(resolvedSchema, isOutput, true /* isType */)
			if err != nil {
				return nil, err
			}

			// Don't generate a type definition for a typed object with zero properties.
			if len(props.specs) == 0 {
				delete(m.visitedTypes, tok)
				return nil, nil
			}

			m.pkg.Types[tok] = pschema.ObjectTypeSpec{
				Description: resolvedSchema.Description,
				Type:        "object",
				Properties:  props.specs,
				Required:    props.requiredSpecs.SortedValues(),
			}

			m.metadata.Types[tok] = provider.AzureApiType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}
		}
	}

	// Define the type of maps (untyped objects).
	var additionalProperties *pschema.TypeSpec
	if primitiveTypeName == "object" {
		if resolvedSchema.AdditionalProperties != nil && resolvedSchema.AdditionalProperties.Schema != nil {
			additionalProperties, err = m.genTypeSpec(propertyName, resolvedSchema.AdditionalProperties.Schema, resolvedSchema.ReferenceContext, isOutput)
			if err != nil {
				return nil, err
			}

			// Don't generate a type definition for a typed dictionary with empty value type.
			if additionalProperties == nil {
				return nil, nil
			}
		} else {
			additionalProperties = &pschema.TypeSpec{
				Ref: "pulumi.json#/Any",
			}
		}
	}

	result := pschema.TypeSpec{
		Type:                 primitiveTypeName,
		AdditionalProperties: additionalProperties,
		Ref:                  referencedTypeName,
	}

	// Resolve the element type for array types.
	if resolvedSchema.Items != nil && resolvedSchema.Items.Schema != nil {
		itemsSpec, err := m.genProperty(propertyName, resolvedSchema.Items.Schema, context, isOutput)
		if err != nil {
			return nil, err
		}

		// Don't generate a type definition for a typed array with empty item type.
		if itemsSpec == nil {
			return nil, nil
		}

		result.Items = &itemsSpec.TypeSpec
	}

	return &result, nil
}

func (m *moduleGenerator) typeName(ctx *openapi.ReferenceContext, isOutput bool) string {
	suffix := ""
	if isOutput {
		suffix = "Response"
	}
	return fmt.Sprintf("azurerm:%s:%s%s", m.module, makeLegalIdentifier(ctx.ReferenceName), suffix)
}

// parameterBag keeps the schema and metadata parameters for a single resource or invocation.
type parameterBag struct {
	description   string
	specs         map[string]pschema.PropertySpec
	parameters    []provider.AzureApiParameter
	requiredSpecs codegen.StringSet
}

func newParameterBag() *parameterBag {
	return &parameterBag{
		specs:         map[string]pschema.PropertySpec{},
		requiredSpecs: codegen.NewStringSet(),
	}
}

func (bag *parameterBag) merge(other *propertyBag) {
	for key, value := range other.specs {
		bag.specs[key] = value
	}
	for key := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
}

// propertyBag keeps the schema and metadata properties for a single API type.
type propertyBag struct {
	description        string
	specs              map[string]pschema.PropertySpec
	properties         map[string]provider.AzureApiProperty
	requiredSpecs      codegen.StringSet
	requiredProperties codegen.StringSet
}

func newPropertyBag() *propertyBag {
	return &propertyBag{
		specs:              map[string]pschema.PropertySpec{},
		properties:         map[string]provider.AzureApiProperty{},
		requiredSpecs:      codegen.NewStringSet(),
		requiredProperties: codegen.NewStringSet(),
	}
}

func (bag *propertyBag) merge(other *propertyBag) {
	for key, value := range other.specs {
		bag.specs[key] = value
	}
	for key, value := range other.properties {
		bag.properties[key] = value
	}
	for key := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
	for key := range other.requiredProperties {
		bag.requiredProperties.Add(key)
	}
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
