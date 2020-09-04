// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview
{
    public static class GetSqlResourceSqlDatabase
    {
        public static Task<GetSqlResourceSqlDatabaseResult> InvokeAsync(GetSqlResourceSqlDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSqlResourceSqlDatabaseResult>("azurerm:documentdb/v20200601preview:getSqlResourceSqlDatabase", args ?? new GetSqlResourceSqlDatabaseArgs(), options.WithVersion());
    }


    public sealed class GetSqlResourceSqlDatabaseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSqlResourceSqlDatabaseArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSqlResourceSqlDatabaseResult
    {
        /// <summary>
        /// Identity for the resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponseResult? Identity;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.SqlDatabaseGetPropertiesResponseOptionsResult? Options;
        public readonly Outputs.SqlDatabaseGetPropertiesResponseResourceResult? Resource;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSqlResourceSqlDatabaseResult(
            Outputs.ManagedServiceIdentityResponseResult? identity,

            string? location,

            string name,

            Outputs.SqlDatabaseGetPropertiesResponseOptionsResult? options,

            Outputs.SqlDatabaseGetPropertiesResponseResourceResult? resource,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Options = options;
            Resource = resource;
            Tags = tags;
            Type = type;
        }
    }
}
