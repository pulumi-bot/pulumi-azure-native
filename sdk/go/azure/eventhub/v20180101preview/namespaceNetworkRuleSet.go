// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of topic resource.
type NamespaceNetworkRuleSet struct {
	pulumi.CustomResourceState

	// Default Action for Network Rule Set
	DefaultAction pulumi.StringPtrOutput `pulumi:"defaultAction"`
	// List of IpRules
	IpRules NWRuleSetIpRulesResponseArrayOutput `pulumi:"ipRules"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNamespaceNetworkRuleSet registers a new resource with the given unique name, arguments, and options.
func NewNamespaceNetworkRuleSet(ctx *pulumi.Context,
	name string, args *NamespaceNetworkRuleSetArgs, opts ...pulumi.ResourceOption) (*NamespaceNetworkRuleSet, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NamespaceNetworkRuleSetArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventhub/latest:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20170401:NamespaceNetworkRuleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceNetworkRuleSet
	err := ctx.RegisterResource("azure-nextgen:eventhub/v20180101preview:NamespaceNetworkRuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceNetworkRuleSet gets an existing NamespaceNetworkRuleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceNetworkRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceNetworkRuleSetState, opts ...pulumi.ResourceOption) (*NamespaceNetworkRuleSet, error) {
	var resource NamespaceNetworkRuleSet
	err := ctx.ReadResource("azure-nextgen:eventhub/v20180101preview:NamespaceNetworkRuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceNetworkRuleSet resources.
type namespaceNetworkRuleSetState struct {
	// Default Action for Network Rule Set
	DefaultAction *string `pulumi:"defaultAction"`
	// List of IpRules
	IpRules []NWRuleSetIpRulesResponse `pulumi:"ipRules"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type NamespaceNetworkRuleSetState struct {
	// Default Action for Network Rule Set
	DefaultAction pulumi.StringPtrInput
	// List of IpRules
	IpRules NWRuleSetIpRulesResponseArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (NamespaceNetworkRuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceNetworkRuleSetState)(nil)).Elem()
}

type namespaceNetworkRuleSetArgs struct {
	// Default Action for Network Rule Set
	DefaultAction *string `pulumi:"defaultAction"`
	// List of IpRules
	IpRules []NWRuleSetIpRules `pulumi:"ipRules"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NamespaceNetworkRuleSet resource.
type NamespaceNetworkRuleSetArgs struct {
	// Default Action for Network Rule Set
	DefaultAction pulumi.StringPtrInput
	// List of IpRules
	IpRules NWRuleSetIpRulesArrayInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (NamespaceNetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceNetworkRuleSetArgs)(nil)).Elem()
}
