// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ApiRelease details.
//
// ## Example Usage
// ### ApiManagementCreateApiRelease
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewApiRelease(ctx, "apiRelease", &apimanagement.ApiReleaseArgs{
// 			ApiId:             pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/a1"),
// 			Notes:             pulumi.String("yahooagain"),
// 			ReleaseId:         pulumi.String("testrev"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApiRelease struct {
	pulumi.CustomResourceState

	// Identifier of the API the release belongs to.
	ApiId pulumi.StringPtrOutput `pulumi:"apiId"`
	// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
	CreatedDateTime pulumi.StringOutput `pulumi:"createdDateTime"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Release Notes
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the API release was updated.
	UpdatedDateTime pulumi.StringOutput `pulumi:"updatedDateTime"`
}

// NewApiRelease registers a new resource with the given unique name, arguments, and options.
func NewApiRelease(ctx *pulumi.Context,
	name string, args *ApiReleaseArgs, opts ...pulumi.ResourceOption) (*ApiRelease, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.ReleaseId == nil {
		return nil, errors.New("missing required argument 'ReleaseId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ApiReleaseArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:ApiRelease"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiRelease"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiRelease"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ApiRelease"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:ApiRelease"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ApiRelease"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiRelease
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201:ApiRelease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiRelease gets an existing ApiRelease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiReleaseState, opts ...pulumi.ResourceOption) (*ApiRelease, error) {
	var resource ApiRelease
	err := ctx.ReadResource("azurerm:apimanagement/v20191201:ApiRelease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiRelease resources.
type apiReleaseState struct {
	// Identifier of the API the release belongs to.
	ApiId *string `pulumi:"apiId"`
	// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
	CreatedDateTime *string `pulumi:"createdDateTime"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Release Notes
	Notes *string `pulumi:"notes"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// The time the API release was updated.
	UpdatedDateTime *string `pulumi:"updatedDateTime"`
}

type ApiReleaseState struct {
	// Identifier of the API the release belongs to.
	ApiId pulumi.StringPtrInput
	// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
	CreatedDateTime pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Release Notes
	Notes pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// The time the API release was updated.
	UpdatedDateTime pulumi.StringPtrInput
}

func (ApiReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiReleaseState)(nil)).Elem()
}

type apiReleaseArgs struct {
	// Identifier of the API the release belongs to.
	ApiId string `pulumi:"apiId"`
	// Release Notes
	Notes *string `pulumi:"notes"`
	// Release identifier within an API. Must be unique in the current API Management service instance.
	ReleaseId string `pulumi:"releaseId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ApiRelease resource.
type ApiReleaseArgs struct {
	// Identifier of the API the release belongs to.
	ApiId pulumi.StringInput
	// Release Notes
	Notes pulumi.StringPtrInput
	// Release identifier within an API. Must be unique in the current API Management service instance.
	ReleaseId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ApiReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiReleaseArgs)(nil)).Elem()
}
