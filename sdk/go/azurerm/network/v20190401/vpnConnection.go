// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnConnection Resource.
//
// ## Example Usage
// ### VpnConnectionPut
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20190401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewVpnConnection(ctx, "vpnConnection", &network.VpnConnectionArgs{
// 			ConnectionName: pulumi.String("vpnConnection1"),
// 			GatewayName:    pulumi.String("gateway1"),
// 			RemoteVpnSite: &network.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
// 			},
// 			ResourceGroupName:         pulumi.String("rg1"),
// 			SharedKey:                 pulumi.String("key"),
// 			VpnConnectionProtocolType: pulumi.String("IKEv1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type VpnConnection struct {
	pulumi.CustomResourceState

	// Expected bandwidth in MBPS.
	ConnectionBandwidth pulumi.IntPtrOutput `pulumi:"connectionBandwidth"`
	// The connection status.
	ConnectionStatus pulumi.StringPtrOutput `pulumi:"connectionStatus"`
	// Egress bytes transferred.
	EgressBytesTransferred pulumi.IntOutput `pulumi:"egressBytesTransferred"`
	// EnableBgp flag.
	EnableBgp pulumi.BoolPtrOutput `pulumi:"enableBgp"`
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrOutput `pulumi:"enableInternetSecurity"`
	// EnableBgp flag.
	EnableRateLimiting pulumi.BoolPtrOutput `pulumi:"enableRateLimiting"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Ingress bytes transferred.
	IngressBytesTransferred pulumi.IntOutput `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyResponseArrayOutput `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite SubResourceResponsePtrOutput `pulumi:"remoteVpnSite"`
	// Routing weight for vpn connection.
	RoutingWeight pulumi.IntPtrOutput `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
	// Use local azure ip to initiate connection.
	UseLocalAzureIpAddress pulumi.BoolPtrOutput `pulumi:"useLocalAzureIpAddress"`
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrOutput `pulumi:"usePolicyBasedTrafficSelectors"`
	// Connection protocol used for this connection.
	VpnConnectionProtocolType pulumi.StringPtrOutput `pulumi:"vpnConnectionProtocolType"`
}

// NewVpnConnection registers a new resource with the given unique name, arguments, and options.
func NewVpnConnection(ctx *pulumi.Context,
	name string, args *VpnConnectionArgs, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	if args == nil || args.ConnectionName == nil {
		return nil, errors.New("missing required argument 'ConnectionName'")
	}
	if args == nil || args.GatewayName == nil {
		return nil, errors.New("missing required argument 'GatewayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VpnConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:VpnConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:VpnConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnConnection
	err := ctx.RegisterResource("azurerm:network/v20190401:VpnConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnConnection gets an existing VpnConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnConnectionState, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	var resource VpnConnection
	err := ctx.ReadResource("azurerm:network/v20190401:VpnConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnConnection resources.
type vpnConnectionState struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth *int `pulumi:"connectionBandwidth"`
	// The connection status.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// Egress bytes transferred.
	EgressBytesTransferred *int `pulumi:"egressBytesTransferred"`
	// EnableBgp flag.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// EnableBgp flag.
	EnableRateLimiting *bool `pulumi:"enableRateLimiting"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Ingress bytes transferred.
	IngressBytesTransferred *int `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResourceResponse `pulumi:"remoteVpnSite"`
	// Routing weight for vpn connection.
	RoutingWeight *int `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey *string `pulumi:"sharedKey"`
	// Use local azure ip to initiate connection.
	UseLocalAzureIpAddress *bool `pulumi:"useLocalAzureIpAddress"`
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors *bool `pulumi:"usePolicyBasedTrafficSelectors"`
	// Connection protocol used for this connection.
	VpnConnectionProtocolType *string `pulumi:"vpnConnectionProtocolType"`
}

type VpnConnectionState struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth pulumi.IntPtrInput
	// The connection status.
	ConnectionStatus pulumi.StringPtrInput
	// Egress bytes transferred.
	EgressBytesTransferred pulumi.IntPtrInput
	// EnableBgp flag.
	EnableBgp pulumi.BoolPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// EnableBgp flag.
	EnableRateLimiting pulumi.BoolPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Ingress bytes transferred.
	IngressBytesTransferred pulumi.IntPtrInput
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyResponseArrayInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Id of the connected vpn site.
	RemoteVpnSite SubResourceResponsePtrInput
	// Routing weight for vpn connection.
	RoutingWeight pulumi.IntPtrInput
	// SharedKey for the vpn connection.
	SharedKey pulumi.StringPtrInput
	// Use local azure ip to initiate connection.
	UseLocalAzureIpAddress pulumi.BoolPtrInput
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput
	// Connection protocol used for this connection.
	VpnConnectionProtocolType pulumi.StringPtrInput
}

func (VpnConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionState)(nil)).Elem()
}

type vpnConnectionArgs struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth *int `pulumi:"connectionBandwidth"`
	// The name of the connection.
	ConnectionName string `pulumi:"connectionName"`
	// The connection status.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// EnableBgp flag.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// EnableBgp flag.
	EnableRateLimiting *bool `pulumi:"enableRateLimiting"`
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicy `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResource `pulumi:"remoteVpnSite"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Routing weight for vpn connection.
	RoutingWeight *int `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey *string `pulumi:"sharedKey"`
	// Use local azure ip to initiate connection.
	UseLocalAzureIpAddress *bool `pulumi:"useLocalAzureIpAddress"`
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors *bool `pulumi:"usePolicyBasedTrafficSelectors"`
	// Connection protocol used for this connection.
	VpnConnectionProtocolType *string `pulumi:"vpnConnectionProtocolType"`
}

// The set of arguments for constructing a VpnConnection resource.
type VpnConnectionArgs struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth pulumi.IntPtrInput
	// The name of the connection.
	ConnectionName pulumi.StringInput
	// The connection status.
	ConnectionStatus pulumi.StringPtrInput
	// EnableBgp flag.
	EnableBgp pulumi.BoolPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// EnableBgp flag.
	EnableRateLimiting pulumi.BoolPtrInput
	// The name of the gateway.
	GatewayName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyArrayInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Id of the connected vpn site.
	RemoteVpnSite SubResourcePtrInput
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput
	// Routing weight for vpn connection.
	RoutingWeight pulumi.IntPtrInput
	// SharedKey for the vpn connection.
	SharedKey pulumi.StringPtrInput
	// Use local azure ip to initiate connection.
	UseLocalAzureIpAddress pulumi.BoolPtrInput
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput
	// Connection protocol used for this connection.
	VpnConnectionProtocolType pulumi.StringPtrInput
}

func (VpnConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionArgs)(nil)).Elem()
}
