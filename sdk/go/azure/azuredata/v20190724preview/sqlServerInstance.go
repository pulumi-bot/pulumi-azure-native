// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190724preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A SqlServerInstance.
type SqlServerInstance struct {
	pulumi.CustomResourceState

	// ARM Resource id of the container resource (Azure Arc for Servers)
	ContainerResourceId pulumi.StringOutput `pulumi:"containerResourceId"`
	// The time when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// SQL Server edition.
	Edition pulumi.StringOutput `pulumi:"edition"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The cloud connectivity status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time when the resource was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The number of logical processors used by the SQL Server instance.
	VCore pulumi.StringOutput `pulumi:"vCore"`
	// SQL Server version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSqlServerInstance registers a new resource with the given unique name, arguments, and options.
func NewSqlServerInstance(ctx *pulumi.Context,
	name string, args *SqlServerInstanceArgs, opts ...pulumi.ResourceOption) (*SqlServerInstance, error) {
	if args == nil || args.ContainerResourceId == nil {
		return nil, errors.New("missing required argument 'ContainerResourceId'")
	}
	if args == nil || args.Edition == nil {
		return nil, errors.New("missing required argument 'Edition'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SqlServerInstanceName == nil {
		return nil, errors.New("missing required argument 'SqlServerInstanceName'")
	}
	if args == nil || args.Status == nil {
		return nil, errors.New("missing required argument 'Status'")
	}
	if args == nil || args.VCore == nil {
		return nil, errors.New("missing required argument 'VCore'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	if args == nil {
		args = &SqlServerInstanceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:azuredata/v20200908preview:SqlServerInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlServerInstance
	err := ctx.RegisterResource("azure-nextgen:azuredata/v20190724preview:SqlServerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlServerInstance gets an existing SqlServerInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlServerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerInstanceState, opts ...pulumi.ResourceOption) (*SqlServerInstance, error) {
	var resource SqlServerInstance
	err := ctx.ReadResource("azure-nextgen:azuredata/v20190724preview:SqlServerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlServerInstance resources.
type sqlServerInstanceState struct {
	// ARM Resource id of the container resource (Azure Arc for Servers)
	ContainerResourceId *string `pulumi:"containerResourceId"`
	// The time when the resource was created.
	CreateTime *string `pulumi:"createTime"`
	// SQL Server edition.
	Edition *string `pulumi:"edition"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The cloud connectivity status.
	Status *string `pulumi:"status"`
	// Read only system data
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// The time when the resource was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// The number of logical processors used by the SQL Server instance.
	VCore *string `pulumi:"vCore"`
	// SQL Server version.
	Version *string `pulumi:"version"`
}

type SqlServerInstanceState struct {
	// ARM Resource id of the container resource (Azure Arc for Servers)
	ContainerResourceId pulumi.StringPtrInput
	// The time when the resource was created.
	CreateTime pulumi.StringPtrInput
	// SQL Server edition.
	Edition pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The cloud connectivity status.
	Status pulumi.StringPtrInput
	// Read only system data
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// The time when the resource was last updated.
	UpdateTime pulumi.StringPtrInput
	// The number of logical processors used by the SQL Server instance.
	VCore pulumi.StringPtrInput
	// SQL Server version.
	Version pulumi.StringPtrInput
}

func (SqlServerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerInstanceState)(nil)).Elem()
}

type sqlServerInstanceArgs struct {
	// ARM Resource id of the container resource (Azure Arc for Servers)
	ContainerResourceId string `pulumi:"containerResourceId"`
	// SQL Server edition.
	Edition string `pulumi:"edition"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of SQL Server Instance
	SqlServerInstanceName string `pulumi:"sqlServerInstanceName"`
	// The cloud connectivity status.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The number of logical processors used by the SQL Server instance.
	VCore string `pulumi:"vCore"`
	// SQL Server version.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a SqlServerInstance resource.
type SqlServerInstanceArgs struct {
	// ARM Resource id of the container resource (Azure Arc for Servers)
	ContainerResourceId pulumi.StringInput
	// SQL Server edition.
	Edition pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput
	// The name of SQL Server Instance
	SqlServerInstanceName pulumi.StringInput
	// The cloud connectivity status.
	Status pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The number of logical processors used by the SQL Server instance.
	VCore pulumi.StringInput
	// SQL Server version.
	Version pulumi.StringInput
}

func (SqlServerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerInstanceArgs)(nil)).Elem()
}

type SqlServerInstanceInput interface {
	pulumi.Input

	ToSqlServerInstanceOutput() SqlServerInstanceOutput
	ToSqlServerInstanceOutputWithContext(ctx context.Context) SqlServerInstanceOutput
}

func (SqlServerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstance)(nil)).Elem()
}

func (i SqlServerInstance) ToSqlServerInstanceOutput() SqlServerInstanceOutput {
	return i.ToSqlServerInstanceOutputWithContext(context.Background())
}

func (i SqlServerInstance) ToSqlServerInstanceOutputWithContext(ctx context.Context) SqlServerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstanceOutput)
}

type SqlServerInstanceOutput struct {
	*pulumi.OutputState
}

func (SqlServerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstanceOutput)(nil)).Elem()
}

func (o SqlServerInstanceOutput) ToSqlServerInstanceOutput() SqlServerInstanceOutput {
	return o
}

func (o SqlServerInstanceOutput) ToSqlServerInstanceOutputWithContext(ctx context.Context) SqlServerInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlServerInstanceOutput{})
}
