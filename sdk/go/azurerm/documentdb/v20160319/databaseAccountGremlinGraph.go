// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160319

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Gremlin graph.
//
// ## Example Usage
// ### CosmosDBGremlinGraphCreateUpdate
//
// ```go
// package main
//
// import (
// 	documentdb "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/documentdb/v20160319"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := documentdb.NewDatabaseAccountGremlinGraph(ctx, "databaseAccountGremlinGraph", &documentdb.DatabaseAccountGremlinGraphArgs{
// 			AccountName:  pulumi.String("ddb1"),
// 			DatabaseName: pulumi.String("databaseName"),
// 			GraphName:    pulumi.String("graphName"),
// 			Options:      nil,
// 			Resource: &documentdb.GremlinGraphResourceArgs{
// 				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicyArgs{
// 					ConflictResolutionPath: pulumi.String("/path"),
// 					Mode:                   pulumi.String("LastWriterWins"),
// 				},
// 				DefaultTtl: pulumi.Int(100),
// 				Id:         pulumi.String("graphName"),
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
type DatabaseAccountGremlinGraph struct {
	pulumi.CustomResourceState

	// The conflict resolution policy for the graph.
	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrOutput `pulumi:"conflictResolutionPolicy"`
	// Default time to live
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the graph
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

// NewDatabaseAccountGremlinGraph registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountGremlinGraph(ctx *pulumi.Context,
	name string, args *DatabaseAccountGremlinGraphArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinGraph, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.GraphName == nil {
		return nil, errors.New("missing required argument 'GraphName'")
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
		args = &DatabaseAccountGremlinGraphArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20150401:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20150408:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20151106:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20160331:DatabaseAccountGremlinGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountGremlinGraph
	err := ctx.RegisterResource("azurerm:documentdb/v20160319:DatabaseAccountGremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountGremlinGraph gets an existing DatabaseAccountGremlinGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountGremlinGraphState, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinGraph, error) {
	var resource DatabaseAccountGremlinGraph
	err := ctx.ReadResource("azurerm:documentdb/v20160319:DatabaseAccountGremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountGremlinGraph resources.
type databaseAccountGremlinGraphState struct {
	// The conflict resolution policy for the graph.
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	// Default time to live
	DefaultTtl *int `pulumi:"defaultTtl"`
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `pulumi:"etag"`
	// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the graph
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

type DatabaseAccountGremlinGraphState struct {
	// The conflict resolution policy for the graph.
	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrInput
	// Default time to live
	DefaultTtl pulumi.IntPtrInput
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag pulumi.StringPtrInput
	// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the graph
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

func (DatabaseAccountGremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinGraphState)(nil)).Elem()
}

type databaseAccountGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Cosmos DB graph name.
	GraphName string `pulumi:"graphName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a Gremlin graph
	Resource GremlinGraphResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountGremlinGraph resource.
type DatabaseAccountGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// Cosmos DB graph name.
	GraphName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a Gremlin graph
	Resource GremlinGraphResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountGremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinGraphArgs)(nil)).Elem()
}
