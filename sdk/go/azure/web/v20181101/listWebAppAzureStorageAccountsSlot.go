// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppAzureStorageAccountsSlot(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsSlotResult, error) {
	var rv ListWebAppAzureStorageAccountsSlotResult
	err := ctx.Invoke("azure-nextgen:web/v20181101:listWebAppAzureStorageAccountsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAzureStorageAccountsSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
	Slot string `pulumi:"slot"`
}

// AzureStorageInfo dictionary resource.
type ListWebAppAzureStorageAccountsSlotResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Azure storage accounts.
	Properties map[string]AzureStorageInfoValueResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
