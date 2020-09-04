// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20200101Preview
{
    public static class GetSecurityContact
    {
        public static Task<GetSecurityContactResult> InvokeAsync(GetSecurityContactArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityContactResult>("azurerm:security/v20200101preview:getSecurityContact", args ?? new GetSecurityContactArgs(), options.WithVersion());
    }


    public sealed class GetSecurityContactArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the security contact object
        /// </summary>
        [Input("securityContactName", required: true)]
        public string SecurityContactName { get; set; } = null!;

        public GetSecurityContactArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecurityContactResult
    {
        /// <summary>
        /// Defines whether to send email notifications about new security alerts
        /// </summary>
        public readonly Outputs.SecurityContactPropertiesResponseAlertNotificationsResult? AlertNotifications;
        /// <summary>
        /// List of email addresses which will get notifications from Azure Security Center by the configurations defined in this security contact.
        /// </summary>
        public readonly string? Emails;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines whether to send email notifications from Azure Security Center to persons with specific RBAC roles on the subscription.
        /// </summary>
        public readonly Outputs.SecurityContactPropertiesResponseNotificationsByRoleResult? NotificationsByRole;
        /// <summary>
        /// The security contact's phone number
        /// </summary>
        public readonly string? Phone;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSecurityContactResult(
            Outputs.SecurityContactPropertiesResponseAlertNotificationsResult? alertNotifications,

            string? emails,

            string name,

            Outputs.SecurityContactPropertiesResponseNotificationsByRoleResult? notificationsByRole,

            string? phone,

            string type)
        {
            AlertNotifications = alertNotifications;
            Emails = emails;
            Name = name;
            NotificationsByRole = notificationsByRole;
            Phone = phone;
            Type = type;
        }
    }
}
