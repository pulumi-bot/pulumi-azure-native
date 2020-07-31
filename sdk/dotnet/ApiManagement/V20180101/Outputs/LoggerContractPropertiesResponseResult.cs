// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180101.Outputs
{

    [OutputType]
    public sealed class LoggerContractPropertiesResponseResult
    {
        /// <summary>
        /// The name and SendRule connection string of the event hub for azureEventHub logger.
        /// Instrumentation key for applicationInsights logger.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Credentials;
        /// <summary>
        /// Logger description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Whether records are buffered in the logger before publishing. Default is assumed to be true.
        /// </summary>
        public readonly bool? IsBuffered;
        /// <summary>
        /// Logger type.
        /// </summary>
        public readonly string LoggerType;

        [OutputConstructor]
        private LoggerContractPropertiesResponseResult(
            ImmutableDictionary<string, string> credentials,

            string? description,

            bool? isBuffered,

            string loggerType)
        {
            Credentials = credentials;
            Description = description;
            IsBuffered = isBuffered;
            LoggerType = loggerType;
        }
    }
}