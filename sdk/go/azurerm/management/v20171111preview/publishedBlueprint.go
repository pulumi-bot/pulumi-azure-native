// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a published Blueprint.
//
// ## Example Usage
// ### PublishedBlueprint_Publish
//
// ```go
// package main
//
// import (
// 	management "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/management/v20171111preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := management.NewPublishedBlueprint(ctx, "publishedBlueprint", &management.PublishedBlueprintArgs{
// 			BlueprintName:       pulumi.String("simpleBlueprint"),
// 			ManagementGroupName: pulumi.String("ContosoOnlineGroup"),
// 			VersionId:           pulumi.String("v2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type PublishedBlueprint struct {
	pulumi.CustomResourceState

	// Name of the Blueprint definition.
	BlueprintName pulumi.StringPtrOutput `pulumi:"blueprintName"`
	// Version-specific change notes
	ChangeNotes pulumi.StringPtrOutput `pulumi:"changeNotes"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters required by this Blueprint definition.
	Parameters ParameterDefinitionResponseMapOutput `pulumi:"parameters"`
	// Resource group placeholders defined by this Blueprint definition.
	ResourceGroups ResourceGroupDefinitionResponseMapOutput `pulumi:"resourceGroups"`
	// Status of the Blueprint. This field is readonly.
	Status BlueprintStatusResponseOutput `pulumi:"status"`
	// The scope where this Blueprint can be applied.
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
	if args == nil || args.ManagementGroupName == nil {
		return nil, errors.New("missing required argument 'ManagementGroupName'")
	}
	if args == nil || args.VersionId == nil {
		return nil, errors.New("missing required argument 'VersionId'")
	}
	if args == nil {
		args = &PublishedBlueprintArgs{}
	}
	var resource PublishedBlueprint
	err := ctx.RegisterResource("azurerm:management/v20171111preview:PublishedBlueprint", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:management/v20171111preview:PublishedBlueprint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublishedBlueprint resources.
type publishedBlueprintState struct {
	// Name of the Blueprint definition.
	BlueprintName *string `pulumi:"blueprintName"`
	// Version-specific change notes
	ChangeNotes *string `pulumi:"changeNotes"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Name of this resource.
	Name *string `pulumi:"name"`
	// Parameters required by this Blueprint definition.
	Parameters map[string]ParameterDefinitionResponse `pulumi:"parameters"`
	// Resource group placeholders defined by this Blueprint definition.
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	// Status of the Blueprint. This field is readonly.
	Status *BlueprintStatusResponse `pulumi:"status"`
	// The scope where this Blueprint can be applied.
	TargetScope *string `pulumi:"targetScope"`
	// Type of this resource.
	Type *string `pulumi:"type"`
}

type PublishedBlueprintState struct {
	// Name of the Blueprint definition.
	BlueprintName pulumi.StringPtrInput
	// Version-specific change notes
	ChangeNotes pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Name of this resource.
	Name pulumi.StringPtrInput
	// Parameters required by this Blueprint definition.
	Parameters ParameterDefinitionResponseMapInput
	// Resource group placeholders defined by this Blueprint definition.
	ResourceGroups ResourceGroupDefinitionResponseMapInput
	// Status of the Blueprint. This field is readonly.
	Status BlueprintStatusResponsePtrInput
	// The scope where this Blueprint can be applied.
	TargetScope pulumi.StringPtrInput
	// Type of this resource.
	Type pulumi.StringPtrInput
}

func (PublishedBlueprintState) ElementType() reflect.Type {
	return reflect.TypeOf((*publishedBlueprintState)(nil)).Elem()
}

type publishedBlueprintArgs struct {
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// version of the published blueprint.
	VersionId string `pulumi:"versionId"`
}

// The set of arguments for constructing a PublishedBlueprint resource.
type PublishedBlueprintArgs struct {
	// name of the blueprint.
	BlueprintName pulumi.StringInput
	// ManagementGroup where blueprint stores.
	ManagementGroupName pulumi.StringInput
	// version of the published blueprint.
	VersionId pulumi.StringInput
}

func (PublishedBlueprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publishedBlueprintArgs)(nil)).Elem()
}
