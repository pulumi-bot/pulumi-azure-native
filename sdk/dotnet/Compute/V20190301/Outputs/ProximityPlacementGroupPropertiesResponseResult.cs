// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301.Outputs
{

    [OutputType]
    public sealed class ProximityPlacementGroupPropertiesResponseResult
    {
        /// <summary>
        /// A list of references to all availability sets in the proximity placement group.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> AvailabilitySets;
        /// <summary>
        /// Specifies the type of the proximity placement group. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **Standard** : Co-locate resources within an Azure region or Availability Zone. &lt;br&gt;&lt;br&gt; **Ultra** : For future use.
        /// </summary>
        public readonly string? ProximityPlacementGroupType;
        /// <summary>
        /// A list of references to all virtual machine scale sets in the proximity placement group.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> VirtualMachineScaleSets;
        /// <summary>
        /// A list of references to all virtual machines in the proximity placement group.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> VirtualMachines;

        [OutputConstructor]
        private ProximityPlacementGroupPropertiesResponseResult(
            ImmutableArray<Outputs.SubResourceResponseResult> availabilitySets,

            string? proximityPlacementGroupType,

            ImmutableArray<Outputs.SubResourceResponseResult> virtualMachineScaleSets,

            ImmutableArray<Outputs.SubResourceResponseResult> virtualMachines)
        {
            AvailabilitySets = availabilitySets;
            ProximityPlacementGroupType = proximityPlacementGroupType;
            VirtualMachineScaleSets = virtualMachineScaleSets;
            VirtualMachines = virtualMachines;
        }
    }
}