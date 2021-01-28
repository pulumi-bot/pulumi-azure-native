// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Workload group operations for a sql pool
type SqlPoolWorkloadGroup struct {
	pulumi.CustomResourceState

	// The workload group importance level.
	Importance pulumi.StringPtrOutput `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntOutput `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrOutput `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntOutput `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64Output `pulumi:"minResourcePercentPerRequest"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrOutput `pulumi:"queryExecutionTimeout"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlPoolWorkloadGroup registers a new resource with the given unique name, arguments, and options.
func NewSqlPoolWorkloadGroup(ctx *pulumi.Context,
	name string, args *SqlPoolWorkloadGroupArgs, opts ...pulumi.ResourceOption) (*SqlPoolWorkloadGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxResourcePercent == nil {
		return nil, errors.New("invalid value for required argument 'MaxResourcePercent'")
	}
	if args.MinResourcePercent == nil {
		return nil, errors.New("invalid value for required argument 'MinResourcePercent'")
	}
	if args.MinResourcePercentPerRequest == nil {
		return nil, errors.New("invalid value for required argument 'MinResourcePercentPerRequest'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.WorkloadGroupName == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20201201:SqlPoolWorkloadGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolWorkloadGroup
	err := ctx.RegisterResource("azure-nextgen:synapse/v20190601preview:SqlPoolWorkloadGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlPoolWorkloadGroup gets an existing SqlPoolWorkloadGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlPoolWorkloadGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolWorkloadGroupState, opts ...pulumi.ResourceOption) (*SqlPoolWorkloadGroup, error) {
	var resource SqlPoolWorkloadGroup
	err := ctx.ReadResource("azure-nextgen:synapse/v20190601preview:SqlPoolWorkloadGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlPoolWorkloadGroup resources.
type sqlPoolWorkloadGroupState struct {
	// The workload group importance level.
	Importance *string `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent *int `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent *int `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest *float64 `pulumi:"minResourcePercentPerRequest"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type SqlPoolWorkloadGroupState struct {
	// The workload group importance level.
	Importance pulumi.StringPtrInput
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntPtrInput
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntPtrInput
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64PtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (SqlPoolWorkloadGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolWorkloadGroupState)(nil)).Elem()
}

type sqlPoolWorkloadGroupArgs struct {
	// The workload group importance level.
	Importance *string `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent int `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent int `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest float64 `pulumi:"minResourcePercentPerRequest"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the workload group.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SqlPoolWorkloadGroup resource.
type SqlPoolWorkloadGroupArgs struct {
	// The workload group importance level.
	Importance pulumi.StringPtrInput
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntInput
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntInput
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64Input
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// SQL pool name
	SqlPoolName pulumi.StringInput
	// The name of the workload group.
	WorkloadGroupName pulumi.StringInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (SqlPoolWorkloadGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolWorkloadGroupArgs)(nil)).Elem()
}

type SqlPoolWorkloadGroupInput interface {
	pulumi.Input

	ToSqlPoolWorkloadGroupOutput() SqlPoolWorkloadGroupOutput
	ToSqlPoolWorkloadGroupOutputWithContext(ctx context.Context) SqlPoolWorkloadGroupOutput
}

func (*SqlPoolWorkloadGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolWorkloadGroup)(nil))
}

func (i *SqlPoolWorkloadGroup) ToSqlPoolWorkloadGroupOutput() SqlPoolWorkloadGroupOutput {
	return i.ToSqlPoolWorkloadGroupOutputWithContext(context.Background())
}

func (i *SqlPoolWorkloadGroup) ToSqlPoolWorkloadGroupOutputWithContext(ctx context.Context) SqlPoolWorkloadGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolWorkloadGroupOutput)
}

type SqlPoolWorkloadGroupOutput struct {
	*pulumi.OutputState
}

func (SqlPoolWorkloadGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolWorkloadGroup)(nil))
}

func (o SqlPoolWorkloadGroupOutput) ToSqlPoolWorkloadGroupOutput() SqlPoolWorkloadGroupOutput {
	return o
}

func (o SqlPoolWorkloadGroupOutput) ToSqlPoolWorkloadGroupOutputWithContext(ctx context.Context) SqlPoolWorkloadGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlPoolWorkloadGroupOutput{})
}
