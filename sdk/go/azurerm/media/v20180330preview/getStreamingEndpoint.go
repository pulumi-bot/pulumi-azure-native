// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180330preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStreamingEndpoint(ctx *pulumi.Context, args *LookupStreamingEndpointArgs, opts ...pulumi.InvokeOption) (*LookupStreamingEndpointResult, error) {
	var rv LookupStreamingEndpointResult
	err := ctx.Invoke("azurerm:media/v20180330preview:getStreamingEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingEndpointArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the StreamingEndpoint.
	StreamingEndpointName string `pulumi:"streamingEndpointName"`
}

// The StreamingEndpoint.
type LookupStreamingEndpointResult struct {
	// The access control definition of the StreamingEndpoint.
	AccessControl *StreamingEndpointAccessControlResponse `pulumi:"accessControl"`
	// AvailabilitySet name
	AvailabilitySetName *string `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled *bool `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile *string `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider *string `pulumi:"cdnProvider"`
	// The exact time the StreamingEndpoint was created.
	Created string `pulumi:"created"`
	// The StreamingEndpoint access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the StreamingEndpoint
	CustomHostNames []string `pulumi:"customHostNames"`
	// The StreamingEndpoint description.
	Description *string `pulumi:"description"`
	// The free trial expiration time.
	FreeTrialEndTime string `pulumi:"freeTrialEndTime"`
	// The StreamingEndpoint host name.
	HostName string `pulumi:"hostName"`
	// The exact time the StreamingEndpoint was last modified.
	LastModified string `pulumi:"lastModified"`
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// Max cache age
	MaxCacheAge *int `pulumi:"maxCacheAge"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of the StreamingEndpoint.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource state of the StreamingEndpoint.
	ResourceState string `pulumi:"resourceState"`
	// The number of scale units.
	ScaleUnits *int `pulumi:"scaleUnits"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
