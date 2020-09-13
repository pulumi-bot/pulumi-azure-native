// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Defines web application firewall policy.
//
// ## Example Usage
// ### Creates specific policy
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20180801"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewPolicy(ctx, "policy", &network.PolicyArgs{
// 			CustomRules: &network.CustomRulesArgs{
// 				Rules: network.CustomRuleArray{
// 					&network.CustomRuleArgs{
// 						Action: pulumi.String("Block"),
// 						MatchConditions: network.MatchConditionArray{
// 							&network.MatchConditionArgs{
// 								MatchValue: pulumi.StringArray{
// 									pulumi.String("192.168.1.0/24"),
// 									pulumi.String("10.0.0.0/24"),
// 								},
// 								MatchVariable: pulumi.String("RemoteAddr"),
// 								Operator:      pulumi.String("IPMatch"),
// 							},
// 						},
// 						Name:               pulumi.String("Rule1"),
// 						Priority:           pulumi.Int(1),
// 						RateLimitThreshold: pulumi.Int(1000),
// 						RuleType:           pulumi.String("RateLimitRule"),
// 					},
// 					&network.CustomRuleArgs{
// 						Action: pulumi.String("Block"),
// 						MatchConditions: network.MatchConditionArray{
// 							&network.MatchConditionArgs{
// 								MatchValue: pulumi.StringArray{
// 									pulumi.String("CH"),
// 								},
// 								MatchVariable: pulumi.String("RemoteAddr"),
// 								Operator:      pulumi.String("GeoMatch"),
// 							},
// 							&network.MatchConditionArgs{
// 								MatchValue: pulumi.StringArray{
// 									pulumi.String("Windows"),
// 								},
// 								MatchVariable: pulumi.String("RequestHeader"),
// 								Operator:      pulumi.String("Contains"),
// 								Selector:      pulumi.String("UserAgent"),
// 							},
// 						},
// 						Name:     pulumi.String("Rule2"),
// 						Priority: pulumi.Int(2),
// 						RuleType: pulumi.String("MatchRule"),
// 					},
// 				},
// 			},
// 			ManagedRules: &network.ManagedRuleSetsArgs{
// 				RuleSets: network.ManagedRuleSetArray{
// 					&network.ManagedRuleSetArgs{
// 						Priority: pulumi.Int(1),
// 						RuleGroupOverrides: pulumi.StringMapArray{
// 							pulumi.StringMap{
// 								"action":            pulumi.String("Block"),
// 								"ruleGroupOverride": pulumi.String("SqlInjection"),
// 							},
// 							pulumi.StringMap{
// 								"action":            pulumi.String("Log"),
// 								"ruleGroupOverride": pulumi.String("XSS"),
// 							},
// 						},
// 						RuleSetType: pulumi.String("AzureManagedRuleSet"),
// 					},
// 				},
// 			},
// 			PolicyName:        pulumi.String("Policy1"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Describes custom rules inside the policy
	CustomRules CustomRulesResponsePtrOutput `pulumi:"customRules"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Describes managed rules inside the policy
	ManagedRules ManagedRuleSetsResponsePtrOutput `pulumi:"managedRules"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes  policySettings for policy
	PolicySettings PolicySettingsResponsePtrOutput `pulumi:"policySettings"`
	// Provisioning state of the WebApplicationFirewallPolicy.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	ResourceState     pulumi.StringOutput `pulumi:"resourceState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.PolicyName == nil {
		return nil, errors.New("missing required argument 'PolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:Policy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190301:Policy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191001:Policy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azurerm:network/v20180801:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azurerm:network/v20180801:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Describes custom rules inside the policy
	CustomRules *CustomRulesResponse `pulumi:"customRules"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Describes managed rules inside the policy
	ManagedRules *ManagedRuleSetsResponse `pulumi:"managedRules"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Describes  policySettings for policy
	PolicySettings *PolicySettingsResponse `pulumi:"policySettings"`
	// Provisioning state of the WebApplicationFirewallPolicy.
	ProvisioningState *string `pulumi:"provisioningState"`
	ResourceState     *string `pulumi:"resourceState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type PolicyState struct {
	// Describes custom rules inside the policy
	CustomRules CustomRulesResponsePtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Describes managed rules inside the policy
	ManagedRules ManagedRuleSetsResponsePtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Describes  policySettings for policy
	PolicySettings PolicySettingsResponsePtrInput
	// Provisioning state of the WebApplicationFirewallPolicy.
	ProvisioningState pulumi.StringPtrInput
	ResourceState     pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Describes custom rules inside the policy
	CustomRules *CustomRules `pulumi:"customRules"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Describes managed rules inside the policy
	ManagedRules *ManagedRuleSets `pulumi:"managedRules"`
	// The name of the resource group.
	PolicyName string `pulumi:"policyName"`
	// Describes  policySettings for policy
	PolicySettings *PolicySettings `pulumi:"policySettings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Describes custom rules inside the policy
	CustomRules CustomRulesPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Describes managed rules inside the policy
	ManagedRules ManagedRuleSetsPtrInput
	// The name of the resource group.
	PolicyName pulumi.StringInput
	// Describes  policySettings for policy
	PolicySettings PolicySettingsPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
