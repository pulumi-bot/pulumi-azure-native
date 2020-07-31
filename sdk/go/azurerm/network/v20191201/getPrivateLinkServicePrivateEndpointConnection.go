// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPrivateLinkServicePrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateLinkServicePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServicePrivateEndpointConnectionResult, error) {
	var rv LookupPrivateLinkServicePrivateEndpointConnectionResult
	err := ctx.Invoke("azurerm:network/v20191201:getPrivateLinkServicePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServicePrivateEndpointConnectionArgs struct {
	// The name of the private end point connection.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName string `pulumi:"serviceName"`
}

// PrivateEndpointConnection resource.
type LookupPrivateLinkServicePrivateEndpointConnectionResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the private end point connection.
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}