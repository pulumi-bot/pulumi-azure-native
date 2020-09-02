// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180419

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A task resource
type Task struct {
	pulumi.CustomResourceState

	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom task properties
	Properties ProjectTaskPropertiesResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTask registers a new resource with the given unique name, arguments, and options.
func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.TaskName == nil {
		return nil, errors.New("missing required argument 'TaskName'")
	}
	if args == nil {
		args = &TaskArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datamigration/latest:Task"),
		},
		{
			Type: pulumi.String("azurerm:datamigration/v20171115preview:Task"),
		},
		{
			Type: pulumi.String("azurerm:datamigration/v20180315preview:Task"),
		},
		{
			Type: pulumi.String("azurerm:datamigration/v20180331preview:Task"),
		},
		{
			Type: pulumi.String("azurerm:datamigration/v20180715preview:Task"),
		},
	})
	opts = append(opts, aliases)
	var resource Task
	err := ctx.RegisterResource("azurerm:datamigration/v20180419:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTask gets an existing Task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("azurerm:datamigration/v20180419:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Task resources.
type taskState struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag *string `pulumi:"etag"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Custom task properties
	Properties *ProjectTaskPropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type TaskState struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Custom task properties
	Properties ProjectTaskPropertiesResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag *string `pulumi:"etag"`
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the project
	ProjectName string `pulumi:"projectName"`
	// Custom task properties
	Properties *ProjectTaskProperties `pulumi:"properties"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
	// Name of the Task
	TaskName string `pulumi:"taskName"`
}

// The set of arguments for constructing a Task resource.
type TaskArgs struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrInput
	// Name of the resource group
	GroupName pulumi.StringInput
	// Name of the project
	ProjectName pulumi.StringInput
	// Custom task properties
	Properties ProjectTaskPropertiesPtrInput
	// Name of the service
	ServiceName pulumi.StringInput
	// Name of the Task
	TaskName pulumi.StringInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}
