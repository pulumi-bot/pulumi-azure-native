// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20190401.Outputs
{

    [OutputType]
    public sealed class AutoScaleRunResponseResult
    {
        public readonly Outputs.AutoScaleRunErrorResponseResult? Error;
        public readonly string EvaluationTime;
        /// <summary>
        /// Each variable value is returned in the form $variable=value, and variables are separated by semicolons.
        /// </summary>
        public readonly string? Results;

        [OutputConstructor]
        private AutoScaleRunResponseResult(
            Outputs.AutoScaleRunErrorResponseResult? error,

            string evaluationTime,

            string? results)
        {
            Error = error;
            EvaluationTime = evaluationTime;
            Results = results;
        }
    }
}