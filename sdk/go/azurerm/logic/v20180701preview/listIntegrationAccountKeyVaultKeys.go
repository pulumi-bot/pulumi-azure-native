// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListIntegrationAccountKeyVaultKeys(ctx *pulumi.Context, args *ListIntegrationAccountKeyVaultKeysArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountKeyVaultKeysResult, error) {
	var rv ListIntegrationAccountKeyVaultKeysResult
	err := ctx.Invoke("azurerm:logic/v20180701preview:listIntegrationAccountKeyVaultKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountKeyVaultKeysArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The key vault reference.
	KeyVault KeyVaultReference `pulumi:"keyVault"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The skip token.
	SkipToken *string `pulumi:"skipToken"`
}

// Collection of key vault keys.
type ListIntegrationAccountKeyVaultKeysResult struct {
	// The skip token.
	SkipToken *string `pulumi:"skipToken"`
	// The key vault keys.
	Value []KeyVaultKeyResponse `pulumi:"value"`
}
