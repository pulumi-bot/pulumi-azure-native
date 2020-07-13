// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStorageAccountQueueServiceQueue(ctx *pulumi.Context, args *LookupStorageAccountQueueServiceQueueArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountQueueServiceQueueResult, error) {
	var rv LookupStorageAccountQueueServiceQueueResult
	err := ctx.Invoke("azurerm:storage:getStorageAccountQueueServiceQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountQueueServiceQueueArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.
	Name string `pulumi:"name"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupStorageAccountQueueServiceQueueResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// Queue resource properties.
	Properties QueuePropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}