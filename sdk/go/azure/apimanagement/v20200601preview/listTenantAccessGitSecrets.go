// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListTenantAccessGitSecrets(ctx *pulumi.Context, args *ListTenantAccessGitSecretsArgs, opts ...pulumi.InvokeOption) (*ListTenantAccessGitSecretsResult, error) {
	var rv ListTenantAccessGitSecretsResult
	err := ctx.Invoke("azure-nextgen:apimanagement/v20200601preview:listTenantAccessGitSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTenantAccessGitSecretsArgs struct {
	// The identifier of the Access configuration.
	AccessName string `pulumi:"accessName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Tenant access information contract of the API Management service.
type ListTenantAccessGitSecretsResult struct {
	// Determines whether direct access is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Primary access key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Secondary access key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	SecondaryKey *string `pulumi:"secondaryKey"`
}
