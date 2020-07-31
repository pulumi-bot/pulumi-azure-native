// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppSitePushSettingsSlot(ctx *pulumi.Context, args *ListWebAppSitePushSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSitePushSettingsSlotResult, error) {
	var rv ListWebAppSitePushSettingsSlotResult
	err := ctx.Invoke("azurerm:web/v20190801:listWebAppSitePushSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSitePushSettingsSlotArgs struct {
	// Name of web app slot. If not specified then will default to production slot.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Push settings for the App.
type ListWebAppSitePushSettingsSlotResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// PushSettings resource specific properties
	Properties PushSettingsResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}