// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a namespace authorization rule.
type TopicAuthorizationRule struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// AuthorizationRule properties.
	Properties SBAuthorizationRuleResponsePropertiesOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopicAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewTopicAuthorizationRule(ctx *pulumi.Context,
	name string, args *TopicAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Rights == nil {
		return nil, errors.New("missing required argument 'Rights'")
	}
	if args == nil || args.TopicName == nil {
		return nil, errors.New("missing required argument 'TopicName'")
	}
	if args == nil {
		args = &TopicAuthorizationRuleArgs{}
	}
	var resource TopicAuthorizationRule
	err := ctx.RegisterResource("azurerm:servicebus/v20170401:TopicAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicAuthorizationRule gets an existing TopicAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicAuthorizationRuleState, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	var resource TopicAuthorizationRule
	err := ctx.ReadResource("azurerm:servicebus/v20170401:TopicAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicAuthorizationRule resources.
type topicAuthorizationRuleState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// AuthorizationRule properties.
	Properties *SBAuthorizationRuleResponseProperties `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type TopicAuthorizationRuleState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// AuthorizationRule properties.
	Properties SBAuthorizationRuleResponsePropertiesPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (TopicAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleState)(nil)).Elem()
}

type topicAuthorizationRuleArgs struct {
	// The authorization rule name.
	Name string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a TopicAuthorizationRule resource.
type TopicAuthorizationRuleArgs struct {
	// The authorization rule name.
	Name pulumi.StringInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
	// The topic name.
	TopicName pulumi.StringInput
}

func (TopicAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleArgs)(nil)).Elem()
}
