// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180930

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azurerm:compute/v20180930:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Disk resource.
type LookupDiskResult struct {
	// Resource location
	Location string `pulumi:"location"`
	// A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy string `pulumi:"managedBy"`
	// Resource name
	Name string `pulumi:"name"`
	// Disk resource properties.
	Properties DiskPropertiesResponse `pulumi:"properties"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku *DiskSkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// The Logical zone list for Disk.
	Zones []string `pulumi:"zones"`
}