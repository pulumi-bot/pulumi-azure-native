// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Route Filter Rule Resource
type RouteFilterRule struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Route Filter Rule Resource
	Properties RouteFilterRulePropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRouteFilterRule registers a new resource with the given unique name, arguments, and options.
func NewRouteFilterRule(ctx *pulumi.Context,
	name string, args *RouteFilterRuleArgs, opts ...pulumi.ResourceOption) (*RouteFilterRule, error) {
	if args == nil || args.Access == nil {
		return nil, errors.New("missing required argument 'Access'")
	}
	if args == nil || args.Communities == nil {
		return nil, errors.New("missing required argument 'Communities'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RouteFilterName == nil {
		return nil, errors.New("missing required argument 'RouteFilterName'")
	}
	if args == nil || args.RouteFilterRuleType == nil {
		return nil, errors.New("missing required argument 'RouteFilterRuleType'")
	}
	if args == nil {
		args = &RouteFilterRuleArgs{}
	}
	var resource RouteFilterRule
	err := ctx.RegisterResource("azurerm:network/v20170301:RouteFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteFilterRule gets an existing RouteFilterRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteFilterRuleState, opts ...pulumi.ResourceOption) (*RouteFilterRule, error) {
	var resource RouteFilterRule
	err := ctx.ReadResource("azurerm:network/v20170301:RouteFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteFilterRule resources.
type routeFilterRuleState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Route Filter Rule Resource
	Properties *RouteFilterRulePropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type RouteFilterRuleState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Route Filter Rule Resource
	Properties RouteFilterRulePropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RouteFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterRuleState)(nil)).Elem()
}

type routeFilterRuleArgs struct {
	// The access type of the rule. Valid values are: 'Allow', 'Deny'
	Access string `pulumi:"access"`
	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
	Communities []string `pulumi:"communities"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the route filter rule.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName string `pulumi:"routeFilterName"`
	// The rule type of the rule. Valid value is: 'Community'
	RouteFilterRuleType string `pulumi:"routeFilterRuleType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RouteFilterRule resource.
type RouteFilterRuleArgs struct {
	// The access type of the rule. Valid values are: 'Allow', 'Deny'
	Access pulumi.StringInput
	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
	Communities pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the route filter rule.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the route filter.
	RouteFilterName pulumi.StringInput
	// The rule type of the rule. Valid value is: 'Community'
	RouteFilterRuleType pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RouteFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterRuleArgs)(nil)).Elem()
}
