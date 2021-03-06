// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Function secrets.
// Latest API Version: 2020-10-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:web:listWebAppSyncFunctionTriggersSlot'.
func ListWebAppSyncFunctionTriggersSlot(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersSlotResult, error) {
	var rv ListWebAppSyncFunctionTriggersSlotResult
	err := ctx.Invoke("azure-native:web/latest:listWebAppSyncFunctionTriggersSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot string `pulumi:"slot"`
}

// Function secrets.
type ListWebAppSyncFunctionTriggersSlotResult struct {
	// Secret key.
	Key *string `pulumi:"key"`
	// Trigger URL.
	TriggerUrl *string `pulumi:"triggerUrl"`
}
