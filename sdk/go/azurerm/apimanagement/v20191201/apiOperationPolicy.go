// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy Contract details.
//
// ## Example Usage
// ### ApiManagementCreateApiOperationPolicy
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
// 		_, err := apimanagement.NewApiOperationPolicy(ctx, "apiOperationPolicy", &apimanagement.ApiOperationPolicyArgs{
// 			ApiId:             pulumi.String("5600b57e7e8880006a040001"),
// 			Format:            pulumi.String("xml"),
// 			OperationId:       pulumi.String("5600b57e7e8880006a080001"),
// 			PolicyId:          pulumi.String("policy"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Value:             pulumi.String("<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApiOperationPolicy struct {
	pulumi.CustomResourceState

	// Format of the policyContent.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApiOperationPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiOperationPolicy(ctx *pulumi.Context,
	name string, args *ApiOperationPolicyArgs, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.OperationId == nil {
		return nil, errors.New("missing required argument 'OperationId'")
	}
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &ApiOperationPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ApiOperationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiOperationPolicy
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201:ApiOperationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperationPolicy gets an existing ApiOperationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationPolicyState, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	var resource ApiOperationPolicy
	err := ctx.ReadResource("azurerm:apimanagement/v20191201:ApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperationPolicy resources.
type apiOperationPolicyState struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value *string `pulumi:"value"`
}

type ApiOperationPolicyState struct {
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringPtrInput
}

func (ApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyState)(nil)).Elem()
}

type apiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApiOperationPolicy resource.
type ApiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringInput
}

func (ApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyArgs)(nil)).Elem()
}
