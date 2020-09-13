// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration service environment.
//
// ## Example Usage
// ### Create or update an integration service environment
//
// ```go
// package main
//
// import (
// 	logic "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/logic/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logic.NewIntegrationServiceEnvironment(ctx, "integrationServiceEnvironment", &logic.IntegrationServiceEnvironmentArgs{
// 			IntegrationServiceEnvironmentName: pulumi.String("testIntegrationServiceEnvironment"),
// 			Location:                          pulumi.String("brazilsouth"),
// 			ResourceGroup:                     pulumi.String("testResourceGroup"),
// 			Sku: &logic.IntegrationServiceEnvironmentSkuArgs{
// 				Capacity: pulumi.Int(2),
// 				Name:     pulumi.String("Premium"),
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
type IntegrationServiceEnvironment struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The integration service environment properties.
	Properties IntegrationServiceEnvironmentPropertiesResponseOutput `pulumi:"properties"`
	// The sku.
	Sku IntegrationServiceEnvironmentSkuResponsePtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationServiceEnvironment registers a new resource with the given unique name, arguments, and options.
func NewIntegrationServiceEnvironment(ctx *pulumi.Context,
	name string, args *IntegrationServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironment, error) {
	if args == nil || args.IntegrationServiceEnvironmentName == nil {
		return nil, errors.New("missing required argument 'IntegrationServiceEnvironmentName'")
	}
	if args == nil || args.ResourceGroup == nil {
		return nil, errors.New("missing required argument 'ResourceGroup'")
	}
	if args == nil {
		args = &IntegrationServiceEnvironmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/v20190501:IntegrationServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationServiceEnvironment
	err := ctx.RegisterResource("azurerm:logic/latest:IntegrationServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationServiceEnvironment gets an existing IntegrationServiceEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationServiceEnvironmentState, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironment, error) {
	var resource IntegrationServiceEnvironment
	err := ctx.ReadResource("azurerm:logic/latest:IntegrationServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationServiceEnvironment resources.
type integrationServiceEnvironmentState struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The integration service environment properties.
	Properties *IntegrationServiceEnvironmentPropertiesResponse `pulumi:"properties"`
	// The sku.
	Sku *IntegrationServiceEnvironmentSkuResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type IntegrationServiceEnvironmentState struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The integration service environment properties.
	Properties IntegrationServiceEnvironmentPropertiesResponsePtrInput
	// The sku.
	Sku IntegrationServiceEnvironmentSkuResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentState)(nil)).Elem()
}

type integrationServiceEnvironmentArgs struct {
	// The integration service environment name.
	IntegrationServiceEnvironmentName string `pulumi:"integrationServiceEnvironmentName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration service environment properties.
	Properties *IntegrationServiceEnvironmentProperties `pulumi:"properties"`
	// The resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The sku.
	Sku *IntegrationServiceEnvironmentSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationServiceEnvironment resource.
type IntegrationServiceEnvironmentArgs struct {
	// The integration service environment name.
	IntegrationServiceEnvironmentName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration service environment properties.
	Properties IntegrationServiceEnvironmentPropertiesPtrInput
	// The resource group.
	ResourceGroup pulumi.StringInput
	// The sku.
	Sku IntegrationServiceEnvironmentSkuPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentArgs)(nil)).Elem()
}
