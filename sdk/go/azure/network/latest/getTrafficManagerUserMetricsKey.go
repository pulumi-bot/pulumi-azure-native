// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing Traffic Manager User Metrics.
// Latest API Version: 2018-04-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:network:getTrafficManagerUserMetricsKey'.
func LookupTrafficManagerUserMetricsKey(ctx *pulumi.Context, args *LookupTrafficManagerUserMetricsKeyArgs, opts ...pulumi.InvokeOption) (*LookupTrafficManagerUserMetricsKeyResult, error) {
	var rv LookupTrafficManagerUserMetricsKeyResult
	err := ctx.Invoke("azure-native:network/latest:getTrafficManagerUserMetricsKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficManagerUserMetricsKeyArgs struct {
}

// Class representing Traffic Manager User Metrics.
type LookupTrafficManagerUserMetricsKeyResult struct {
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `pulumi:"id"`
	// The key returned by the User Metrics operation.
	Key *string `pulumi:"key"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
}
