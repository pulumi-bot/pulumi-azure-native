// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20190701
{
    public static class ListRedisKeys
    {
        public static Task<ListRedisKeysResult> InvokeAsync(ListRedisKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListRedisKeysResult>("azurerm:cache/v20190701:listRedisKeys", args ?? new ListRedisKeysArgs(), options.WithVersion());
    }


    public sealed class ListRedisKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListRedisKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListRedisKeysResult
    {
        /// <summary>
        /// The current primary key that clients can use to authenticate with Redis cache.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// The current secondary key that clients can use to authenticate with Redis cache.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListRedisKeysResult(
            string primaryKey,

            string secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}
