// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Gremlin database.
type DatabaseAccountGremlinDatabase struct {
	pulumi.CustomResourceState

	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// A system generated property. A unique identifier.
	Rid pulumi.StringPtrOutput `pulumi:"rid"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A system generated property that denotes the last updated timestamp of the resource.
	Ts pulumi.AnyOutput `pulumi:"ts"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountGremlinDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountGremlinDatabase(ctx *pulumi.Context,
	name string, args *DatabaseAccountGremlinDatabaseArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinDatabase, error) {
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
		args = &DatabaseAccountGremlinDatabaseArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:DatabaseAccountGremlinDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountGremlinDatabase
	err := ctx.RegisterResource("azure-nextgen:documentdb/latest:DatabaseAccountGremlinDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountGremlinDatabase gets an existing DatabaseAccountGremlinDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountGremlinDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountGremlinDatabaseState, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinDatabase, error) {
	var resource DatabaseAccountGremlinDatabase
	err := ctx.ReadResource("azure-nextgen:documentdb/latest:DatabaseAccountGremlinDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountGremlinDatabase resources.
type databaseAccountGremlinDatabaseState struct {
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `pulumi:"etag"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// A system generated property. A unique identifier.
	Rid *string `pulumi:"rid"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// A system generated property that denotes the last updated timestamp of the resource.
	Ts interface{} `pulumi:"ts"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type DatabaseAccountGremlinDatabaseState struct {
	// A system generated property representing the resource etag required for optimistic concurrency control.
	Etag pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// A system generated property. A unique identifier.
	Rid pulumi.StringPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// A system generated property that denotes the last updated timestamp of the resource.
	Ts pulumi.Input
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (DatabaseAccountGremlinDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinDatabaseState)(nil)).Elem()
}

type databaseAccountGremlinDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a Gremlin database
	Resource GremlinDatabaseResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountGremlinDatabase resource.
type DatabaseAccountGremlinDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a Gremlin database
	Resource GremlinDatabaseResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountGremlinDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinDatabaseArgs)(nil)).Elem()
}

type DatabaseAccountGremlinDatabaseInput interface {
	pulumi.Input

	ToDatabaseAccountGremlinDatabaseOutput() DatabaseAccountGremlinDatabaseOutput
	ToDatabaseAccountGremlinDatabaseOutputWithContext(ctx context.Context) DatabaseAccountGremlinDatabaseOutput
}

func (DatabaseAccountGremlinDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountGremlinDatabase)(nil)).Elem()
}

func (i DatabaseAccountGremlinDatabase) ToDatabaseAccountGremlinDatabaseOutput() DatabaseAccountGremlinDatabaseOutput {
	return i.ToDatabaseAccountGremlinDatabaseOutputWithContext(context.Background())
}

func (i DatabaseAccountGremlinDatabase) ToDatabaseAccountGremlinDatabaseOutputWithContext(ctx context.Context) DatabaseAccountGremlinDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountGremlinDatabaseOutput)
}

type DatabaseAccountGremlinDatabaseOutput struct {
	*pulumi.OutputState
}

func (DatabaseAccountGremlinDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountGremlinDatabaseOutput)(nil)).Elem()
}

func (o DatabaseAccountGremlinDatabaseOutput) ToDatabaseAccountGremlinDatabaseOutput() DatabaseAccountGremlinDatabaseOutput {
	return o
}

func (o DatabaseAccountGremlinDatabaseOutput) ToDatabaseAccountGremlinDatabaseOutputWithContext(ctx context.Context) DatabaseAccountGremlinDatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountGremlinDatabaseOutput{})
}
