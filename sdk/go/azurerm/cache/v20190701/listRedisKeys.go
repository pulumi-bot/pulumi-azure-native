// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListRedisKeys(ctx *pulumi.Context, args *ListRedisKeysArgs, opts ...pulumi.InvokeOption) (*ListRedisKeysResult, error) {
	var rv ListRedisKeysResult
	err := ctx.Invoke("azurerm:cache/v20190701:listRedisKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRedisKeysArgs struct {
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Redis cache access keys.
type ListRedisKeysResult struct {
	// The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey string `pulumi:"primaryKey"`
	// The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey string `pulumi:"secondaryKey"`
}
