// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200515

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An access policy is used to grant users and applications access to the environment. Roles are assigned to service principals in Azure Active Directory. These roles define the actions the principal can perform through the Time Series Insights data plane APIs.
type AccessPolicy struct {
	pulumi.CustomResourceState

	// An description of the access policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringPtrOutput `pulumi:"principalObjectId"`
	// The list of roles the principal is assigned on the environment.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil || args.AccessPolicyName == nil {
		return nil, errors.New("missing required argument 'AccessPolicyName'")
	}
	if args == nil || args.EnvironmentName == nil {
		return nil, errors.New("missing required argument 'EnvironmentName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AccessPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:timeseriesinsights/latest:AccessPolicy"),
		},
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20170228preview:AccessPolicy"),
		},
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20171115:AccessPolicy"),
		},
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20180815preview:AccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessPolicy
	err := ctx.RegisterResource("azurerm:timeseriesinsights/v20200515:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicy gets an existing AccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("azurerm:timeseriesinsights/v20200515:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicy resources.
type accessPolicyState struct {
	// An description of the access policy.
	Description *string `pulumi:"description"`
	// Resource name
	Name *string `pulumi:"name"`
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// The list of roles the principal is assigned on the environment.
	Roles []string `pulumi:"roles"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AccessPolicyState struct {
	// An description of the access policy.
	Description pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringPtrInput
	// The list of roles the principal is assigned on the environment.
	Roles pulumi.StringArrayInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	// Name of the access policy.
	AccessPolicyName string `pulumi:"accessPolicyName"`
	// An description of the access policy.
	Description *string `pulumi:"description"`
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of roles the principal is assigned on the environment.
	Roles []string `pulumi:"roles"`
}

// The set of arguments for constructing a AccessPolicy resource.
type AccessPolicyArgs struct {
	// Name of the access policy.
	AccessPolicyName pulumi.StringInput
	// An description of the access policy.
	Description pulumi.StringPtrInput
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The list of roles the principal is assigned on the environment.
	Roles pulumi.StringArrayInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}
