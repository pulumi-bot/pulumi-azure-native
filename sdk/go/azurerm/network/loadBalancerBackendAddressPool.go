// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Pool of backend IP addresses.
type LoadBalancerBackendAddressPool struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of load balancer backend address pool.
	Properties BackendAddressPoolPropertiesFormatResponseOutput `pulumi:"properties"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLoadBalancerBackendAddressPool registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendAddressPoolArgs, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LoadBalancerBackendAddressPoolArgs{}
	}
	var resource LoadBalancerBackendAddressPool
	err := ctx.RegisterResource("azurerm:network:LoadBalancerBackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerBackendAddressPool gets an existing LoadBalancerBackendAddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerBackendAddressPoolState, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	var resource LoadBalancerBackendAddressPool
	err := ctx.ReadResource("azurerm:network:LoadBalancerBackendAddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerBackendAddressPool resources.
type loadBalancerBackendAddressPoolState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of load balancer backend address pool.
	Properties *BackendAddressPoolPropertiesFormatResponse `pulumi:"properties"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type LoadBalancerBackendAddressPoolState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of load balancer backend address pool.
	Properties BackendAddressPoolPropertiesFormatResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (LoadBalancerBackendAddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolState)(nil)).Elem()
}

type loadBalancerBackendAddressPoolArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the backend address pool.
	Name string `pulumi:"name"`
	// Properties of load balancer backend address pool.
	Properties *BackendAddressPoolPropertiesFormat `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LoadBalancerBackendAddressPool resource.
type LoadBalancerBackendAddressPoolArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput
	// The name of the backend address pool.
	Name pulumi.StringInput
	// Properties of load balancer backend address pool.
	Properties BackendAddressPoolPropertiesFormatPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (LoadBalancerBackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolArgs)(nil)).Elem()
}
