// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListOpenIdConnectProviderSecrets(ctx *pulumi.Context, args *ListOpenIdConnectProviderSecretsArgs, opts ...pulumi.InvokeOption) (*ListOpenIdConnectProviderSecretsResult, error) {
	var rv ListOpenIdConnectProviderSecretsResult
	err := ctx.Invoke("azurerm:apimanagement/latest:listOpenIdConnectProviderSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenIdConnectProviderSecretsArgs struct {
	// Identifier of the OpenID Connect Provider.
	Opid string `pulumi:"opid"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
type ListOpenIdConnectProviderSecretsResult struct {
	// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
	ClientSecret *string `pulumi:"clientSecret"`
}