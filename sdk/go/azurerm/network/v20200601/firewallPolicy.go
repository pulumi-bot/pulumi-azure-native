// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FirewallPolicy Resource.
//
// ## Example Usage
// ### Create FirewallPolicy
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewFirewallPolicy(ctx, "firewallPolicy", &network.FirewallPolicyArgs{
// 			DnsSettings: &network.DnsSettingsArgs{
// 				EnableProxy:                 pulumi.Bool(true),
// 				RequireProxyForNetworkRules: pulumi.Bool(false),
// 				Servers: pulumi.StringArray{
// 					pulumi.String("30.3.4.5"),
// 				},
// 			},
// 			FirewallPolicyName: pulumi.String("firewallPolicy"),
// 			Location:           pulumi.String("West US"),
// 			ResourceGroupName:  pulumi.String("rg1"),
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 			},
// 			ThreatIntelMode: pulumi.String("Alert"),
// 			ThreatIntelWhitelist: &network.FirewallPolicyThreatIntelWhitelistArgs{
// 				Fqdns: pulumi.StringArray{
// 					pulumi.String("*.microsoft.com"),
// 				},
// 				IpAddresses: pulumi.StringArray{
// 					pulumi.String("20.3.4.5"),
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
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// The parent firewall policy from which rules are inherited.
	BasePolicy SubResourceResponsePtrOutput `pulumi:"basePolicy"`
	// List of references to Child Firewall Policies.
	ChildPolicies SubResourceResponseArrayOutput `pulumi:"childPolicies"`
	// DNS Proxy Settings definition.
	DnsSettings DnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls SubResourceResponseArrayOutput `pulumi:"firewalls"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the firewall policy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of references to FirewallPolicyRuleCollectionGroups.
	RuleCollectionGroups SubResourceResponseArrayOutput `pulumi:"ruleCollectionGroups"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrOutput `pulumi:"threatIntelMode"`
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistResponsePtrOutput `pulumi:"threatIntelWhitelist"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil || args.FirewallPolicyName == nil {
		return nil, errors.New("missing required argument 'FirewallPolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FirewallPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:FirewallPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallPolicy
	err := ctx.RegisterResource("azurerm:network/v20200601:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("azurerm:network/v20200601:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// The parent firewall policy from which rules are inherited.
	BasePolicy *SubResourceResponse `pulumi:"basePolicy"`
	// List of references to Child Firewall Policies.
	ChildPolicies []SubResourceResponse `pulumi:"childPolicies"`
	// DNS Proxy Settings definition.
	DnsSettings *DnsSettingsResponse `pulumi:"dnsSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// List of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls []SubResourceResponse `pulumi:"firewalls"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the firewall policy resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// List of references to FirewallPolicyRuleCollectionGroups.
	RuleCollectionGroups []SubResourceResponse `pulumi:"ruleCollectionGroups"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist *FirewallPolicyThreatIntelWhitelistResponse `pulumi:"threatIntelWhitelist"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type FirewallPolicyState struct {
	// The parent firewall policy from which rules are inherited.
	BasePolicy SubResourceResponsePtrInput
	// List of references to Child Firewall Policies.
	ChildPolicies SubResourceResponseArrayInput
	// DNS Proxy Settings definition.
	DnsSettings DnsSettingsResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// List of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls SubResourceResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the firewall policy resource.
	ProvisioningState pulumi.StringPtrInput
	// List of references to FirewallPolicyRuleCollectionGroups.
	RuleCollectionGroups SubResourceResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// The parent firewall policy from which rules are inherited.
	BasePolicy *SubResource `pulumi:"basePolicy"`
	// DNS Proxy Settings definition.
	DnsSettings *DnsSettings `pulumi:"dnsSettings"`
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist *FirewallPolicyThreatIntelWhitelist `pulumi:"threatIntelWhitelist"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// The parent firewall policy from which rules are inherited.
	BasePolicy SubResourcePtrInput
	// DNS Proxy Settings definition.
	DnsSettings DnsSettingsPtrInput
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistPtrInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}
