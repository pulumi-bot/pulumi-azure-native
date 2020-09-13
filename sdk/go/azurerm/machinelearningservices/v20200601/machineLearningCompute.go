// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Machine Learning compute object wrapped into ARM resource envelope.
//
// ## Example Usage
// ### Create AKS Compute
//
// ```go
// package main
//
// import (
// 	machinelearningservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/machinelearningservices/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := machinelearningservices.NewMachineLearningCompute(ctx, "machineLearningCompute", &machinelearningservices.MachineLearningComputeArgs{
// 			ComputeName:       pulumi.String("compute123"),
// 			Location:          pulumi.String("eastus"),
// 			ResourceGroupName: pulumi.String("testrg123"),
// 			WorkspaceName:     pulumi.String("workspaces123"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a AML Compute
//
// ```go
//
// ```
// ### Create a ComputeInstance Compute
//
// ```go
//
// ```
// ### Create a ComputeInstance Compute with minimal inputs
//
// ```go
//
// ```
// ### Create a DataFactory Compute
//
// ```go
// package main
//
// import (
// 	machinelearningservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/machinelearningservices/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := machinelearningservices.NewMachineLearningCompute(ctx, "machineLearningCompute", &machinelearningservices.MachineLearningComputeArgs{
// 			ComputeName:       pulumi.String("compute123"),
// 			Location:          pulumi.String("eastus"),
// 			ResourceGroupName: pulumi.String("testrg123"),
// 			WorkspaceName:     pulumi.String("workspaces123"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Update a AKS Compute
//
// ```go
//
// ```
// ### Update a AML Compute
//
// ```go
//
// ```
type MachineLearningCompute struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Compute properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The sku of the workspace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachineLearningCompute registers a new resource with the given unique name, arguments, and options.
func NewMachineLearningCompute(ctx *pulumi.Context,
	name string, args *MachineLearningComputeArgs, opts ...pulumi.ResourceOption) (*MachineLearningCompute, error) {
	if args == nil || args.ComputeName == nil {
		return nil, errors.New("missing required argument 'ComputeName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &MachineLearningComputeArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:machinelearningservices/latest:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20180301preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20181119:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20190501:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20190601:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20191101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200101:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200218preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200301:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200401:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200501preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200515preview:MachineLearningCompute"),
		},
		{
			Type: pulumi.String("azurerm:machinelearningservices/v20200901preview:MachineLearningCompute"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineLearningCompute
	err := ctx.RegisterResource("azurerm:machinelearningservices/v20200601:MachineLearningCompute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineLearningCompute gets an existing MachineLearningCompute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineLearningCompute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningComputeState, opts ...pulumi.ResourceOption) (*MachineLearningCompute, error) {
	var resource MachineLearningCompute
	err := ctx.ReadResource("azurerm:machinelearningservices/v20200601:MachineLearningCompute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineLearningCompute resources.
type machineLearningComputeState struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// Compute properties
	Properties interface{} `pulumi:"properties"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type MachineLearningComputeState struct {
	// The identity of the resource.
	Identity IdentityResponsePtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// Compute properties
	Properties pulumi.Input
	// The sku of the workspace.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (MachineLearningComputeState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningComputeState)(nil)).Elem()
}

type machineLearningComputeArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName string `pulumi:"computeName"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Compute properties
	Properties interface{} `pulumi:"properties"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the workspace.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MachineLearningCompute resource.
type MachineLearningComputeArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName pulumi.StringInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Compute properties
	Properties pulumi.Input
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// The sku of the workspace.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (MachineLearningComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningComputeArgs)(nil)).Elem()
}
