// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20180710.Inputs
{

    /// <summary>
    /// VMwareCbt specific enable migration input.
    /// </summary>
    public sealed class VMwareCbtEnableMigrationInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data mover RunAs account Id.
        /// </summary>
        [Input("dataMoverRunAsAccountId", required: true)]
        public Input<string> DataMoverRunAsAccountId { get; set; } = null!;

        [Input("disksToInclude", required: true)]
        private InputList<Inputs.VMwareCbtDiskInputArgs>? _disksToInclude;

        /// <summary>
        /// The disks to include list.
        /// </summary>
        public InputList<Inputs.VMwareCbtDiskInputArgs> DisksToInclude
        {
            get => _disksToInclude ?? (_disksToInclude = new InputList<Inputs.VMwareCbtDiskInputArgs>());
            set => _disksToInclude = value;
        }

        /// <summary>
        /// The class type.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// License type.
        /// </summary>
        [Input("licenseType")]
        public InputUnion<string, Pulumi.AzureNextGen.RecoveryServices.V20180710.LicenseType>? LicenseType { get; set; }

        /// <summary>
        /// The snapshot RunAs account Id.
        /// </summary>
        [Input("snapshotRunAsAccountId", required: true)]
        public Input<string> SnapshotRunAsAccountId { get; set; } = null!;

        /// <summary>
        /// The target availability set ARM Id.
        /// </summary>
        [Input("targetAvailabilitySetId")]
        public Input<string>? TargetAvailabilitySetId { get; set; }

        /// <summary>
        /// The target boot diagnostics storage account ARM Id.
        /// </summary>
        [Input("targetBootDiagnosticsStorageAccountId")]
        public Input<string>? TargetBootDiagnosticsStorageAccountId { get; set; }

        /// <summary>
        /// The target network ARM Id.
        /// </summary>
        [Input("targetNetworkId", required: true)]
        public Input<string> TargetNetworkId { get; set; } = null!;

        /// <summary>
        /// The target resource group ARM Id.
        /// </summary>
        [Input("targetResourceGroupId", required: true)]
        public Input<string> TargetResourceGroupId { get; set; } = null!;

        /// <summary>
        /// The target subnet name.
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
        /// The ARM Id of the VM discovered in VMware.
        /// </summary>
        [Input("vmwareMachineId", required: true)]
        public Input<string> VmwareMachineId { get; set; } = null!;

        public VMwareCbtEnableMigrationInputArgs()
        {
        }
    }
}
