// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Rule Collection Group resource.
type FirewallPolicyRuleCollectionGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The provisioning state of the firewall policy rule collection group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Group of Firewall Policy rule collections.
	RuleCollections pulumi.ArrayOutput `pulumi:"ruleCollections"`
	// Rule Group type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallPolicyRuleCollectionGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	if args == nil || args.FirewallPolicyName == nil {
		return nil, errors.New("missing required argument 'FirewallPolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RuleCollectionGroupName == nil {
		return nil, errors.New("missing required argument 'RuleCollectionGroupName'")
	}
	if args == nil {
		args = &FirewallPolicyRuleCollectionGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:FirewallPolicyRuleCollectionGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.RegisterResource("azure-nextgen:network/latest:FirewallPolicyRuleCollectionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyRuleCollectionGroup gets an existing FirewallPolicyRuleCollectionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleCollectionGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.ReadResource("azure-nextgen:network/latest:FirewallPolicyRuleCollectionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyRuleCollectionGroup resources.
type firewallPolicyRuleCollectionGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority *int `pulumi:"priority"`
	// The provisioning state of the firewall policy rule collection group resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Group of Firewall Policy rule collections.
	RuleCollections []interface{} `pulumi:"ruleCollections"`
	// Rule Group type.
	Type *string `pulumi:"type"`
}

type FirewallPolicyRuleCollectionGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority pulumi.IntPtrInput
	// The provisioning state of the firewall policy rule collection group resource.
	ProvisioningState pulumi.StringPtrInput
	// Group of Firewall Policy rule collections.
	RuleCollections pulumi.ArrayInput
	// Rule Group type.
	Type pulumi.StringPtrInput
}

func (FirewallPolicyRuleCollectionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupState)(nil)).Elem()
}

type firewallPolicyRuleCollectionGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority *int `pulumi:"priority"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the FirewallPolicyRuleCollectionGroup.
	RuleCollectionGroupName string `pulumi:"ruleCollectionGroupName"`
	// Group of Firewall Policy rule collections.
	RuleCollections []interface{} `pulumi:"ruleCollections"`
}

// The set of arguments for constructing a FirewallPolicyRuleCollectionGroup resource.
type FirewallPolicyRuleCollectionGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the FirewallPolicyRuleCollectionGroup.
	RuleCollectionGroupName pulumi.StringInput
	// Group of Firewall Policy rule collections.
	RuleCollections pulumi.ArrayInput
}

func (FirewallPolicyRuleCollectionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupArgs)(nil)).Elem()
}

type FirewallPolicyRuleCollectionGroupInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput
	ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput
}

func (FirewallPolicyRuleCollectionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionGroup)(nil)).Elem()
}

func (i FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return i.ToFirewallPolicyRuleCollectionGroupOutputWithContext(context.Background())
}

func (i FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupOutput)
}

type FirewallPolicyRuleCollectionGroupOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyRuleCollectionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionGroupOutput)(nil)).Elem()
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupOutput{})
}
