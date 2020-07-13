// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A hostname binding object.
type AppServiceHostNameBinding struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// HostNameBinding resource specific properties
	Properties HostNameBindingResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServiceHostNameBinding registers a new resource with the given unique name, arguments, and options.
func NewAppServiceHostNameBinding(ctx *pulumi.Context,
	name string, args *AppServiceHostNameBindingArgs, opts ...pulumi.ResourceOption) (*AppServiceHostNameBinding, error) {
	if args == nil || args.HostName == nil {
		return nil, errors.New("missing required argument 'HostName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AppServiceHostNameBindingArgs{}
	}
	var resource AppServiceHostNameBinding
	err := ctx.RegisterResource("azurerm:web:AppServiceHostNameBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceHostNameBinding gets an existing AppServiceHostNameBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceHostNameBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceHostNameBindingState, opts ...pulumi.ResourceOption) (*AppServiceHostNameBinding, error) {
	var resource AppServiceHostNameBinding
	err := ctx.ReadResource("azurerm:web:AppServiceHostNameBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceHostNameBinding resources.
type appServiceHostNameBindingState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// HostNameBinding resource specific properties
	Properties *HostNameBindingResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServiceHostNameBindingState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// HostNameBinding resource specific properties
	Properties HostNameBindingResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServiceHostNameBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceHostNameBindingState)(nil)).Elem()
}

type appServiceHostNameBindingArgs struct {
	// Hostname in the hostname binding.
	HostName string `pulumi:"hostName"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// HostNameBinding resource specific properties
	Properties *HostNameBindingProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AppServiceHostNameBinding resource.
type AppServiceHostNameBindingArgs struct {
	// Hostname in the hostname binding.
	HostName pulumi.StringInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// HostNameBinding resource specific properties
	Properties HostNameBindingPropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (AppServiceHostNameBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceHostNameBindingArgs)(nil)).Elem()
}