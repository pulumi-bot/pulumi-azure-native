// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20150501Preview
{
    public static class GetSyncAgent
    {
        public static Task<GetSyncAgentResult> InvokeAsync(GetSyncAgentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSyncAgentResult>("azurerm:sql/v20150501preview:getSyncAgent", args ?? new GetSyncAgentArgs(), options.WithVersion());
    }


    public sealed class GetSyncAgentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server on which the sync agent is hosted.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the sync agent.
        /// </summary>
        [Input("syncAgentName", required: true)]
        public string SyncAgentName { get; set; } = null!;

        public GetSyncAgentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSyncAgentResult
    {
        /// <summary>
        /// Expiration time of the sync agent version.
        /// </summary>
        public readonly string ExpiryTime;
        /// <summary>
        /// If the sync agent version is up to date.
        /// </summary>
        public readonly bool IsUpToDate;
        /// <summary>
        /// Last alive time of the sync agent.
        /// </summary>
        public readonly string LastAliveTime;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the sync agent.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// ARM resource id of the sync database in the sync agent.
        /// </summary>
        public readonly string? SyncDatabaseId;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the sync agent.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetSyncAgentResult(
            string expiryTime,

            bool isUpToDate,

            string lastAliveTime,

            string name,

            string state,

            string? syncDatabaseId,

            string type,

            string version)
        {
            ExpiryTime = expiryTime;
            IsUpToDate = isUpToDate;
            LastAliveTime = lastAliveTime;
            Name = name;
            State = state;
            SyncDatabaseId = syncDatabaseId;
            Type = type;
            Version = version;
        }
    }
}
