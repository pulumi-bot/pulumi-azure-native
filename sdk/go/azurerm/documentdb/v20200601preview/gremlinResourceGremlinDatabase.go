// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Gremlin database.
//
// ## Example Usage
// ### CosmosDBGremlinDatabaseCreateUpdate
//
// ```go
// package main
//
// import (
// 	documentdb "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/documentdb/v20200601preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := documentdb.NewGremlinResourceGremlinDatabase(ctx, "gremlinResourceGremlinDatabase", &documentdb.GremlinResourceGremlinDatabaseArgs{
// 			AccountName:  pulumi.String("ddb1"),
// 			DatabaseName: pulumi.String("databaseName"),
// 			Location:     pulumi.String("West US"),
// 			Options:      nil,
// 			Resource: &documentdb.GremlinDatabaseResourceArgs{
// 				Id: pulumi.String("databaseName"),
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
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
type GremlinResourceGremlinDatabase struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                   `pulumi:"name"`
	Options  GremlinDatabaseGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource GremlinDatabaseGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGremlinResourceGremlinDatabase registers a new resource with the given unique name, arguments, and options.
func NewGremlinResourceGremlinDatabase(ctx *pulumi.Context,
	name string, args *GremlinResourceGremlinDatabaseArgs, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinDatabase, error) {
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
		args = &GremlinResourceGremlinDatabaseArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20190801:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20191212:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200301:GremlinResourceGremlinDatabase"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200401:GremlinResourceGremlinDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource GremlinResourceGremlinDatabase
	err := ctx.RegisterResource("azurerm:documentdb/v20200601preview:GremlinResourceGremlinDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGremlinResourceGremlinDatabase gets an existing GremlinResourceGremlinDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGremlinResourceGremlinDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinResourceGremlinDatabaseState, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinDatabase, error) {
	var resource GremlinResourceGremlinDatabase
	err := ctx.ReadResource("azurerm:documentdb/v20200601preview:GremlinResourceGremlinDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GremlinResourceGremlinDatabase resources.
type gremlinResourceGremlinDatabaseState struct {
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     *string                                       `pulumi:"name"`
	Options  *GremlinDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type GremlinResourceGremlinDatabaseState struct {
	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name     pulumi.StringPtrInput
	Options  GremlinDatabaseGetPropertiesResponseOptionsPtrInput
	Resource GremlinDatabaseGetPropertiesResponseResourcePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (GremlinResourceGremlinDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinDatabaseState)(nil)).Elem()
}

type gremlinResourceGremlinDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Gremlin database
	Resource GremlinDatabaseResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GremlinResourceGremlinDatabase resource.
type GremlinResourceGremlinDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsInput
	// The standard JSON format of a Gremlin database
	Resource GremlinDatabaseResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (GremlinResourceGremlinDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinDatabaseArgs)(nil)).Elem()
}
