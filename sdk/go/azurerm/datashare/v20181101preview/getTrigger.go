// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTrigger(ctx *pulumi.Context, args *LookupTriggerArgs, opts ...pulumi.InvokeOption) (*LookupTriggerResult, error) {
	var rv LookupTriggerResult
	err := ctx.Invoke("azurerm:datashare/v20181101preview:getTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTriggerArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	// The name of the trigger.
	TriggerName string `pulumi:"triggerName"`
}

// A Trigger data transfer object.
type LookupTriggerResult struct {
	// Kind of synchronization
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}
