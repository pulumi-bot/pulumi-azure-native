// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListBatchAccountKeys(ctx *pulumi.Context, args *ListBatchAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListBatchAccountKeysResult, error) {
	var rv ListBatchAccountKeysResult
	err := ctx.Invoke("azure-nextgen:batch/v20210101:listBatchAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBatchAccountKeysArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A set of Azure Batch account keys.
type ListBatchAccountKeysResult struct {
	// The Batch account name.
	AccountName string `pulumi:"accountName"`
	// The primary key associated with the account.
	Primary string `pulumi:"primary"`
	// The secondary key associated with the account.
	Secondary string `pulumi:"secondary"`
}
