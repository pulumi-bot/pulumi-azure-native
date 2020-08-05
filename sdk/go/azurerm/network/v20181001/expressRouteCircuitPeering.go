// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peering in an ExpressRouteCircuit resource.
type ExpressRouteCircuitPeering struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       pulumi.StringPtrOutput                                   `pulumi:"name"`
	Properties ExpressRouteCircuitPeeringPropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewExpressRouteCircuitPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	if args == nil || args.CircuitName == nil {
		return nil, errors.New("missing required argument 'CircuitName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCircuitPeeringArgs{}
	}
	var resource ExpressRouteCircuitPeering
	err := ctx.RegisterResource("azurerm:network/v20181001:ExpressRouteCircuitPeering", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20181001:ExpressRouteCircuitPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitPeering resources.
type expressRouteCircuitPeeringState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       *string                                             `pulumi:"name"`
	Properties *ExpressRouteCircuitPeeringPropertiesFormatResponse `pulumi:"properties"`
}

type ExpressRouteCircuitPeeringState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       pulumi.StringPtrInput
	Properties ExpressRouteCircuitPeeringPropertiesFormatResponsePtrInput
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
	// The ExpressRoute connection.
	ExpressRouteConnection *ExpressRouteConnectionId `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	// Gets whether the provider or the customer last modified the peering.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The name of the peering.
	Name string `pulumi:"name"`
	// The peer ASN.
	PeerASN *int `pulumi:"peerASN"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference of the RouteFilter resource.
	RouteFilter *RouteFilterType `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// Gets peering stats.
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
	// The ExpressRoute connection.
	ExpressRouteConnection ExpressRouteConnectionIdPtrInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	// Gets whether the provider or the customer last modified the peering.
	LastModifiedBy pulumi.StringPtrInput
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigPtrInput
	// The name of the peering.
	Name pulumi.StringInput
	// The peer ASN.
	PeerASN pulumi.IntPtrInput
	// The peering type.
	PeeringType pulumi.StringPtrInput
	// The primary port.
	PrimaryAzurePort pulumi.StringPtrInput
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference of the RouteFilter resource.
	RouteFilter RouteFilterTypePtrInput
	// The secondary port.
	SecondaryAzurePort pulumi.StringPtrInput
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key.
	SharedKey pulumi.StringPtrInput
	// The peering state.
	State pulumi.StringPtrInput
	// Gets peering stats.
	Stats ExpressRouteCircuitStatsPtrInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCircuitPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringArgs)(nil)).Elem()
}
