// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701privatepreview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a server maintenance window.
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	// The day of week of the maintenance window to start
	DayOfWeek pulumi.IntOutput `pulumi:"dayOfWeek"`
	// The duration of the maintenance window to run.
	DurationInMinutes pulumi.IntPtrOutput `pulumi:"durationInMinutes"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The starting hour of the maintenance window.
	StartHour pulumi.IntOutput `pulumi:"startHour"`
	// The starting minutes of the maintenance window.
	StartMinute pulumi.IntOutput `pulumi:"startMinute"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil || args.DayOfWeek == nil {
		return nil, errors.New("missing required argument 'DayOfWeek'")
	}
	if args == nil || args.MaintenanceWindowName == nil {
		return nil, errors.New("missing required argument 'MaintenanceWindowName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.StartHour == nil {
		return nil, errors.New("missing required argument 'StartHour'")
	}
	if args == nil || args.StartMinute == nil {
		return nil, errors.New("missing required argument 'StartMinute'")
	}
	if args == nil {
		args = &MaintenanceWindowArgs{}
	}
	var resource MaintenanceWindow
	err := ctx.RegisterResource("azurerm:dbformysql/v20200701privatepreview:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("azurerm:dbformysql/v20200701privatepreview:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
	// The day of week of the maintenance window to start
	DayOfWeek *int `pulumi:"dayOfWeek"`
	// The duration of the maintenance window to run.
	DurationInMinutes *int `pulumi:"durationInMinutes"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The starting hour of the maintenance window.
	StartHour *int `pulumi:"startHour"`
	// The starting minutes of the maintenance window.
	StartMinute *int `pulumi:"startMinute"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type MaintenanceWindowState struct {
	// The day of week of the maintenance window to start
	DayOfWeek pulumi.IntPtrInput
	// The duration of the maintenance window to run.
	DurationInMinutes pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The starting hour of the maintenance window.
	StartHour pulumi.IntPtrInput
	// The starting minutes of the maintenance window.
	StartMinute pulumi.IntPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	// The day of week of the maintenance window to start
	DayOfWeek int `pulumi:"dayOfWeek"`
	// The duration of the maintenance window to run.
	DurationInMinutes *int `pulumi:"durationInMinutes"`
	// The name of the maintenance window.
	MaintenanceWindowName string `pulumi:"maintenanceWindowName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The starting hour of the maintenance window.
	StartHour int `pulumi:"startHour"`
	// The starting minutes of the maintenance window.
	StartMinute int `pulumi:"startMinute"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	// The day of week of the maintenance window to start
	DayOfWeek pulumi.IntInput
	// The duration of the maintenance window to run.
	DurationInMinutes pulumi.IntPtrInput
	// The name of the maintenance window.
	MaintenanceWindowName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The starting hour of the maintenance window.
	StartHour pulumi.IntInput
	// The starting minutes of the maintenance window.
	StartMinute pulumi.IntInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}
