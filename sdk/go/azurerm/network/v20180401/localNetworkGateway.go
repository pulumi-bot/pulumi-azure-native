// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A common class for general resource information
//
// ## Example Usage
// ### CreateLocalNetworkGateway
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20180401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewLocalNetworkGateway(ctx, "localNetworkGateway", &network.LocalNetworkGatewayArgs{
// 			GatewayIpAddress: pulumi.String("x.x.x.x"),
// 			LocalNetworkAddressSpace: &network.AddressSpaceArgs{
// 				AddressPrefixes: pulumi.StringArray{
// 					pulumi.String("10.1.0.0/16"),
// 				},
// 			},
// 			LocalNetworkGatewayName: pulumi.String("localgw"),
// 			Location:                pulumi.String("Central US"),
// 			ResourceGroupName:       pulumi.String("rg1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type LocalNetworkGateway struct {
	pulumi.CustomResourceState

	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsResponsePtrOutput `pulumi:"bgpSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// IP address of local network gateway.
	GatewayIpAddress pulumi.StringPtrOutput `pulumi:"gatewayIpAddress"`
	// Local network site address space.
	LocalNetworkAddressSpace AddressSpaceResponsePtrOutput `pulumi:"localNetworkAddressSpace"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the LocalNetworkGateway resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLocalNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewLocalNetworkGateway(ctx *pulumi.Context,
	name string, args *LocalNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	if args == nil || args.LocalNetworkGatewayName == nil {
		return nil, errors.New("missing required argument 'LocalNetworkGatewayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LocalNetworkGatewayArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:LocalNetworkGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource LocalNetworkGateway
	err := ctx.RegisterResource("azurerm:network/v20180401:LocalNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNetworkGateway gets an existing LocalNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNetworkGatewayState, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	var resource LocalNetworkGateway
	err := ctx.ReadResource("azurerm:network/v20180401:LocalNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNetworkGateway resources.
type localNetworkGatewayState struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettingsResponse `pulumi:"bgpSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// IP address of local network gateway.
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Local network site address space.
	LocalNetworkAddressSpace *AddressSpaceResponse `pulumi:"localNetworkAddressSpace"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource GUID property of the LocalNetworkGateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type LocalNetworkGatewayState struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// IP address of local network gateway.
	GatewayIpAddress pulumi.StringPtrInput
	// Local network site address space.
	LocalNetworkAddressSpace AddressSpaceResponsePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The resource GUID property of the LocalNetworkGateway resource.
	ResourceGuid pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (LocalNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayState)(nil)).Elem()
}

type localNetworkGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// IP address of local network gateway.
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Local network site address space.
	LocalNetworkAddressSpace *AddressSpace `pulumi:"localNetworkAddressSpace"`
	// The name of the local network gateway.
	LocalNetworkGatewayName string `pulumi:"localNetworkGatewayName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the LocalNetworkGateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LocalNetworkGateway resource.
type LocalNetworkGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// IP address of local network gateway.
	GatewayIpAddress pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Local network site address space.
	LocalNetworkAddressSpace AddressSpacePtrInput
	// The name of the local network gateway.
	LocalNetworkGatewayName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the LocalNetworkGateway resource.
	ResourceGuid pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LocalNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayArgs)(nil)).Elem()
}
