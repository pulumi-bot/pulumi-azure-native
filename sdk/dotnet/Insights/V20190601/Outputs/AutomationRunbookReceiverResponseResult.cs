// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20190601.Outputs
{

    [OutputType]
    public sealed class AutomationRunbookReceiverResponseResult
    {
        /// <summary>
        /// The Azure automation account Id which holds this runbook and authenticate to Azure resource.
        /// </summary>
        public readonly string AutomationAccountId;
        /// <summary>
        /// Indicates whether this instance is global runbook.
        /// </summary>
        public readonly bool IsGlobalRunbook;
        /// <summary>
        /// Indicates name of the webhook.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The name for this runbook.
        /// </summary>
        public readonly string RunbookName;
        /// <summary>
        /// The URI where webhooks should be sent.
        /// </summary>
        public readonly string? ServiceUri;
        /// <summary>
        /// Indicates whether to use common alert schema.
        /// </summary>
        public readonly bool UseCommonAlertSchema;
        /// <summary>
        /// The resource id for webhook linked to this runbook.
        /// </summary>
        public readonly string WebhookResourceId;

        [OutputConstructor]
        private AutomationRunbookReceiverResponseResult(
            string automationAccountId,

            bool isGlobalRunbook,

            string? name,

            string runbookName,

            string? serviceUri,

            bool useCommonAlertSchema,

            string webhookResourceId)
        {
            AutomationAccountId = automationAccountId;
            IsGlobalRunbook = isGlobalRunbook;
            Name = name;
            RunbookName = runbookName;
            ServiceUri = serviceUri;
            UseCommonAlertSchema = useCommonAlertSchema;
            WebhookResourceId = webhookResourceId;
        }
    }
}