// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.LabServices.Outputs
{

    [OutputType]
    public sealed class LatestOperationResultResponseResult
    {
        /// <summary>
        /// Error code on failure.
        /// </summary>
        public readonly string ErrorCode;
        /// <summary>
        /// The error message.
        /// </summary>
        public readonly string ErrorMessage;
        /// <summary>
        /// The HttpMethod - PUT/POST/DELETE for the operation.
        /// </summary>
        public readonly string HttpMethod;
        /// <summary>
        /// The URL to use to check long-running operation status
        /// </summary>
        public readonly string OperationUrl;
        /// <summary>
        /// Request URI of the operation.
        /// </summary>
        public readonly string RequestUri;
        /// <summary>
        /// The current status of the operation.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private LatestOperationResultResponseResult(
            string errorCode,

            string errorMessage,

            string httpMethod,

            string operationUrl,

            string requestUri,

            string status)
        {
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            HttpMethod = httpMethod;
            OperationUrl = operationUrl;
            RequestUri = requestUri;
            Status = status;
        }
    }
}