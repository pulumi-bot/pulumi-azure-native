// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagecache

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCachStorageTarget(ctx *pulumi.Context, args *LookupCachStorageTargetArgs, opts ...pulumi.InvokeOption) (*LookupCachStorageTargetResult, error) {
	var rv LookupCachStorageTargetResult
	err := ctx.Invoke("azurerm:storagecache:getCachStorageTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCachStorageTargetArgs struct {
	// Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
	CacheName string `pulumi:"cacheName"`
	// Name of the Storage Target. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
	Name string `pulumi:"name"`
	// Target resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Type of the Storage Target.
type LookupCachStorageTargetResult struct {
	// Name of the Storage Target.
	Name string `pulumi:"name"`
	// StorageTarget properties
	Properties StorageTargetPropertiesResponse `pulumi:"properties"`
	// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type string `pulumi:"type"`
}