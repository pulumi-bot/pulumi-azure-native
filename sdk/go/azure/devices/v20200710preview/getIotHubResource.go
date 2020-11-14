// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200710preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIotHubResource(ctx *pulumi.Context, args *LookupIotHubResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceResult, error) {
	var rv LookupIotHubResourceResult
	err := ctx.Invoke("azure-nextgen:devices/v20200710preview:getIotHubResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubResourceArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The description of the IoT hub.
type LookupIotHubResourceResult struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The managed identities for the IotHub.
	Identity *ArmIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// IotHub properties
	Properties IotHubPropertiesResponse `pulumi:"properties"`
	// IotHub SKU info
	Sku IotHubSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}
