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
    /// Describes the trigger for performing auto scaling operation.
    /// </summary>
    public sealed class AutoScalingTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of auto scaling trigger
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        public AutoScalingTriggerArgs()
        {
        }
    }
}
