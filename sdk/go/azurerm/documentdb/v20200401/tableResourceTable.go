// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Table.
type TableResourceTable struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an Azure Cosmos DB Table
	Properties TableGetPropertiesResponseOutput `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTableResourceTable registers a new resource with the given unique name, arguments, and options.
func NewTableResourceTable(ctx *pulumi.Context,
	name string, args *TableResourceTableArgs, opts ...pulumi.ResourceOption) (*TableResourceTable, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
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
		args = &TableResourceTableArgs{}
	}
	var resource TableResourceTable
	err := ctx.RegisterResource("azurerm:documentdb/v20200401:TableResourceTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableResourceTable gets an existing TableResourceTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableResourceTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableResourceTableState, opts ...pulumi.ResourceOption) (*TableResourceTable, error) {
	var resource TableResourceTable
	err := ctx.ReadResource("azurerm:documentdb/v20200401:TableResourceTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableResourceTable resources.
type tableResourceTableState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name *string `pulumi:"name"`
	// The properties of an Azure Cosmos DB Table
	Properties *TableGetPropertiesResponse `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *TagsResponse `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type TableResourceTableState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name pulumi.StringPtrInput
	// The properties of an Azure Cosmos DB Table
	Properties TableGetPropertiesResponsePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsResponsePtrInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (TableResourceTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableResourceTableState)(nil)).Elem()
}

type tableResourceTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// Cosmos DB table name.
	Name string `pulumi:"name"`
	// Properties to create and update Azure Cosmos DB Table.
	Properties TableCreateUpdateProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags *Tags `pulumi:"tags"`
}

// The set of arguments for constructing a TableResourceTable resource.
type TableResourceTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// Cosmos DB table name.
	Name pulumi.StringInput
	// Properties to create and update Azure Cosmos DB Table.
	Properties TableCreateUpdatePropertiesInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags TagsPtrInput
}

func (TableResourceTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableResourceTableArgs)(nil)).Elem()
}