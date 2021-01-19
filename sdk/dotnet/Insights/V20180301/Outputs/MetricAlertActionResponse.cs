// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Insights.V20180301.Outputs
{

    [OutputType]
    public sealed class MetricAlertActionResponse
    {
        /// <summary>
        /// the id of the action group to use.
        /// </summary>
        public readonly string? ActionGroupId;
        /// <summary>
        /// This field allows specifying custom properties, which would be appended to the alert payload sent as input to the webhook.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? WebHookProperties;

        [OutputConstructor]
        private MetricAlertActionResponse(
            string? actionGroupId,

            ImmutableDictionary<string, string>? webHookProperties)
        {
            ActionGroupId = actionGroupId;
            WebHookProperties = webHookProperties;
        }
    }
}
