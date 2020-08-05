// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Identity Provider details.
type IdentityProvider struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Identity Provider contract properties.
	Properties IdentityProviderContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &IdentityProviderArgs{}
	}
	var resource IdentityProvider
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("azurerm:apimanagement/v20191201:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Identity Provider contract properties.
	Properties *IdentityProviderContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type IdentityProviderState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Identity Provider contract properties.
	Properties IdentityProviderContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	// List of Allowed Tenants when configuring Azure Active Directory login.
	AllowedTenants []string `pulumi:"allowedTenants"`
	// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
	Authority *string `pulumi:"authority"`
	// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
	ClientId string `pulumi:"clientId"`
	// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret string `pulumi:"clientSecret"`
	// Identity Provider Type identifier.
	Name string `pulumi:"name"`
	// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
	PasswordResetPolicyName *string `pulumi:"passwordResetPolicyName"`
	// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
	ProfileEditingPolicyName *string `pulumi:"profileEditingPolicyName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Signin Policy Name. Only applies to AAD B2C Identity Provider.
	SigninPolicyName *string `pulumi:"signinPolicyName"`
	// The TenantId to use instead of Common when logging into Active Directory
	SigninTenant *string `pulumi:"signinTenant"`
	// Signup Policy Name. Only applies to AAD B2C Identity Provider.
	SignupPolicyName *string `pulumi:"signupPolicyName"`
	// Identity Provider Type identifier.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// List of Allowed Tenants when configuring Azure Active Directory login.
	AllowedTenants pulumi.StringArrayInput
	// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
	Authority pulumi.StringPtrInput
	// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
	ClientId pulumi.StringInput
	// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret pulumi.StringInput
	// Identity Provider Type identifier.
	Name pulumi.StringInput
	// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
	PasswordResetPolicyName pulumi.StringPtrInput
	// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
	ProfileEditingPolicyName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Signin Policy Name. Only applies to AAD B2C Identity Provider.
	SigninPolicyName pulumi.StringPtrInput
	// The TenantId to use instead of Common when logging into Active Directory
	SigninTenant pulumi.StringPtrInput
	// Signup Policy Name. Only applies to AAD B2C Identity Provider.
	SignupPolicyName pulumi.StringPtrInput
	// Identity Provider Type identifier.
	Type pulumi.StringPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}
