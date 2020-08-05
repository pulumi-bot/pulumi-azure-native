// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151106

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Cassandra table.
type DatabaseAccountCassandraTable struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an Azure Cosmos DB Cassandra table
	Properties CassandraTablePropertiesResponseOutput `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountCassandraTable registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountCassandraTable(ctx *pulumi.Context,
	name string, args *DatabaseAccountCassandraTableArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraTable, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.KeyspaceName == nil {
		return nil, errors.New("missing required argument 'KeyspaceName'")
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
		args = &DatabaseAccountCassandraTableArgs{}
	}
	var resource DatabaseAccountCassandraTable
	err := ctx.RegisterResource("azurerm:documentdb/v20151106:DatabaseAccountCassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountCassandraTable gets an existing DatabaseAccountCassandraTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountCassandraTableState, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraTable, error) {
	var resource DatabaseAccountCassandraTable
	err := ctx.ReadResource("azurerm:documentdb/v20151106:DatabaseAccountCassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountCassandraTable resources.
type databaseAccountCassandraTableState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// The properties of an Azure Cosmos DB Cassandra table
	Properties *CassandraTablePropertiesResponse `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *TagsResponse `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type DatabaseAccountCassandraTableState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// The properties of an Azure Cosmos DB Cassandra table
	Properties CassandraTablePropertiesResponsePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (DatabaseAccountCassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraTableState)(nil)).Elem()
}

type databaseAccountCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName string `pulumi:"keyspaceName"`
	// Cosmos DB table name.
	Name string `pulumi:"name"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountCassandraTable resource.
type DatabaseAccountCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringInput
	// Cosmos DB table name.
	Name pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsInput
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraTableArgs)(nil)).Elem()
}
