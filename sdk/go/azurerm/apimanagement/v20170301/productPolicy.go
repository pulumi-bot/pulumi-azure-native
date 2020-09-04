// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy Contract details.
type ProductPolicy struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringOutput `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProductPolicy registers a new resource with the given unique name, arguments, and options.
func NewProductPolicy(ctx *pulumi.Context,
	name string, args *ProductPolicyArgs, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	if args == nil || args.PolicyContent == nil {
		return nil, errors.New("missing required argument 'PolicyContent'")
	}
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ProductPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:ProductPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ProductPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:ProductPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:ProductPolicy"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:ProductPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ProductPolicy
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:ProductPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductPolicy gets an existing ProductPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductPolicyState, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	var resource ProductPolicy
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:ProductPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductPolicy resources.
type productPolicyState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent *string `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ProductPolicyState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ProductPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyState)(nil)).Elem()
}

type productPolicyArgs struct {
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent string `pulumi:"policyContent"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductPolicy resource.
type ProductPolicyArgs struct {
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyArgs)(nil)).Elem()
}
