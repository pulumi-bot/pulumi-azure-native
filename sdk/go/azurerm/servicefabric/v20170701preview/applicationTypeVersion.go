// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An application type version resource for the specified application type name resource.
//
// ## Example Usage
// ### Put an application type version
//
// ```go
// package main
//
// import (
// 	servicefabric "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/servicefabric/v20170701preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicefabric.NewApplicationTypeVersion(ctx, "applicationTypeVersion", &servicefabric.ApplicationTypeVersionArgs{
// 			AppPackageUrl:       pulumi.String("http://fakelink.test.com/MyAppType"),
// 			ApplicationTypeName: pulumi.String("myAppType"),
// 			ClusterName:         pulumi.String("myCluster"),
// 			Location:            pulumi.String("eastus"),
// 			ResourceGroupName:   pulumi.String("resRg"),
// 			Version:             pulumi.String("1.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApplicationTypeVersion struct {
	pulumi.CustomResourceState

	// The URL to the application package
	AppPackageUrl pulumi.StringOutput `pulumi:"appPackageUrl"`
	// List of application type parameters that can be overridden when creating or updating the application.
	DefaultParameterList pulumi.StringMapOutput `pulumi:"defaultParameterList"`
	// Azure resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationTypeVersion registers a new resource with the given unique name, arguments, and options.
func NewApplicationTypeVersion(ctx *pulumi.Context,
	name string, args *ApplicationTypeVersionArgs, opts ...pulumi.ResourceOption) (*ApplicationTypeVersion, error) {
	if args == nil || args.AppPackageUrl == nil {
		return nil, errors.New("missing required argument 'AppPackageUrl'")
	}
	if args == nil || args.ApplicationTypeName == nil {
		return nil, errors.New("missing required argument 'ApplicationTypeName'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	if args == nil {
		args = &ApplicationTypeVersionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:servicefabric/latest:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190601preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20191101preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20200301:ApplicationTypeVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationTypeVersion
	err := ctx.RegisterResource("azurerm:servicefabric/v20170701preview:ApplicationTypeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationTypeVersion gets an existing ApplicationTypeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationTypeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationTypeVersionState, opts ...pulumi.ResourceOption) (*ApplicationTypeVersion, error) {
	var resource ApplicationTypeVersion
	err := ctx.ReadResource("azurerm:servicefabric/v20170701preview:ApplicationTypeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationTypeVersion resources.
type applicationTypeVersionState struct {
	// The URL to the application package
	AppPackageUrl *string `pulumi:"appPackageUrl"`
	// List of application type parameters that can be overridden when creating or updating the application.
	DefaultParameterList map[string]string `pulumi:"defaultParameterList"`
	// Azure resource location.
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState *string `pulumi:"provisioningState"`
	// Azure resource type.
	Type *string `pulumi:"type"`
}

type ApplicationTypeVersionState struct {
	// The URL to the application package
	AppPackageUrl pulumi.StringPtrInput
	// List of application type parameters that can be overridden when creating or updating the application.
	DefaultParameterList pulumi.StringMapInput
	// Azure resource location.
	Location pulumi.StringPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringPtrInput
	// Azure resource type.
	Type pulumi.StringPtrInput
}

func (ApplicationTypeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeVersionState)(nil)).Elem()
}

type applicationTypeVersionArgs struct {
	// The URL to the application package
	AppPackageUrl string `pulumi:"appPackageUrl"`
	// The name of the application type name resource.
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Azure resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The application type version.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a ApplicationTypeVersion resource.
type ApplicationTypeVersionArgs struct {
	// The URL to the application package
	AppPackageUrl pulumi.StringInput
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Azure resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The application type version.
	Version pulumi.StringInput
}

func (ApplicationTypeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeVersionArgs)(nil)).Elem()
}
