// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class JobErrorResponseResult
    {
        /// <summary>
        /// Helps with categorization of errors.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Error code describing the error.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// An array of details about specific errors that led to this reported error.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobErrorDetailResponseResult> Details;
        /// <summary>
        /// A human-readable language-dependent representation of the error.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Indicates that it may be possible to retry the Job. If retry is unsuccessful, please contact Azure support via Azure Portal.
        /// </summary>
        public readonly string Retry;

        [OutputConstructor]
        private JobErrorResponseResult(
            string category,

            string code,

            ImmutableArray<Outputs.JobErrorDetailResponseResult> details,

            string message,

            string retry)
        {
            Category = category;
            Code = code;
            Details = details;
            Message = message;
            Retry = retry;
        }
    }
}
