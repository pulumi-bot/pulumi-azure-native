// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.Outputs
{

    [OutputType]
    public sealed class NotificationChannelPropertiesResponse
    {
        /// <summary>
        /// The creation date of the notification channel.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// Description of notification.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        /// </summary>
        public readonly string? EmailRecipient;
        /// <summary>
        /// The list of event for which this notification is enabled.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventResponse> Events;
        /// <summary>
        /// The locale to use when sending a notification (fallback for unsupported languages is EN).
        /// </summary>
        public readonly string? NotificationLocale;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string UniqueIdentifier;
        /// <summary>
        /// The webhook URL to send notifications to.
        /// </summary>
        public readonly string? WebHookUrl;

        [OutputConstructor]
        private NotificationChannelPropertiesResponse(
            string createdDate,

            string? description,

            string? emailRecipient,

            ImmutableArray<Outputs.EventResponse> events,

            string? notificationLocale,

            string provisioningState,

            string uniqueIdentifier,

            string? webHookUrl)
        {
            CreatedDate = createdDate;
            Description = description;
            EmailRecipient = emailRecipient;
            Events = events;
            NotificationLocale = notificationLocale;
            ProvisioningState = provisioningState;
            UniqueIdentifier = uniqueIdentifier;
            WebHookUrl = webHookUrl;
        }
    }
}