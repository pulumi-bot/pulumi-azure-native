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
    /// VMware Azure specific enable protection input.
    /// </summary>
    public sealed class InMageAzureV2EnableProtectionInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DiskEncryptionSet ARM ID.
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// The DiskType.
        /// </summary>
        [Input("diskType")]
        public InputUnion<string, Pulumi.AzureNextGen.RecoveryServices.V20180710.DiskAccountType>? DiskType { get; set; }

        [Input("disksToInclude")]
        private InputList<Inputs.InMageAzureV2DiskInputDetailsArgs>? _disksToInclude;

        /// <summary>
        /// The disks to include list.
        /// </summary>
        public InputList<Inputs.InMageAzureV2DiskInputDetailsArgs> DisksToInclude
        {
            get => _disksToInclude ?? (_disksToInclude = new InputList<Inputs.InMageAzureV2DiskInputDetailsArgs>());
            set => _disksToInclude = value;
        }

        /// <summary>
        /// The selected option to enable RDP\SSH on target vm after failover. String value of {SrsDataContract.EnableRDPOnTargetOption} enum.
        /// </summary>
        [Input("enableRdpOnTargetOption")]
        public Input<string>? EnableRdpOnTargetOption { get; set; }

        /// <summary>
        /// The class type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The storage account to be used for logging during replication.
        /// </summary>
        [Input("logStorageAccountId")]
        public Input<string>? LogStorageAccountId { get; set; }

        /// <summary>
        /// The Master target Id.
        /// </summary>
        [Input("masterTargetId")]
        public Input<string>? MasterTargetId { get; set; }

        /// <summary>
        /// The multi vm group Id.
        /// </summary>
        [Input("multiVmGroupId")]
        public Input<string>? MultiVmGroupId { get; set; }

        /// <summary>
        /// The multi vm group name.
        /// </summary>
        [Input("multiVmGroupName")]
        public Input<string>? MultiVmGroupName { get; set; }

        /// <summary>
        /// The Process Server Id.
        /// </summary>
        [Input("processServerId")]
        public Input<string>? ProcessServerId { get; set; }

        /// <summary>
        /// The CS account Id.
        /// </summary>
        [Input("runAsAccountId")]
        public Input<string>? RunAsAccountId { get; set; }

        /// <summary>
        /// The storage account name.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        /// <summary>
        /// The target availability zone.
        /// </summary>
        [Input("targetAvailabilityZone")]
        public Input<string>? TargetAvailabilityZone { get; set; }

        /// <summary>
        /// The selected target Azure network Id.
        /// </summary>
        [Input("targetAzureNetworkId")]
        public Input<string>? TargetAzureNetworkId { get; set; }

        /// <summary>
        /// The selected target Azure subnet Id.
        /// </summary>
        [Input("targetAzureSubnetId")]
        public Input<string>? TargetAzureSubnetId { get; set; }

        /// <summary>
        /// The Id of the target resource group (for classic deployment) in which the failover VM is to be created.
        /// </summary>
        [Input("targetAzureV1ResourceGroupId")]
        public Input<string>? TargetAzureV1ResourceGroupId { get; set; }

        /// <summary>
        /// The Id of the target resource group (for resource manager deployment) in which the failover VM is to be created.
        /// </summary>
        [Input("targetAzureV2ResourceGroupId")]
        public Input<string>? TargetAzureV2ResourceGroupId { get; set; }

        /// <summary>
        /// The target azure Vm Name.
        /// </summary>
        [Input("targetAzureVmName")]
        public Input<string>? TargetAzureVmName { get; set; }

        /// <summary>
        /// The proximity placement group ARM Id.
        /// </summary>
        [Input("targetProximityPlacementGroupId")]
        public Input<string>? TargetProximityPlacementGroupId { get; set; }

        public InMageAzureV2EnableProtectionInputArgs()
        {
        }
    }
}
