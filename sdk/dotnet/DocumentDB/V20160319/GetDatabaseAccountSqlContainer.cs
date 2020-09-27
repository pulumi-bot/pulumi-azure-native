// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20160319
{
    public static class GetDatabaseAccountSqlContainer
    {
        public static Task<GetDatabaseAccountSqlContainerResult> InvokeAsync(GetDatabaseAccountSqlContainerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAccountSqlContainerResult>("azure-nextgen:documentdb/v20160319:getDatabaseAccountSqlContainer", args ?? new GetDatabaseAccountSqlContainerArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseAccountSqlContainerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB container name.
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDatabaseAccountSqlContainerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseAccountSqlContainerResult
    {
        /// <summary>
        /// The conflict resolution policy for the container.
        /// </summary>
        public readonly Outputs.ConflictResolutionPolicyResponse? ConflictResolutionPolicy;
        /// <summary>
        /// Default time to live
        /// </summary>
        public readonly int? DefaultTtl;
        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
        /// </summary>
        public readonly Outputs.IndexingPolicyResponse? IndexingPolicy;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The configuration of the partition key to be used for partitioning data into multiple partitions
        /// </summary>
        public readonly Outputs.ContainerPartitionKeyResponse? PartitionKey;
        /// <summary>
        /// A system generated property. A unique identifier.
        /// </summary>
        public readonly string? Rid;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
        /// </summary>
        public readonly Outputs.UniqueKeyPolicyResponse? UniqueKeyPolicy;

        [OutputConstructor]
        private GetDatabaseAccountSqlContainerResult(
            Outputs.ConflictResolutionPolicyResponse? conflictResolutionPolicy,

            int? defaultTtl,

            string? etag,

            Outputs.IndexingPolicyResponse? indexingPolicy,

            string? location,

            string name,

            Outputs.ContainerPartitionKeyResponse? partitionKey,

            string? rid,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.UniqueKeyPolicyResponse? uniqueKeyPolicy)
        {
            ConflictResolutionPolicy = conflictResolutionPolicy;
            DefaultTtl = defaultTtl;
            Etag = etag;
            IndexingPolicy = indexingPolicy;
            Location = location;
            Name = name;
            PartitionKey = partitionKey;
            Rid = rid;
            Tags = tags;
            Type = type;
            UniqueKeyPolicy = uniqueKeyPolicy;
        }
    }
}
