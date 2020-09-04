// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetVMProfileResponseResult
    {
        /// <summary>
        /// The virtual machine scale set extension profile.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetExtensionProfileResponseResult? ExtensionProfile;
        /// <summary>
        /// The virtual machine scale set network profile.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetNetworkProfileResponseResult? NetworkProfile;
        /// <summary>
        /// The virtual machine scale set OS profile.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetOSProfileResponseResult? OsProfile;
        /// <summary>
        /// The virtual machine scale set storage profile.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetStorageProfileResponseResult? StorageProfile;

        [OutputConstructor]
        private VirtualMachineScaleSetVMProfileResponseResult(
            Outputs.VirtualMachineScaleSetExtensionProfileResponseResult? extensionProfile,

            Outputs.VirtualMachineScaleSetNetworkProfileResponseResult? networkProfile,

            Outputs.VirtualMachineScaleSetOSProfileResponseResult? osProfile,

            Outputs.VirtualMachineScaleSetStorageProfileResponseResult? storageProfile)
        {
            ExtensionProfile = extensionProfile;
            NetworkProfile = networkProfile;
            OsProfile = osProfile;
            StorageProfile = storageProfile;
        }
    }
}
