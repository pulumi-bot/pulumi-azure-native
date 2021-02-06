// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"context"
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
	// A collection of references to express route circuit peerings.
	Peerings ExpressRouteCircuitPeeringResponseArrayOutput `pulumi:"peerings"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Collection of RouteFilterRules contained within a route filter.
	Rules RouteFilterRuleResponseArrayOutput `pulumi:"rules"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRouteFilter registers a new resource with the given unique name, arguments, and options.
func NewRouteFilter(ctx *pulumi.Context,
	name string, args *RouteFilterArgs, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RouteFilterName == nil {
		return nil, errors.New("invalid value for required argument 'RouteFilterName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:RouteFilter"),
		},
	})
	opts = append(opts, aliases)
	var resource RouteFilter
	err := ctx.RegisterResource("azure-nextgen:network/v20180601:RouteFilter", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:network/v20180601:RouteFilter", name, id, state, &resource, opts...)
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
	// A collection of references to express route circuit peerings.
	Peerings []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Collection of RouteFilterRules contained within a route filter.
	Rules []RouteFilterRuleResponse `pulumi:"rules"`
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
	// A collection of references to express route circuit peerings.
	Peerings ExpressRouteCircuitPeeringResponseArrayInput
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// Collection of RouteFilterRules contained within a route filter.
	Rules RouteFilterRuleResponseArrayInput
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
	Location *string `pulumi:"location"`
	// A collection of references to express route circuit peerings.
	Peerings []ExpressRouteCircuitPeeringType `pulumi:"peerings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName string `pulumi:"routeFilterName"`
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
	Location pulumi.StringPtrInput
	// A collection of references to express route circuit peerings.
	Peerings ExpressRouteCircuitPeeringTypeArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the route filter.
	RouteFilterName pulumi.StringInput
	// Collection of RouteFilterRules contained within a route filter.
	Rules RouteFilterRuleTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RouteFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterArgs)(nil)).Elem()
}

type RouteFilterInput interface {
	pulumi.Input

	ToRouteFilterOutput() RouteFilterOutput
	ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput
}

func (*RouteFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilter)(nil))
}

func (i *RouteFilter) ToRouteFilterOutput() RouteFilterOutput {
	return i.ToRouteFilterOutputWithContext(context.Background())
}

func (i *RouteFilter) ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterOutput)
}

type RouteFilterOutput struct {
	*pulumi.OutputState
}

func (RouteFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilter)(nil))
}

func (o RouteFilterOutput) ToRouteFilterOutput() RouteFilterOutput {
	return o
}

func (o RouteFilterOutput) ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteFilterOutput{})
}
