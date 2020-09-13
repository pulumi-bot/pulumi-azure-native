// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single item in a List or Get IpFilterRules operation
//
// ## Example Usage
// ### NameSpaceIpFilterRuleCreate
//
// ```go
// package main
//
// import (
// 	servicebus "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/servicebus/v20180101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicebus.NewNamespaceIpFilterRule(ctx, "namespaceIpFilterRule", &servicebus.NamespaceIpFilterRuleArgs{
// 			Action:            pulumi.String("Accept"),
// 			FilterName:        pulumi.String("sdk-IPFilterRules-7337"),
// 			IpFilterRuleName:  pulumi.String("sdk-IPFilterRules-7337"),
// 			IpMask:            pulumi.String("13.78.143.246/32"),
// 			NamespaceName:     pulumi.String("sdk-Namespace-5232"),
// 			ResourceGroupName: pulumi.String("ResourceGroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type NamespaceIpFilterRule struct {
	pulumi.CustomResourceState

	// The IP Filter Action
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// IP Filter name
	FilterName pulumi.StringPtrOutput `pulumi:"filterName"`
	// IP Mask
	IpMask pulumi.StringPtrOutput `pulumi:"ipMask"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNamespaceIpFilterRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceIpFilterRule(ctx *pulumi.Context,
	name string, args *NamespaceIpFilterRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceIpFilterRule, error) {
	if args == nil || args.IpFilterRuleName == nil {
		return nil, errors.New("missing required argument 'IpFilterRuleName'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NamespaceIpFilterRuleArgs{}
	}
	var resource NamespaceIpFilterRule
	err := ctx.RegisterResource("azurerm:servicebus/v20180101preview:NamespaceIpFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceIpFilterRule gets an existing NamespaceIpFilterRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceIpFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIpFilterRuleState, opts ...pulumi.ResourceOption) (*NamespaceIpFilterRule, error) {
	var resource NamespaceIpFilterRule
	err := ctx.ReadResource("azurerm:servicebus/v20180101preview:NamespaceIpFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceIpFilterRule resources.
type namespaceIpFilterRuleState struct {
	// The IP Filter Action
	Action *string `pulumi:"action"`
	// IP Filter name
	FilterName *string `pulumi:"filterName"`
	// IP Mask
	IpMask *string `pulumi:"ipMask"`
	// Resource name
	Name *string `pulumi:"name"`
	// Resource type
	Type *string `pulumi:"type"`
}

type NamespaceIpFilterRuleState struct {
	// The IP Filter Action
	Action pulumi.StringPtrInput
	// IP Filter name
	FilterName pulumi.StringPtrInput
	// IP Mask
	IpMask pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (NamespaceIpFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIpFilterRuleState)(nil)).Elem()
}

type namespaceIpFilterRuleArgs struct {
	// The IP Filter Action
	Action *string `pulumi:"action"`
	// IP Filter name
	FilterName *string `pulumi:"filterName"`
	// The IP Filter Rule name.
	IpFilterRuleName string `pulumi:"ipFilterRuleName"`
	// IP Mask
	IpMask *string `pulumi:"ipMask"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NamespaceIpFilterRule resource.
type NamespaceIpFilterRuleArgs struct {
	// The IP Filter Action
	Action pulumi.StringPtrInput
	// IP Filter name
	FilterName pulumi.StringPtrInput
	// The IP Filter Rule name.
	IpFilterRuleName pulumi.StringInput
	// IP Mask
	IpMask pulumi.StringPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (NamespaceIpFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIpFilterRuleArgs)(nil)).Elem()
}
