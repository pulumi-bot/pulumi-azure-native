// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a namespace resource.
type Namespace struct {
	pulumi.CustomResourceState

	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the namespace.
	Properties NamespacePropertiesResponseOutput `pulumi:"properties"`
	// SKU of the namespace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
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
		args = &NamespaceArgs{}
	}
	var resource Namespace
	err := ctx.RegisterResource("azurerm:servicebus/v20150801:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azurerm:servicebus/v20150801:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Properties of the namespace.
	Properties *NamespacePropertiesResponse `pulumi:"properties"`
	// SKU of the namespace.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type NamespaceState struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Properties of the namespace.
	Properties NamespacePropertiesResponsePtrInput
	// SKU of the namespace.
	Sku SkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Indicates whether to create an ACS namespace.
	CreateACSNamespace *bool `pulumi:"createACSNamespace"`
	// Specifies whether this instance is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Namespace location.
	Location string `pulumi:"location"`
	// The namespace name.
	Name string `pulumi:"name"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of the namespace.
	Sku *Sku `pulumi:"sku"`
	// State of the namespace.
	Status *string `pulumi:"status"`
	// Namespace tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Indicates whether to create an ACS namespace.
	CreateACSNamespace pulumi.BoolPtrInput
	// Specifies whether this instance is enabled.
	Enabled pulumi.BoolPtrInput
	// Namespace location.
	Location pulumi.StringInput
	// The namespace name.
	Name pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// SKU of the namespace.
	Sku SkuPtrInput
	// State of the namespace.
	Status pulumi.StringPtrInput
	// Namespace tags.
	Tags pulumi.StringMapInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}
