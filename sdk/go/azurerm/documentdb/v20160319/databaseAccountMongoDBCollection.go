// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160319

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB MongoDB collection.
type DatabaseAccountMongoDBCollection struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an Azure Cosmos DB MongoDB collection
	Properties MongoDBCollectionPropertiesResponseOutput `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountMongoDBCollection registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DatabaseAccountMongoDBCollectionArgs{}
	}
	var resource DatabaseAccountMongoDBCollection
	err := ctx.RegisterResource("azurerm:documentdb/v20160319:DatabaseAccountMongoDBCollection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:documentdb/v20160319:DatabaseAccountMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountMongoDBCollection resources.
type databaseAccountMongoDBCollectionState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// The properties of an Azure Cosmos DB MongoDB collection
	Properties *MongoDBCollectionPropertiesResponse `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *TagsResponse `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type DatabaseAccountMongoDBCollectionState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// The properties of an Azure Cosmos DB MongoDB collection
	Properties MongoDBCollectionPropertiesResponsePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (DatabaseAccountMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionState)(nil)).Elem()
}

type databaseAccountMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Cosmos DB collection name.
	Name string `pulumi:"name"`
	// Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties MongoDBCollectionCreateUpdateProperties `pulumi:"properties"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountMongoDBCollection resource.
type DatabaseAccountMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// Cosmos DB collection name.
	Name pulumi.StringInput
	// Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties MongoDBCollectionCreateUpdatePropertiesInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionArgs)(nil)).Elem()
}