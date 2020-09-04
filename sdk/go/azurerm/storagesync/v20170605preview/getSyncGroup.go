// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170605preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSyncGroup(ctx *pulumi.Context, args *LookupSyncGroupArgs, opts ...pulumi.InvokeOption) (*LookupSyncGroupResult, error) {
	var rv LookupSyncGroupResult
	err := ctx.Invoke("azurerm:storagesync/v20170605preview:getSyncGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncGroupArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// Sync Group object.
type LookupSyncGroupResult struct {
	// The name of the resource.
	Name string `pulumi:"name"`
	// Sync group status
	SyncGroupStatus string `pulumi:"syncGroupStatus"`
	// The type of the resource
	Type string `pulumi:"type"`
	// Unique Id
	UniqueId *string `pulumi:"uniqueId"`
}
