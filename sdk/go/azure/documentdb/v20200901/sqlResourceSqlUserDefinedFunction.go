// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB userDefinedFunction.
type SqlResourceSqlUserDefinedFunction struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                          `pulumi:"name"`
	Resource SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlUserDefinedFunction registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context,
	name string, args *SqlResourceSqlUserDefinedFunctionArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlUserDefinedFunction, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ContainerName == nil {
		return nil, errors.New("missing required argument 'ContainerName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.Resource == nil {
		return nil, errors.New("missing required argument 'Resource'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserDefinedFunctionName == nil {
		return nil, errors.New("missing required argument 'UserDefinedFunctionName'")
	}
	if args == nil {
		args = &SqlResourceSqlUserDefinedFunctionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/latest:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:SqlResourceSqlUserDefinedFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlUserDefinedFunction
	err := ctx.RegisterResource("azure-nextgen:documentdb/v20200901:SqlResourceSqlUserDefinedFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlUserDefinedFunction gets an existing SqlResourceSqlUserDefinedFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlUserDefinedFunctionState, opts ...pulumi.ResourceOption) (*SqlResourceSqlUserDefinedFunction, error) {
	var resource SqlResourceSqlUserDefinedFunction
	err := ctx.ReadResource("azure-nextgen:documentdb/v20200901:SqlResourceSqlUserDefinedFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlUserDefinedFunction resources.
type sqlResourceSqlUserDefinedFunctionState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     *string                                              `pulumi:"name"`
	Resource *SqlUserDefinedFunctionGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type SqlResourceSqlUserDefinedFunctionState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name     pulumi.StringPtrInput
	Resource SqlUserDefinedFunctionGetPropertiesResponseResourcePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (SqlResourceSqlUserDefinedFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlUserDefinedFunctionState)(nil)).Elem()
}

type sqlResourceSqlUserDefinedFunctionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a userDefinedFunction
	Resource SqlUserDefinedFunctionResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// Cosmos DB userDefinedFunction name.
	UserDefinedFunctionName string `pulumi:"userDefinedFunctionName"`
}

// The set of arguments for constructing a SqlResourceSqlUserDefinedFunction resource.
type SqlResourceSqlUserDefinedFunctionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB container name.
	ContainerName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a userDefinedFunction
	Resource SqlUserDefinedFunctionResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// Cosmos DB userDefinedFunction name.
	UserDefinedFunctionName pulumi.StringInput
}

func (SqlResourceSqlUserDefinedFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlUserDefinedFunctionArgs)(nil)).Elem()
}

type SqlResourceSqlUserDefinedFunctionInput interface {
	pulumi.Input

	ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput
	ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput
}

func (SqlResourceSqlUserDefinedFunction) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlUserDefinedFunction)(nil)).Elem()
}

func (i SqlResourceSqlUserDefinedFunction) ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput {
	return i.ToSqlResourceSqlUserDefinedFunctionOutputWithContext(context.Background())
}

func (i SqlResourceSqlUserDefinedFunction) ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlUserDefinedFunctionOutput)
}

type SqlResourceSqlUserDefinedFunctionOutput struct {
	*pulumi.OutputState
}

func (SqlResourceSqlUserDefinedFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlUserDefinedFunctionOutput)(nil)).Elem()
}

func (o SqlResourceSqlUserDefinedFunctionOutput) ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput {
	return o
}

func (o SqlResourceSqlUserDefinedFunctionOutput) ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlUserDefinedFunctionOutput{})
}
