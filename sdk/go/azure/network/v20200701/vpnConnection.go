// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnConnection Resource.
type VpnConnection struct {
	pulumi.CustomResourceState

	// Expected bandwidth in MBPS.
	ConnectionBandwidth pulumi.IntPtrOutput `pulumi:"connectionBandwidth"`
	// The connection status.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// DPD timeout in seconds for vpn connection.
	DpdTimeoutSeconds pulumi.IntPtrOutput `pulumi:"dpdTimeoutSeconds"`
	// Egress bytes transferred.
	EgressBytesTransferred pulumi.Float64Output `pulumi:"egressBytesTransferred"`
	// EnableBgp flag.
	EnableBgp pulumi.BoolPtrOutput `pulumi:"enableBgp"`
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrOutput `pulumi:"enableInternetSecurity"`
	// EnableBgp flag.
	EnableRateLimiting pulumi.BoolPtrOutput `pulumi:"enableRateLimiting"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Ingress bytes transferred.
	IngressBytesTransferred pulumi.Float64Output `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyResponseArrayOutput `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the VPN connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite SubResourceResponsePtrOutput `pulumi:"remoteVpnSite"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationResponsePtrOutput `pulumi:"routingConfiguration"`
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
	// List of all vpn site link connections to the gateway.
	VpnLinkConnections VpnSiteLinkConnectionResponseArrayOutput `pulumi:"vpnLinkConnections"`
}

// NewVpnConnection registers a new resource with the given unique name, arguments, and options.
func NewVpnConnection(ctx *pulumi.Context,
	name string, args *VpnConnectionArgs, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VpnConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnConnection
	err := ctx.RegisterResource("azure-nextgen:network/v20200701:VpnConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:network/v20200701:VpnConnection", name, id, state, &resource, opts...)
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
	// DPD timeout in seconds for vpn connection.
	DpdTimeoutSeconds *int `pulumi:"dpdTimeoutSeconds"`
	// Egress bytes transferred.
	EgressBytesTransferred *float64 `pulumi:"egressBytesTransferred"`
	// EnableBgp flag.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// EnableBgp flag.
	EnableRateLimiting *bool `pulumi:"enableRateLimiting"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Ingress bytes transferred.
	IngressBytesTransferred *float64 `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the VPN connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResourceResponse `pulumi:"remoteVpnSite"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
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
	VpnLinkConnections []VpnSiteLinkConnectionResponse `pulumi:"vpnLinkConnections"`
}

type VpnConnectionState struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth pulumi.IntPtrInput
	// The connection status.
	ConnectionStatus pulumi.StringPtrInput
	// DPD timeout in seconds for vpn connection.
	DpdTimeoutSeconds pulumi.IntPtrInput
	// Egress bytes transferred.
	EgressBytesTransferred pulumi.Float64PtrInput
	// EnableBgp flag.
	EnableBgp pulumi.BoolPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// EnableBgp flag.
	EnableRateLimiting pulumi.BoolPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Ingress bytes transferred.
	IngressBytesTransferred pulumi.Float64PtrInput
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyResponseArrayInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the VPN connection resource.
	ProvisioningState pulumi.StringPtrInput
	// Id of the connected vpn site.
	RemoteVpnSite SubResourceResponsePtrInput
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationResponsePtrInput
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
	VpnLinkConnections VpnSiteLinkConnectionResponseArrayInput
}

func (VpnConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionState)(nil)).Elem()
}

type vpnConnectionArgs struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidth *int `pulumi:"connectionBandwidth"`
	// The name of the connection.
	ConnectionName string `pulumi:"connectionName"`
	// DPD timeout in seconds for vpn connection.
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
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResource `pulumi:"remoteVpnSite"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfiguration `pulumi:"routingConfiguration"`
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
	// The name of the connection.
	ConnectionName pulumi.StringInput
	// DPD timeout in seconds for vpn connection.
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
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Id of the connected vpn site.
	RemoteVpnSite SubResourcePtrInput
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationPtrInput
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

type VpnConnectionInput interface {
	pulumi.Input

	ToVpnConnectionOutput() VpnConnectionOutput
	ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput
}

func (*VpnConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnection)(nil))
}

func (i *VpnConnection) ToVpnConnectionOutput() VpnConnectionOutput {
	return i.ToVpnConnectionOutputWithContext(context.Background())
}

func (i *VpnConnection) ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionOutput)
}

type VpnConnectionOutput struct {
	*pulumi.OutputState
}

func (VpnConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnection)(nil))
}

func (o VpnConnectionOutput) ToVpnConnectionOutput() VpnConnectionOutput {
	return o
}

func (o VpnConnectionOutput) ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnConnectionOutput{})
}
