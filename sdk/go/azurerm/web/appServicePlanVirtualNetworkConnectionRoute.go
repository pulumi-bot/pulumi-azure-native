// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Virtual Network route contract used to pass routing information for a Virtual Network.
type AppServicePlanVirtualNetworkConnectionRoute struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VnetRoute resource specific properties
	Properties VnetRouteResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServicePlanVirtualNetworkConnectionRoute registers a new resource with the given unique name, arguments, and options.
func NewAppServicePlanVirtualNetworkConnectionRoute(ctx *pulumi.Context,
	name string, args *AppServicePlanVirtualNetworkConnectionRouteArgs, opts ...pulumi.ResourceOption) (*AppServicePlanVirtualNetworkConnectionRoute, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VnetName == nil {
		return nil, errors.New("missing required argument 'VnetName'")
	}
	if args == nil {
		args = &AppServicePlanVirtualNetworkConnectionRouteArgs{}
	}
	var resource AppServicePlanVirtualNetworkConnectionRoute
	err := ctx.RegisterResource("azurerm:web:AppServicePlanVirtualNetworkConnectionRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServicePlanVirtualNetworkConnectionRoute gets an existing AppServicePlanVirtualNetworkConnectionRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServicePlanVirtualNetworkConnectionRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServicePlanVirtualNetworkConnectionRouteState, opts ...pulumi.ResourceOption) (*AppServicePlanVirtualNetworkConnectionRoute, error) {
	var resource AppServicePlanVirtualNetworkConnectionRoute
	err := ctx.ReadResource("azurerm:web:AppServicePlanVirtualNetworkConnectionRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServicePlanVirtualNetworkConnectionRoute resources.
type appServicePlanVirtualNetworkConnectionRouteState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// VnetRoute resource specific properties
	Properties *VnetRouteResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServicePlanVirtualNetworkConnectionRouteState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// VnetRoute resource specific properties
	Properties VnetRouteResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServicePlanVirtualNetworkConnectionRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanVirtualNetworkConnectionRouteState)(nil)).Elem()
}

type appServicePlanVirtualNetworkConnectionRouteArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the Virtual Network route.
	Name string `pulumi:"name"`
	// VnetRoute resource specific properties
	Properties *VnetRouteProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Virtual Network.
	VnetName string `pulumi:"vnetName"`
}

// The set of arguments for constructing a AppServicePlanVirtualNetworkConnectionRoute resource.
type AppServicePlanVirtualNetworkConnectionRouteArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the Virtual Network route.
	Name pulumi.StringInput
	// VnetRoute resource specific properties
	Properties VnetRoutePropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the Virtual Network.
	VnetName pulumi.StringInput
}

func (AppServicePlanVirtualNetworkConnectionRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanVirtualNetworkConnectionRouteArgs)(nil)).Elem()
}
