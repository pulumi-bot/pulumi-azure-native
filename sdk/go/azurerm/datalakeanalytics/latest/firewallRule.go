// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Data Lake Analytics firewall rule information.
//
// ## Example Usage
// ### Creates or updates the specified firewall rule
//
// ```go
// package main
//
// import (
// 	datalakeanalytics "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/datalakeanalytics/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datalakeanalytics.NewFirewallRule(ctx, "firewallRule", &datalakeanalytics.FirewallRuleArgs{
// 			AccountName:       pulumi.String("contosoadla"),
// 			EndIpAddress:      pulumi.String("2.2.2.2"),
// 			FirewallRuleName:  pulumi.String("test_rule"),
// 			ResourceGroupName: pulumi.String("contosorg"),
// 			StartIpAddress:    pulumi.String("1.1.1.1"),
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

	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.EndIpAddress == nil {
		return nil, errors.New("missing required argument 'EndIpAddress'")
	}
	if args == nil || args.FirewallRuleName == nil {
		return nil, errors.New("missing required argument 'FirewallRuleName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StartIpAddress == nil {
		return nil, errors.New("missing required argument 'StartIpAddress'")
	}
	if args == nil {
		args = &FirewallRuleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datalakeanalytics/v20161101:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azurerm:datalakeanalytics/latest:FirewallRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:datalakeanalytics/latest:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress *string `pulumi:"startIpAddress"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type FirewallRuleState struct {
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress pulumi.StringPtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The name of the firewall rule to create or update.
	FirewallRuleName string `pulumi:"firewallRuleName"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName pulumi.StringInput
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress pulumi.StringInput
	// The name of the firewall rule to create or update.
	FirewallRuleName pulumi.StringInput
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}
