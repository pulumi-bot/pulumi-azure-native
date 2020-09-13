// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Dataset resource type.
//
// ## Example Usage
// ### Datasets_Create
//
// ```go
// package main
//
// import (
// 	datafactory "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/datafactory/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datafactory.NewDataset(ctx, "dataset", &datafactory.DatasetArgs{
// 			DatasetName:       pulumi.String("exampleDataset"),
// 			FactoryName:       pulumi.String("exampleFactoryName"),
// 			ResourceGroupName: pulumi.String("exampleResourceGroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Datasets_Update
//
// ```go
// package main
//
// import (
// 	datafactory "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/datafactory/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datafactory.NewDataset(ctx, "dataset", &datafactory.DatasetArgs{
// 			DatasetName:       pulumi.String("exampleDataset"),
// 			FactoryName:       pulumi.String("exampleFactoryName"),
// 			ResourceGroupName: pulumi.String("exampleResourceGroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Dataset struct {
	pulumi.CustomResourceState

	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Dataset properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil || args.DatasetName == nil {
		return nil, errors.New("missing required argument 'DatasetName'")
	}
	if args == nil || args.FactoryName == nil {
		return nil, errors.New("missing required argument 'FactoryName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DatasetArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datafactory/v20170901preview:Dataset"),
		},
		{
			Type: pulumi.String("azurerm:datafactory/v20180601:Dataset"),
		},
	})
	opts = append(opts, aliases)
	var resource Dataset
	err := ctx.RegisterResource("azurerm:datafactory/latest:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("azurerm:datafactory/latest:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
	// Etag identifies change in the resource.
	Etag *string `pulumi:"etag"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Dataset properties.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type DatasetState struct {
	// Etag identifies change in the resource.
	Etag pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Dataset properties.
	Properties pulumi.Input
	// The resource type.
	Type pulumi.StringPtrInput
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// The dataset name.
	DatasetName string `pulumi:"datasetName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Dataset properties.
	Properties interface{} `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// The dataset name.
	DatasetName pulumi.StringInput
	// The factory name.
	FactoryName pulumi.StringInput
	// Dataset properties.
	Properties pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}
