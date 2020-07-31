// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Application Insights web test definition.
type WebTest struct {
	pulumi.CustomResourceState

	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Metadata describing a web test for an Azure resource.
	Properties WebTestPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebTest registers a new resource with the given unique name, arguments, and options.
func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WebTestArgs{}
	}
	var resource WebTest
	err := ctx.RegisterResource("azurerm:insights/v20150501:WebTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTest gets an existing WebTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTestState, opts ...pulumi.ResourceOption) (*WebTest, error) {
	var resource WebTest
	err := ctx.ReadResource("azurerm:insights/v20150501:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTest resources.
type webTestState struct {
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Metadata describing a web test for an Azure resource.
	Properties *WebTestPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type WebTestState struct {
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Metadata describing a web test for an Azure resource.
	Properties WebTestPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (WebTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestState)(nil)).Elem()
}

type webTestArgs struct {
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the Application Insights webtest resource.
	Name string `pulumi:"name"`
	// Metadata describing a web test for an Azure resource.
	Properties *WebTestProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebTest resource.
type WebTestArgs struct {
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringInput
	// The name of the Application Insights webtest resource.
	Name pulumi.StringInput
	// Metadata describing a web test for an Azure resource.
	Properties WebTestPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (WebTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestArgs)(nil)).Elem()
}