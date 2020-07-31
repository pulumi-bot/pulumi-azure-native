// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200401
{
    public static class GetDatabaseAccount
    {
        public static Task<GetDatabaseAccountResult> InvokeAsync(GetDatabaseAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAccountResult>("azurerm:documentdb/v20200401:getDatabaseAccount", args ?? new GetDatabaseAccountArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDatabaseAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseAccountResult
    {
        /// <summary>
        /// Indicates the type of database account. This can only be set at database account creation.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties for the database account.
        /// </summary>
        public readonly Outputs.DatabaseAccountGetPropertiesResponseResult Properties;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly Outputs.TagsResponseResult? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDatabaseAccountResult(
            string? kind,

            string? location,

            string name,

            Outputs.DatabaseAccountGetPropertiesResponseResult properties,

            Outputs.TagsResponseResult? tags,

            string type)
        {
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}