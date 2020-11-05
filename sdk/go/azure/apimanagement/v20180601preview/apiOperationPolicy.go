// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy Contract details.
type ApiOperationPolicy struct {
	pulumi.CustomResourceState

	// Format of the policyContent.
	ContentFormat pulumi.StringPtrOutput `pulumi:"contentFormat"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringOutput `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
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
		args = &ApiOperationPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/latest:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:ApiOperationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiOperationPolicy
	err := ctx.RegisterResource("azure-nextgen:apimanagement/v20180601preview:ApiOperationPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:apimanagement/v20180601preview:ApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperationPolicy resources.
type apiOperationPolicyState struct {
	// Format of the policyContent.
	ContentFormat *string `pulumi:"contentFormat"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent *string `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ApiOperationPolicyState struct {
	// Format of the policyContent.
	ContentFormat pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyState)(nil)).Elem()
}

type apiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Format of the policyContent.
	ContentFormat *string `pulumi:"contentFormat"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent string `pulumi:"policyContent"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ApiOperationPolicy resource.
type ApiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Format of the policyContent.
	ContentFormat pulumi.StringPtrInput
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyArgs)(nil)).Elem()
}
