// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.SecurityInsights.V20190101Preview.Outputs
{

    [OutputType]
    public sealed class SecurityAlertTimelineItemResponseResult
    {
        /// <summary>
        /// The name of the alert type.
        /// </summary>
        public readonly string AlertType;
        /// <summary>
        /// The alert azure resource id.
        /// </summary>
        public readonly string AzureResourceId;
        /// <summary>
        /// The alert name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The alert end time.
        /// </summary>
        public readonly string EndTimeUtc;
        /// <summary>
        /// The entity query kind type.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The alert product name.
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// The alert severity.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The alert start time.
        /// </summary>
        public readonly string StartTimeUtc;
        /// <summary>
        /// The alert generated time.
        /// </summary>
        public readonly string TimeGenerated;

        [OutputConstructor]
        private SecurityAlertTimelineItemResponseResult(
            string alertType,

            string azureResourceId,

            string displayName,

            string endTimeUtc,

            string kind,

            string productName,

            string severity,

            string startTimeUtc,

            string timeGenerated)
        {
            AlertType = alertType;
            AzureResourceId = azureResourceId;
            DisplayName = displayName;
            EndTimeUtc = endTimeUtc;
            Kind = kind;
            ProductName = productName;
            Severity = severity;
            StartTimeUtc = startTimeUtc;
            TimeGenerated = timeGenerated;
        }
    }
}
