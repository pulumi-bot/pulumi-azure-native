// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181203

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// EnterpriseKnowledgeGraph resource definition
//
// ## Example Usage
// ### Create EnterpriseKnowledgeGraph
//
// ```go
// package main
//
// import (
// 	enterpriseknowledgegraph "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/enterpriseknowledgegraph/v20181203"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := enterpriseknowledgegraph.NewEnterpriseKnowledgeGraph(ctx, "enterpriseKnowledgeGraph", &enterpriseknowledgegraph.EnterpriseKnowledgeGraphArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("OneResourceGroupName"),
// 			ResourceName:      pulumi.String("sampleekgname"),
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type EnterpriseKnowledgeGraph struct {
	pulumi.CustomResourceState

	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnterpriseKnowledgeGraph registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseKnowledgeGraph(ctx *pulumi.Context,
	name string, args *EnterpriseKnowledgeGraphArgs, opts ...pulumi.ResourceOption) (*EnterpriseKnowledgeGraph, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &EnterpriseKnowledgeGraphArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:enterpriseknowledgegraph/latest:EnterpriseKnowledgeGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource EnterpriseKnowledgeGraph
	err := ctx.RegisterResource("azurerm:enterpriseknowledgegraph/v20181203:EnterpriseKnowledgeGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseKnowledgeGraph gets an existing EnterpriseKnowledgeGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseKnowledgeGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseKnowledgeGraphState, opts ...pulumi.ResourceOption) (*EnterpriseKnowledgeGraph, error) {
	var resource EnterpriseKnowledgeGraph
	err := ctx.ReadResource("azurerm:enterpriseknowledgegraph/v20181203:EnterpriseKnowledgeGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseKnowledgeGraph resources.
type enterpriseKnowledgeGraphState struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties *EnterpriseKnowledgeGraphPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type EnterpriseKnowledgeGraphState struct {
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesResponsePtrInput
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (EnterpriseKnowledgeGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseKnowledgeGraphState)(nil)).Elem()
}

type enterpriseKnowledgeGraphArgs struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties *EnterpriseKnowledgeGraphProperties `pulumi:"properties"`
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EnterpriseKnowledgeGraph resource.
type EnterpriseKnowledgeGraphArgs struct {
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesPtrInput
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName pulumi.StringInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (EnterpriseKnowledgeGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseKnowledgeGraphArgs)(nil)).Elem()
}
