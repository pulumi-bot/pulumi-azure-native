// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20150408
{
    public static class GetDatabaseAccountSqlDatabase
    {
        public static Task<GetDatabaseAccountSqlDatabaseResult> InvokeAsync(GetDatabaseAccountSqlDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAccountSqlDatabaseResult>("azure-nextgen:documentdb/v20150408:getDatabaseAccountSqlDatabase", args ?? new GetDatabaseAccountSqlDatabaseArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseAccountSqlDatabaseArgs : Pulumi.InvokeArgs
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
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDatabaseAccountSqlDatabaseArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseAccountSqlDatabaseResult
    {
        /// <summary>
        /// A system generated property that specified the addressable path of the collections resource.
        /// </summary>
        public readonly string? Colls;
        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
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
        /// A system generated property that specifies the addressable path of the users resource.
        /// </summary>
        public readonly string? Users;

        [OutputConstructor]
        private GetDatabaseAccountSqlDatabaseResult(
            string? colls,

            string? etag,

            string? location,

            string name,

            string? rid,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? users)
        {
            Colls = colls;
            Etag = etag;
            Location = location;
            Name = name;
            Rid = rid;
            Tags = tags;
            Type = type;
            Users = users;
        }
    }
}
