// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPrivateEndpointPrivateDnsZoneGroup(ctx *pulumi.Context, args *LookupPrivateEndpointPrivateDnsZoneGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointPrivateDnsZoneGroupResult, error) {
	var rv LookupPrivateEndpointPrivateDnsZoneGroupResult
	err := ctx.Invoke("azurerm:network:getPrivateEndpointPrivateDnsZoneGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointPrivateDnsZoneGroupArgs struct {
	// The name of the private dns zone group.
	Name string `pulumi:"name"`
	// The name of the private endpoint.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Private dns zone group resource.
type LookupPrivateEndpointPrivateDnsZoneGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the private dns zone group.
	Properties PrivateDnsZoneGroupPropertiesFormatResponse `pulumi:"properties"`
}