// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Route resource
type Route struct {
	pulumi.CustomResourceState

	// The destination CIDR to which the route applies.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress pulumi.StringPtrOutput `pulumi:"nextHopIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
	NextHopType pulumi.StringOutput `pulumi:"nextHopType"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil || args.NextHopType == nil {
		return nil, errors.New("missing required argument 'NextHopType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RouteName == nil {
		return nil, errors.New("missing required argument 'RouteName'")
	}
	if args == nil || args.RouteTableName == nil {
		return nil, errors.New("missing required argument 'RouteTableName'")
	}
	if args == nil {
		args = &RouteArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150501preview:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160601:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:Route"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:Route"),
		},
	})
	opts = append(opts, aliases)
	var resource Route
	err := ctx.RegisterResource("azurerm:network/v20181201:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azurerm:network/v20181201:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The destination CIDR to which the route applies.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
	NextHopType *string `pulumi:"nextHopType"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
}

type RouteState struct {
	// The destination CIDR to which the route applies.
	AddressPrefix pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress pulumi.StringPtrInput
	// The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
	NextHopType pulumi.StringPtrInput
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The destination CIDR to which the route applies.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
	// The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
	NextHopType string `pulumi:"nextHopType"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route.
	RouteName string `pulumi:"routeName"`
	// The name of the route table.
	RouteTableName string `pulumi:"routeTableName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The destination CIDR to which the route applies.
	AddressPrefix pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress pulumi.StringPtrInput
	// The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
	NextHopType pulumi.StringInput
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the route.
	RouteName pulumi.StringInput
	// The name of the route table.
	RouteTableName pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}
