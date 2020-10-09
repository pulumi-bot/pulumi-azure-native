package arm2pulumi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/tle"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

// TemplateElement is a top-level item in the template, e.g. parameters, variables,
// resources and outputs.
type TemplateElement interface {
	Name() string
	Args() interface{}
	SetArgs(interface{})
	// FinishInitializing is called after all template evaluation phases are complete
	// but before PCLExpression. Most relevant usecase is to inject
	// implicit default variables and validate against metadata. Returns a general
	// error if a fatal issue is hit. Returned error may be a Diagnostic in which
	// case the error may be non-fatal but may result in the element being excluded
	// in the IR format.
	FinishInitializing(ctx *pclRenderContext, implicits implicitVariables) error
	PCLExpression(ctx *pclRenderContext) ([]model.BodyItem, error)
}

type implicitVariables interface {
	defaultResourceGroupName() (*model.Variable, *dependencyTracking, error)
	defaultResourceGroup() (*model.Variable, *dependencyTracking, error)
}

type pclRenderContext struct {
	metadata               *provider.AzureAPIMetadata
	pkgSpec                *schema.PackageSpec
	dep                    *dependencyTracking
	resourceTokenConverter *resourceTokenConverter
}

func newParameter(name string, rawBody map[string]interface{}) *parameter {
	return &parameter{
		paramName: name,
		rawBody:   rawBody,
	}
}

type parameter struct {
	paramName string
	rawBody   map[string]interface{}
}

func (p *parameter) FinishInitializing(ctx *pclRenderContext, implicits implicitVariables) error {
	return nil
}

func (p *parameter) Name() string {
	return p.paramName
}

func (p *parameter) Args() interface{} {
	return p.rawBody
}

func (p *parameter) SetArgs(body interface{}) {
	p.rawBody = body.(map[string]interface{})
}

func (p *parameter) Validate(ctx *pclRenderContext) (bool, *Diagnostic) {
	return true, nil
}

func (p *parameter) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	typ, err := extractType(p.rawBody)
	if err != nil {
		return nil, fmt.Errorf("invalid parameter %s: %w", p.paramName, err)
	}

	labels := []string{p.paramName}
	if typ != "" {
		labels = append(labels, typ)
	}

	var defaultValue model.Expression
	if def, ok := p.rawBody["defaultValue"]; ok {
		v, err := pcl.RenderValue(def)
		if err != nil {
			return nil, err
		}
		defaultValue = v
	}

	comment := extractDescription(p.rawBody)
	configDef := &model.Block{
		Type:   "config",
		Labels: labels,
		Body:   &model.Body{},
		Tokens: getLeadingTriviaTokens(comment),
	}

	if defaultValue != nil {
		configDef.Body.Items = append(configDef.Body.Items, &model.Attribute{
			Name:  "default",
			Value: defaultValue,
		})
	}
	return []model.BodyItem{configDef}, nil
}

func newVariable(name string, rawBody interface{}) *variable {
	return &variable{
		varName: name,
		rawBody: rawBody,
	}
}

type variable struct {
	varName string
	rawBody interface{}
}

func (v *variable) FinishInitializing(ctx *pclRenderContext, implicits implicitVariables) error {
	return nil
}

func (v *variable) Name() string {
	return v.varName
}

func (v *variable) Args() interface{} {
	return v.rawBody
}

func (v *variable) SetArgs(body interface{}) {
	v.rawBody = body
}

func (v *variable) Validate(ctx *pclRenderContext) (bool, *Diagnostic) {
	return true, nil
}

func (v *variable) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	fObj := v.rawBody
	m, err := pcl.RenderValue(fObj)
	if err != nil {
		return nil, err
	}

	return []model.BodyItem{&model.Attribute{
		Name:  v.varName,
		Value: m,
	}}, nil
}

func newOutput(name string, rawBody map[string]interface{}) *output {
	return &output{
		outputName: name,
		rawBody:    rawBody,
	}
}

type output struct {
	outputName string
	rawBody    map[string]interface{}
}

func (o *output) FinishInitializing(ctx *pclRenderContext, implicits implicitVariables) error {
	return nil
}

func (o *output) Name() string {
	return o.outputName
}

func (o *output) Args() interface{} {
	return o.rawBody
}

func (o *output) SetArgs(body interface{}) {
	o.rawBody = body.(map[string]interface{})
}

func (o *output) Validate(ctx *pclRenderContext) (bool, *Diagnostic) {
	return true, nil
}

