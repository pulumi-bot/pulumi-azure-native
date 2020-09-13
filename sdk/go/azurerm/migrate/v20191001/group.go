// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A group created in a Migration project.
//
// ## Example Usage
// ### Groups_Create
//
// ```go
// package main
//
// import (
// 	migrate "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/migrate/v20191001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := migrate.NewGroup(ctx, "group", &migrate.GroupArgs{
// 			ETag:              pulumi.String("\"1e000c2c-0000-0d00-0000-5cdaa4190000\""),
// 			GroupName:         pulumi.String("Group2"),
// 			ProjectName:       pulumi.String("abgoyalWEselfhostb72bproject"),
// 			ResourceGroupName: pulumi.String("abgoyal-westEurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Name of the group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the group.
	Properties GroupPropertiesResponseOutput `pulumi:"properties"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects/groups].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:migrate/latest:Group"),
		},
	})
	opts = append(opts, aliases)
	var resource Group
	err := ctx.RegisterResource("azurerm:migrate/v20191001:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azurerm:migrate/v20191001:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Name of the group.
	Name *string `pulumi:"name"`
	// Properties of the group.
	Properties *GroupPropertiesResponse `pulumi:"properties"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects/groups].
	Type *string `pulumi:"type"`
}

type GroupState struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Name of the group.
	Name pulumi.StringPtrInput
	// Properties of the group.
	Properties GroupPropertiesResponsePtrInput
	// Type of the object = [Microsoft.Migrate/assessmentProjects/groups].
	Type pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Unique name of a group within a project.
	GroupName string `pulumi:"groupName"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Unique name of a group within a project.
	GroupName pulumi.StringInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
