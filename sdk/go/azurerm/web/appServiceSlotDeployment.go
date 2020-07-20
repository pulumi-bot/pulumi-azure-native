// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// User credentials used for publishing activity.
type AppServiceSlotDeployment struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deployment resource specific properties
	Properties DeploymentResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServiceSlotDeployment registers a new resource with the given unique name, arguments, and options.
func NewAppServiceSlotDeployment(ctx *pulumi.Context,
	name string, args *AppServiceSlotDeploymentArgs, opts ...pulumi.ResourceOption) (*AppServiceSlotDeployment, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &AppServiceSlotDeploymentArgs{}
	}
	var resource AppServiceSlotDeployment
	err := ctx.RegisterResource("azurerm:web:AppServiceSlotDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceSlotDeployment gets an existing AppServiceSlotDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceSlotDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceSlotDeploymentState, opts ...pulumi.ResourceOption) (*AppServiceSlotDeployment, error) {
	var resource AppServiceSlotDeployment
	err := ctx.ReadResource("azurerm:web:AppServiceSlotDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceSlotDeployment resources.
type appServiceSlotDeploymentState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Deployment resource specific properties
	Properties *DeploymentResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServiceSlotDeploymentState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Deployment resource specific properties
	Properties DeploymentResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServiceSlotDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceSlotDeploymentState)(nil)).Elem()
}

type appServiceSlotDeploymentArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// ID of an existing deployment.
	Name string `pulumi:"name"`
	// Deployment resource specific properties
	Properties *DeploymentProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API creates a deployment for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a AppServiceSlotDeployment resource.
type AppServiceSlotDeploymentArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// ID of an existing deployment.
	Name pulumi.StringInput
	// Deployment resource specific properties
	Properties DeploymentPropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API creates a deployment for the production slot.
	Slot pulumi.StringInput
}

func (AppServiceSlotDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceSlotDeploymentArgs)(nil)).Elem()
}
