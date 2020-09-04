// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Machine Extension.
type MachineExtension struct {
	pulumi.CustomResourceState

	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The machine extension instance view.
	InstanceView MachineExtensionPropertiesResponseInstanceViewPtrOutput `pulumi:"instanceView"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.MapOutput `pulumi:"protectedSettings"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.MapOutput `pulumi:"settings"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewMachineExtension registers a new resource with the given unique name, arguments, and options.
func NewMachineExtension(ctx *pulumi.Context,
	name string, args *MachineExtensionArgs, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	if args == nil || args.ExtensionName == nil {
		return nil, errors.New("missing required argument 'ExtensionName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MachineExtensionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:hybridcompute/v20190802preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20191212:MachineExtension"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20200730preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20200802:MachineExtension"),
		},
		{
			Type: pulumi.String("azurerm:hybridcompute/v20200815preview:MachineExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineExtension
	err := ctx.RegisterResource("azurerm:hybridcompute/latest:MachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineExtension gets an existing MachineExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineExtensionState, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	var resource MachineExtension
	err := ctx.ReadResource("azurerm:hybridcompute/latest:MachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineExtension resources.
type machineExtensionState struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The machine extension instance view.
	InstanceView *MachineExtensionPropertiesResponseInstanceView `pulumi:"instanceView"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings map[string]interface{} `pulumi:"protectedSettings"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings map[string]interface{} `pulumi:"settings"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

type MachineExtensionState struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The machine extension instance view.
	InstanceView MachineExtensionPropertiesResponseInstanceViewPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.MapInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// Json formatted public settings for the extension.
	Settings pulumi.MapInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
}

func (MachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionState)(nil)).Elem()
}

type machineExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// The name of the machine extension.
	ExtensionName string `pulumi:"extensionName"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the machine where the extension should be created or updated.
	Name string `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings map[string]interface{} `pulumi:"protectedSettings"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings map[string]interface{} `pulumi:"settings"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

// The set of arguments for constructing a MachineExtension resource.
type MachineExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// The name of the machine extension.
	ExtensionName pulumi.StringInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the machine where the extension should be created or updated.
	Name pulumi.StringInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.MapInput
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.MapInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
}

func (MachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionArgs)(nil)).Elem()
}
