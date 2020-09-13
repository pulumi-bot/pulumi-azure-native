// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Tag Contract details.
//
// ## Example Usage
// ### ApiManagementCreateApiTag
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
// 		_, err := apimanagement.NewTagByApi(ctx, "tagByApi", &apimanagement.TagByApiArgs{
// 			ApiId:             pulumi.String("5931a75ae4bbd512a88c680b"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			TagId:             pulumi.String("tagId1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type TagByApi struct {
	pulumi.CustomResourceState

	// Tag name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagByApi registers a new resource with the given unique name, arguments, and options.
func NewTagByApi(ctx *pulumi.Context,
	name string, args *TagByApiArgs, opts ...pulumi.ResourceOption) (*TagByApi, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.TagId == nil {
		return nil, errors.New("missing required argument 'TagId'")
	}
	if args == nil {
		args = &TagByApiArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:TagByApi"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:TagByApi"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:TagByApi"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:TagByApi"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:TagByApi"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:TagByApi"),
		},
	})
	opts = append(opts, aliases)
	var resource TagByApi
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201:TagByApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagByApi gets an existing TagByApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagByApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByApiState, opts ...pulumi.ResourceOption) (*TagByApi, error) {
	var resource TagByApi
	err := ctx.ReadResource("azurerm:apimanagement/v20191201:TagByApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagByApi resources.
type tagByApiState struct {
	// Tag name.
	DisplayName *string `pulumi:"displayName"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type TagByApiState struct {
	// Tag name.
	DisplayName pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (TagByApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByApiState)(nil)).Elem()
}

type tagByApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// The set of arguments for constructing a TagByApi resource.
type TagByApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput
}

func (TagByApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByApiArgs)(nil)).Elem()
}
