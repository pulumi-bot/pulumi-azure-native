// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The policy definition.
type PolicyDefinition struct {
	pulumi.CustomResourceState

	// The name of the policy definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy definition properties.
	Properties PolicyDefinitionPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicyDefinition registers a new resource with the given unique name, arguments, and options.
func NewPolicyDefinition(ctx *pulumi.Context,
	name string, args *PolicyDefinitionArgs, opts ...pulumi.ResourceOption) (*PolicyDefinition, error) {
	if args == nil || args.ManagementGroupId == nil {
		return nil, errors.New("missing required argument 'ManagementGroupId'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &PolicyDefinitionArgs{}
	}
	var resource PolicyDefinition
	err := ctx.RegisterResource("azurerm:authorization:PolicyDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyDefinition gets an existing PolicyDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionState, opts ...pulumi.ResourceOption) (*PolicyDefinition, error) {
	var resource PolicyDefinition
	err := ctx.ReadResource("azurerm:authorization:PolicyDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyDefinition resources.
type policyDefinitionState struct {
	// The name of the policy definition.
	Name *string `pulumi:"name"`
	// The policy definition properties.
	Properties *PolicyDefinitionPropertiesResponse `pulumi:"properties"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type *string `pulumi:"type"`
}

type PolicyDefinitionState struct {
	// The name of the policy definition.
	Name pulumi.StringPtrInput
	// The policy definition properties.
	Properties PolicyDefinitionPropertiesResponsePtrInput
	// The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type pulumi.StringPtrInput
}

func (PolicyDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionState)(nil)).Elem()
}

type policyDefinitionArgs struct {
	// The ID of the management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The name of the policy definition to create.
	Name string `pulumi:"name"`
	// The policy definition properties.
	Properties *PolicyDefinitionProperties `pulumi:"properties"`
}

// The set of arguments for constructing a PolicyDefinition resource.
type PolicyDefinitionArgs struct {
	// The ID of the management group.
	ManagementGroupId pulumi.StringInput
	// The name of the policy definition to create.
	Name pulumi.StringInput
	// The policy definition properties.
	Properties PolicyDefinitionPropertiesPtrInput
}

func (PolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionArgs)(nil)).Elem()
}