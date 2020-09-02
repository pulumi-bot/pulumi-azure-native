// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peering in an ExpressRouteCircuit resource.
type ExpressRouteCircuitPeering struct {
	pulumi.CustomResourceState

	// The Azure ASN.
	AzureASN pulumi.IntPtrOutput `pulumi:"azureASN"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections ExpressRouteCircuitConnectionResponseArrayOutput `pulumi:"connections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The ExpressRoute connection.
	ExpressRouteConnection ExpressRouteConnectionIdResponsePtrOutput `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrOutput `pulumi:"gatewayManagerEtag"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The peer ASN.
	PeerASN pulumi.IntPtrOutput `pulumi:"peerASN"`
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections PeerExpressRouteCircuitConnectionResponseArrayOutput `pulumi:"peeredConnections"`
	// The peering type.
	PeeringType pulumi.StringPtrOutput `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort pulumi.StringPtrOutput `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrOutput `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The reference of the RouteFilter resource.
	RouteFilter SubResourceResponsePtrOutput `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort pulumi.StringPtrOutput `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrOutput `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
	// The peering state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The peering stats of express route circuit.
	Stats ExpressRouteCircuitStatsResponsePtrOutput `pulumi:"stats"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VLAN ID.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
}

// NewExpressRouteCircuitPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	if args == nil || args.CircuitName == nil {
		return nil, errors.New("missing required argument 'CircuitName'")
	}
	if args == nil || args.PeeringName == nil {
		return nil, errors.New("missing required argument 'PeeringName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCircuitPeeringArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150501preview:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:ExpressRouteCircuitPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuitPeering
	err := ctx.RegisterResource("azurerm:network/v20190701:ExpressRouteCircuitPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitPeering gets an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	var resource ExpressRouteCircuitPeering
	err := ctx.ReadResource("azurerm:network/v20190701:ExpressRouteCircuitPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitPeering resources.
type expressRouteCircuitPeeringState struct {
	// The Azure ASN.
	AzureASN *int `pulumi:"azureASN"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections []ExpressRouteCircuitConnectionResponse `pulumi:"connections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The ExpressRoute connection.
	ExpressRouteConnection *ExpressRouteConnectionIdResponse `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *int `pulumi:"peerASN"`
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections []PeerExpressRouteCircuitConnectionResponse `pulumi:"peeredConnections"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The reference of the RouteFilter resource.
	RouteFilter *SubResourceResponse `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The peering stats of express route circuit.
	Stats *ExpressRouteCircuitStatsResponse `pulumi:"stats"`
	// Type of the resource.
	Type *string `pulumi:"type"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

type ExpressRouteCircuitPeeringState struct {
	// The Azure ASN.
	AzureASN pulumi.IntPtrInput
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections ExpressRouteCircuitConnectionResponseArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The ExpressRoute connection.
	ExpressRouteConnection ExpressRouteConnectionIdResponsePtrInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigResponsePtrInput
	// Who was the last to modify the peering.
	LastModifiedBy pulumi.StringPtrInput
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigResponsePtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The peer ASN.
	PeerASN pulumi.IntPtrInput
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections PeerExpressRouteCircuitConnectionResponseArrayInput
	// The peering type.
	PeeringType pulumi.StringPtrInput
	// The primary port.
	PrimaryAzurePort pulumi.StringPtrInput
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState pulumi.StringPtrInput
	// The reference of the RouteFilter resource.
	RouteFilter SubResourceResponsePtrInput
	// The secondary port.
	SecondaryAzurePort pulumi.StringPtrInput
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key.
	SharedKey pulumi.StringPtrInput
	// The peering state.
	State pulumi.StringPtrInput
	// The peering stats of express route circuit.
	Stats ExpressRouteCircuitStatsResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCircuitPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringState)(nil)).Elem()
}

type expressRouteCircuitPeeringArgs struct {
	// The Azure ASN.
	AzureASN *int `pulumi:"azureASN"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections []ExpressRouteCircuitConnectionType `pulumi:"connections"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *int `pulumi:"peerASN"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference of the RouteFilter resource.
	RouteFilter *SubResource `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The peering stats of express route circuit.
	Stats *ExpressRouteCircuitStats `pulumi:"stats"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

// The set of arguments for constructing a ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitPeeringArgs struct {
	// The Azure ASN.
	AzureASN pulumi.IntPtrInput
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections ExpressRouteCircuitConnectionTypeArrayInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	// Who was the last to modify the peering.
	LastModifiedBy pulumi.StringPtrInput
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The peer ASN.
	PeerASN pulumi.IntPtrInput
	// The name of the peering.
	PeeringName pulumi.StringInput
	// The peering type.
	PeeringType pulumi.StringPtrInput
	// The primary port.
	PrimaryAzurePort pulumi.StringPtrInput
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference of the RouteFilter resource.
	RouteFilter SubResourcePtrInput
	// The secondary port.
	SecondaryAzurePort pulumi.StringPtrInput
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key.
	SharedKey pulumi.StringPtrInput
	// The peering state.
	State pulumi.StringPtrInput
	// The peering stats of express route circuit.
	Stats ExpressRouteCircuitStatsPtrInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCircuitPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringArgs)(nil)).Elem()
}
