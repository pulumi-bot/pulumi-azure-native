// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListStreamingLocatorPaths(ctx *pulumi.Context, args *ListStreamingLocatorPathsArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorPathsResult, error) {
	var rv ListStreamingLocatorPathsResult
	err := ctx.Invoke("azurerm:media/v20180601preview:listStreamingLocatorPaths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorPathsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}

// Class of response for listPaths action
type ListStreamingLocatorPathsResult struct {
	// Download Paths supported by current Streaming Locator
	DownloadPaths []string `pulumi:"downloadPaths"`
	// Streaming Paths supported by current Streaming Locator
	StreamingPaths []StreamingPathResponse `pulumi:"streamingPaths"`
}
