// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of Rule Resource.
type Rule struct {
	pulumi.CustomResourceState

	// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
	Action ActionResponsePtrOutput `pulumi:"action"`
	// Properties of correlationFilter
	CorrelationFilter CorrelationFilterResponsePtrOutput `pulumi:"correlationFilter"`
	// Filter type that is evaluated against a BrokeredMessage.
	FilterType pulumi.StringPtrOutput `pulumi:"filterType"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of sqlFilter
	SqlFilter SqlFilterResponsePtrOutput `pulumi:"sqlFilter"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus/latest:Rule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:Rule"),
		},
	})
	opts = append(opts, aliases)
	var resource Rule
	err := ctx.RegisterResource("azure-nextgen:servicebus/v20170401:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("azure-nextgen:servicebus/v20170401:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
	Action *ActionResponse `pulumi:"action"`
	// Properties of correlationFilter
	CorrelationFilter *CorrelationFilterResponse `pulumi:"correlationFilter"`
	// Filter type that is evaluated against a BrokeredMessage.
	FilterType *string `pulumi:"filterType"`
	// Resource name
	Name *string `pulumi:"name"`
	// Properties of sqlFilter
	SqlFilter *SqlFilterResponse `pulumi:"sqlFilter"`
	// Resource type
	Type *string `pulumi:"type"`
}

type RuleState struct {
	// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
	Action ActionResponsePtrInput
	// Properties of correlationFilter
	CorrelationFilter CorrelationFilterResponsePtrInput
	// Filter type that is evaluated against a BrokeredMessage.
	FilterType pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Properties of sqlFilter
	SqlFilter SqlFilterResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
	Action *Action `pulumi:"action"`
	// Properties of correlationFilter
	CorrelationFilter *CorrelationFilter `pulumi:"correlationFilter"`
	// Filter type that is evaluated against a BrokeredMessage.
	FilterType *string `pulumi:"filterType"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rule name.
	RuleName string `pulumi:"ruleName"`
	// Properties of sqlFilter
	SqlFilter *SqlFilter `pulumi:"sqlFilter"`
	// The subscription name.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
	Action ActionPtrInput
	// Properties of correlationFilter
	CorrelationFilter CorrelationFilterPtrInput
	// Filter type that is evaluated against a BrokeredMessage.
	FilterType FilterType
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rule name.
	RuleName pulumi.StringInput
	// Properties of sqlFilter
	SqlFilter SqlFilterPtrInput
	// The subscription name.
	SubscriptionName pulumi.StringInput
	// The topic name.
	TopicName pulumi.StringInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (Rule) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil)).Elem()
}

func (i Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

type RuleOutput struct {
	*pulumi.OutputState
}

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleOutput)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RuleOutput{})
}
