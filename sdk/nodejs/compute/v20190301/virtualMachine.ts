// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Describes a Virtual Machine.
 *
 * ## Example Usage
 * ### Create a custom-image vm from an unmanaged generalized os image.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             image: {
 *                 uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd",
 *             },
 *             name: "myVMosdisk",
 *             osType: "Windows",
 *             vhd: {
 *                 uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
 *             },
 *         },
 *     },
 *     vmName: "{vm-name}",
 * });
 *
 * ```
 * ### Create a platform-image vm with unmanaged os and data disks.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D2_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         dataDisks: [
 *             {
 *                 createOption: "Empty",
 *                 diskSizeGB: 1023,
 *                 lun: 0,
 *                 vhd: {
 *                     uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd",
 *                 },
 *             },
 *             {
 *                 createOption: "Empty",
 *                 diskSizeGB: 1023,
 *                 lun: 1,
 *                 vhd: {
 *                     uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd",
 *                 },
 *             },
 *         ],
 *         imageReference: {
 *             offer: "WindowsServer",
 *             publisher: "MicrosoftWindowsServer",
 *             sku: "2016-Datacenter",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             name: "myVMosdisk",
 *             vhd: {
 *                 uri: "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
 *             },
 *         },
 *     },
 *     vmName: "{vm-name}",
 * });
 *
 * ```
 * ### Create a vm from a custom image.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm in an availability set.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     availabilitySet: {
 *         id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/{existing-availability-set-name}",
 *     },
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "WindowsServer",
 *             publisher: "MicrosoftWindowsServer",
 *             sku: "2016-Datacenter",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with a marketplace image plan.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     plan: {
 *         name: "windows2016",
 *         product: "windows-data-science-vm",
 *         publisher: "microsoft-ads",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "windows-data-science-vm",
 *             publisher: "microsoft-ads",
 *             sku: "windows2016",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with boot diagnostics.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     diagnosticsProfile: {
 *         bootDiagnostics: {
 *             enabled: true,
 *             storageUri: "http://{existing-storage-account-name}.blob.core.windows.net",
 *         },
 *     },
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "WindowsServer",
 *             publisher: "MicrosoftWindowsServer",
 *             sku: "2016-Datacenter",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with empty data disks.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D2_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         dataDisks: [
 *             {
 *                 createOption: "Empty",
 *                 diskSizeGB: 1023,
 *                 lun: 0,
 *             },
 *             {
 *                 createOption: "Empty",
 *                 diskSizeGB: 1023,
 *                 lun: 1,
 *             },
 *         ],
 *         imageReference: {
 *             offer: "WindowsServer",
 *             publisher: "MicrosoftWindowsServer",
 *             sku: "2016-Datacenter",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with ephemeral os disk.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_DS1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     plan: {
 *         name: "windows2016",
 *         product: "windows-data-science-vm",
 *         publisher: "microsoft-ads",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "windows-data-science-vm",
 *             publisher: "microsoft-ads",
 *             sku: "windows2016",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadOnly",
 *             createOption: "FromImage",
 *             diffDiskSettings: {
 *                 option: "Local",
 *             },
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with password authentication.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "WindowsServer",
 *             publisher: "MicrosoftWindowsServer",
 *             sku: "2016-Datacenter",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with premium storage.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminPassword: "{your-password}",
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "WindowsServer",
 *             publisher: "MicrosoftWindowsServer",
 *             sku: "2016-Datacenter",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Premium_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 * ### Create a vm with ssh authentication.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.compute.v20190301.VirtualMachine("virtualMachine", {
 *     hardwareProfile: {
 *         vmSize: "Standard_D1_v2",
 *     },
 *     location: "westus",
 *     networkProfile: {
 *         networkInterfaces: [{
 *             id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
 *         }],
 *     },
 *     osProfile: {
 *         adminUsername: "{your-username}",
 *         computerName: "myVM",
 *         linuxConfiguration: {
 *             disablePasswordAuthentication: true,
 *             ssh: {
 *                 publicKeys: [{
 *                     keyData: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1",
 *                     path: "/home/{your-username}/.ssh/authorized_keys",
 *                 }],
 *             },
 *         },
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         imageReference: {
 *             offer: "{image_offer}",
 *             publisher: "{image_publisher}",
 *             sku: "{image_sku}",
 *             version: "latest",
 *         },
 *         osDisk: {
 *             caching: "ReadWrite",
 *             createOption: "FromImage",
 *             managedDisk: {
 *                 storageAccountType: "Standard_LRS",
 *             },
 *             name: "myVMosdisk",
 *         },
 *     },
 *     vmName: "myVM",
 * });
 *
 * ```
 */
export class VirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachine {
        return new VirtualMachine(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20190301:VirtualMachine';

    /**
     * Returns true if the given object is an instance of VirtualMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachine.__pulumiType;
    }

    /**
     * Specifies additional capabilities enabled or disabled on the virtual machine.
     */
    public readonly additionalCapabilities!: pulumi.Output<outputs.compute.v20190301.AdditionalCapabilitiesResponse | undefined>;
    /**
     * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
     */
    public readonly availabilitySet!: pulumi.Output<outputs.compute.v20190301.SubResourceResponse | undefined>;
    /**
     * Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
     */
    public readonly billingProfile!: pulumi.Output<outputs.compute.v20190301.BillingProfileResponse | undefined>;
    /**
     * Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
     */
    public readonly diagnosticsProfile!: pulumi.Output<outputs.compute.v20190301.DiagnosticsProfileResponse | undefined>;
    /**
     * Specifies the eviction policy for the Azure Spot virtual machine. Only supported value is 'Deallocate'. <br><br>Minimum api-version: 2019-03-01
     */
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * Specifies the hardware settings for the virtual machine.
     */
    public readonly hardwareProfile!: pulumi.Output<outputs.compute.v20190301.HardwareProfileResponse | undefined>;
    /**
     * Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
     */
    public readonly host!: pulumi.Output<outputs.compute.v20190301.SubResourceResponse | undefined>;
    /**
     * The identity of the virtual machine, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.compute.v20190301.VirtualMachineIdentityResponse | undefined>;
    /**
     * The virtual machine instance view.
     */
    public /*out*/ readonly instanceView!: pulumi.Output<outputs.compute.v20190301.VirtualMachineInstanceViewResponse>;
    /**
     * Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the network interfaces of the virtual machine.
     */
    public readonly networkProfile!: pulumi.Output<outputs.compute.v20190301.NetworkProfileResponse | undefined>;
    /**
     * Specifies the operating system settings for the virtual machine.
     */
    public readonly osProfile!: pulumi.Output<outputs.compute.v20190301.OSProfileResponse | undefined>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    public readonly plan!: pulumi.Output<outputs.compute.v20190301.PlanResponse | undefined>;
    /**
     * Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
     */
    public readonly proximityPlacementGroup!: pulumi.Output<outputs.compute.v20190301.SubResourceResponse | undefined>;
    /**
     * The virtual machine child extension resources.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.compute.v20190301.VirtualMachineExtensionResponse[]>;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    public readonly storageProfile!: pulumi.Output<outputs.compute.v20190301.StorageProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
     */
    public readonly virtualMachineScaleSet!: pulumi.Output<outputs.compute.v20190301.SubResourceResponse | undefined>;
    /**
     * Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
     */
    public /*out*/ readonly vmId!: pulumi.Output<string>;
    /**
     * The virtual machine zones.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualMachineArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vmName === undefined) {
                throw new Error("Missing required property 'vmName'");
            }
            inputs["additionalCapabilities"] = args ? args.additionalCapabilities : undefined;
            inputs["availabilitySet"] = args ? args.availabilitySet : undefined;
            inputs["billingProfile"] = args ? args.billingProfile : undefined;
            inputs["diagnosticsProfile"] = args ? args.diagnosticsProfile : undefined;
            inputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            inputs["hardwareProfile"] = args ? args.hardwareProfile : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["osProfile"] = args ? args.osProfile : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["proximityPlacementGroup"] = args ? args.proximityPlacementGroup : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachineScaleSet"] = args ? args.virtualMachineScaleSet : undefined;
            inputs["vmName"] = args ? args.vmName : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["instanceView"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resources"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:VirtualMachine" }, { type: "azurerm:compute/v20150615:VirtualMachine" }, { type: "azurerm:compute/v20160330:VirtualMachine" }, { type: "azurerm:compute/v20170330:VirtualMachine" }, { type: "azurerm:compute/v20171201:VirtualMachine" }, { type: "azurerm:compute/v20180401:VirtualMachine" }, { type: "azurerm:compute/v20180601:VirtualMachine" }, { type: "azurerm:compute/v20181001:VirtualMachine" }, { type: "azurerm:compute/v20190701:VirtualMachine" }, { type: "azurerm:compute/v20191201:VirtualMachine" }, { type: "azurerm:compute/v20200601:VirtualMachine" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualMachine.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    /**
     * Specifies additional capabilities enabled or disabled on the virtual machine.
     */
    readonly additionalCapabilities?: pulumi.Input<inputs.compute.v20190301.AdditionalCapabilities>;
    /**
     * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
     */
    readonly availabilitySet?: pulumi.Input<inputs.compute.v20190301.SubResource>;
    /**
     * Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
     */
    readonly billingProfile?: pulumi.Input<inputs.compute.v20190301.BillingProfile>;
    /**
     * Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
     */
    readonly diagnosticsProfile?: pulumi.Input<inputs.compute.v20190301.DiagnosticsProfile>;
    /**
     * Specifies the eviction policy for the Azure Spot virtual machine. Only supported value is 'Deallocate'. <br><br>Minimum api-version: 2019-03-01
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * Specifies the hardware settings for the virtual machine.
     */
    readonly hardwareProfile?: pulumi.Input<inputs.compute.v20190301.HardwareProfile>;
    /**
     * Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
     */
    readonly host?: pulumi.Input<inputs.compute.v20190301.SubResource>;
    /**
     * The identity of the virtual machine, if configured.
     */
    readonly identity?: pulumi.Input<inputs.compute.v20190301.VirtualMachineIdentity>;
    /**
     * Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies the network interfaces of the virtual machine.
     */
    readonly networkProfile?: pulumi.Input<inputs.compute.v20190301.NetworkProfile>;
    /**
     * Specifies the operating system settings for the virtual machine.
     */
    readonly osProfile?: pulumi.Input<inputs.compute.v20190301.OSProfile>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: pulumi.Input<inputs.compute.v20190301.Plan>;
    /**
     * Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
     */
    readonly proximityPlacementGroup?: pulumi.Input<inputs.compute.v20190301.SubResource>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    readonly storageProfile?: pulumi.Input<inputs.compute.v20190301.StorageProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
     */
    readonly virtualMachineScaleSet?: pulumi.Input<inputs.compute.v20190301.SubResource>;
    /**
     * The name of the virtual machine.
     */
    readonly vmName: pulumi.Input<string>;
    /**
     * The virtual machine zones.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
