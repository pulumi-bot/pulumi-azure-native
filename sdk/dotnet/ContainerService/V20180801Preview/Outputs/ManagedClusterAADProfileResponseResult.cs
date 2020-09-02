// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20180801Preview.Outputs
{

    [OutputType]
    public sealed class ManagedClusterAADProfileResponseResult
    {
        /// <summary>
        /// The client AAD application ID.
        /// </summary>
        public readonly string ClientAppID;
        /// <summary>
        /// The server AAD application ID.
        /// </summary>
        public readonly string ServerAppID;
        /// <summary>
        /// The server AAD application secret.
        /// </summary>
        public readonly string? ServerAppSecret;
        /// <summary>
        /// The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.
        /// </summary>
        public readonly string? TenantID;

        [OutputConstructor]
        private ManagedClusterAADProfileResponseResult(
            string clientAppID,

            string serverAppID,

            string? serverAppSecret,

            string? tenantID)
        {
            ClientAppID = clientAppID;
            ServerAppID = serverAppID;
            ServerAppSecret = serverAppSecret;
            TenantID = tenantID;
        }
    }
}
