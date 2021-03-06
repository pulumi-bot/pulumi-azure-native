// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// IoT Defender settings
// API Version: 2020-08-06-preview.
func LookupIotDefenderSetting(ctx *pulumi.Context, args *LookupIotDefenderSettingArgs, opts ...pulumi.InvokeOption) (*LookupIotDefenderSettingResult, error) {
	var rv LookupIotDefenderSettingResult
	err := ctx.Invoke("azure-native:security:getIotDefenderSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDefenderSettingArgs struct {
}

// IoT Defender settings
type LookupIotDefenderSettingResult struct {
	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota int `pulumi:"deviceQuota"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds []string `pulumi:"sentinelWorkspaceResourceIds"`
	// Resource type
	Type string `pulumi:"type"`
}
