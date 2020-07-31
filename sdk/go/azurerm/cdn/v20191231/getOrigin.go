// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191231

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupOrigin(ctx *pulumi.Context, args *LookupOriginArgs, opts ...pulumi.InvokeOption) (*LookupOriginResult, error) {
	var rv LookupOriginResult
	err := ctx.Invoke("azurerm:cdn/v20191231:getOrigin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Name of the origin which is unique within the endpoint.
	Name string `pulumi:"name"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
type LookupOriginResult struct {
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The JSON object that contains the properties of the origin.
	Properties OriginPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}