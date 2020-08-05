// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnConnection Resource.
type VpnConnection struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the VPN connection.
	Properties VpnConnectionPropertiesResponseOutput `pulumi:"properties"`
}

// NewVpnConnection registers a new resource with the given unique name, arguments, and options.
func NewVpnConnection(ctx *pulumi.Context,
	name string, args *VpnConnectionArgs, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	if args == nil || args.GatewayName == nil {
		return nil, errors.New("missing required argument 'GatewayName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VpnConnectionArgs{}
	}
	var resource VpnConnection
	err := ctx.RegisterResource("azurerm:network/v20200301:VpnConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20200301:VpnConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnConnection resources.
type vpnConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the VPN connection.
	Properties *VpnConnectionPropertiesResponse `pulumi:"properties"`
}

type VpnConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the VPN connection.
	Properties VpnConnectionPropertiesResponsePtrInput
}

func (VpnConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionState)(nil)).Elem()
}

type vpnConnectionArgs struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth *int `pulumi:"connectionBandwidth"`
	// The connection status.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// The dead peer detection timeout for a vpn connection in seconds.
	DpdTimeoutSeconds *int `pulumi:"dpdTimeoutSeconds"`
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
	// The name of the connection.
	Name string `pulumi:"name"`
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
	// List of all vpn site link connections to the gateway.
	VpnLinkConnections []VpnSiteLinkConnection `pulumi:"vpnLinkConnections"`
}

// The set of arguments for constructing a VpnConnection resource.
type VpnConnectionArgs struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth pulumi.IntPtrInput
	// The connection status.
	ConnectionStatus pulumi.StringPtrInput
	// The dead peer detection timeout for a vpn connection in seconds.
	DpdTimeoutSeconds pulumi.IntPtrInput
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
	// The name of the connection.
	Name pulumi.StringInput
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
	// List of all vpn site link connections to the gateway.
	VpnLinkConnections VpnSiteLinkConnectionArrayInput
}

func (VpnConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionArgs)(nil)).Elem()
}
