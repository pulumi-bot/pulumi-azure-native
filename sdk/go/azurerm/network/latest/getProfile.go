// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azurerm:network/latest:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// The name of the Traffic Manager profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group containing the Traffic Manager profile.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Traffic Manager profile.
type LookupProfileResult struct {
	// The DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfigResponse `pulumi:"dnsConfig"`
	// The list of endpoints in the Traffic Manager profile.
	Endpoints []EndpointResponse `pulumi:"endpoints"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn *int `pulumi:"maxReturn"`
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfigResponse `pulumi:"monitorConfig"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The status of the Traffic Manager profile.
	ProfileStatus *string `pulumi:"profileStatus"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *string `pulumi:"trafficViewEnrollmentStatus"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
}