// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NetApp.V20200701.Inputs
{

    /// <summary>
    /// Hourly Schedule properties
    /// </summary>
    public sealed class HourlyScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates which minute snapshot should be taken
        /// </summary>
        [Input("minute")]
        public Input<int>? Minute { get; set; }

        /// <summary>
        /// Hourly snapshot count to keep
        /// </summary>
        [Input("snapshotsToKeep")]
        public Input<int>? SnapshotsToKeep { get; set; }

        /// <summary>
        /// Resource size in bytes, current storage usage for the volume in bytes
        /// </summary>
        [Input("usedBytes")]
        public Input<int>? UsedBytes { get; set; }

        public HourlyScheduleArgs()
        {
        }
    }
}
