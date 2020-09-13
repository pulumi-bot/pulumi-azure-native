// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Service resource
//
// ## Example Usage
// ### Services_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	appplatform "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/appplatform/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appplatform.NewService(ctx, "service", &appplatform.ServiceArgs{
// 			Location:          pulumi.String("eastus"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			ServiceName:       pulumi.String("myservice"),
// 			Sku: &appplatform.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
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
// ### Services_CreateOrUpdate_VNetInjection
//
// ```go
// package main
//
// import (
// 	appplatform "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/appplatform/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appplatform.NewService(ctx, "service", &appplatform.ServiceArgs{
// 			Location:          pulumi.String("eastus"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			ServiceName:       pulumi.String("myservice"),
// 			Sku: &appplatform.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
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
type Service struct {
	pulumi.CustomResourceState

	// The GEO location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Service resource
	Properties ClusterResourcePropertiesResponseOutput `pulumi:"properties"`
	// Sku of the Service resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:appplatform/v20190501preview:Service"),
		},
		{
			Type: pulumi.String("azurerm:appplatform/v20200701:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azurerm:appplatform/latest:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azurerm:appplatform/latest:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The GEO location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the Service resource
	Properties *ClusterResourcePropertiesResponse `pulumi:"properties"`
	// Sku of the Service resource
	Sku *SkuResponse `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ServiceState struct {
	// The GEO location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the Service resource
	Properties ClusterResourcePropertiesResponsePtrInput
	// Sku of the Service resource
	Sku SkuResponsePtrInput
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The GEO location of the resource.
	Location *string `pulumi:"location"`
	// Properties of the Service resource
	Properties *ClusterResourceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Sku of the Service resource
	Sku *Sku `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The GEO location of the resource.
	Location pulumi.StringPtrInput
	// Properties of the Service resource
	Properties ClusterResourcePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// Sku of the Service resource
	Sku SkuPtrInput
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
