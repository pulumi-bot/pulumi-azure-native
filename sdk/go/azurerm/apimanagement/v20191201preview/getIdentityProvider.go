// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIdentityProvider(ctx *pulumi.Context, args *LookupIdentityProviderArgs, opts ...pulumi.InvokeOption) (*LookupIdentityProviderResult, error) {
	var rv LookupIdentityProviderResult
	err := ctx.Invoke("azurerm:apimanagement/v20191201preview:getIdentityProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdentityProviderArgs struct {
	// Identity Provider Type identifier.
	IdentityProviderName string `pulumi:"identityProviderName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Identity Provider details.
type LookupIdentityProviderResult struct {
	// List of Allowed Tenants when configuring Azure Active Directory login.
	AllowedTenants []string `pulumi:"allowedTenants"`
	// OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
	Authority *string `pulumi:"authority"`
	// Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
	ClientId string `pulumi:"clientId"`
	// Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret *string `pulumi:"clientSecret"`
	// Resource name.
	Name string `pulumi:"name"`
	// Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
	PasswordResetPolicyName *string `pulumi:"passwordResetPolicyName"`
	// Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
	ProfileEditingPolicyName *string `pulumi:"profileEditingPolicyName"`
	// Signin Policy Name. Only applies to AAD B2C Identity Provider.
	SigninPolicyName *string `pulumi:"signinPolicyName"`
	// The TenantId to use instead of Common when logging into Active Directory
	SigninTenant *string `pulumi:"signinTenant"`
	// Signup Policy Name. Only applies to AAD B2C Identity Provider.
	SignupPolicyName *string `pulumi:"signupPolicyName"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
