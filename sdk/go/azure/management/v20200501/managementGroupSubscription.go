// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The details of subscription under management group.
type ManagementGroupSubscription struct {
	pulumi.CustomResourceState

	// The friendly name of the subscription.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the parent management group.
	Parent DescendantParentGroupInfoResponsePtrOutput `pulumi:"parent"`
	// The state of the subscription.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant pulumi.StringPtrOutput `pulumi:"tenant"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementGroupSubscription registers a new resource with the given unique name, arguments, and options.
func NewManagementGroupSubscription(ctx *pulumi.Context,
	name string, args *ManagementGroupSubscriptionArgs, opts ...pulumi.ResourceOption) (*ManagementGroupSubscription, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil {
		args = &ManagementGroupSubscriptionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:management/latest:ManagementGroupSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementGroupSubscription
	err := ctx.RegisterResource("azure-nextgen:management/v20200501:ManagementGroupSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementGroupSubscription gets an existing ManagementGroupSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementGroupSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupSubscriptionState, opts ...pulumi.ResourceOption) (*ManagementGroupSubscription, error) {
	var resource ManagementGroupSubscription
	err := ctx.ReadResource("azure-nextgen:management/v20200501:ManagementGroupSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementGroupSubscription resources.
type managementGroupSubscriptionState struct {
	// The friendly name of the subscription.
	DisplayName *string `pulumi:"displayName"`
	// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name *string `pulumi:"name"`
	// The ID of the parent management group.
	Parent *DescendantParentGroupInfoResponse `pulumi:"parent"`
	// The state of the subscription.
	State *string `pulumi:"state"`
	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant *string `pulumi:"tenant"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
	Type *string `pulumi:"type"`
}

type ManagementGroupSubscriptionState struct {
	// The friendly name of the subscription.
	DisplayName pulumi.StringPtrInput
	// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringPtrInput
	// The ID of the parent management group.
	Parent DescendantParentGroupInfoResponsePtrInput
	// The state of the subscription.
	State pulumi.StringPtrInput
	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant pulumi.StringPtrInput
	// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
	Type pulumi.StringPtrInput
}

func (ManagementGroupSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupSubscriptionState)(nil)).Elem()
}

type managementGroupSubscriptionArgs struct {
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
}

// The set of arguments for constructing a ManagementGroupSubscription resource.
type ManagementGroupSubscriptionArgs struct {
	// Management Group ID.
	GroupId pulumi.StringInput
}

func (ManagementGroupSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupSubscriptionArgs)(nil)).Elem()
}

type ManagementGroupSubscriptionInput interface {
	pulumi.Input

	ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput
	ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput
}

func (ManagementGroupSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupSubscription)(nil)).Elem()
}

func (i ManagementGroupSubscription) ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput {
	return i.ToManagementGroupSubscriptionOutputWithContext(context.Background())
}

func (i ManagementGroupSubscription) ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupSubscriptionOutput)
}

type ManagementGroupSubscriptionOutput struct {
	*pulumi.OutputState
}

func (ManagementGroupSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupSubscriptionOutput)(nil)).Elem()
}

func (o ManagementGroupSubscriptionOutput) ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput {
	return o
}

func (o ManagementGroupSubscriptionOutput) ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupSubscriptionOutput{})
}
