// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupActivityLogAlert(ctx *pulumi.Context, args *LookupActivityLogAlertArgs, opts ...pulumi.InvokeOption) (*LookupActivityLogAlertResult, error) {
	var rv LookupActivityLogAlertResult
	err := ctx.Invoke("azurerm:insights/v20170301preview:getActivityLogAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActivityLogAlertArgs struct {
	// The name of the activity log alert.
	ActivityLogAlertName string `pulumi:"activityLogAlertName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An activity log alert resource.
type LookupActivityLogAlertResult struct {
	// The actions that will activate when the condition is met.
	Actions ActivityLogAlertActionListResponse `pulumi:"actions"`
	// The condition that will cause this alert to activate.
	Condition ActivityLogAlertAllOfConditionResponse `pulumi:"condition"`
	// A description of this activity log alert.
	Description *string `pulumi:"description"`
	// Indicates whether this activity log alert is enabled. If an activity log alert is not enabled, then none of its actions will be activated.
	Enabled *bool `pulumi:"enabled"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// A list of resourceIds that will be used as prefixes. The alert will only apply to activityLogs with resourceIds that fall under one of these prefixes. This list must include at least one item.
	Scopes []string `pulumi:"scopes"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
}
