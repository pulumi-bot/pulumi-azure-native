// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.Latest.Inputs
{

    /// <summary>
    /// InMageRcm specific enable protection input.
    /// </summary>
    public sealed class InMageRcmEnableProtectionInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default disk input.
        /// </summary>
        [Input("disksDefault")]
        public Input<Inputs.InMageRcmDisksDefaultInputArgs>? DisksDefault { get; set; }

        [Input("disksToInclude")]
        private InputList<Inputs.InMageRcmDiskInputArgs>? _disksToInclude;

        /// <summary>
        /// The disks to include list.
        /// </summary>
        public InputList<Inputs.InMageRcmDiskInputArgs> DisksToInclude
        {
            get => _disksToInclude ?? (_disksToInclude = new InputList<Inputs.InMageRcmDiskInputArgs>());
            set => _disksToInclude = value;
        }

        /// <summary>
        /// The ARM Id of discovered machine.
        /// </summary>
        [Input("fabricDiscoveryMachineId")]
        public Input<string>? FabricDiscoveryMachineId { get; set; }

        /// <summary>
        /// The class type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The license type.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The multi VM group name.
        /// </summary>
        [Input("multiVmGroupName")]
        public Input<string>? MultiVmGroupName { get; set; }

        /// <summary>
        /// The process server Id.
        /// </summary>
        [Input("processServerId")]
        public Input<string>? ProcessServerId { get; set; }

        /// <summary>
        /// The run-as account Id.
        /// </summary>
        [Input("runAsAccountId")]
        public Input<string>? RunAsAccountId { get; set; }

        /// <summary>
        /// The target availability set ARM Id.
        /// </summary>
        [Input("targetAvailabilitySetId")]
        public Input<string>? TargetAvailabilitySetId { get; set; }

        /// <summary>
        /// The target availability zone.
        /// </summary>
        [Input("targetAvailabilityZone")]
        public Input<string>? TargetAvailabilityZone { get; set; }

        /// <summary>
        /// The target boot diagnostics storage account ARM Id.
        /// </summary>
        [Input("targetBootDiagnosticsStorageAccountId")]
        public Input<string>? TargetBootDiagnosticsStorageAccountId { get; set; }

        /// <summary>
        /// The selected target network ARM Id.
        /// </summary>
        [Input("targetNetworkId")]
        public Input<string>? TargetNetworkId { get; set; }

        /// <summary>
        /// The target proximity placement group Id.
        /// </summary>
        [Input("targetProximityPlacementGroupId")]
        public Input<string>? TargetProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The target resource group ARM Id.
        /// </summary>
        [Input("targetResourceGroupId")]
        public Input<string>? TargetResourceGroupId { get; set; }

        /// <summary>
        /// The selected target subnet name.
        /// </summary>
        [Input("targetSubnetName")]
        public Input<string>? TargetSubnetName { get; set; }

        /// <summary>
        /// The target VM name.
        /// </summary>
        [Input("targetVmName")]
        public Input<string>? TargetVmName { get; set; }

        /// <summary>
        /// The target VM size.
        /// </summary>
        [Input("targetVmSize")]
        public Input<string>? TargetVmSize { get; set; }

        /// <summary>
        /// The selected test network ARM Id.
        /// </summary>
        [Input("testNetworkId")]
        public Input<string>? TestNetworkId { get; set; }

        /// <summary>
        /// The selected test subnet name.
        /// </summary>
        [Input("testSubnetName")]
        public Input<string>? TestSubnetName { get; set; }

        public InMageRcmEnableProtectionInputArgs()
        {
        }
    }
}
