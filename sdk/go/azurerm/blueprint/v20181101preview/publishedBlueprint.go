// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a published blueprint.
type PublishedBlueprint struct {
	pulumi.CustomResourceState

	// Name of the published blueprint definition.
	BlueprintName pulumi.StringPtrOutput `pulumi:"blueprintName"`
	// Version-specific change notes.
	ChangeNotes pulumi.StringPtrOutput `pulumi:"changeNotes"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters required by this blueprint definition.
	Parameters ParameterDefinitionResponseMapOutput `pulumi:"parameters"`
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups ResourceGroupDefinitionResponseMapOutput `pulumi:"resourceGroups"`
	// Status of the blueprint. This field is readonly.
	Status BlueprintStatusResponseOutput `pulumi:"status"`
	// The scope where this blueprint definition can be assigned.
	TargetScope pulumi.StringPtrOutput `pulumi:"targetScope"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPublishedBlueprint registers a new resource with the given unique name, arguments, and options.
func NewPublishedBlueprint(ctx *pulumi.Context,
	name string, args *PublishedBlueprintArgs, opts ...pulumi.ResourceOption) (*PublishedBlueprint, error) {
	if args == nil || args.BlueprintName == nil {
		return nil, errors.New("missing required argument 'BlueprintName'")
	}
	if args == nil || args.ResourceScope == nil {
		return nil, errors.New("missing required argument 'ResourceScope'")
	}
	if args == nil || args.VersionId == nil {
		return nil, errors.New("missing required argument 'VersionId'")
	}
	if args == nil {
		args = &PublishedBlueprintArgs{}
	}
	var resource PublishedBlueprint
	err := ctx.RegisterResource("azurerm:blueprint/v20181101preview:PublishedBlueprint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublishedBlueprint gets an existing PublishedBlueprint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublishedBlueprint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublishedBlueprintState, opts ...pulumi.ResourceOption) (*PublishedBlueprint, error) {
	var resource PublishedBlueprint
	err := ctx.ReadResource("azurerm:blueprint/v20181101preview:PublishedBlueprint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublishedBlueprint resources.
type publishedBlueprintState struct {
	// Name of the published blueprint definition.
	BlueprintName *string `pulumi:"blueprintName"`
	// Version-specific change notes.
	ChangeNotes *string `pulumi:"changeNotes"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Name of this resource.
	Name *string `pulumi:"name"`
	// Parameters required by this blueprint definition.
	Parameters map[string]ParameterDefinitionResponse `pulumi:"parameters"`
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	// Status of the blueprint. This field is readonly.
	Status *BlueprintStatusResponse `pulumi:"status"`
	// The scope where this blueprint definition can be assigned.
	TargetScope *string `pulumi:"targetScope"`
	// Type of this resource.
	Type *string `pulumi:"type"`
}

type PublishedBlueprintState struct {
	// Name of the published blueprint definition.
	BlueprintName pulumi.StringPtrInput
	// Version-specific change notes.
	ChangeNotes pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Name of this resource.
	Name pulumi.StringPtrInput
	// Parameters required by this blueprint definition.
	Parameters ParameterDefinitionResponseMapInput
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups ResourceGroupDefinitionResponseMapInput
	// Status of the blueprint. This field is readonly.
	Status BlueprintStatusResponsePtrInput
	// The scope where this blueprint definition can be assigned.
	TargetScope pulumi.StringPtrInput
	// Type of this resource.
	Type pulumi.StringPtrInput
}

func (PublishedBlueprintState) ElementType() reflect.Type {
	return reflect.TypeOf((*publishedBlueprintState)(nil)).Elem()
}

type publishedBlueprintArgs struct {
	// Name of the published blueprint definition.
	BlueprintName string `pulumi:"blueprintName"`
	// Version-specific change notes.
	ChangeNotes *string `pulumi:"changeNotes"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Parameters required by this blueprint definition.
	Parameters map[string]ParameterDefinition `pulumi:"parameters"`
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]ResourceGroupDefinition `pulumi:"resourceGroups"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
	// The scope where this blueprint definition can be assigned.
	TargetScope *string `pulumi:"targetScope"`
	// Version of the published blueprint definition.
	VersionId string `pulumi:"versionId"`
}

// The set of arguments for constructing a PublishedBlueprint resource.
type PublishedBlueprintArgs struct {
	// Name of the published blueprint definition.
	BlueprintName pulumi.StringInput
	// Version-specific change notes.
	ChangeNotes pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Parameters required by this blueprint definition.
	Parameters ParameterDefinitionMapInput
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups ResourceGroupDefinitionMapInput
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput
	// The scope where this blueprint definition can be assigned.
	TargetScope pulumi.StringPtrInput
	// Version of the published blueprint definition.
	VersionId pulumi.StringInput
}

func (PublishedBlueprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publishedBlueprintArgs)(nil)).Elem()
}
