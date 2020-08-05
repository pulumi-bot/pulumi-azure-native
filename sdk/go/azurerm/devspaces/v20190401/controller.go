// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Controller struct {
	pulumi.CustomResourceState

	// Region where the Azure resource is located.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ControllerPropertiesResponseOutput `pulumi:"properties"`
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Tags for the Azure resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewController registers a new resource with the given unique name, arguments, and options.
func NewController(ctx *pulumi.Context,
	name string, args *ControllerArgs, opts ...pulumi.ResourceOption) (*Controller, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil || args.TargetContainerHostCredentialsBase64 == nil {
		return nil, errors.New("missing required argument 'TargetContainerHostCredentialsBase64'")
	}
	if args == nil || args.TargetContainerHostResourceId == nil {
		return nil, errors.New("missing required argument 'TargetContainerHostResourceId'")
	}
	if args == nil {
		args = &ControllerArgs{}
	}
	var resource Controller
	err := ctx.RegisterResource("azurerm:devspaces/v20190401:Controller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetController gets an existing Controller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControllerState, opts ...pulumi.ResourceOption) (*Controller, error) {
	var resource Controller
	err := ctx.ReadResource("azurerm:devspaces/v20190401:Controller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Controller resources.
type controllerState struct {
	// Region where the Azure resource is located.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name       *string                       `pulumi:"name"`
	Properties *ControllerPropertiesResponse `pulumi:"properties"`
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku *SkuResponse `pulumi:"sku"`
	// Tags for the Azure resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ControllerState struct {
	// Region where the Azure resource is located.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name       pulumi.StringPtrInput
	Properties ControllerPropertiesResponsePtrInput
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku SkuResponsePtrInput
	// Tags for the Azure resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerState)(nil)).Elem()
}

type controllerArgs struct {
	// Region where the Azure resource is located.
	Location string `pulumi:"location"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku Sku `pulumi:"sku"`
	// Tags for the Azure resource.
	Tags map[string]string `pulumi:"tags"`
	// Credentials of the target container host (base64).
	TargetContainerHostCredentialsBase64 string `pulumi:"targetContainerHostCredentialsBase64"`
	// Resource ID of the target container host
	TargetContainerHostResourceId string `pulumi:"targetContainerHostResourceId"`
}

// The set of arguments for constructing a Controller resource.
type ControllerArgs struct {
	// Region where the Azure resource is located.
	Location pulumi.StringInput
	// Name of the resource.
	Name pulumi.StringInput
	// Resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku SkuInput
	// Tags for the Azure resource.
	Tags pulumi.StringMapInput
	// Credentials of the target container host (base64).
	TargetContainerHostCredentialsBase64 pulumi.StringInput
	// Resource ID of the target container host
	TargetContainerHostResourceId pulumi.StringInput
}

func (ControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerArgs)(nil)).Elem()
}
