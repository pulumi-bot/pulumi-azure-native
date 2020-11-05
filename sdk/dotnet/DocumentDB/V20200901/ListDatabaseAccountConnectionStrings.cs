// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200901
{
    public static class ListDatabaseAccountConnectionStrings
    {
        public static Task<ListDatabaseAccountConnectionStringsResult> InvokeAsync(ListDatabaseAccountConnectionStringsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListDatabaseAccountConnectionStringsResult>("azure-nextgen:documentdb/v20200901:listDatabaseAccountConnectionStrings", args ?? new ListDatabaseAccountConnectionStringsArgs(), options.WithVersion());
    }


    public sealed class ListDatabaseAccountConnectionStringsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListDatabaseAccountConnectionStringsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListDatabaseAccountConnectionStringsResult
    {
        /// <summary>
        /// An array that contains the connection strings for the Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseAccountConnectionStringResponseResult> ConnectionStrings;

        [OutputConstructor]
        private ListDatabaseAccountConnectionStringsResult(ImmutableArray<Outputs.DatabaseAccountConnectionStringResponseResult> connectionStrings)
        {
            ConnectionStrings = connectionStrings;
        }
    }
}
