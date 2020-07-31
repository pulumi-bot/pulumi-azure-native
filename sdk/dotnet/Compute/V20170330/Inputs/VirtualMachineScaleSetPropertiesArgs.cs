// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20170330.Inputs
{

    /// <summary>
    /// Describes the properties of a Virtual Machine Scale Set.
    /// </summary>
    public sealed class VirtualMachineScaleSetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        /// </summary>
        [Input("overprovision")]
        public Input<bool>? Overprovision { get; set; }

        /// <summary>
        /// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        /// </summary>
        [Input("singlePlacementGroup")]
        public Input<bool>? SinglePlacementGroup { get; set; }

        /// <summary>
        /// The upgrade policy.
        /// </summary>
        [Input("upgradePolicy")]
        public Input<Inputs.UpgradePolicyArgs>? UpgradePolicy { get; set; }

        /// <summary>
        /// The virtual machine profile.
        /// </summary>
        [Input("virtualMachineProfile")]
        public Input<Inputs.VirtualMachineScaleSetVMProfileArgs>? VirtualMachineProfile { get; set; }

        public VirtualMachineScaleSetPropertiesArgs()
        {
        }
    }
}