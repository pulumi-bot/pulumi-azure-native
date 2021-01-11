// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200202

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	var rv LookupVaultResult
	err := ctx.Invoke("azure-nextgen:recoveryservices/v20200202:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVaultArgs struct {
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// Resource information, as returned by the resource provider.
type LookupVaultResult struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityDataResponse `pulumi:"identity"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// Properties of the vault.
	Properties VaultPropertiesResponse `pulumi:"properties"`
	// Identifies the unique system identifier for each Azure resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}
