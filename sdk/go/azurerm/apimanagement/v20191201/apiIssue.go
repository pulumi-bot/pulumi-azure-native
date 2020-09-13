// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Issue Contract details.
//
// ## Example Usage
// ### ApiManagementCreateApiIssue
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
// 		_, err := apimanagement.NewApiIssue(ctx, "apiIssue", &apimanagement.ApiIssueArgs{
// 			ApiId:             pulumi.String("57d1f7558aa04f15146d9d8a"),
// 			CreatedDate:       pulumi.String("2018-02-01T22:21:20.467Z"),
// 			Description:       pulumi.String("New API issue description"),
// 			IssueId:           pulumi.String("57d2ef278aa04f0ad01d6cdc"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			State:             pulumi.String("open"),
// 			Title:             pulumi.String("New API issue"),
// 			UserId:            pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApiIssue struct {
	pulumi.CustomResourceState

	// A resource identifier for the API the issue was created for.
	ApiId pulumi.StringPtrOutput `pulumi:"apiId"`
	// Date and time when the issue was created.
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	// Text describing the issue.
	Description pulumi.StringOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Status of the issue.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The issue title.
	Title pulumi.StringOutput `pulumi:"title"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// A resource identifier for the user created the issue.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewApiIssue registers a new resource with the given unique name, arguments, and options.
func NewApiIssue(ctx *pulumi.Context,
	name string, args *ApiIssueArgs, opts ...pulumi.ResourceOption) (*ApiIssue, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.IssueId == nil {
		return nil, errors.New("missing required argument 'IssueId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil {
		args = &ApiIssueArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:ApiIssue"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiIssue"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiIssue"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:ApiIssue"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ApiIssue"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiIssue
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201:ApiIssue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIssue gets an existing ApiIssue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIssue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueState, opts ...pulumi.ResourceOption) (*ApiIssue, error) {
	var resource ApiIssue
	err := ctx.ReadResource("azurerm:apimanagement/v20191201:ApiIssue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIssue resources.
type apiIssueState struct {
	// A resource identifier for the API the issue was created for.
	ApiId *string `pulumi:"apiId"`
	// Date and time when the issue was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Text describing the issue.
	Description *string `pulumi:"description"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Status of the issue.
	State *string `pulumi:"state"`
	// The issue title.
	Title *string `pulumi:"title"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// A resource identifier for the user created the issue.
	UserId *string `pulumi:"userId"`
}

type ApiIssueState struct {
	// A resource identifier for the API the issue was created for.
	ApiId pulumi.StringPtrInput
	// Date and time when the issue was created.
	CreatedDate pulumi.StringPtrInput
	// Text describing the issue.
	Description pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Status of the issue.
	State pulumi.StringPtrInput
	// The issue title.
	Title pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// A resource identifier for the user created the issue.
	UserId pulumi.StringPtrInput
}

func (ApiIssueState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueState)(nil)).Elem()
}

type apiIssueArgs struct {
	// A resource identifier for the API the issue was created for.
	ApiId string `pulumi:"apiId"`
	// Date and time when the issue was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Text describing the issue.
	Description string `pulumi:"description"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Status of the issue.
	State *string `pulumi:"state"`
	// The issue title.
	Title string `pulumi:"title"`
	// A resource identifier for the user created the issue.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ApiIssue resource.
type ApiIssueArgs struct {
	// A resource identifier for the API the issue was created for.
	ApiId pulumi.StringInput
	// Date and time when the issue was created.
	CreatedDate pulumi.StringPtrInput
	// Text describing the issue.
	Description pulumi.StringInput
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Status of the issue.
	State pulumi.StringPtrInput
	// The issue title.
	Title pulumi.StringInput
	// A resource identifier for the user created the issue.
	UserId pulumi.StringInput
}

func (ApiIssueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueArgs)(nil)).Elem()
}
