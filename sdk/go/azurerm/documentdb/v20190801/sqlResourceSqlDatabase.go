// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB SQL database.
type SqlResourceSqlDatabase struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an Azure Cosmos DB SQL database
	Properties SqlDatabaseGetPropertiesResponseOutput `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlDatabase registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlDatabase(ctx *pulumi.Context,
	name string, args *SqlResourceSqlDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlDatabase, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
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
		args = &SqlResourceSqlDatabaseArgs{}
	}
	var resource SqlResourceSqlDatabase
	err := ctx.RegisterResource("azurerm:documentdb/v20190801:SqlResourceSqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlDatabase gets an existing SqlResourceSqlDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlDatabaseState, opts ...pulumi.ResourceOption) (*SqlResourceSqlDatabase, error) {
	var resource SqlResourceSqlDatabase
	err := ctx.ReadResource("azurerm:documentdb/v20190801:SqlResourceSqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlDatabase resources.
type sqlResourceSqlDatabaseState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name *string `pulumi:"name"`
	// The properties of an Azure Cosmos DB SQL database
	Properties *SqlDatabaseGetPropertiesResponse `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *TagsResponse `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type SqlResourceSqlDatabaseState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name pulumi.StringPtrInput
	// The properties of an Azure Cosmos DB SQL database
	Properties SqlDatabaseGetPropertiesResponsePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (SqlResourceSqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlDatabaseState)(nil)).Elem()
}

type sqlResourceSqlDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// Cosmos DB database name.
	Name string `pulumi:"name"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a SQL database
	Resource SqlDatabaseResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *Tags `pulumi:"tags"`
}

// The set of arguments for constructing a SqlResourceSqlDatabase resource.
type SqlResourceSqlDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// Cosmos DB database name.
	Name pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsInput
	// The standard JSON format of a SQL database
	Resource SqlDatabaseResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsPtrInput
}

func (SqlResourceSqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlDatabaseArgs)(nil)).Elem()
}
