// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20170901Preview.Inputs
{

    /// <summary>
    /// Input directory for the job.
    /// </summary>
    public sealed class InputDirectoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// It will be available for the job as an environment variable under AZ_BATCHAI_INPUT_id. The service will also provide the following  environment variable: AZ_BATCHAI_PREV_OUTPUT_Name. The value of the variable will be populated if the job is being retried after a previous failure, otherwise it will be set to nothing.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public InputDirectoryArgs()
        {
        }
    }
}
