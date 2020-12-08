// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200401.Inputs
{

    /// <summary>
    /// The conflict resolution policy for the container.
    /// </summary>
    public sealed class ConflictResolutionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The conflict resolution path in the case of LastWriterWins mode.
        /// </summary>
        [Input("conflictResolutionPath")]
        public Input<string>? ConflictResolutionPath { get; set; }

        /// <summary>
        /// The procedure to resolve conflicts in the case of custom mode.
        /// </summary>
        [Input("conflictResolutionProcedure")]
        public Input<string>? ConflictResolutionProcedure { get; set; }

        /// <summary>
        /// Indicates the conflict resolution mode.
        /// </summary>
        [Input("mode")]
        public InputUnion<string, Pulumi.AzureNextGen.DocumentDB.V20200401.ConflictResolutionMode>? Mode { get; set; }

        public ConflictResolutionPolicyArgs()
        {
        }
    }
}
