// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170907privatepreview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a Kusto database.
type Database struct {
	pulumi.CustomResourceState

	// An ETag of the resource created.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The number of days of data that should be kept in cache for fast queries.
	HotCachePeriodInDays pulumi.IntPtrOutput `pulumi:"hotCachePeriodInDays"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The number of days data should be kept before it stops being accessible to queries.
	SoftDeletePeriodInDays pulumi.IntOutput `pulumi:"softDeletePeriodInDays"`
	// The statistics of the database.
	Statistics DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SoftDeletePeriodInDays == nil {
		return nil, errors.New("missing required argument 'SoftDeletePeriodInDays'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:kusto/latest:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20180907preview:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20190121:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20190515:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20190907:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20191109:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200215:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200614:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200918:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-nextgen:kusto/v20170907privatepreview:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-nextgen:kusto/v20170907privatepreview:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// An ETag of the resource created.
	Etag *string `pulumi:"etag"`
	// The number of days of data that should be kept in cache for fast queries.
	HotCachePeriodInDays *int `pulumi:"hotCachePeriodInDays"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The number of days data should be kept before it stops being accessible to queries.
	SoftDeletePeriodInDays *int `pulumi:"softDeletePeriodInDays"`
	// The statistics of the database.
	Statistics *DatabaseStatisticsResponse `pulumi:"statistics"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type DatabaseState struct {
	// An ETag of the resource created.
	Etag pulumi.StringPtrInput
	// The number of days of data that should be kept in cache for fast queries.
	HotCachePeriodInDays pulumi.IntPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The number of days data should be kept before it stops being accessible to queries.
	SoftDeletePeriodInDays pulumi.IntPtrInput
	// The statistics of the database.
	Statistics DatabaseStatisticsResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The number of days of data that should be kept in cache for fast queries.
	HotCachePeriodInDays *int `pulumi:"hotCachePeriodInDays"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of days data should be kept before it stops being accessible to queries.
	SoftDeletePeriodInDays int `pulumi:"softDeletePeriodInDays"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// The number of days of data that should be kept in cache for fast queries.
	HotCachePeriodInDays pulumi.IntPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The number of days data should be kept before it stops being accessible to queries.
	SoftDeletePeriodInDays pulumi.IntInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}
