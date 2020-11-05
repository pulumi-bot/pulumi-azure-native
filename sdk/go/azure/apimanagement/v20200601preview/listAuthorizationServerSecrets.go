// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListAuthorizationServerSecrets(ctx *pulumi.Context, args *ListAuthorizationServerSecretsArgs, opts ...pulumi.InvokeOption) (*ListAuthorizationServerSecretsResult, error) {
	var rv ListAuthorizationServerSecretsResult
	err := ctx.Invoke("azure-nextgen:apimanagement/v20200601preview:listAuthorizationServerSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAuthorizationServerSecretsArgs struct {
	// Identifier of the authorization server.
	Authsid string `pulumi:"authsid"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// OAuth Server Secrets Contract.
type ListAuthorizationServerSecretsResult struct {
	// oAuth Authorization Server Secrets.
	ClientSecret *string `pulumi:"clientSecret"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
}
