// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20170601Preview
{
    public static class GetReplication
    {
        public static Task<GetReplicationResult> InvokeAsync(GetReplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReplicationResult>("azurerm:containerregistry/v20170601preview:getReplication", args ?? new GetReplicationArgs(), options.WithVersion());
    }


    public sealed class GetReplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the replication.
        /// </summary>
        [Input("replicationName", required: true)]
        public string ReplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetReplicationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReplicationResult
    {
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the replication at the time the operation was called.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The status of the replication at the time the operation was called.
        /// </summary>
        public readonly Outputs.StatusResponseResult Status;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetReplicationResult(
            string location,

            string name,

            string provisioningState,

            Outputs.StatusResponseResult status,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Status = status;
            Tags = tags;
            Type = type;
        }
    }
}
