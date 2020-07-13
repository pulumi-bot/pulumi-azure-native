// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy Contract details.
type ServiceProductPolicy struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Policy.
	Properties PolicyContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceProductPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceProductPolicy(ctx *pulumi.Context,
	name string, args *ServiceProductPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceProductPolicy, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
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
		args = &ServiceProductPolicyArgs{}
	}
	var resource ServiceProductPolicy
	err := ctx.RegisterResource("azurerm:apimanagement:ServiceProductPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceProductPolicy gets an existing ServiceProductPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceProductPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceProductPolicyState, opts ...pulumi.ResourceOption) (*ServiceProductPolicy, error) {
	var resource ServiceProductPolicy
	err := ctx.ReadResource("azurerm:apimanagement:ServiceProductPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceProductPolicy resources.
type serviceProductPolicyState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the Policy.
	Properties *PolicyContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ServiceProductPolicyState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the Policy.
	Properties PolicyContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ServiceProductPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceProductPolicyState)(nil)).Elem()
}

type serviceProductPolicyArgs struct {
	// The identifier of the Policy.
	Name string `pulumi:"name"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// Properties of the Policy.
	Properties *PolicyContractProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServiceProductPolicy resource.
type ServiceProductPolicyArgs struct {
	// The identifier of the Policy.
	Name pulumi.StringInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// Properties of the Policy.
	Properties PolicyContractPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ServiceProductPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceProductPolicyArgs)(nil)).Elem()
}