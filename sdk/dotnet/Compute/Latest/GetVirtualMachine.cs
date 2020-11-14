// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.Latest
{
    public static class GetVirtualMachine
    {
        public static Task<GetVirtualMachineResult> InvokeAsync(GetVirtualMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("azure-nextgen:compute/latest:getVirtualMachine", args ?? new GetVirtualMachineArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The expand expression to apply on the operation.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("vmName", required: true)]
        public string VmName { get; set; } = null!;

        public GetVirtualMachineArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// Specifies additional capabilities enabled or disabled on the virtual machine.
        /// </summary>
        public readonly Outputs.AdditionalCapabilitiesResponse? AdditionalCapabilities;
        /// <summary>
        /// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). &lt;br&gt;&lt;br&gt; For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) &lt;br&gt;&lt;br&gt; Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. &lt;br&gt;&lt;br&gt;This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
        /// </summary>
        public readonly Outputs.SubResourceResponse? AvailabilitySet;
        /// <summary>
        /// Specifies the billing related details of a Azure Spot virtual machine. &lt;br&gt;&lt;br&gt;Minimum api-version: 2019-03-01.
        /// </summary>
        public readonly Outputs.BillingProfileResponse? BillingProfile;
        /// <summary>
        /// Specifies the boot diagnostic settings state. &lt;br&gt;&lt;br&gt;Minimum api-version: 2015-06-15.
        /// </summary>
        public readonly Outputs.DiagnosticsProfileResponse? DiagnosticsProfile;
        /// <summary>
        /// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. &lt;br&gt;&lt;br&gt;For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. &lt;br&gt;&lt;br&gt;For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
        /// </summary>
        public readonly string? EvictionPolicy;
        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). &lt;br&gt;&lt;br&gt; Minimum api-version: 2020-06-01
        /// </summary>
        public readonly string? ExtensionsTimeBudget;
        /// <summary>
        /// Specifies the hardware settings for the virtual machine.
        /// </summary>
        public readonly Outputs.HardwareProfileResponse? HardwareProfile;
        /// <summary>
        /// Specifies information about the dedicated host that the virtual machine resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-10-01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? Host;
        /// <summary>
        /// Specifies information about the dedicated host group that the virtual machine resides in. &lt;br&gt;&lt;br&gt;Minimum api-version: 2020-06-01. &lt;br&gt;&lt;br&gt;NOTE: User cannot specify both host and hostGroup properties.
        /// </summary>
        public readonly Outputs.SubResourceResponse? HostGroup;
        /// <summary>
        /// The identity of the virtual machine, if configured.
        /// </summary>
        public readonly Outputs.VirtualMachineIdentityResponse? Identity;
        /// <summary>
        /// The virtual machine instance view.
        /// </summary>
        public readonly Outputs.VirtualMachineInstanceViewResponse InstanceView;
        /// <summary>
        /// Specifies that the image or disk that is being used was licensed on-premises. &lt;br&gt;&lt;br&gt; Possible values for Windows Server operating system are: &lt;br&gt;&lt;br&gt; Windows_Client &lt;br&gt;&lt;br&gt; Windows_Server &lt;br&gt;&lt;br&gt; Possible values for Linux Server operating system are: &lt;br&gt;&lt;br&gt; RHEL_BYOS (for RHEL) &lt;br&gt;&lt;br&gt; SLES_BYOS (for SUSE) &lt;br&gt;&lt;br&gt; For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) &lt;br&gt;&lt;br&gt; [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) &lt;br&gt;&lt;br&gt; Minimum api-version: 2015-06-15
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the network interfaces of the virtual machine.
        /// </summary>
        public readonly Outputs.NetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
        /// </summary>
        public readonly Outputs.OSProfileResponse? OsProfile;
        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        public readonly Outputs.PlanResponse? Plan;
        /// <summary>
        /// Specifies the priority for the virtual machine. &lt;br&gt;&lt;br&gt;Minimum api-version: 2019-03-01
        /// </summary>
        public readonly string? Priority;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Specifies information about the proximity placement group that the virtual machine should be assigned to. &lt;br&gt;&lt;br&gt;Minimum api-version: 2018-04-01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? ProximityPlacementGroup;
        /// <summary>
        /// The virtual machine child extension resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineExtensionResponse> Resources;
        /// <summary>
        /// Specifies the Security related profile settings for the virtual machine.
        /// </summary>
        public readonly Outputs.SecurityProfileResponse? SecurityProfile;
        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        public readonly Outputs.StorageProfileResponse? StorageProfile;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. &lt;br&gt;&lt;br&gt;This property cannot exist along with a non-null properties.availabilitySet reference. &lt;br&gt;&lt;br&gt;Minimum api‐version: 2019‐03‐01
        /// </summary>
        public readonly Outputs.SubResourceResponse? VirtualMachineScaleSet;
        /// <summary>
        /// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
        /// </summary>
        public readonly string VmId;
        /// <summary>
        /// The virtual machine zones.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetVirtualMachineResult(
            Outputs.AdditionalCapabilitiesResponse? additionalCapabilities,

            Outputs.SubResourceResponse? availabilitySet,

            Outputs.BillingProfileResponse? billingProfile,

            Outputs.DiagnosticsProfileResponse? diagnosticsProfile,

            string? evictionPolicy,

            string? extensionsTimeBudget,

            Outputs.HardwareProfileResponse? hardwareProfile,

            Outputs.SubResourceResponse? host,

            Outputs.SubResourceResponse? hostGroup,

            Outputs.VirtualMachineIdentityResponse? identity,

            Outputs.VirtualMachineInstanceViewResponse instanceView,

            string? licenseType,

            string location,

            string name,

            Outputs.NetworkProfileResponse? networkProfile,

            Outputs.OSProfileResponse? osProfile,

            Outputs.PlanResponse? plan,

            string? priority,

            string provisioningState,

            Outputs.SubResourceResponse? proximityPlacementGroup,

            ImmutableArray<Outputs.VirtualMachineExtensionResponse> resources,

            Outputs.SecurityProfileResponse? securityProfile,

            Outputs.StorageProfileResponse? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.SubResourceResponse? virtualMachineScaleSet,

            string vmId,

            ImmutableArray<string> zones)
        {
            AdditionalCapabilities = additionalCapabilities;
            AvailabilitySet = availabilitySet;
            BillingProfile = billingProfile;
            DiagnosticsProfile = diagnosticsProfile;
            EvictionPolicy = evictionPolicy;
            ExtensionsTimeBudget = extensionsTimeBudget;
            HardwareProfile = hardwareProfile;
            Host = host;
            HostGroup = hostGroup;
            Identity = identity;
            InstanceView = instanceView;
            LicenseType = licenseType;
            Location = location;
            Name = name;
            NetworkProfile = networkProfile;
            OsProfile = osProfile;
            Plan = plan;
            Priority = priority;
            ProvisioningState = provisioningState;
            ProximityPlacementGroup = proximityPlacementGroup;
            Resources = resources;
            SecurityProfile = securityProfile;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            VirtualMachineScaleSet = virtualMachineScaleSet;
            VmId = vmId;
            Zones = zones;
        }
    }
}
