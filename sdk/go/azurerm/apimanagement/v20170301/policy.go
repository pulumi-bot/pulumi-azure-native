// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy Contract details.
type Policy struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringOutput `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.PolicyContent == nil {
		return nil, errors.New("missing required argument 'PolicyContent'")
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
	if args == nil {
		args = &PolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:Policy"),
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
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201preview:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:Policy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent *string `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type PolicyState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent string `pulumi:"policyContent"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
