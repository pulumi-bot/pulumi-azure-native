// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.Inputs
{

    /// <summary>
    /// Custom MPI job settings.
    /// </summary>
    public sealed class CustomMpiSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command line to be executed by mpi runtime on each compute node.
        /// </summary>
        [Input("commandLine", required: true)]
        public Input<string> CommandLine { get; set; } = null!;

        /// <summary>
        /// Number of processes to launch for the job execution. The default value for this property is equal to nodeCount property
        /// </summary>
        [Input("processCount")]
        public Input<int>? ProcessCount { get; set; }

        public CustomMpiSettingsArgs()
        {
        }
    }
}
