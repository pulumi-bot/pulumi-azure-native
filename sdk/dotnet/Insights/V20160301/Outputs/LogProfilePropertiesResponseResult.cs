// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20160301.Outputs
{

    [OutputType]
    public sealed class LogProfilePropertiesResponseResult
    {
        /// <summary>
        /// the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
        /// </summary>
        public readonly ImmutableArray<string> Categories;
        /// <summary>
        /// List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// the retention policy for the events in the log.
        /// </summary>
        public readonly Outputs.RetentionPolicyResponseResult RetentionPolicy;
        /// <summary>
        /// The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
        /// </summary>
        public readonly string? ServiceBusRuleId;
        /// <summary>
        /// the resource id of the storage account to which you would like to send the Activity Log.
        /// </summary>
        public readonly string? StorageAccountId;

        [OutputConstructor]
        private LogProfilePropertiesResponseResult(
            ImmutableArray<string> categories,

            ImmutableArray<string> locations,

            Outputs.RetentionPolicyResponseResult retentionPolicy,

            string? serviceBusRuleId,

            string? storageAccountId)
        {
            Categories = categories;
            Locations = locations;
            RetentionPolicy = retentionPolicy;
            ServiceBusRuleId = serviceBusRuleId;
            StorageAccountId = storageAccountId;
        }
    }
}