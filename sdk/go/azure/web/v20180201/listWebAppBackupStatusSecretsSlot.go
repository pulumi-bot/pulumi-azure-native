// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppBackupStatusSecretsSlot(ctx *pulumi.Context, args *ListWebAppBackupStatusSecretsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupStatusSecretsSlotResult, error) {
	var rv ListWebAppBackupStatusSecretsSlotResult
	err := ctx.Invoke("azure-nextgen:web/v20180201:listWebAppBackupStatusSecretsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupStatusSecretsSlotArgs struct {
	// ID of backup.
	BackupId string `pulumi:"backupId"`
	// Name of the backup.
	BackupName *string `pulumi:"backupName"`
	// Schedule for the backup if it is executed periodically.
	BackupSchedule *BackupSchedule `pulumi:"backupSchedule"`
	// Databases included in the backup.
	Databases []DatabaseBackupSetting `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
	Enabled *bool `pulumi:"enabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of web app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
	// SAS URL to the container.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
}

// Backup description.
type ListWebAppBackupStatusSecretsSlotResult struct {
	// Id of the backup.
	BackupId int `pulumi:"backupId"`
	// Name of the blob which contains data for this backup.
	BlobName string `pulumi:"blobName"`
	// Unique correlation identifier. Please use this along with the timestamp while communicating with Azure support.
	CorrelationId string `pulumi:"correlationId"`
	// Timestamp of the backup creation.
	Created string `pulumi:"created"`
	// List of databases included in the backup.
	Databases []DatabaseBackupSettingResponse `pulumi:"databases"`
	// Timestamp when this backup finished.
	FinishedTimeStamp string `pulumi:"finishedTimeStamp"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Timestamp of a last restore operation which used this backup.
	LastRestoreTimeStamp string `pulumi:"lastRestoreTimeStamp"`
	// Details regarding this backup. Might contain an error message.
	Log string `pulumi:"log"`
	// Resource Name.
	Name string `pulumi:"name"`
	// True if this backup has been created due to a schedule being triggered.
	Scheduled bool `pulumi:"scheduled"`
	// Size of the backup in bytes.
	SizeInBytes float64 `pulumi:"sizeInBytes"`
	// Backup status.
	Status string `pulumi:"status"`
	// SAS URL for the storage account container which contains this backup.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
	// Resource type.
	Type string `pulumi:"type"`
	// Size of the original web app which has been backed up.
	WebsiteSizeInBytes float64 `pulumi:"websiteSizeInBytes"`
}
