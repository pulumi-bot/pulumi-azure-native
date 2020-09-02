// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20170301Preview
{
    public static class GetJobAgent
    {
        public static Task<GetJobAgentResult> InvokeAsync(GetJobAgentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobAgentResult>("azurerm:sql/v20170301preview:getJobAgent", args ?? new GetJobAgentArgs(), options.WithVersion());
    }


    public sealed class GetJobAgentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the job agent to be retrieved.
        /// </summary>
        [Input("jobAgentName", required: true)]
        public string JobAgentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetJobAgentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobAgentResult
    {
        /// <summary>
        /// Resource ID of the database to store job metadata in.
        /// </summary>
        public readonly string DatabaseId;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name and tier of the SKU.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// The state of the job agent.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetJobAgentResult(
            string databaseId,

            string location,

            string name,

            Outputs.SkuResponseResult? sku,

            string state,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            DatabaseId = databaseId;
            Location = location;
            Name = name;
            Sku = sku;
            State = state;
            Tags = tags;
            Type = type;
        }
    }
}
