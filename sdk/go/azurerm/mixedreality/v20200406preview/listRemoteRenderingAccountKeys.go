// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200406preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListRemoteRenderingAccountKeys(ctx *pulumi.Context, args *ListRemoteRenderingAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListRemoteRenderingAccountKeysResult, error) {
	var rv ListRemoteRenderingAccountKeysResult
	err := ctx.Invoke("azurerm:mixedreality/v20200406preview:listRemoteRenderingAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemoteRenderingAccountKeysArgs struct {
	// Name of an Mixed Reality Account.
	AccountName string `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Developer Keys of account
type ListRemoteRenderingAccountKeysResult struct {
	// value of primary key.
	PrimaryKey string `pulumi:"primaryKey"`
	// value of secondary key.
	SecondaryKey string `pulumi:"secondaryKey"`
}
