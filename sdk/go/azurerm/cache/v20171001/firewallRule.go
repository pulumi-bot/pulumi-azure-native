// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A firewall rule on a redis cache has a name, and describes a contiguous range of IP addresses permitted to connect
//
// ## Example Usage
// ### RedisCacheFirewallRuleCreate
//
// ```go
// package main
//
// import (
// 	cache "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/cache/v20171001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cache.NewFirewallRule(ctx, "firewallRule", &cache.FirewallRuleArgs{
// 			CacheName:         pulumi.String("cache1"),
// 			EndIP:             pulumi.String("192.168.1.4"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			RuleName:          pulumi.String("rule1"),
// 			StartIP:           pulumi.String("192.168.1.1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// highest IP address included in the range
	EndIP pulumi.StringOutput `pulumi:"endIP"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// lowest IP address included in the range
	StartIP pulumi.StringOutput `pulumi:"startIP"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil || args.CacheName == nil {
		return nil, errors.New("missing required argument 'CacheName'")
	}
	if args == nil || args.EndIP == nil {
		return nil, errors.New("missing required argument 'EndIP'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RuleName == nil {
		return nil, errors.New("missing required argument 'RuleName'")
	}
	if args == nil || args.StartIP == nil {
		return nil, errors.New("missing required argument 'StartIP'")
	}
	if args == nil {
		args = &FirewallRuleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:cache/latest:FirewallRule"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20160401:FirewallRule"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20170201:FirewallRule"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20180301:FirewallRule"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20190701:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azurerm:cache/v20171001:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azurerm:cache/v20171001:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// highest IP address included in the range
	EndIP *string `pulumi:"endIP"`
	// Resource name.
	Name *string `pulumi:"name"`
	// lowest IP address included in the range
	StartIP *string `pulumi:"startIP"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type FirewallRuleState struct {
	// highest IP address included in the range
	EndIP pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// lowest IP address included in the range
	StartIP pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The name of the Redis cache.
	CacheName string `pulumi:"cacheName"`
	// highest IP address included in the range
	EndIP string `pulumi:"endIP"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the firewall rule.
	RuleName string `pulumi:"ruleName"`
	// lowest IP address included in the range
	StartIP string `pulumi:"startIP"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The name of the Redis cache.
	CacheName pulumi.StringInput
	// highest IP address included in the range
	EndIP pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the firewall rule.
	RuleName pulumi.StringInput
	// lowest IP address included in the range
	StartIP pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}
