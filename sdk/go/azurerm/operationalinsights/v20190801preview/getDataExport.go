// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDataExport(ctx *pulumi.Context, args *LookupDataExportArgs, opts ...pulumi.InvokeOption) (*LookupDataExportResult, error) {
	var rv LookupDataExportResult
	err := ctx.Invoke("azurerm:operationalinsights/v20190801preview:getDataExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataExportArgs struct {
	// The data export rule name.
	DataExportName string `pulumi:"dataExportName"`
	// The workspace's resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Log Analytics workspace name.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The top level data export resource container.
type LookupDataExportResult struct {
	// When ‘true’, all workspace's tables are exported.
	AllTables *bool `pulumi:"allTables"`
	// The latest data export rule modification time.
	CreatedDate *string `pulumi:"createdDate"`
	// The data export rule ID.
	DataExportId *string `pulumi:"dataExportId"`
	// Active when enabled.
	Enable *bool `pulumi:"enable"`
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName *string `pulumi:"eventHubName"`
	// Date and time when the export was last modified.
	LastModifiedDate *string `pulumi:"lastModifiedDate"`
	// Resource name.
	Name string `pulumi:"name"`
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId string `pulumi:"resourceId"`
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames []string `pulumi:"tableNames"`
	// Resource type.
	Type string `pulumi:"type"`
}
