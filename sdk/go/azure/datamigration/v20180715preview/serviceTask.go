// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180715preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A task resource
type ServiceTask struct {
	pulumi.CustomResourceState

	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom task properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceTask registers a new resource with the given unique name, arguments, and options.
func NewServiceTask(ctx *pulumi.Context,
	name string, args *ServiceTaskArgs, opts ...pulumi.ResourceOption) (*ServiceTask, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.TaskName == nil {
		return nil, errors.New("missing required argument 'TaskName'")
	}
	if args == nil {
		args = &ServiceTaskArgs{}
	}
	var resource ServiceTask
	err := ctx.RegisterResource("azure-nextgen:datamigration/v20180715preview:ServiceTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceTask gets an existing ServiceTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTaskState, opts ...pulumi.ResourceOption) (*ServiceTask, error) {
	var resource ServiceTask
	err := ctx.ReadResource("azure-nextgen:datamigration/v20180715preview:ServiceTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceTask resources.
type serviceTaskState struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag *string `pulumi:"etag"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Custom task properties
	Properties interface{} `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ServiceTaskState struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Custom task properties
	Properties pulumi.Input
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ServiceTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTaskState)(nil)).Elem()
}

type serviceTaskArgs struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag *string `pulumi:"etag"`
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Custom task properties
	Properties interface{} `pulumi:"properties"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
	// Name of the Task
	TaskName string `pulumi:"taskName"`
}

// The set of arguments for constructing a ServiceTask resource.
type ServiceTaskArgs struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrInput
	// Name of the resource group
	GroupName pulumi.StringInput
	// Custom task properties
	Properties pulumi.Input
	// Name of the service
	ServiceName pulumi.StringInput
	// Name of the Task
	TaskName pulumi.StringInput
}

func (ServiceTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTaskArgs)(nil)).Elem()
}

type ServiceTaskInput interface {
	pulumi.Input

	ToServiceTaskOutput() ServiceTaskOutput
	ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput
}

func (ServiceTask) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTask)(nil)).Elem()
}

func (i ServiceTask) ToServiceTaskOutput() ServiceTaskOutput {
	return i.ToServiceTaskOutputWithContext(context.Background())
}

func (i ServiceTask) ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTaskOutput)
}

type ServiceTaskOutput struct {
	*pulumi.OutputState
}

func (ServiceTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTaskOutput)(nil)).Elem()
}

func (o ServiceTaskOutput) ToServiceTaskOutput() ServiceTaskOutput {
	return o
}

func (o ServiceTaskOutput) ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceTaskOutput{})
}
