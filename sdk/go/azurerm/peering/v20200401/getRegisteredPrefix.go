// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRegisteredPrefix(ctx *pulumi.Context, args *LookupRegisteredPrefixArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredPrefixResult, error) {
	var rv LookupRegisteredPrefixResult
	err := ctx.Invoke("azurerm:peering/v20200401:getRegisteredPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredPrefixArgs struct {
	// The name of the registered prefix.
	Name string `pulumi:"name"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The customer's prefix that is registered by the peering service provider.
type LookupRegisteredPrefixResult struct {
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties that define a registered prefix.
	Properties PeeringRegisteredPrefixPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type string `pulumi:"type"`
}