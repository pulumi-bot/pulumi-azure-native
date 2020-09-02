// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview.Inputs
{

    /// <summary>
    /// Describes the auto scaling policy
    /// </summary>
    public sealed class AutoScalingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mechanism that is used to scale when auto scaling operation is invoked.
        /// </summary>
        [Input("mechanism", required: true)]
        public Input<Inputs.AutoScalingMechanismArgs> Mechanism { get; set; } = null!;

        /// <summary>
        /// The name of the auto scaling policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Determines when auto scaling operation will be invoked.
        /// </summary>
        [Input("trigger", required: true)]
        public Input<Inputs.AutoScalingTriggerArgs> Trigger { get; set; } = null!;

        public AutoScalingPolicyArgs()
        {
        }
    }
}
