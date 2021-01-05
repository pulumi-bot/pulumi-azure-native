// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListSiteBackupStatusSecrets(ctx *pulumi.Context, args *ListSiteBackupStatusSecretsArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupStatusSecretsResult, error) {
	var rv ListSiteBackupStatusSecretsResult
	err := ctx.Invoke("azure-nextgen:web/v20150801:listSiteBackupStatusSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupStatusSecretsArgs struct {
	// Id of backup
	BackupId string `pulumi:"backupId"`
	// Schedule for the backup if it is executed periodically
	BackupSchedule *BackupSchedule `pulumi:"backupSchedule"`
	// Databases included in the backup
	Databases []DatabaseBackupSetting `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
	Enabled *bool `pulumi:"enabled"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SAS URL to the container
	StorageAccountUrl *string `pulumi:"storageAccountUrl"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

// Backup description
type ListSiteBackupStatusSecretsResult struct {
	// Name of the blob which contains data for this backup
	BlobName *string `pulumi:"blobName"`
	// Unique correlation identifier. Please use this along with the timestamp while communicating with Azure support.
	CorrelationId *string `pulumi:"correlationId"`
	// Timestamp of the backup creation
	Created *string `pulumi:"created"`
	// List of databases included in the backup
	Databases []DatabaseBackupSettingResponse `pulumi:"databases"`
	// Timestamp when this backup finished.
	FinishedTimeStamp *string `pulumi:"finishedTimeStamp"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Timestamp of a last restore operation which used this backup.
	LastRestoreTimeStamp *string `pulumi:"lastRestoreTimeStamp"`
	// Resource Location
	Location string `pulumi:"location"`
	// Details regarding this backup. Might contain an error message.
	Log *string `pulumi:"log"`
	// Resource Name
	Name *string `pulumi:"name"`
	// True if this backup has been created due to a schedule being triggered.
	Scheduled *bool `pulumi:"scheduled"`
	// Size of the backup in bytes
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Backup status
	Status string `pulumi:"status"`
	// SAS URL for the storage account container which contains this backup
	StorageAccountUrl *string `pulumi:"storageAccountUrl"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Size of the original web app which has been backed up
	WebsiteSizeInBytes *float64 `pulumi:"websiteSizeInBytes"`
}
