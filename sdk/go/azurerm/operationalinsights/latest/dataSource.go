// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Datasources under OMS Workspace.
type DataSource struct {
	pulumi.CustomResourceState

	// The ETag of the data source.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the DataSource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties pulumi.MapOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil || args.DataSourceName == nil {
		return nil, errors.New("missing required argument 'DataSourceName'")
	}
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &DataSourceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:operationalinsights/v20151101preview:DataSource"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200301preview:DataSource"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200801:DataSource"),
		},
	})
	opts = append(opts, aliases)
	var resource DataSource
	err := ctx.RegisterResource("azurerm:operationalinsights/latest:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("azurerm:operationalinsights/latest:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
	// The ETag of the data source.
	Etag *string `pulumi:"etag"`
	// The kind of the DataSource.
	Kind *string `pulumi:"kind"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties map[string]interface{} `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type DataSourceState struct {
	// The ETag of the data source.
	Etag pulumi.StringPtrInput
	// The kind of the DataSource.
	Kind pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties pulumi.MapInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// The name of the datasource resource.
	DataSourceName string `pulumi:"dataSourceName"`
	// The ETag of the data source.
	Etag *string `pulumi:"etag"`
	// The kind of the DataSource.
	Kind string `pulumi:"kind"`
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties map[string]interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// The name of the datasource resource.
	DataSourceName pulumi.StringInput
	// The ETag of the data source.
	Etag pulumi.StringPtrInput
	// The kind of the DataSource.
	Kind pulumi.StringInput
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties pulumi.MapInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}
