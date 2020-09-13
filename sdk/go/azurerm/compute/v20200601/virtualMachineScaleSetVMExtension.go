// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Virtual Machine Extension.
//
// ## Example Usage
// ### Create VirtualMachineScaleSet VM extension.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewVirtualMachineScaleSetVMExtension(ctx, "virtualMachineScaleSetVMExtension", &compute.VirtualMachineScaleSetVMExtensionArgs{
// 			AutoUpgradeMinorVersion: pulumi.Bool(true),
// 			InstanceId:              pulumi.String("0"),
// 			Location:                pulumi.String("westus"),
// 			Publisher:               pulumi.String("extPublisher"),
// 			ResourceGroupName:       pulumi.String("myResourceGroup"),
// 			Settings: pulumi.StringMap{
// 				"UserName": pulumi.String("xyz@microsoft.com"),
// 			},
// 			Type:               pulumi.String("extType"),
// 			TypeHandlerVersion: pulumi.String("1.2"),
// 			VmExtensionName:    pulumi.String("myVMExtension"),
// 			VmScaleSetName:     pulumi.String("myvmScaleSet"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type VirtualMachineScaleSetVMExtension struct {
	pulumi.CustomResourceState

	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewResponsePtrOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.MapOutput `pulumi:"protectedSettings"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.MapOutput `pulumi:"settings"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewVirtualMachineScaleSetVMExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetVMExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMExtension, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VmExtensionName == nil {
		return nil, errors.New("missing required argument 'VmExtensionName'")
	}
	if args == nil || args.VmScaleSetName == nil {
		return nil, errors.New("missing required argument 'VmScaleSetName'")
	}
	if args == nil {
		args = &VirtualMachineScaleSetVMExtensionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/latest:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190701:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20191201:VirtualMachineScaleSetVMExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetVMExtension
	err := ctx.RegisterResource("azurerm:compute/v20200601:VirtualMachineScaleSetVMExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetVMExtension gets an existing VirtualMachineScaleSetVMExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetVMExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMExtension, error) {
	var resource VirtualMachineScaleSetVMExtension
	err := ctx.ReadResource("azurerm:compute/v20200601:VirtualMachineScaleSetVMExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetVMExtension resources.
type virtualMachineScaleSetVMExtensionState struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The virtual machine extension instance view.
	InstanceView *VirtualMachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings map[string]interface{} `pulumi:"protectedSettings"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings map[string]interface{} `pulumi:"settings"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

type VirtualMachineScaleSetVMExtensionState struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.MapInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// Json formatted public settings for the extension.
	Settings pulumi.MapInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
}

func (VirtualMachineScaleSetVMExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetVMExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The instance ID of the virtual machine.
	InstanceId string `pulumi:"instanceId"`
	// The virtual machine extension instance view.
	InstanceView *VirtualMachineExtensionInstanceView `pulumi:"instanceView"`
	// Resource location
	Location string `pulumi:"location"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings map[string]interface{} `pulumi:"protectedSettings"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings map[string]interface{} `pulumi:"settings"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The name of the virtual machine extension.
	VmExtensionName string `pulumi:"vmExtensionName"`
	// The name of the VM scale set.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// The set of arguments for constructing a VirtualMachineScaleSetVMExtension resource.
type VirtualMachineScaleSetVMExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The instance ID of the virtual machine.
	InstanceId pulumi.StringInput
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewPtrInput
	// Resource location
	Location pulumi.StringInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.MapInput
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.MapInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
	// The name of the virtual machine extension.
	VmExtensionName pulumi.StringInput
	// The name of the VM scale set.
	VmScaleSetName pulumi.StringInput
}

func (VirtualMachineScaleSetVMExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMExtensionArgs)(nil)).Elem()
}
