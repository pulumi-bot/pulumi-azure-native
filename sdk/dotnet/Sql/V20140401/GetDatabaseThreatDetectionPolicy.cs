// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Sql.V20140401
{
    public static class GetDatabaseThreatDetectionPolicy
    {
        public static Task<GetDatabaseThreatDetectionPolicyResult> InvokeAsync(GetDatabaseThreatDetectionPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseThreatDetectionPolicyResult>("azure-nextgen:sql/v20140401:getDatabaseThreatDetectionPolicy", args ?? new GetDatabaseThreatDetectionPolicyArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseThreatDetectionPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database for which database Threat Detection policy is defined.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the security alert policy.
        /// </summary>
        [Input("securityAlertPolicyName", required: true)]
        public string SecurityAlertPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetDatabaseThreatDetectionPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseThreatDetectionPolicyResult
    {
        /// <summary>
        /// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
        /// </summary>
        public readonly string? DisabledAlerts;
        /// <summary>
        /// Specifies that the alert is sent to the account administrators.
        /// </summary>
        public readonly string? EmailAccountAdmins;
        /// <summary>
        /// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
        /// </summary>
        public readonly string? EmailAddresses;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource kind.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the number of days to keep in the Threat Detection audit logs.
        /// </summary>
        public readonly int? RetentionDays;
        /// <summary>
        /// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Specifies the identifier key of the Threat Detection audit storage account. If state is Enabled, storageAccountAccessKey is required.
        /// </summary>
        public readonly string? StorageAccountAccessKey;
        /// <summary>
        /// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
        /// </summary>
        public readonly string? StorageEndpoint;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies whether to use the default server policy.
        /// </summary>
        public readonly string? UseServerDefault;

        [OutputConstructor]
        private GetDatabaseThreatDetectionPolicyResult(
            string? disabledAlerts,

            string? emailAccountAdmins,

            string? emailAddresses,

            string id,

            string kind,

            string? location,

            string name,

            int? retentionDays,

            string state,

            string? storageAccountAccessKey,

            string? storageEndpoint,

            string type,

            string? useServerDefault)
        {
            DisabledAlerts = disabledAlerts;
            EmailAccountAdmins = emailAccountAdmins;
            EmailAddresses = emailAddresses;
            Id = id;
            Kind = kind;
            Location = location;
            Name = name;
            RetentionDays = retentionDays;
            State = state;
            StorageAccountAccessKey = storageAccountAccessKey;
            StorageEndpoint = storageEndpoint;
            Type = type;
            UseServerDefault = useServerDefault;
        }
    }
}
