// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetsnapshotPolicy(ctx *pulumi.Context, args *GetsnapshotPolicyArgs, opts ...pulumi.InvokeOption) (*GetsnapshotPolicyResult, error) {
	var rv GetsnapshotPolicyResult
	err := ctx.Invoke("azurerm:netapp/v20200601:getsnapshotPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetsnapshotPolicyArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the snapshot policy target
	SnapshotPolicyName string `pulumi:"snapshotPolicyName"`
}

// Snapshot policy information
type GetsnapshotPolicyResult struct {
	// Schedule for daily snapshots
	DailySchedule map[string]interface{} `pulumi:"dailySchedule"`
	// The property to decide policy is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// Schedule for hourly snapshots
	HourlySchedule map[string]interface{} `pulumi:"hourlySchedule"`
	// Resource location
	Location string `pulumi:"location"`
	// Schedule for monthly snapshots
	MonthlySchedule map[string]interface{} `pulumi:"monthlySchedule"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// Schedule for weekly snapshots
	WeeklySchedule map[string]interface{} `pulumi:"weeklySchedule"`
}