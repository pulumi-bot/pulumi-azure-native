// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a namespace authorization rule.
type TopicAuthorizationRule struct {
	pulumi.CustomResourceState

	// A string that describes Claim Type for authorization rule.
	ClaimType pulumi.StringPtrOutput `pulumi:"claimType"`
	// A string that describes Claim Value of authorization rule.
	ClaimValue pulumi.StringPtrOutput `pulumi:"claimValue"`
	// The time the namespace was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// A string that describes the Key Name of authorization rule.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The time the namespace was updated.
	ModifiedTime pulumi.StringOutput `pulumi:"modifiedTime"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringPtrOutput `pulumi:"primaryKey"`
	// The rights associated with the rule.
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringPtrOutput `pulumi:"secondaryKey"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopicAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewTopicAuthorizationRule(ctx *pulumi.Context,
	name string, args *TopicAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	if args == nil || args.AuthorizationRuleName == nil {
		return nil, errors.New("missing required argument 'AuthorizationRuleName'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus/latest:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20150801:TopicAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:TopicAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource TopicAuthorizationRule
	err := ctx.RegisterResource("azure-nextgen:servicebus/v20140901:TopicAuthorizationRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:servicebus/v20140901:TopicAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicAuthorizationRule resources.
type topicAuthorizationRuleState struct {
	// A string that describes Claim Type for authorization rule.
	ClaimType *string `pulumi:"claimType"`
	// A string that describes Claim Value of authorization rule.
	ClaimValue *string `pulumi:"claimValue"`
	// The time the namespace was created.
	CreatedTime *string `pulumi:"createdTime"`
	// A string that describes the Key Name of authorization rule.
	KeyName *string `pulumi:"keyName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The time the namespace was updated.
	ModifiedTime *string `pulumi:"modifiedTime"`
	// Resource name
	Name *string `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Resource type
	Type *string `pulumi:"type"`
}

type TopicAuthorizationRuleState struct {
	// A string that describes Claim Type for authorization rule.
	ClaimType pulumi.StringPtrInput
	// A string that describes Claim Value of authorization rule.
	ClaimValue pulumi.StringPtrInput
	// The time the namespace was created.
	CreatedTime pulumi.StringPtrInput
	// A string that describes the Key Name of authorization rule.
	KeyName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The time the namespace was updated.
	ModifiedTime pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringPtrInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (TopicAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleState)(nil)).Elem()
}

type topicAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// A string that describes Claim Type for authorization rule.
	ClaimType *string `pulumi:"claimType"`
	// A string that describes Claim Value of authorization rule.
	ClaimValue *string `pulumi:"claimValue"`
	// A string that describes the Key Name of authorization rule.
	KeyName *string `pulumi:"keyName"`
	// data center location.
	Location *string `pulumi:"location"`
	// Name of the authorization rule.
	Name *string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a TopicAuthorizationRule resource.
type TopicAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput
	// A string that describes Claim Type for authorization rule.
	ClaimType pulumi.StringPtrInput
	// A string that describes Claim Value of authorization rule.
	ClaimValue pulumi.StringPtrInput
	// A string that describes the Key Name of authorization rule.
	KeyName pulumi.StringPtrInput
	// data center location.
	Location pulumi.StringPtrInput
	// Name of the authorization rule.
	Name pulumi.StringPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringPtrInput
	// The topic name.
	TopicName pulumi.StringInput
}

func (TopicAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleArgs)(nil)).Elem()
}

type TopicAuthorizationRuleInput interface {
	pulumi.Input

	ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput
	ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput
}

func (TopicAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicAuthorizationRule)(nil)).Elem()
}

func (i TopicAuthorizationRule) ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput {
	return i.ToTopicAuthorizationRuleOutputWithContext(context.Background())
}

func (i TopicAuthorizationRule) ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRuleOutput)
}

type TopicAuthorizationRuleOutput struct {
	*pulumi.OutputState
}

func (TopicAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicAuthorizationRuleOutput)(nil)).Elem()
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput {
	return o
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TopicAuthorizationRuleOutput{})
}
