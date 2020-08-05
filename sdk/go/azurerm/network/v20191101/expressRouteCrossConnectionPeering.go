// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peering in an ExpressRoute Cross Connection resource.
type ExpressRouteCrossConnectionPeering struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the express route cross connection peering.
	Properties ExpressRouteCrossConnectionPeeringPropertiesResponseOutput `pulumi:"properties"`
}

// NewExpressRouteCrossConnectionPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCrossConnectionPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCrossConnectionPeering, error) {
	if args == nil || args.CrossConnectionName == nil {
		return nil, errors.New("missing required argument 'CrossConnectionName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCrossConnectionPeeringArgs{}
	}
	var resource ExpressRouteCrossConnectionPeering
	err := ctx.RegisterResource("azurerm:network/v20191101:ExpressRouteCrossConnectionPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCrossConnectionPeering gets an existing ExpressRouteCrossConnectionPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCrossConnectionPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCrossConnectionPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCrossConnectionPeering, error) {
	var resource ExpressRouteCrossConnectionPeering
	err := ctx.ReadResource("azurerm:network/v20191101:ExpressRouteCrossConnectionPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCrossConnectionPeering resources.
type expressRouteCrossConnectionPeeringState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the express route cross connection peering.
	Properties *ExpressRouteCrossConnectionPeeringPropertiesResponse `pulumi:"properties"`
}

type ExpressRouteCrossConnectionPeeringState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the express route cross connection peering.
	Properties ExpressRouteCrossConnectionPeeringPropertiesResponsePtrInput
}

func (ExpressRouteCrossConnectionPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCrossConnectionPeeringState)(nil)).Elem()
}

type expressRouteCrossConnectionPeeringArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName string `pulumi:"crossConnectionName"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The name of the peering.
	Name string `pulumi:"name"`
	// The peer ASN.
	PeerASN *int `pulumi:"peerASN"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

// The set of arguments for constructing a ExpressRouteCrossConnectionPeering resource.
type ExpressRouteCrossConnectionPeeringArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName pulumi.StringInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigPtrInput
	// The name of the peering.
	Name pulumi.StringInput
	// The peer ASN.
	PeerASN pulumi.IntPtrInput
	// The peering type.
	PeeringType pulumi.StringPtrInput
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key.
	SharedKey pulumi.StringPtrInput
	// The peering state.
	State pulumi.StringPtrInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCrossConnectionPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCrossConnectionPeeringArgs)(nil)).Elem()
}
