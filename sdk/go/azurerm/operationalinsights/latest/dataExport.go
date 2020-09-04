// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The top level data export resource container.
type DataExport struct {
	pulumi.CustomResourceState

	// When ‘true’, all workspace's tables are exported.
	AllTables pulumi.BoolPtrOutput `pulumi:"allTables"`
	// The latest data export rule modification time.
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	// The data export rule ID.
	DataExportId pulumi.StringPtrOutput `pulumi:"dataExportId"`
	// Active when enabled.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName pulumi.StringPtrOutput `pulumi:"eventHubName"`
	// Date and time when the export was last modified.
	LastModifiedDate pulumi.StringPtrOutput `pulumi:"lastModifiedDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames pulumi.StringArrayOutput `pulumi:"tableNames"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataExport registers a new resource with the given unique name, arguments, and options.
func NewDataExport(ctx *pulumi.Context,
	name string, args *DataExportArgs, opts ...pulumi.ResourceOption) (*DataExport, error) {
	if args == nil || args.DataExportName == nil {
		return nil, errors.New("missing required argument 'DataExportName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &DataExportArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:operationalinsights/v20190801preview:DataExport"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200301preview:DataExport"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200801:DataExport"),
		},
	})
	opts = append(opts, aliases)
	var resource DataExport
	err := ctx.RegisterResource("azurerm:operationalinsights/latest:DataExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataExport gets an existing DataExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExportState, opts ...pulumi.ResourceOption) (*DataExport, error) {
	var resource DataExport
	err := ctx.ReadResource("azurerm:operationalinsights/latest:DataExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataExport resources.
type dataExportState struct {
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
	// The name of the resource
	Name *string `pulumi:"name"`
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId *string `pulumi:"resourceId"`
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames []string `pulumi:"tableNames"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type DataExportState struct {
	// When ‘true’, all workspace's tables are exported.
	AllTables pulumi.BoolPtrInput
	// The latest data export rule modification time.
	CreatedDate pulumi.StringPtrInput
	// The data export rule ID.
	DataExportId pulumi.StringPtrInput
	// Active when enabled.
	Enable pulumi.BoolPtrInput
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName pulumi.StringPtrInput
	// Date and time when the export was last modified.
	LastModifiedDate pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId pulumi.StringPtrInput
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames pulumi.StringArrayInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (DataExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExportState)(nil)).Elem()
}

type dataExportArgs struct {
	// When ‘true’, all workspace's tables are exported.
	AllTables *bool `pulumi:"allTables"`
	// The latest data export rule modification time.
	CreatedDate *string `pulumi:"createdDate"`
	// The data export rule ID.
	DataExportId *string `pulumi:"dataExportId"`
	// The data export rule name.
	DataExportName string `pulumi:"dataExportName"`
	// Active when enabled.
	Enable *bool `pulumi:"enable"`
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName *string `pulumi:"eventHubName"`
	// Date and time when the export was last modified.
	LastModifiedDate *string `pulumi:"lastModifiedDate"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId string `pulumi:"resourceId"`
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames []string `pulumi:"tableNames"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataExport resource.
type DataExportArgs struct {
	// When ‘true’, all workspace's tables are exported.
	AllTables pulumi.BoolPtrInput
	// The latest data export rule modification time.
	CreatedDate pulumi.StringPtrInput
	// The data export rule ID.
	DataExportId pulumi.StringPtrInput
	// The data export rule name.
	DataExportName pulumi.StringInput
	// Active when enabled.
	Enable pulumi.BoolPtrInput
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName pulumi.StringPtrInput
	// Date and time when the export was last modified.
	LastModifiedDate pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId pulumi.StringInput
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames pulumi.StringArrayInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DataExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExportArgs)(nil)).Elem()
}
