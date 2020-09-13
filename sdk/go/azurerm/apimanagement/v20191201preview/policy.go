// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy Contract details.
//
// ## Example Usage
// ### ApiManagementCreatePolicy
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewPolicy(ctx, "policy", &apimanagement.PolicyArgs{
// 			Format:            pulumi.String("xml"),
// 			PolicyId:          pulumi.String("policy"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 			Value: pulumi.String("<policies>\n  <inbound />\n  <backend>\n    <forward-request />\n  </backend>\n  <outbound />\n</policies>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Policy struct {
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

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
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
		args = &PolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:Policy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:Policy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:Policy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:Policy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:Policy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201preview:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azurerm:apimanagement/v20191201preview:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value *string `pulumi:"value"`
}

type PolicyState struct {
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// The identifier of the Policy.
	PolicyId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
