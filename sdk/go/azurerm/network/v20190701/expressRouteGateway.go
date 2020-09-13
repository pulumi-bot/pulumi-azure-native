// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ExpressRoute gateway resource.
//
// ## Example Usage
// ### ExpressRouteGatewayCreate
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20190701"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewExpressRouteGateway(ctx, "expressRouteGateway", &network.ExpressRouteGatewayArgs{
// 			AutoScaleConfiguration: &network.ExpressRouteGatewayPropertiesAutoScaleConfigurationArgs{
// 				Bounds: &network.ExpressRouteGatewayPropertiesBoundsArgs{
// 					Min: pulumi.Int(3),
// 				},
// 			},
// 			ExpressRouteGatewayName: pulumi.String("gateway-2"),
// 			Location:                pulumi.String("westus"),
// 			ResourceGroupName:       pulumi.String("resourceGroupName"),
// 			VirtualHub: &network.VirtualHubIdArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupId/providers/Microsoft.Network/virtualHubs/virtualHubName"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ExpressRouteGateway struct {
	pulumi.CustomResourceState

	// Configuration for auto scaling.
	AutoScaleConfiguration ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput `pulumi:"autoScaleConfiguration"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of ExpressRoute connections to the ExpressRoute gateway.
	ExpressRouteConnections ExpressRouteConnectionResponseArrayOutput `pulumi:"expressRouteConnections"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the express route gateway resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubIdResponseOutput `pulumi:"virtualHub"`
}

// NewExpressRouteGateway registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteGateway(ctx *pulumi.Context,
	name string, args *ExpressRouteGatewayArgs, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	if args == nil || args.ExpressRouteGatewayName == nil {
		return nil, errors.New("missing required argument 'ExpressRouteGatewayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualHub == nil {
		return nil, errors.New("missing required argument 'VirtualHub'")
	}
	if args == nil {
		args = &ExpressRouteGatewayArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:ExpressRouteGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteGateway
	err := ctx.RegisterResource("azurerm:network/v20190701:ExpressRouteGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteGateway gets an existing ExpressRouteGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteGatewayState, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	var resource ExpressRouteGateway
	err := ctx.ReadResource("azurerm:network/v20190701:ExpressRouteGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteGateway resources.
type expressRouteGatewayState struct {
	// Configuration for auto scaling.
	AutoScaleConfiguration *ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// List of ExpressRoute connections to the ExpressRoute gateway.
	ExpressRouteConnections []ExpressRouteConnectionResponse `pulumi:"expressRouteConnections"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the express route gateway resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub *VirtualHubIdResponse `pulumi:"virtualHub"`
}

type ExpressRouteGatewayState struct {
	// Configuration for auto scaling.
	AutoScaleConfiguration ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// List of ExpressRoute connections to the ExpressRoute gateway.
	ExpressRouteConnections ExpressRouteConnectionResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the express route gateway resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubIdResponsePtrInput
}

func (ExpressRouteGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayState)(nil)).Elem()
}

type expressRouteGatewayArgs struct {
	// Configuration for auto scaling.
	AutoScaleConfiguration *ExpressRouteGatewayPropertiesAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The provisioning state of the express route gateway resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubId `pulumi:"virtualHub"`
}

// The set of arguments for constructing a ExpressRouteGateway resource.
type ExpressRouteGatewayArgs struct {
	// Configuration for auto scaling.
	AutoScaleConfiguration ExpressRouteGatewayPropertiesAutoScaleConfigurationPtrInput
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The provisioning state of the express route gateway resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubIdInput
}

func (ExpressRouteGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayArgs)(nil)).Elem()
}
