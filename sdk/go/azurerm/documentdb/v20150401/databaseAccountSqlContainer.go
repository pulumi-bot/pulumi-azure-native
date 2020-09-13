// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB container.
//
// ## Example Usage
// ### CosmosDBSqlContainerCreateUpdate
//
// ```go
// package main
//
// import (
// 	documentdb "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/documentdb/v20150401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := documentdb.NewDatabaseAccountSqlContainer(ctx, "databaseAccountSqlContainer", &documentdb.DatabaseAccountSqlContainerArgs{
// 			AccountName:   pulumi.String("ddb1"),
// 			ContainerName: pulumi.String("containerName"),
// 			DatabaseName:  pulumi.String("databaseName"),
// 			Options:       nil,
// 			Resource: &documentdb.SqlContainerResourceArgs{
// 				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicyArgs{
// 					ConflictResolutionPath: pulumi.String("/path"),
// 					Mode:                   pulumi.String("LastWriterWins"),
// 				},
// 				DefaultTtl: pulumi.Int(100),
// 				Id:         pulumi.String("containerName"),
// 				IndexingPolicy: &documentdb.IndexingPolicyArgs{
// 					Automatic:     pulumi.Bool(true),
// 					ExcludedPaths: documentdb.ExcludedPathArray{},
// 					IncludedPaths: documentdb.IncludedPathArray{
// 						&documentdb.IncludedPathArgs{
// 							Indexes: documentdb.IndexesArray{
// 								&documentdb.IndexesArgs{
// 									DataType:  pulumi.String("String"),
// 									Kind:      pulumi.String("Range"),
// 									Precision: pulumi.Int(-1),
// 								},
// 								&documentdb.IndexesArgs{
// 									DataType:  pulumi.String("Number"),
// 									Kind:      pulumi.String("Range"),
// 									Precision: pulumi.Int(-1),
// 								},
// 							},
// 							Path: pulumi.String("/*"),
// 						},
// 					},
// 					IndexingMode: pulumi.String("Consistent"),
// 				},
// 				PartitionKey: &documentdb.ContainerPartitionKeyArgs{
// 					Kind: pulumi.String("Hash"),
// 					Paths: pulumi.StringArray{
// 						pulumi.String("/AccountNumber"),
// 					},
// 				},
// 				UniqueKeyPolicy: &documentdb.UniqueKeyPolicyArgs{
// 					UniqueKeys: documentdb.UniqueKeyArray{
// 						&documentdb.UniqueKeyArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/testPath"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type DatabaseAccountSqlContainer struct {
	pulumi.CustomResourceState

	// The conflict resolution policy for the container.
	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrOutput `pulumi:"conflictResolutionPolicy"`
	// Default time to live
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
	IndexingPolicy IndexingPolicyResponsePtrOutput `pulumi:"indexingPolicy"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey ContainerPartitionKeyResponsePtrOutput `pulumi:"partitionKey"`
	// A system generated property. A unique identifier.
	Rid pulumi.StringPtrOutput `pulumi:"rid"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A system generated property that denotes the last updated timestamp of the resource.
	Ts pulumi.MapOutput `pulumi:"ts"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
	UniqueKeyPolicy UniqueKeyPolicyResponsePtrOutput `pulumi:"uniqueKeyPolicy"`
}

// NewDatabaseAccountSqlContainer registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountSqlContainer(ctx *pulumi.Context,
	name string, args *DatabaseAccountSqlContainerArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountSqlContainer, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ContainerName == nil {
		return nil, errors.New("missing required argument 'ContainerName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
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
	if args == nil {
		args = &DatabaseAccountSqlContainerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20150408:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20151106:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20160319:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20160331:DatabaseAccountSqlContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountSqlContainer
	err := ctx.RegisterResource("azurerm:documentdb/v20150401:DatabaseAccountSqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountSqlContainer gets an existing DatabaseAccountSqlContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountSqlContainerState, opts ...pulumi.ResourceOption) (*DatabaseAccountSqlContainer, error) {
	var resource DatabaseAccountSqlContainer
	err := ctx.ReadResource("azurerm:documentdb/v20150401:DatabaseAccountSqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountSqlContainer resources.
type databaseAccountSqlContainerState struct {
	// The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	// Default time to live
	DefaultTtl *int `pulumi:"defaultTtl"`
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `pulumi:"etag"`
	// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
	IndexingPolicy *IndexingPolicyResponse `pulumi:"indexingPolicy"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey *ContainerPartitionKeyResponse `pulumi:"partitionKey"`
	// A system generated property. A unique identifier.
	Rid *string `pulumi:"rid"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// A system generated property that denotes the last updated timestamp of the resource.
	Ts map[string]interface{} `pulumi:"ts"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
	// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicyResponse `pulumi:"uniqueKeyPolicy"`
}

type DatabaseAccountSqlContainerState struct {
	// The conflict resolution policy for the container.
	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrInput
	// Default time to live
	DefaultTtl pulumi.IntPtrInput
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag pulumi.StringPtrInput
	// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
	IndexingPolicy IndexingPolicyResponsePtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey ContainerPartitionKeyResponsePtrInput
	// A system generated property. A unique identifier.
	Rid pulumi.StringPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// A system generated property that denotes the last updated timestamp of the resource.
	Ts pulumi.MapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
	// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
	UniqueKeyPolicy UniqueKeyPolicyResponsePtrInput
}

func (DatabaseAccountSqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountSqlContainerState)(nil)).Elem()
}

type databaseAccountSqlContainerArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a container
	Resource SqlContainerResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountSqlContainer resource.
type DatabaseAccountSqlContainerArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB container name.
	ContainerName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a container
	Resource SqlContainerResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountSqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountSqlContainerArgs)(nil)).Elem()
}
