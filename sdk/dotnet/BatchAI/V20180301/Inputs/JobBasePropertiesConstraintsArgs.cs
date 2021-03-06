// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20180301.Inputs
{

    /// <summary>
    /// Constraints associated with the Job.
    /// </summary>
    public sealed class JobBasePropertiesConstraintsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default Value = 1 week.
        /// </summary>
        [Input("maxWallClockTime")]
        public Input<string>? MaxWallClockTime { get; set; }

        public JobBasePropertiesConstraintsArgs()
        {
            MaxWallClockTime = "7.00:00:00";
        }
    }
}
