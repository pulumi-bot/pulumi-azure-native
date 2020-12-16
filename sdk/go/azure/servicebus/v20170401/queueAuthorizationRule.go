// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a namespace authorization rule.
type QueueAuthorizationRule struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The rights associated with the rule.
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQueueAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewQueueAuthorizationRule(ctx *pulumi.Context,
	name string, args *QueueAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationRuleName == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationRuleName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.QueueName == nil {
		return nil, errors.New("invalid value for required argument 'QueueName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus/latest:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20140901:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20150801:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:QueueAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource QueueAuthorizationRule
	err := ctx.RegisterResource("azure-nextgen:servicebus/v20170401:QueueAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueAuthorizationRule gets an existing QueueAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueAuthorizationRuleState, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	var resource QueueAuthorizationRule
	err := ctx.ReadResource("azure-nextgen:servicebus/v20170401:QueueAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueAuthorizationRule resources.
type queueAuthorizationRuleState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// Resource type
	Type *string `pulumi:"type"`
}

type QueueAuthorizationRuleState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (QueueAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleState)(nil)).Elem()
}

type queueAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The queue name.
	QueueName string `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
}

// The set of arguments for constructing a QueueAuthorizationRule resource.
type QueueAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// The queue name.
	QueueName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights AccessRightsArrayInput
}

func (QueueAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleArgs)(nil)).Elem()
}

type QueueAuthorizationRuleInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput
	ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput
}

func (QueueAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueAuthorizationRule)(nil)).Elem()
}

func (i QueueAuthorizationRule) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return i.ToQueueAuthorizationRuleOutputWithContext(context.Background())
}

func (i QueueAuthorizationRule) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleOutput)
}

type QueueAuthorizationRuleOutput struct {
	*pulumi.OutputState
}

func (QueueAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueAuthorizationRuleOutput)(nil)).Elem()
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(QueueAuthorizationRuleOutput{})
}
