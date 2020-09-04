// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200501Preview.Outputs
{

    [OutputType]
    public sealed class MachineLearningServiceErrorResponseResult
    {
        /// <summary>
        /// The error response.
        /// </summary>
        public readonly Outputs.ErrorResponseResponseResult Error;

        [OutputConstructor]
        private MachineLearningServiceErrorResponseResult(Outputs.ErrorResponseResponseResult error)
        {
            Error = error;
        }
    }
}
