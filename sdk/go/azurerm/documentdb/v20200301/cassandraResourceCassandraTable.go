// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Cassandra table.
type CassandraResourceCassandraTable struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an Azure Cosmos DB Cassandra table
	Properties CassandraTableGetPropertiesResponseOutput `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrOutput `pulumi:"tags"`
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
		args = &CassandraResourceCassandraTableArgs{}
	}
	var resource CassandraResourceCassandraTable
	err := ctx.RegisterResource("azurerm:documentdb/v20200301:CassandraResourceCassandraTable", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:documentdb/v20200301:CassandraResourceCassandraTable", name, id, state, &resource, opts...)
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
	Name *string `pulumi:"name"`
	// The properties of an Azure Cosmos DB Cassandra table
	Properties *CassandraTableGetPropertiesResponse `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *TagsResponse `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type CassandraResourceCassandraTableState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name pulumi.StringPtrInput
	// The properties of an Azure Cosmos DB Cassandra table
	Properties CassandraTableGetPropertiesResponsePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrInput
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
	// Cosmos DB table name.
	Name string `pulumi:"name"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *Tags `pulumi:"tags"`
}

// The set of arguments for constructing a CassandraResourceCassandraTable resource.
type CassandraResourceCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// Cosmos DB table name.
	Name pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsInput
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsPtrInput
}

func (CassandraResourceCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableArgs)(nil)).Elem()
}
