// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIscsiServer(ctx *pulumi.Context, args *LookupIscsiServerArgs, opts ...pulumi.InvokeOption) (*LookupIscsiServerResult, error) {
	var rv LookupIscsiServerResult
	err := ctx.Invoke("azurerm:storsimple/v20161001:getIscsiServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiServerArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The iSCSI server name.
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The iSCSI server.
type LookupIscsiServerResult struct {
	// The name.
	Name string `pulumi:"name"`
	// The properties.
	Properties ISCSIServerPropertiesResponse `pulumi:"properties"`
	// The type.
	Type string `pulumi:"type"`
}