// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Latest
{
    /// <summary>
    /// Describes a Virtual Machine.
    /// </summary>
    public partial class VirtualMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies additional capabilities enabled or disabled on the virtual machine.
        /// </summary>
        [Output("additionalCapabilities")]
        public Output<Outputs.AdditionalCapabilitiesResponseResult?> AdditionalCapabilities { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). &lt;br&gt;&lt;br&gt; For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) &lt;br&gt;&lt;br&gt; Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. &lt;br&gt;&lt;br&gt;This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
        /// </summary>
        [Output("availabilitySet")]
        public Output<Outputs.SubResourceResponseResult?> AvailabilitySet { get; private set; } = null!;

        /// <summary>
        /// Specifies the billing related details of a Azure Spot virtual machine. &lt;br&gt;&lt;br&gt;Minimum api-version: 2019-03-01.
        /// </summary>
        [Output("billingProfile")]
        public Output<Outputs.BillingProfileResponseResult?> BillingProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the boot diagnostic settings state. &lt;br&gt;&lt;br&gt;Minimum api-version: 2015-06-15.
        /// </summary>
        [Output("diagnosticsProfile")]
        public Output<Outputs.DiagnosticsProfileResponseResult?> DiagnosticsProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. &lt;br&gt;&lt;br&gt;For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. &lt;br&gt;&lt;br&gt;For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
        /// </summary>
        [Output("evictionPolicy")]
        public Output<string?> EvictionPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). &lt;br&gt;&lt;br&gt; Minimum api-version: 2020-06-01
        /// </summary>
        [Output("extensionsTimeBudget")]
        public Output<string?> ExtensionsTimeBudget { get; private set; } = null!;

        /// <summary>
        /// Specifies the hardware settings for the virtual machine.
        /// </summary>
        [Output("hardwareProfile")]
        public Output<Outputs.HardwareProfileResponseResult?> HardwareProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the dedicated host that the virtual machine resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-10-01.
        /// </summary>
        [Output("host")]
        public Output<Outputs.SubResourceResponseResult?> Host { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the dedicated host group that the virtual machine resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2020-06-01. &lt;br&gt;&lt;br&gt;NOTE: User cannot specify both host and hostGroup properties.
        /// </summary>
        [Output("hostGroup")]
        public Output<Outputs.SubResourceResponseResult?> HostGroup { get; private set; } = null!;

        /// <summary>
        /// The identity of the virtual machine, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.VirtualMachineIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The virtual machine instance view.
        /// </summary>
        [Output("instanceView")]
        public Output<Outputs.VirtualMachineInstanceViewResponseResult> InstanceView { get; private set; } = null!;

        /// <summary>
        /// Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; Windows_Client &lt;br&gt;&lt;br&gt; Windows_Server &lt;br&gt;&lt;br&gt; If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. &lt;br&gt;&lt;br&gt; For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) &lt;br&gt;&lt;br&gt; Minimum api-version: 2015-06-15
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the network interfaces of the virtual machine.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.NetworkProfileResponseResult?> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
        /// </summary>
        [Output("osProfile")]
        public Output<Outputs.OSProfileResponseResult?> OsProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.PlanResponseResult?> Plan { get; private set; } = null!;

        /// <summary>
        /// Specifies the priority for the virtual machine. &lt;br&gt;&lt;br&gt;Minimum api-version: 2019-03-01
        /// </summary>
        [Output("priority")]
        public Output<string?> Priority { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the proximity placement group that the virtual machine should be assigned to. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-04-01.
        /// </summary>
        [Output("proximityPlacementGroup")]
        public Output<Outputs.SubResourceResponseResult?> ProximityPlacementGroup { get; private set; } = null!;

        /// <summary>
        /// The virtual machine child extension resources.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.VirtualMachineExtensionResponseResult>> Resources { get; private set; } = null!;

        /// <summary>
        /// Specifies the Security related profile settings for the virtual machine.
        /// </summary>
        [Output("securityProfile")]
        public Output<Outputs.SecurityProfileResponseResult?> SecurityProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        [Output("storageProfile")]
        public Output<Outputs.StorageProfileResponseResult?> StorageProfile { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. &lt;br&gt;&lt;br&gt;This property cannot exist along with a non-null properties.availabilitySet reference. &lt;br&gt;&lt;br&gt;Minimum api‐version: 2019‐03‐01
        /// </summary>
        [Output("virtualMachineScaleSet")]
        public Output<Outputs.SubResourceResponseResult?> VirtualMachineScaleSet { get; private set; } = null!;

        /// <summary>
        /// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
        /// </summary>
        [Output("vmId")]
        public Output<string> VmId { get; private set; } = null!;

        /// <summary>
        /// The virtual machine zones.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/latest:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/latest:VirtualMachine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/v20150615:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20160330:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20160430preview:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20170330:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20171201:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20180401:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20180601:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20181001:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190301:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190701:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20191201:VirtualMachine"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20200601:VirtualMachine"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, options);
        }
    }

    public sealed class VirtualMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies additional capabilities enabled or disabled on the virtual machine.
        /// </summary>
        [Input("additionalCapabilities")]
        public Input<Inputs.AdditionalCapabilitiesArgs>? AdditionalCapabilities { get; set; }

        /// <summary>
        /// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). &lt;br&gt;&lt;br&gt; For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) &lt;br&gt;&lt;br&gt; Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. &lt;br&gt;&lt;br&gt;This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
        /// </summary>
        [Input("availabilitySet")]
        public Input<Inputs.SubResourceArgs>? AvailabilitySet { get; set; }

        /// <summary>
        /// Specifies the billing related details of a Azure Spot virtual machine. &lt;br&gt;&lt;br&gt;Minimum api-version: 2019-03-01.
        /// </summary>
        [Input("billingProfile")]
        public Input<Inputs.BillingProfileArgs>? BillingProfile { get; set; }

        /// <summary>
        /// Specifies the boot diagnostic settings state. &lt;br&gt;&lt;br&gt;Minimum api-version: 2015-06-15.
        /// </summary>
        [Input("diagnosticsProfile")]
        public Input<Inputs.DiagnosticsProfileArgs>? DiagnosticsProfile { get; set; }

        /// <summary>
        /// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. &lt;br&gt;&lt;br&gt;For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. &lt;br&gt;&lt;br&gt;For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
        /// </summary>
        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). &lt;br&gt;&lt;br&gt; Minimum api-version: 2020-06-01
        /// </summary>
        [Input("extensionsTimeBudget")]
        public Input<string>? ExtensionsTimeBudget { get; set; }

        /// <summary>
        /// Specifies the hardware settings for the virtual machine.
        /// </summary>
        [Input("hardwareProfile")]
        public Input<Inputs.HardwareProfileArgs>? HardwareProfile { get; set; }

        /// <summary>
        /// Specifies information about the dedicated host that the virtual machine resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-10-01.
        /// </summary>
        [Input("host")]
        public Input<Inputs.SubResourceArgs>? Host { get; set; }

        /// <summary>
        /// Specifies information about the dedicated host group that the virtual machine resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2020-06-01. &lt;br&gt;&lt;br&gt;NOTE: User cannot specify both host and hostGroup properties.
        /// </summary>
        [Input("hostGroup")]
        public Input<Inputs.SubResourceArgs>? HostGroup { get; set; }

        /// <summary>
        /// The identity of the virtual machine, if configured.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.VirtualMachineIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; Windows_Client &lt;br&gt;&lt;br&gt; Windows_Server &lt;br&gt;&lt;br&gt; If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. &lt;br&gt;&lt;br&gt; For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) &lt;br&gt;&lt;br&gt; Minimum api-version: 2015-06-15
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Specifies the network interfaces of the virtual machine.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.NetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
        /// </summary>
        [Input("osProfile")]
        public Input<Inputs.OSProfileArgs>? OsProfile { get; set; }

        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.PlanArgs>? Plan { get; set; }

        /// <summary>
        /// Specifies the priority for the virtual machine. &lt;br&gt;&lt;br&gt;Minimum api-version: 2019-03-01
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Specifies information about the proximity placement group that the virtual machine should be assigned to. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-04-01.
        /// </summary>
        [Input("proximityPlacementGroup")]
        public Input<Inputs.SubResourceArgs>? ProximityPlacementGroup { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the Security related profile settings for the virtual machine.
        /// </summary>
        [Input("securityProfile")]
        public Input<Inputs.SecurityProfileArgs>? SecurityProfile { get; set; }

        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.StorageProfileArgs>? StorageProfile { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. &lt;br&gt;&lt;br&gt;This property cannot exist along with a non-null properties.availabilitySet reference. &lt;br&gt;&lt;br&gt;Minimum api‐version: 2019‐03‐01
        /// </summary>
        [Input("virtualMachineScaleSet")]
        public Input<Inputs.SubResourceArgs>? VirtualMachineScaleSet { get; set; }

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("vmName", required: true)]
        public Input<string> VmName { get; set; } = null!;

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// The virtual machine zones.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public VirtualMachineArgs()
        {
        }
    }
}