func (o *output) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	value := o.rawBody

	if cond, ok := value["condition"]; ok {
		if keep, ok := cond.(bool); !ok {
			return nil, fmt.Errorf("expect condition field in output %s to be boolean, received: %T",
				o.outputName, cond)
		} else if !keep {
			return nil, nil
		}
	}

	// TODO: Need to add support for copy
	val, ok := value["value"]
	if !ok {
		return nil, fmt.Errorf("output '%v' has no \"value\" attribute", o.outputName)
	}

	x, err := pcl.RenderValue(val)
	if err != nil {
		return nil, err
	}

	// TODO: Add support for secure variants as type

	comment := extractDescription(o.rawBody)
	return []model.BodyItem{&model.Block{
		Type:   "output",
		Labels: []string{o.outputName},
		Body: &model.Body{
			Items: []model.BodyItem{
				&model.Attribute{
					Name:  "value",
					Value: x,
				},
			},
		},
		Tokens: getLeadingTriviaTokens(comment),
	}}, nil
}

func NewTemplateElements() *TemplateElements {
	return &TemplateElements{
		elements:       map[string]*dependencyTracking{},
		canonicalNames: map[string]string{},
		diagnostics:    map[string][]Diagnostic{},
	}
}

// TemplateElements keeps an internal state of template elements to be rendered to PCL,
// including maintaining dependency ordering of the elements
type TemplateElements struct {
	elements       map[string]*dependencyTracking
	canonicalNames map[string]string
	diagnostics    map[string][]Diagnostic
}

func (t *TemplateElements) lookup(name string) *dependencyTracking {
	if el, ok := t.elements[name]; ok {
		return el
	}

	cname := gen.ToLowerCamel(gen.MakeLegalIdentifier(name))
	_, ok := t.canonicalNames[cname]
	if ok {
		return t.elements[cname]
	}
	return nil
}

func (t *TemplateElements) recordCanonicalizedName(name string) string {
	cName := gen.ToLowerCamel(gen.MakeLegalIdentifier(name))
	t.canonicalNames[cName] = name
	return cName
}

// EvaluateExpressions evaluates template expressions. By default, it tries
// to evaluate all templates if preliminaryPhase is true, it will skip resolving
// evaluating templates with certain functions. This is required for instance,
// with functions like reference. Since we may add additional elements due to
// template evaluation and the order of evaluation is random, we wait to perform
// referential evaluations till after all other evaluations are complete.
func (t *TemplateElements) EvaluateExpressions(preliminaryPhase bool) error {
	for name, el := range t.elements {
		args := el.TemplateElement.Args()
		var err error
		if preliminaryPhase {
			args, err = t.evalTemplateExpressions(el, args, "", "reference")
		} else {
			args, err = t.evalTemplateExpressions(el, args, "")
		}
		if err != nil {
			return fmt.Errorf("failed to evaluate template expression for %s: %w", name, err)
		}
		el.SetArgs(args)
	}
	return nil
}

// evalTemplateExpressions evaluates template expressions recursively for an individual template element
func (t *TemplateElements) evalTemplateExpressions(
	d *dependencyTracking,
	args interface{},
	crumbs string,
	excludeFuncs ...string,
) (interface{}, error) {
	val := reflect.ValueOf(args)
	exclusions := codegen.NewStringSet()
	for _, excl := range excludeFuncs {
		exclusions.Add(excl)
	}
	typ := reflect.TypeOf(args)
	if typ == nil {
		return args, nil
	}
	switch typ.Kind() {
	case reflect.Map:
		ret := map[string]interface{}{}
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			curr := fmt.Sprintf(".%s", k.String())
			updated, err := t.evalTemplateExpressions(d, v.Interface(), curr, excludeFuncs...)
			if err != nil {
				return nil, err
			}
			ret[k.String()] = updated
		}
		return ret, nil
	case reflect.Slice, reflect.Array:
		var ret []interface{}
		for i := 0; i < val.Len(); i++ {
			curr := fmt.Sprintf("%s[%d]", crumbs, i)
			updated, err := t.evalTemplateExpressions(d, val.Index(i).Interface(), curr, excludeFuncs...)
			if err != nil {
				return nil, err
			}
			ret = append(ret, updated)
		}
		return ret, nil
	case reflect.String:
		// json.Number is typed as a string so we must disambiguate.
		if num, ok := args.(json.Number); ok {
			fl, err := num.Float64()
			if err != nil {
				return nil, fmt.Errorf("invalid numeric value at %s: %w", crumbs, err)
			}
			return fl, nil
		}
		s := val.String()
		evaled, err := t.eval(d, s, exclusions)
		if err != nil {
			return nil, err
		}
		return evaled, nil
	}

	return args, nil
}

