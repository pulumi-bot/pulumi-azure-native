// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a namespace authorization rule.
type QueueAuthorizationRule struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// AuthorizationRule properties.
	Properties SBAuthorizationRuleResponsePropertiesOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQueueAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewQueueAuthorizationRule(ctx *pulumi.Context,
	name string, args *QueueAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.QueueName == nil {
		return nil, errors.New("missing required argument 'QueueName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Rights == nil {
		return nil, errors.New("missing required argument 'Rights'")
	}
	if args == nil {
		args = &QueueAuthorizationRuleArgs{}
	}
	var resource QueueAuthorizationRule
	err := ctx.RegisterResource("azurerm:servicebus/v20170401:QueueAuthorizationRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:servicebus/v20170401:QueueAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueAuthorizationRule resources.
type queueAuthorizationRuleState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// AuthorizationRule properties.
	Properties *SBAuthorizationRuleResponseProperties `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type QueueAuthorizationRuleState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// AuthorizationRule properties.
	Properties SBAuthorizationRuleResponsePropertiesPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (QueueAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleState)(nil)).Elem()
}

type queueAuthorizationRuleArgs struct {
	// The authorization rule name.
	Name string `pulumi:"name"`
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
	Name pulumi.StringInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// The queue name.
	QueueName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
}

func (QueueAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleArgs)(nil)).Elem()
}
