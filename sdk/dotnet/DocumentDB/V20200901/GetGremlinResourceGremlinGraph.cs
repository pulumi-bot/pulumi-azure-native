// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200901
{
    public static class GetGremlinResourceGremlinGraph
    {
        public static Task<GetGremlinResourceGremlinGraphResult> InvokeAsync(GetGremlinResourceGremlinGraphArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGremlinResourceGremlinGraphResult>("azure-nextgen:documentdb/v20200901:getGremlinResourceGremlinGraph", args ?? new GetGremlinResourceGremlinGraphArgs(), options.WithVersion());
    }


    public sealed class GetGremlinResourceGremlinGraphArgs : Pulumi.InvokeArgs
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
        /// Cosmos DB graph name.
        /// </summary>
        [Input("graphName", required: true)]
        public string GraphName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGremlinResourceGremlinGraphArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGremlinResourceGremlinGraphResult
    {
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.GremlinGraphGetPropertiesResponseOptions? Options;
        public readonly Outputs.GremlinGraphGetPropertiesResponseResource? Resource;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGremlinResourceGremlinGraphResult(
            string? location,

            string name,

            Outputs.GremlinGraphGetPropertiesResponseOptions? options,

            Outputs.GremlinGraphGetPropertiesResponseResource? resource,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Options = options;
            Resource = resource;
            Tags = tags;
            Type = type;
        }
    }
}
