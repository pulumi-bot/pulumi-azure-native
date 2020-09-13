// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160331

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB MongoDB database.
//
// ## Example Usage
// ### CosmosDBMongoDBDatabaseCreateUpdate
//
// ```go
// package main
//
// import (
// 	documentdb "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/documentdb/v20160331"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := documentdb.NewDatabaseAccountMongoDBDatabase(ctx, "databaseAccountMongoDBDatabase", &documentdb.DatabaseAccountMongoDBDatabaseArgs{
// 			AccountName:  pulumi.String("ddb1"),
// 			DatabaseName: pulumi.String("databaseName"),
// 			Options:      nil,
// 			Resource: &documentdb.MongoDBDatabaseResourceArgs{
// 				Id: pulumi.String("updatedDatabaseName"),
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
type DatabaseAccountMongoDBDatabase struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountMongoDBDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountMongoDBDatabase(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBDatabaseArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBDatabase, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
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
		args = &DatabaseAccountMongoDBDatabaseArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20150401:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20150408:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20151106:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20160319:DatabaseAccountMongoDBDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountMongoDBDatabase
	err := ctx.RegisterResource("azurerm:documentdb/v20160331:DatabaseAccountMongoDBDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountMongoDBDatabase gets an existing DatabaseAccountMongoDBDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountMongoDBDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountMongoDBDatabaseState, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBDatabase, error) {
	var resource DatabaseAccountMongoDBDatabase
	err := ctx.ReadResource("azurerm:documentdb/v20160331:DatabaseAccountMongoDBDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountMongoDBDatabase resources.
type databaseAccountMongoDBDatabaseState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type DatabaseAccountMongoDBDatabaseState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (DatabaseAccountMongoDBDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBDatabaseState)(nil)).Elem()
}

type databaseAccountMongoDBDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a MongoDB database
	Resource MongoDBDatabaseResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountMongoDBDatabase resource.
type DatabaseAccountMongoDBDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a MongoDB database
	Resource MongoDBDatabaseResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBDatabaseArgs)(nil)).Elem()
}
