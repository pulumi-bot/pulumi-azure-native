// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Rule Group resource.
//
// ## Example Usage
// ### Create FirewallPolicyRuleGroup
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20191201"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewFirewallPolicyRuleGroup(ctx, "firewallPolicyRuleGroup", &network.FirewallPolicyRuleGroupArgs{
// 			FirewallPolicyName: pulumi.String("firewallPolicy"),
// 			Priority:           pulumi.Int(110),
// 			ResourceGroupName:  pulumi.String("rg1"),
// 			RuleGroupName:      pulumi.String("ruleGroup1"),
// 			Rules: network.FirewallPolicyRuleArray{
// 				&network.FirewallPolicyRuleArgs{
// 					Action: pulumi.StringMap{
// 						"type": pulumi.String("Deny"),
// 					},
// 					Name: pulumi.String("Example-Filter-Rule"),
// 					RuleConditions: pulumi.MapArray{
// 						pulumi.Map{
// 							"destinationAddresses": pulumi.StringArray{
// 								pulumi.String("*"),
// 							},
// 							"destinationPorts": pulumi.StringArray{
// 								pulumi.String("*"),
// 							},
// 							"ipProtocols": pulumi.StringArray{
// 								pulumi.String("TCP"),
// 							},
// 							"name":              pulumi.String("network-condition1"),
// 							"ruleConditionType": pulumi.String("NetworkRuleCondition"),
// 							"sourceAddresses": pulumi.StringArray{
// 								pulumi.String("10.1.25.0/24"),
// 							},
// 						},
// 					},
// 					RuleType: pulumi.String("FirewallPolicyFilterRule"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create FirewallPolicyRuleGroup With IpGroups
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20191201"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewFirewallPolicyRuleGroup(ctx, "firewallPolicyRuleGroup", &network.FirewallPolicyRuleGroupArgs{
// 			FirewallPolicyName: pulumi.String("firewallPolicy"),
// 			Priority:           pulumi.Int(110),
// 			ResourceGroupName:  pulumi.String("rg1"),
// 			RuleGroupName:      pulumi.String("ruleGroup1"),
// 			Rules: network.FirewallPolicyRuleArray{
// 				&network.FirewallPolicyRuleArgs{
// 					Action: pulumi.StringMap{
// 						"type": pulumi.String("Deny"),
// 					},
// 					Name: pulumi.String("Example-Filter-Rule"),
// 					RuleConditions: pulumi.MapArray{
// 						pulumi.Map{
// 							"destinationIpGroups": pulumi.StringArray{
// 								pulumi.String("/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups2"),
// 							},
// 							"destinationPorts": pulumi.StringArray{
// 								pulumi.String("*"),
// 							},
// 							"ipProtocols": pulumi.StringArray{
// 								pulumi.String("TCP"),
// 							},
// 							"name":              pulumi.String("network-condition1"),
// 							"ruleConditionType": pulumi.String("NetworkRuleCondition"),
// 							"sourceIpGroups": pulumi.StringArray{
// 								pulumi.String("/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups1"),
// 							},
// 						},
// 					},
// 					RuleType: pulumi.String("FirewallPolicyFilterRule"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type FirewallPolicyRuleGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Priority of the Firewall Policy Rule Group resource.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The provisioning state of the firewall policy rule group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Group of Firewall Policy rules.
	Rules pulumi.ArrayOutput `pulumi:"rules"`
	// Rule Group type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallPolicyRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyRuleGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleGroup, error) {
	if args == nil || args.FirewallPolicyName == nil {
		return nil, errors.New("missing required argument 'FirewallPolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RuleGroupName == nil {
		return nil, errors.New("missing required argument 'RuleGroupName'")
	}
	if args == nil {
		args = &FirewallPolicyRuleGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:FirewallPolicyRuleGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallPolicyRuleGroup
	err := ctx.RegisterResource("azurerm:network/v20191201:FirewallPolicyRuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyRuleGroup gets an existing FirewallPolicyRuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleGroup, error) {
	var resource FirewallPolicyRuleGroup
	err := ctx.ReadResource("azurerm:network/v20191201:FirewallPolicyRuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyRuleGroup resources.
type firewallPolicyRuleGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Group resource.
	Priority *int `pulumi:"priority"`
	// The provisioning state of the firewall policy rule group resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Group of Firewall Policy rules.
	Rules []interface{} `pulumi:"rules"`
	// Rule Group type.
	Type *string `pulumi:"type"`
}

type FirewallPolicyRuleGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Priority of the Firewall Policy Rule Group resource.
	Priority pulumi.IntPtrInput
	// The provisioning state of the firewall policy rule group resource.
	ProvisioningState pulumi.StringPtrInput
	// Group of Firewall Policy rules.
	Rules pulumi.ArrayInput
	// Rule Group type.
	Type pulumi.StringPtrInput
}

func (FirewallPolicyRuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleGroupState)(nil)).Elem()
}

type firewallPolicyRuleGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Group resource.
	Priority *int `pulumi:"priority"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the FirewallPolicyRuleGroup.
	RuleGroupName string `pulumi:"ruleGroupName"`
	// Group of Firewall Policy rules.
	Rules []interface{} `pulumi:"rules"`
}

// The set of arguments for constructing a FirewallPolicyRuleGroup resource.
type FirewallPolicyRuleGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Priority of the Firewall Policy Rule Group resource.
	Priority pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the FirewallPolicyRuleGroup.
	RuleGroupName pulumi.StringInput
	// Group of Firewall Policy rules.
	Rules pulumi.ArrayInput
}

func (FirewallPolicyRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleGroupArgs)(nil)).Elem()
}