// assignName tries to create a unique name elements based on the
// nesting information in crumbs
func (t *TemplateElements) assignName(crumbs []string) string {
	varName := ""
	for _, c := range crumbs {
		varName += strings.Title(c)
	}
	varName = gen.ToLowerCamel(gen.MakeLegalIdentifier(varName))
	if el := t.lookup(varName); el == nil {
		return varName
	}

	base := varName
	i := 0
	for {
		varName = fmt.Sprintf("%s%d", base, i)
		if el := t.lookup(varName); el == nil {
			break
		}
		i++
	}
	return varName
}

// addParameter adds the parameter with the given name and template args as
// a config item when translated from the template.
// Since names are only unique to a section, if addSuffix is set, a suffix
// of `Param` is added at the end of the variable name.
func (t *TemplateElements) addParameter(name string, args map[string]interface{}, addSuffix bool) error {
	if addSuffix {
		if !strings.HasSuffix(strings.ToLower(name), "param") {
			name = name + "Param"
		}
	}

	if el := t.lookup(name); el != nil {
		// Don't expect name collision with params
		return fmt.Errorf("another item with the same name as parameter %s already defined", name)
	}
	name = t.recordCanonicalizedName(name)
	p := newParameter(name, args)
	dep := newDependencyTracking(p)
	t.elements[name] = dep
	return nil
}

// addVariable adds the variable with the given name and template args.
// Since names are only unique to a section, if addSuffix is set, a suffix
// of `Var` is added at the end of the variable name.
func (t *TemplateElements) addVariable(name string, args interface{}, addSuffix bool) (*dependencyTracking, error) {
	if addSuffix {
		if !strings.HasSuffix(strings.ToLower(name), "var") {
			name = name + "Var"
		}
	}

	if el := t.lookup(name); el != nil {
		return nil, fmt.Errorf("another item with the same name as variable %s already defined", name)
	}

	name = t.recordCanonicalizedName(name)
	v := newVariable(name, args)
	dep := newDependencyTracking(v)
	t.elements[name] = dep

	return dep, nil
}

// addOutput adds the output with the given name and template args.
// Since names are only unique to a section, if addSuffix is set, a suffix
// of `Out` is added at the end of the variable name.
func (t *TemplateElements) addOutput(name string, args map[string]interface{}, addSuffix bool) error {
	if addSuffix {
		if !strings.HasSuffix(strings.ToLower(name), "out") {
			name = name + "Out"
		}
	}

	if el := t.lookup(name); el != nil {
		// Don't expect name collision with outputs
		return fmt.Errorf("another item with the same name as output %s already defined", name)
	}

	name = t.recordCanonicalizedName(name)
	o := newOutput(name, args)
	dep := newDependencyTracking(o)
	t.elements[name] = dep
	return nil
}

// addResource adds the resource with the given template args.
// Resource names are often templated and only resolvable at runtime so
// we generate our own variable names for resources derived from the type.
func (t *TemplateElements) addResource(args map[string]interface{}) error {
	err := t.handleAddResource(args, "", "")
	if err != nil {
		return err
	}
	return nil
}

