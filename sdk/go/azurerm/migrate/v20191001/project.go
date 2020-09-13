// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Azure Migrate Project.
//
// ## Example Usage
// ### Projects_Create
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
// 		_, err := migrate.NewProject(ctx, "project", &migrate.ProjectArgs{
// 			ETag:              pulumi.String(""),
// 			Location:          pulumi.String("West Europe"),
// 			ProjectName:       pulumi.String("abGoyalProject2"),
// 			ResourceGroupName: pulumi.String("abgoyal-westEurope"),
// 			Tags:              nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Project struct {
	pulumi.CustomResourceState

	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Azure location in which project is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the project.
	Properties ProjectPropertiesResponseOutput `pulumi:"properties"`
	// Tags provided by Azure Tagging service.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ProjectArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:migrate/latest:Project"),
		},
	})
	opts = append(opts, aliases)
	var resource Project
	err := ctx.RegisterResource("azurerm:migrate/v20191001:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azurerm:migrate/v20191001:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the project.
	Name *string `pulumi:"name"`
	// Properties of the project.
	Properties *ProjectPropertiesResponse `pulumi:"properties"`
	// Tags provided by Azure Tagging service.
	Tags map[string]interface{} `pulumi:"tags"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects].
	Type *string `pulumi:"type"`
}

type ProjectState struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which project is created.
	Location pulumi.StringPtrInput
	// Name of the project.
	Name pulumi.StringPtrInput
	// Properties of the project.
	Properties ProjectPropertiesResponsePtrInput
	// Tags provided by Azure Tagging service.
	Tags pulumi.MapInput
	// Type of the object = [Microsoft.Migrate/assessmentProjects].
	Type pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Properties of the project.
	Properties *ProjectProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags provided by Azure Tagging service.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which project is created.
	Location pulumi.StringPtrInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	// Properties of the project.
	Properties ProjectPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
	// Tags provided by Azure Tagging service.
	Tags pulumi.MapInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}
