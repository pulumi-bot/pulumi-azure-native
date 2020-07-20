// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitPeeringConnection struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the express route circuit connection.
	Properties ExpressRouteCircuitConnectionPropertiesFormatResponseOutput `pulumi:"properties"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRouteCircuitPeeringConnection registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitPeeringConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeeringConnection, error) {
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
		args = &ExpressRouteCircuitPeeringConnectionArgs{}
	}
	var resource ExpressRouteCircuitPeeringConnection
	err := ctx.RegisterResource("azurerm:network:ExpressRouteCircuitPeeringConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitPeeringConnection gets an existing ExpressRouteCircuitPeeringConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitPeeringConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitPeeringConnectionState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeeringConnection, error) {
	var resource ExpressRouteCircuitPeeringConnection
	err := ctx.ReadResource("azurerm:network:ExpressRouteCircuitPeeringConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitPeeringConnection resources.
type expressRouteCircuitPeeringConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the express route circuit connection.
	Properties *ExpressRouteCircuitConnectionPropertiesFormatResponse `pulumi:"properties"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type ExpressRouteCircuitPeeringConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the express route circuit connection.
	Properties ExpressRouteCircuitConnectionPropertiesFormatResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (ExpressRouteCircuitPeeringConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringConnectionState)(nil)).Elem()
}

type expressRouteCircuitPeeringConnectionArgs struct {
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the express route circuit connection.
	Name string `pulumi:"name"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// Properties of the express route circuit connection.
	Properties *ExpressRouteCircuitConnectionPropertiesFormat `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitPeeringConnection resource.
type ExpressRouteCircuitPeeringConnectionArgs struct {
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the express route circuit connection.
	Name pulumi.StringInput
	// The name of the peering.
	PeeringName pulumi.StringInput
	// Properties of the express route circuit connection.
	Properties ExpressRouteCircuitConnectionPropertiesFormatPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitPeeringConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringConnectionArgs)(nil)).Elem()
}
