// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

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
	ConnectionBandwidthInMbps pulumi.IntOutput `pulumi:"connectionBandwidthInMbps"`
	// The connection status.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// Egress bytes transferred.
	EgressBytesTransferred pulumi.Float64Output `pulumi:"egressBytesTransferred"`
	// EnableBgp flag
	EnableBgp pulumi.BoolPtrOutput `pulumi:"enableBgp"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Ingress bytes transferred.
	IngressBytesTransferred pulumi.Float64Output `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyResponseArrayOutput `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite SubResourceResponsePtrOutput `pulumi:"remoteVpnSite"`
	// routing weight for vpn connection.
	RoutingWeight pulumi.IntPtrOutput `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
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
			Type: pulumi.String("azure-nextgen:network/v20200701:VpnConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnConnection
	err := ctx.RegisterResource("azure-nextgen:network/v20180701:VpnConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:network/v20180701:VpnConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnConnection resources.
type vpnConnectionState struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidthInMbps *int `pulumi:"connectionBandwidthInMbps"`
	// The connection status.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// Egress bytes transferred.
	EgressBytesTransferred *float64 `pulumi:"egressBytesTransferred"`
	// EnableBgp flag
	EnableBgp *bool `pulumi:"enableBgp"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Ingress bytes transferred.
	IngressBytesTransferred *float64 `pulumi:"ingressBytesTransferred"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Id of the connected vpn site.
	RemoteVpnSite *SubResourceResponse `pulumi:"remoteVpnSite"`
	// routing weight for vpn connection.
	RoutingWeight *int `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey *string `pulumi:"sharedKey"`
}

type VpnConnectionState struct {
	// Expected bandwidth in MBPS.
	ConnectionBandwidthInMbps pulumi.IntPtrInput
	// The connection status.
	ConnectionStatus pulumi.StringPtrInput
	// Egress bytes transferred.
	EgressBytesTransferred pulumi.Float64PtrInput
	// EnableBgp flag
	EnableBgp pulumi.BoolPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Ingress bytes transferred.
	IngressBytesTransferred pulumi.Float64PtrInput
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies IpsecPolicyResponseArrayInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Id of the connected vpn site.
	RemoteVpnSite SubResourceResponsePtrInput
	// routing weight for vpn connection.
	RoutingWeight pulumi.IntPtrInput
	// SharedKey for the vpn connection.
	SharedKey pulumi.StringPtrInput
}

func (VpnConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionState)(nil)).Elem()
}

type vpnConnectionArgs struct {
	// The name of the connection.
	ConnectionName string `pulumi:"connectionName"`
	// EnableBgp flag
	EnableBgp *bool `pulumi:"enableBgp"`
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
	// routing weight for vpn connection.
	RoutingWeight *int `pulumi:"routingWeight"`
	// SharedKey for the vpn connection.
	SharedKey *string `pulumi:"sharedKey"`
}

// The set of arguments for constructing a VpnConnection resource.
type VpnConnectionArgs struct {
	// The name of the connection.
	ConnectionName pulumi.StringInput
	// EnableBgp flag
	EnableBgp pulumi.BoolPtrInput
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
	// routing weight for vpn connection.
	RoutingWeight pulumi.IntPtrInput
	// SharedKey for the vpn connection.
	SharedKey pulumi.StringPtrInput
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
