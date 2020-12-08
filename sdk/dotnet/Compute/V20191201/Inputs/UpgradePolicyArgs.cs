// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20191201.Inputs
{

    /// <summary>
    /// Describes an upgrade policy - automatic, manual, or rolling.
    /// </summary>
    public sealed class UpgradePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration parameters used for performing automatic OS Upgrade.
        /// </summary>
        [Input("automaticOSUpgradePolicy")]
        public Input<Inputs.AutomaticOSUpgradePolicyArgs>? AutomaticOSUpgradePolicy { get; set; }

        /// <summary>
        /// Specifies the mode of an upgrade to virtual machines in the scale set.&lt;br /&gt;&lt;br /&gt; Possible values are:&lt;br /&gt;&lt;br /&gt; **Manual** - You  control the application of updates to virtual machines in the scale set. You do this by using the manualUpgrade action.&lt;br /&gt;&lt;br /&gt; **Automatic** - All virtual machines in the scale set are  automatically updated at the same time.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.AzureNextGen.Compute.V20191201.UpgradeMode>? Mode { get; set; }

        /// <summary>
        /// The configuration parameters used while performing a rolling upgrade.
        /// </summary>
        [Input("rollingUpgradePolicy")]
        public Input<Inputs.RollingUpgradePolicyArgs>? RollingUpgradePolicy { get; set; }

        public UpgradePolicyArgs()
        {
        }
    }
}
