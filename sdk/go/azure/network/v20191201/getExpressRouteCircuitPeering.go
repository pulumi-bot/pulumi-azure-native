// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupExpressRouteCircuitPeering(ctx *pulumi.Context, args *LookupExpressRouteCircuitPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitPeeringResult, error) {
	var rv LookupExpressRouteCircuitPeeringResult
	err := ctx.Invoke("azure-nextgen:network/v20191201:getExpressRouteCircuitPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitPeeringArgs struct {
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitPeeringResult struct {
	// The Azure ASN.
	AzureASN *int `pulumi:"azureASN"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections []ExpressRouteCircuitConnectionResponse `pulumi:"connections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The ExpressRoute connection.
	ExpressRouteConnection *ExpressRouteConnectionIdResponse `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *float64 `pulumi:"peerASN"`
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections []PeerExpressRouteCircuitConnectionResponse `pulumi:"peeredConnections"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The reference to the RouteFilter resource.
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
	Type string `pulumi:"type"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}
