// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The policy assignment.
type PolicyAssignment struct {
	pulumi.CustomResourceState

	// The managed identity associated with the policy assignment.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the policy assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties for the policy assignment.
	Properties PolicyAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// The policy sku. This property is optional, obsolete, and will be ignored.
	Sku PolicySkuResponsePtrOutput `pulumi:"sku"`
	// The type of the policy assignment.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAssignment(ctx *pulumi.Context,
	name string, args *PolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &PolicyAssignmentArgs{}
	}
	var resource PolicyAssignment
	err := ctx.RegisterResource("azurerm:authorization/v20190601:PolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAssignment gets an existing PolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentState, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	var resource PolicyAssignment
	err := ctx.ReadResource("azurerm:authorization/v20190601:PolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAssignment resources.
type policyAssignmentState struct {
	// The managed identity associated with the policy assignment.
	Identity *IdentityResponse `pulumi:"identity"`
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location *string `pulumi:"location"`
	// The name of the policy assignment.
	Name *string `pulumi:"name"`
	// Properties for the policy assignment.
	Properties *PolicyAssignmentPropertiesResponse `pulumi:"properties"`
	// The policy sku. This property is optional, obsolete, and will be ignored.
	Sku *PolicySkuResponse `pulumi:"sku"`
	// The type of the policy assignment.
	Type *string `pulumi:"type"`
}

type PolicyAssignmentState struct {
	// The managed identity associated with the policy assignment.
	Identity IdentityResponsePtrInput
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location pulumi.StringPtrInput
	// The name of the policy assignment.
	Name pulumi.StringPtrInput
	// Properties for the policy assignment.
	Properties PolicyAssignmentPropertiesResponsePtrInput
	// The policy sku. This property is optional, obsolete, and will be ignored.
	Sku PolicySkuResponsePtrInput
	// The type of the policy assignment.
	Type pulumi.StringPtrInput
}

func (PolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentState)(nil)).Elem()
}

type policyAssignmentArgs struct {
	// This message will be part of response in case of policy violation.
	Description *string `pulumi:"description"`
	// The display name of the policy assignment.
	DisplayName *string `pulumi:"displayName"`
	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode *string `pulumi:"enforcementMode"`
	// The managed identity associated with the policy assignment.
	Identity *Identity `pulumi:"identity"`
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location *string `pulumi:"location"`
	// The policy assignment metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the policy assignment.
	Name string `pulumi:"name"`
	// The policy's excluded scopes.
	NotScopes []string `pulumi:"notScopes"`
	// Required if a parameter is used in policy rule.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The scope for the policy assignment.
	Scope string `pulumi:"scope"`
	// The policy sku. This property is optional, obsolete, and will be ignored.
	Sku *PolicySku `pulumi:"sku"`
}

// The set of arguments for constructing a PolicyAssignment resource.
type PolicyAssignmentArgs struct {
	// This message will be part of response in case of policy violation.
	Description pulumi.StringPtrInput
	// The display name of the policy assignment.
	DisplayName pulumi.StringPtrInput
	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode pulumi.StringPtrInput
	// The managed identity associated with the policy assignment.
	Identity IdentityPtrInput
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location pulumi.StringPtrInput
	// The policy assignment metadata.
	Metadata pulumi.MapInput
	// The name of the policy assignment.
	Name pulumi.StringInput
	// The policy's excluded scopes.
	NotScopes pulumi.StringArrayInput
	// Required if a parameter is used in policy rule.
	Parameters pulumi.MapInput
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId pulumi.StringPtrInput
	// The scope for the policy assignment.
	Scope pulumi.StringInput
	// The policy sku. This property is optional, obsolete, and will be ignored.
	Sku PolicySkuPtrInput
}

func (PolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArgs)(nil)).Elem()
}
