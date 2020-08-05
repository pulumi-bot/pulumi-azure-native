// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppBackupStatusSecrets(ctx *pulumi.Context, args *ListWebAppBackupStatusSecretsArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupStatusSecretsResult, error) {
	var rv ListWebAppBackupStatusSecretsResult
	err := ctx.Invoke("azurerm:web/v20180201:listWebAppBackupStatusSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupStatusSecretsArgs struct {
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
	// ID of backup.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SAS URL to the container.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
}

// Backup description.
type ListWebAppBackupStatusSecretsResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// BackupItem resource specific properties
	Properties BackupItemResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
