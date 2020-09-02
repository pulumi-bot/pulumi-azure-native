// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Binding resource payload
type Binding struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Binding resource
	Properties BindingResourcePropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil || args.AppName == nil {
		return nil, errors.New("missing required argument 'AppName'")
	}
	if args == nil || args.BindingName == nil {
		return nil, errors.New("missing required argument 'BindingName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &BindingArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:appplatform/v20190501preview:Binding"),
		},
		{
			Type: pulumi.String("azurerm:appplatform/v20200701:Binding"),
		},
	})
	opts = append(opts, aliases)
	var resource Binding
	err := ctx.RegisterResource("azurerm:appplatform/latest:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("azurerm:appplatform/latest:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Binding resources.
type bindingState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the Binding resource
	Properties *BindingResourcePropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type BindingState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the Binding resource
	Properties BindingResourcePropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Binding resource.
	BindingName string `pulumi:"bindingName"`
	// Properties of the Binding resource
	Properties *BindingResourceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput
	// The name of the Binding resource.
	BindingName pulumi.StringInput
	// Properties of the Binding resource
	Properties BindingResourcePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}
