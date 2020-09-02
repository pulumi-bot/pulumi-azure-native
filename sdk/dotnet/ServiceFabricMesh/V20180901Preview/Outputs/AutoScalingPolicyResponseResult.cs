// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class AutoScalingPolicyResponseResult
    {
        /// <summary>
        /// The mechanism that is used to scale when auto scaling operation is invoked.
        /// </summary>
        public readonly Outputs.AutoScalingMechanismResponseResult Mechanism;
        /// <summary>
        /// The name of the auto scaling policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Determines when auto scaling operation will be invoked.
        /// </summary>
        public readonly Outputs.AutoScalingTriggerResponseResult Trigger;

        [OutputConstructor]
        private AutoScalingPolicyResponseResult(
            Outputs.AutoScalingMechanismResponseResult mechanism,

            string name,

            Outputs.AutoScalingTriggerResponseResult trigger)
        {
            Mechanism = mechanism;
            Name = name;
            Trigger = trigger;
        }
    }
}
