// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagerDeviceFileserverShare(ctx *pulumi.Context, args *LookupManagerDeviceFileserverShareArgs, opts ...pulumi.InvokeOption) (*LookupManagerDeviceFileserverShareResult, error) {
	var rv LookupManagerDeviceFileserverShareResult
	err := ctx.Invoke("azurerm:storsimple:getManagerDeviceFileserverShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerDeviceFileserverShareArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The file server name.
	FileServerName string `pulumi:"fileServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The file share name.
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The File Share.
type LookupManagerDeviceFileserverShareResult struct {
	// The name.
	Name string `pulumi:"name"`
	// The properties.
	Properties FileSharePropertiesResponse `pulumi:"properties"`
	// The type.
	Type string `pulumi:"type"`
}