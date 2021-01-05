// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-nextgen:netapp/v20190801:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// Volume resource
type LookupVolumeResult struct {
	// Unique Baremetal Tenant Identifier.
	BaremetalTenantId string `pulumi:"baremetalTenantId"`
	// A unique file path for the volume. Used when creating mount targets
	CreationToken string `pulumi:"creationToken"`
	// DataProtection volume, can have a replication object
	DataProtection *VolumePropertiesResponseDataProtection `pulumi:"dataProtection"`
	// Set of export policy rules
	ExportPolicy *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	// Unique FileSystem Identifier.
	FileSystemId string `pulumi:"fileSystemId"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// List of mount targets
	MountTargets []MountTargetPropertiesResponse `pulumi:"mountTargets"`
	// Resource name
	Name string `pulumi:"name"`
	// Set of protocol types
	ProtocolTypes []string `pulumi:"protocolTypes"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// The service level of the file system
	ServiceLevel *string `pulumi:"serviceLevel"`
	// UUID v4 or resource identifier used to identify the Snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId string `pulumi:"subnetId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
	UsageThreshold float64 `pulumi:"usageThreshold"`
	// What type of volume is this
	VolumeType *string `pulumi:"volumeType"`
}
