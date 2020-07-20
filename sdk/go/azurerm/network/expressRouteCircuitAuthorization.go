// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Authorization in an ExpressRouteCircuit resource.
type ExpressRouteCircuitAuthorization struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the express route circuit authorization.
	Properties AuthorizationPropertiesFormatResponseOutput `pulumi:"properties"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRouteCircuitAuthorization registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
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
		args = &ExpressRouteCircuitAuthorizationArgs{}
	}
	var resource ExpressRouteCircuitAuthorization
	err := ctx.RegisterResource("azurerm:network:ExpressRouteCircuitAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitAuthorization gets an existing ExpressRouteCircuitAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	var resource ExpressRouteCircuitAuthorization
	err := ctx.ReadResource("azurerm:network:ExpressRouteCircuitAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitAuthorization resources.
type expressRouteCircuitAuthorizationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the express route circuit authorization.
	Properties *AuthorizationPropertiesFormatResponse `pulumi:"properties"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type ExpressRouteCircuitAuthorizationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the express route circuit authorization.
	Properties AuthorizationPropertiesFormatResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (ExpressRouteCircuitAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationState)(nil)).Elem()
}

type expressRouteCircuitAuthorizationArgs struct {
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the authorization.
	Name string `pulumi:"name"`
	// Properties of the express route circuit authorization.
	Properties *AuthorizationPropertiesFormat `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitAuthorization resource.
type ExpressRouteCircuitAuthorizationArgs struct {
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the authorization.
	Name pulumi.StringInput
	// Properties of the express route circuit authorization.
	Properties AuthorizationPropertiesFormatPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationArgs)(nil)).Elem()
}
