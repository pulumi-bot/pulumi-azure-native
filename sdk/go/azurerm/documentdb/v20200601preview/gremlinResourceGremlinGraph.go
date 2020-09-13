// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB Gremlin graph.
//
// ## Example Usage
// ### CosmosDBGremlinGraphCreateUpdate
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
// 		_, err := documentdb.NewGremlinResourceGremlinGraph(ctx, "gremlinResourceGremlinGraph", &documentdb.GremlinResourceGremlinGraphArgs{
// 			AccountName:  pulumi.String("ddb1"),
// 			DatabaseName: pulumi.String("databaseName"),
// 			GraphName:    pulumi.String("graphName"),
// 			Location:     pulumi.String("West US"),
// 			Options:      nil,
// 			Resource: &documentdb.GremlinGraphResourceArgs{
// 				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicyArgs{
// 					ConflictResolutionPath: pulumi.String("/path"),
// 					Mode:                   pulumi.String("LastWriterWins"),
// 				},
// 				DefaultTtl: pulumi.Int(100),
// 				Id:         pulumi.String("graphName"),
// 				IndexingPolicy: &documentdb.IndexingPolicyArgs{
// 					Automatic:     pulumi.Bool(true),
// 					ExcludedPaths: documentdb.ExcludedPathArray{},
// 					IncludedPaths: documentdb.IncludedPathArray{
// 						&documentdb.IncludedPathArgs{
// 							Indexes: documentdb.IndexesArray{
// 								&documentdb.IndexesArgs{
// 									DataType:  pulumi.String("String"),
// 									Kind:      pulumi.String("Range"),
// 									Precision: pulumi.Int(-1),
// 								},
// 								&documentdb.IndexesArgs{
// 									DataType:  pulumi.String("Number"),
// 									Kind:      pulumi.String("Range"),
// 									Precision: pulumi.Int(-1),
// 								},
// 							},
// 							Path: pulumi.String("/*"),
// 						},
// 					},
// 					IndexingMode: pulumi.String("Consistent"),
// 				},
// 				PartitionKey: &documentdb.ContainerPartitionKeyArgs{
// 					Kind: pulumi.String("Hash"),
// 					Paths: pulumi.StringArray{
// 						pulumi.String("/AccountNumber"),
// 					},
// 				},
// 				UniqueKeyPolicy: &documentdb.UniqueKeyPolicyArgs{
// 					UniqueKeys: documentdb.UniqueKeyArray{
// 						&documentdb.UniqueKeyArgs{
// 							Paths: pulumi.StringArray{
// 								pulumi.String("/testPath"),
// 							},
// 						},
// 					},
// 				},
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
type GremlinResourceGremlinGraph struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                `pulumi:"name"`
	Options  GremlinGraphGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource GremlinGraphGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGremlinResourceGremlinGraph registers a new resource with the given unique name, arguments, and options.
func NewGremlinResourceGremlinGraph(ctx *pulumi.Context,
	name string, args *GremlinResourceGremlinGraphArgs, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinGraph, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.GraphName == nil {
		return nil, errors.New("missing required argument 'GraphName'")
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
		args = &GremlinResourceGremlinGraphArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:documentdb/latest:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20190801:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20191212:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200301:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azurerm:documentdb/v20200401:GremlinResourceGremlinGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource GremlinResourceGremlinGraph
	err := ctx.RegisterResource("azurerm:documentdb/v20200601preview:GremlinResourceGremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGremlinResourceGremlinGraph gets an existing GremlinResourceGremlinGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGremlinResourceGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinResourceGremlinGraphState, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinGraph, error) {
	var resource GremlinResourceGremlinGraph
	err := ctx.ReadResource("azurerm:documentdb/v20200601preview:GremlinResourceGremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GremlinResourceGremlinGraph resources.
type gremlinResourceGremlinGraphState struct {
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     *string                                    `pulumi:"name"`
	Options  *GremlinGraphGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinGraphGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type GremlinResourceGremlinGraphState struct {
	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name     pulumi.StringPtrInput
	Options  GremlinGraphGetPropertiesResponseOptionsPtrInput
	Resource GremlinGraphGetPropertiesResponseResourcePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (GremlinResourceGremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinGraphState)(nil)).Elem()
}

type gremlinResourceGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Cosmos DB graph name.
	GraphName string `pulumi:"graphName"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Gremlin graph
	Resource GremlinGraphResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GremlinResourceGremlinGraph resource.
type GremlinResourceGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// Cosmos DB graph name.
	GraphName pulumi.StringInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsInput
	// The standard JSON format of a Gremlin graph
	Resource GremlinGraphResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (GremlinResourceGremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinGraphArgs)(nil)).Elem()
}
