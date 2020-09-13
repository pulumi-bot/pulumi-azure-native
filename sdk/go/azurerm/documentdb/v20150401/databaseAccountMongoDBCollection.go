// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB MongoDB collection.
//
// ## Example Usage
// ### CosmosDBMongoDBCollectionCreateUpdate
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
// 		_, err := documentdb.NewDatabaseAccountMongoDBCollection(ctx, "databaseAccountMongoDBCollection", &documentdb.DatabaseAccountMongoDBCollectionArgs{
// 			AccountName:    pulumi.String("ddb1"),
// 			CollectionName: pulumi.String("collectionName"),
// 			DatabaseName:   pulumi.String("databaseName"),
// 			Options:        nil,
// 			Resource: &documentdb.MongoDBCollectionResourceArgs{
// 				Id: pulumi.String("testcoll"),
// 				Indexes: documentdb.MongoIndexArray{
// 					&documentdb.MongoIndexArgs{
// 						Key: &documentdb.MongoIndexKeysArgs{
// 							Keys: pulumi.StringArray{
// 								pulumi.String("testKey"),
// 							},
// 						},
// 						Options: &documentdb.MongoIndexOptionsArgs{
// 							ExpireAfterSeconds: pulumi.Int(100),
// 							Unique:             pulumi.Bool(true),
// 						},
// 					},
// 				},
// 				ShardKey: pulumi.StringMap{
// 					"testKey": pulumi.String("Hash"),
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
type DatabaseAccountMongoDBCollection struct {
	pulumi.CustomResourceState

	// List of index keys
	Indexes MongoIndexResponseArrayOutput `pulumi:"indexes"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// A key-value pair of shard keys to be applied for the request.
	ShardKey pulumi.StringMapOutput `pulumi:"shardKey"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountMongoDBCollection registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.CollectionName == nil {
		return nil, errors.New("missing required argument 'CollectionName'")
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
		args = &DatabaseAccountMongoDBCollectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20150408:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20151106:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20160319:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20160331:DatabaseAccountMongoDBCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountMongoDBCollection
	err := ctx.RegisterResource("azurerm:documentdb/v20150401:DatabaseAccountMongoDBCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountMongoDBCollection gets an existing DatabaseAccountMongoDBCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountMongoDBCollectionState, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	var resource DatabaseAccountMongoDBCollection
	err := ctx.ReadResource("azurerm:documentdb/v20150401:DatabaseAccountMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountMongoDBCollection resources.
type databaseAccountMongoDBCollectionState struct {
	// List of index keys
	Indexes []MongoIndexResponse `pulumi:"indexes"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// A key-value pair of shard keys to be applied for the request.
	ShardKey map[string]string `pulumi:"shardKey"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type DatabaseAccountMongoDBCollectionState struct {
	// List of index keys
	Indexes MongoIndexResponseArrayInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// A key-value pair of shard keys to be applied for the request.
	ShardKey pulumi.StringMapInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (DatabaseAccountMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionState)(nil)).Elem()
}

type databaseAccountMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB collection name.
	CollectionName string `pulumi:"collectionName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountMongoDBCollection resource.
type DatabaseAccountMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB collection name.
	CollectionName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionArgs)(nil)).Elem()
}
