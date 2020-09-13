// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Cassandra table.
//
// ## Example Usage
// ### CosmosDBCassandraTableCreateUpdate
//
// ```go
// package main
//
// import (
// 	documentdb "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/documentdb/v20190801"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := documentdb.NewCassandraResourceCassandraTable(ctx, "cassandraResourceCassandraTable", &documentdb.CassandraResourceCassandraTableArgs{
// 			AccountName:  pulumi.String("ddb1"),
// 			KeyspaceName: pulumi.String("keyspaceName"),
// 			Location:     pulumi.String("West US"),
// 			Options:      nil,
// 			Resource: &documentdb.CassandraTableResourceArgs{
// 				DefaultTtl: pulumi.Int(100),
// 				Id:         pulumi.String("tableName"),
// 				Schema: &documentdb.CassandraSchemaArgs{
// 					ClusterKeys: documentdb.ClusterKeyArray{
// 						&documentdb.ClusterKeyArgs{
// 							Name:    pulumi.String("columnA"),
// 							OrderBy: pulumi.String("Asc"),
// 						},
// 					},
// 					Columns: documentdb.ColumnArray{
// 						&documentdb.ColumnArgs{
// 							Name: pulumi.String("columnA"),
// 							Type: pulumi.String("Ascii"),
// 						},
// 					},
// 					PartitionKeys: documentdb.CassandraPartitionKeyArray{
// 						&documentdb.CassandraPartitionKeyArgs{
// 							Name: pulumi.String("columnA"),
// 						},
// 					},
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			TableName:         pulumi.String("tableName"),
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
type CassandraResourceCassandraTable struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                  `pulumi:"name"`
	Resource CassandraTableGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCassandraResourceCassandraTable registers a new resource with the given unique name, arguments, and options.
func NewCassandraResourceCassandraTable(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraTableArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraTable, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.KeyspaceName == nil {
		return nil, errors.New("missing required argument 'KeyspaceName'")
	}
	if args == nil || args.Options == nil {
		return nil, errors.New("missing required argument 'Options'")
	}
	if args == nil || args.Resource == nil {
		return nil, errors.New("missing required argument 'Resource'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TableName == nil {
		return nil, errors.New("missing required argument 'TableName'")
	}
	if args == nil {
		args = &CassandraResourceCassandraTableArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20191212:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200301:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200601preview:CassandraResourceCassandraTable"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraResourceCassandraTable
	err := ctx.RegisterResource("azurerm:documentdb/v20190801:CassandraResourceCassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraResourceCassandraTable gets an existing CassandraResourceCassandraTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraResourceCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraTableState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraTable, error) {
	var resource CassandraResourceCassandraTable
	err := ctx.ReadResource("azurerm:documentdb/v20190801:CassandraResourceCassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraResourceCassandraTable resources.
type cassandraResourceCassandraTableState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     *string                                      `pulumi:"name"`
	Resource *CassandraTableGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type CassandraResourceCassandraTableState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name     pulumi.StringPtrInput
	Resource CassandraTableGetPropertiesResponseResourcePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (CassandraResourceCassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableState)(nil)).Elem()
}

type cassandraResourceCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName string `pulumi:"keyspaceName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName string `pulumi:"tableName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CassandraResourceCassandraTable resource.
type CassandraResourceCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB table name.
	TableName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (CassandraResourceCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableArgs)(nil)).Elem()
}
