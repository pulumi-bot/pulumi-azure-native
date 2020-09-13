// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Custom domain resource payload.
//
// ## Example Usage
// ### CustomDomains_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	appplatform "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/appplatform/v20190501preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appplatform.NewCustomDomain(ctx, "customDomain", &appplatform.CustomDomainArgs{
// 			AppName:           pulumi.String("myapp"),
// 			DomainName:        pulumi.String("mydomain.com"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			ServiceName:       pulumi.String("myservice"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type CustomDomain struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the custom domain resource.
	Properties CustomDomainPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil || args.AppName == nil {
		return nil, errors.New("missing required argument 'AppName'")
	}
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &CustomDomainArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:appplatform/latest:CustomDomain"),
		},
		{
			Type: pulumi.String("azurerm:appplatform/v20200701:CustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomDomain
	err := ctx.RegisterResource("azurerm:appplatform/v20190501preview:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("azurerm:appplatform/v20190501preview:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the custom domain resource.
	Properties *CustomDomainPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type CustomDomainState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the custom domain resource.
	Properties CustomDomainPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the custom domain resource.
	DomainName string `pulumi:"domainName"`
	// Properties of the custom domain resource.
	Properties *CustomDomainProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput
	// The name of the custom domain resource.
	DomainName pulumi.StringInput
	// Properties of the custom domain resource.
	Properties CustomDomainPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}
