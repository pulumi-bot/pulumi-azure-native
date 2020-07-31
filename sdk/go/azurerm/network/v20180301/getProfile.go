// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azurerm:network/v20180301:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// The name of the Traffic Manager profile.
	Name string `pulumi:"name"`
	// The name of the resource group containing the Traffic Manager profile.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Traffic Manager profile.
type LookupProfileResult struct {
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The properties of the Traffic Manager profile.
	Properties ProfilePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
}