func (t *TemplateElements) handleAddResource(args map[string]interface{}, parentName, parentType string) error {
	name, ok := args["name"]
	if !ok {
		return errors.New("missing required name field in resource")
	}

	var prefix []string

	var typestr string
	typ, ok := args["type"]
	if !ok {
		return errors.New("missing required type field in resource")
	}
	typestr = typ.(string)

	// Microsoft.ContainerService/managedClusters -> managedClusterResource
	if !isTemplateExpression(typestr) {
		resource := strings.Split(typestr, "/")
		res := resource[len(resource)-1]
		res = inflector.Singularize(strings.Title(gen.ToLowerCamel(res)))
		prefix = append(prefix, res)
	}
	prefix = append(prefix, "Resource")

	// we often have names as template expressions. We don't want to rely on evaluating the name
	// as the resource variable name so we create a reasonable variable name to house the resource
	// instead. In this case we use the type to help define the variable name, unless that is also
	// a template expression - in which case we just add a 'Resource' suffix.
	// Otherwise just use the resource name field.
	namestr := name.(string)

	// If this is a subresource, the name of the parent needs to precede the name of this resource,
	// i.e. the actual name is <parent>/<child>. We flatten the subresources by rendering a separate
	// resource but perform the name rewrite by adding a synthetic template expression to concatenate
	// the names of the parent and child.
	if parentName != "" {
		thisNameParse, err := tle.Parse(namestr)
		if err != nil {
			return err
		}
		parentNameParse, err := tle.Parse(parentName)
		if err != nil {
			return err
		}
		val := tle.FunctionCallValue{
			NameToken: tle.NewToken(tle.Literal, 0, "concat"),
			ArgumentExpressions: []tle.Value{parentNameParse.Expression,
				tle.NewStringValue(tle.CreateQuotedString(0, "/")), thisNameParse.Expression},
		}
		namestr, err = tle.ToTemplateExpressionString(&val)
		if err != nil {
			return err
		}
		args["name"] = namestr
	}

	// Similarly, the subresource's type is also derived from the parent by prefixing
	// the parent's type to its type name, unless it starts with the provider definition,
	// i.e. is an absolute ARM type
	if parentType != "" {
		if !strings.HasPrefix(strings.ToLower(typestr), "microsoft.") {
			typestr = fmt.Sprintf("%s/%s", parentType, typestr)
		}
		args["type"] = typestr
	}

	if !isTemplateExpression(namestr) {
		prefix = []string{namestr}
	}

	varName := t.assignName(prefix)
	if el := t.lookup(varName); el != nil {
		// Don't expect name collision with resources
		return fmt.Errorf("another resource with name: %s already defined", name)
	}

	_, ok = args["apiVersion"]
	if !ok {
		return fmt.Errorf("required field 'apiVersion' missing for %s", name)
	}

	varName = t.recordCanonicalizedName(varName)
	// Since anything can be a template expression referring to other resources
	// we can't really take the name or the resource token for granted at
	// at this stage so we use the variable name as resource token temporarily
	r, err := newResource(varName, args, varName)
	if err != nil {
		return err
	}
	dep := newDependencyTracking(r)
	t.elements[varName] = dep

	if subResources, ok := args["resources"]; ok {
		subResourcesSlice, ok := subResources.([]interface{})
		if !ok {
			return errors.New("expect sub-resources to be a slice")
		}

		for _, sub := range subResourcesSlice {
			resMap, ok := sub.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expect resource body to be a map, got: %T", sub)
			}
			if err = t.handleAddResource(resMap, namestr, typestr); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *TemplateElements) GetDiagnostics() map[string][]Diagnostic {
	return t.diagnostics
}

func (t *TemplateElements) Validate(pkgSpec *schema.PackageSpec, metadata *provider.AzureAPIMetadata) error {
	resourceTokenConverter := newResourceTokenConverter(metadata)
	for _, el := range t.elements {
		ctx := pclRenderContext{
			metadata:               metadata,
			pkgSpec:                pkgSpec,
			dep:                    el,
			resourceTokenConverter: resourceTokenConverter,
		}
		if err := el.TemplateElement.FinishInitializing(&ctx, t); err != nil {
			diag := &Diagnostic{}
			if errors.As(err, &diag) {
				t.addDiagnostic(*diag)
			}
		}
	}
	return nil
}

func (t *TemplateElements) RenderPCL(
	metadata *provider.AzureAPIMetadata,
	pkgSpec *schema.PackageSpec,
) (*model.Body, error) {
	resourceTokenConverter := newResourceTokenConverter(metadata)
	elements, err := t.topologicalOrder()
	if err != nil {
		return nil, err
	}

	elements, err = t.filterUnusedElements(elements)
	if err != nil {
		return nil, err
	}

	var items []model.BodyItem
	for _, el := range elements {
		ctx := pclRenderContext{
			metadata:               metadata,
			pkgSpec:                pkgSpec,
			dep:                    el,
			resourceTokenConverter: resourceTokenConverter,
		}
		bodyItems, err := el.PCLExpression(&ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, bodyItems...)
	}

	return &model.Body{
		Items: items,
	}, nil
}

func (t *TemplateElements) topologicalOrder() ([]*dependencyTracking, error) {
	visited := map[*dependencyTracking]bool{}
	var orderedElements []*dependencyTracking
	keys := codegen.SortedKeys(t.elements)
	for _, k := range keys {
		v := t.elements[k]
		var err error
		var ordered []*dependencyTracking
		ordered, err = topoSort(v, visited, nil)
		if err != nil {
			return nil, err
		}

		orderedElements = append(orderedElements, ordered...)
	}
	return orderedElements, nil
}

func (t *TemplateElements) defaultResourceGroupName() (*model.Variable, *dependencyTracking, error) {
	rgName, _, err := t.ensureGlobals()
	if err != nil {
		return nil, nil, err
	}

	resourceGroupName := &model.Variable{
		Name:         "resourceGroupNameParam",
		VariableType: model.DynamicType,
	}
	return resourceGroupName, rgName, nil
}

func (t *TemplateElements) defaultResourceGroup() (*model.Variable, *dependencyTracking, error) {
	_, rgVar, err := t.ensureGlobals()
	if err != nil {
		return nil, nil, err
	}

	resourceGroupVar := &model.Variable{
		Name:         "resourceGroupVar",
		VariableType: model.DynamicType,
	}

	return resourceGroupVar, rgVar, nil
}

// We are primarily focussed on ARM templates targeting a specific resource group. ARM has a bunch of implicit
// behavior around filling in the default resource group name or referring to it through template functions.
// We place these defaults in place optimistically and then decide to eliminate them if they are not used at
// IR rendering time.
func (t *TemplateElements) ensureGlobals() (
	resourceGroupNameParam *dependencyTracking, resourceGroupVar *dependencyTracking, err error) {

	resourceGroupNameParam = t.lookup("resourceGroupNameParam")
	if resourceGroupNameParam == nil {
		if err = t.addParameter("resourceGroupNameParam", map[string]interface{}{
			"type": "string",
		}, false); err != nil {
			return
		}
		resourceGroupNameParam = t.lookup("resourceGroupNameParam")
	}

	resourceGroupName := &model.Variable{
		Name:         "resourceGroupNameParam",
		VariableType: model.DynamicType,
	}

	resourceGroupVar = t.lookup("resourceGroupVar")
	if resourceGroupVar == nil {
		invoke := pcl.Invoke("azure-nextgen:resources/latest:getResourceGroup",
			pcl.ObjectConsItem("resourceGroupName", model.VariableReference(resourceGroupName)))

		resourceGroupVar, err = t.addVariable("resourceGroupVar", invoke, false)
		// Add dependency information to make sure we emit the globals in PCL
		resourceGroupVar.RefersTo(resourceGroupNameParam)
	}
	return
}

func (t *TemplateElements) filterUnusedElements(elements []*dependencyTracking) ([]*dependencyTracking, error) {
	var filtered []*dependencyTracking
	_, rgVar, err := t.ensureGlobals()
	if err != nil {
		return nil, err
	}

	for _, el := range elements {
		if len(el.incoming) == 0 {
			switch el.TemplateElement.(type) {
			case *resource:
				if el.TemplateElement.Name() != rgVar.Name() {
					// Don't filter unused resources unless its the resource group we injected ourselves
					filtered = append(filtered, el)
				}
			case *output:
				filtered = append(filtered, el)
			}
			continue
		}
		filtered = append(filtered, el)
	}
	return filtered, nil
}

func extractDescription(info map[string]interface{}) string {
	if metadata, ok := info["metadata"]; ok {
		if metamap, ok := metadata.(map[string]interface{}); ok {
			if desc, ok := metamap["description"]; ok {
				return fmt.Sprintf("%s", desc)
			}
		}
	}
	return ""
}

// TODO: Looks like trivia is never serialized by HCL when rendered... :(
func getLeadingTriviaTokens(comment string) *syntax.BlockTokens {
	if comment == "" {
		return nil
	}
	return &syntax.BlockTokens{
		Type: syntax.Token{
			LeadingTrivia: syntax.TriviaList{&syntax.Comment{Lines: []string{comment}}},
		},
	}
}

func extractType(info map[string]interface{}) (string, error) {
	typ, ok := info["type"]
	if !ok {
		return "", errors.New("missing required field 'type'")
	}

	typeStr, ok := typ.(string)
	if !ok {
		return "", errors.New("expect type to be a string")
	}

	typeStr = strings.ToLower(typeStr)

	switch typeStr {
	case "int":
		return "number", nil
	case "bool", "string": //nolint:goconst
		return typeStr, nil
	case "securestring":
		return "string", nil
	case "secureobject":
		return "", nil
	default:
		return "", nil
	}
}
