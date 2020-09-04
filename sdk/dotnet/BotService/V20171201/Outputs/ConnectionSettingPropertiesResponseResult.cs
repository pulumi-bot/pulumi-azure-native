// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BotService.V20171201.Outputs
{

    [OutputType]
    public sealed class ConnectionSettingPropertiesResponseResult
    {
        /// <summary>
        /// Client Id associated with the Connection Setting.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Client Secret associated with the Connection Setting
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// Service Provider Parameters associated with the Connection Setting
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionSettingParameterResponseResult> Parameters;
        /// <summary>
        /// Scopes associated with the Connection Setting
        /// </summary>
        public readonly string? Scopes;
        /// <summary>
        /// Service Provider Display Name associated with the Connection Setting
        /// </summary>
        public readonly string? ServiceProviderDisplayName;
        /// <summary>
        /// Service Provider Id associated with the Connection Setting
        /// </summary>
        public readonly string? ServiceProviderId;
        /// <summary>
        /// Setting Id set by the service for the Connection Setting.
        /// </summary>
        public readonly string SettingId;

        [OutputConstructor]
        private ConnectionSettingPropertiesResponseResult(
            string? clientId,

            string? clientSecret,

            ImmutableArray<Outputs.ConnectionSettingParameterResponseResult> parameters,

            string? scopes,

            string? serviceProviderDisplayName,

            string? serviceProviderId,

            string settingId)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            Parameters = parameters;
            Scopes = scopes;
            ServiceProviderDisplayName = serviceProviderDisplayName;
            ServiceProviderId = serviceProviderId;
            SettingId = settingId;
        }
    }
}
