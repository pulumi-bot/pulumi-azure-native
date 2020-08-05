// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Route Filter Resource.
type RouteFilter struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Route Filter Resource
	Properties RouteFilterPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRouteFilter registers a new resource with the given unique name, arguments, and options.
func NewRouteFilter(ctx *pulumi.Context,
	name string, args *RouteFilterArgs, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &RouteFilterArgs{}
	}
	var resource RouteFilter
	err := ctx.RegisterResource("azurerm:network/v20180201:RouteFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteFilter gets an existing RouteFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteFilterState, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	var resource RouteFilter
	err := ctx.ReadResource("azurerm:network/v20180201:RouteFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteFilter resources.
type routeFilterState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Route Filter Resource
	Properties *RouteFilterPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type RouteFilterState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Route Filter Resource
	Properties RouteFilterPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (RouteFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterState)(nil)).Elem()
}

type routeFilterArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the route filter.
	Name string `pulumi:"name"`
	// A collection of references to express route circuit peerings.
	Peerings []ExpressRouteCircuitPeeringType `pulumi:"peerings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Collection of RouteFilterRules contained within a route filter.
	Rules []RouteFilterRuleType `pulumi:"rules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RouteFilter resource.
type RouteFilterArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the route filter.
	Name pulumi.StringInput
	// A collection of references to express route circuit peerings.
	Peerings ExpressRouteCircuitPeeringTypeArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Collection of RouteFilterRules contained within a route filter.
	Rules RouteFilterRuleTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RouteFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterArgs)(nil)).Elem()
}
