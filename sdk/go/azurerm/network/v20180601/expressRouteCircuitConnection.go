// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitConnection struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       pulumi.StringPtrOutput                                      `pulumi:"name"`
	Properties ExpressRouteCircuitConnectionPropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewExpressRouteCircuitConnection registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitConnection, error) {
	if args == nil || args.CircuitName == nil {
		return nil, errors.New("missing required argument 'CircuitName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PeeringName == nil {
		return nil, errors.New("missing required argument 'PeeringName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCircuitConnectionArgs{}
	}
	var resource ExpressRouteCircuitConnection
	err := ctx.RegisterResource("azurerm:network/v20180601:ExpressRouteCircuitConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitConnection gets an existing ExpressRouteCircuitConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitConnectionState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitConnection, error) {
	var resource ExpressRouteCircuitConnection
	err := ctx.ReadResource("azurerm:network/v20180601:ExpressRouteCircuitConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitConnection resources.
type expressRouteCircuitConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       *string                                                `pulumi:"name"`
	Properties *ExpressRouteCircuitConnectionPropertiesFormatResponse `pulumi:"properties"`
}

type ExpressRouteCircuitConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       pulumi.StringPtrInput
	Properties ExpressRouteCircuitConnectionPropertiesFormatResponsePtrInput
}

func (ExpressRouteCircuitConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitConnectionState)(nil)).Elem()
}

type expressRouteCircuitConnectionArgs struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering *SubResource `pulumi:"expressRouteCircuitPeering"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the express route circuit connection.
	Name string `pulumi:"name"`
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering *SubResource `pulumi:"peerExpressRouteCircuitPeering"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitConnection resource.
type ExpressRouteCircuitConnectionArgs struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix pulumi.StringPtrInput
	// The authorization key.
	AuthorizationKey pulumi.StringPtrInput
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the express route circuit connection.
	Name pulumi.StringInput
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering SubResourcePtrInput
	// The name of the peering.
	PeeringName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitConnectionArgs)(nil)).Elem()
}
