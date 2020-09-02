// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Maps.V20200201Preview
{
    public static class ListAccountKeys
    {
        public static Task<ListAccountKeysResult> InvokeAsync(ListAccountKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListAccountKeysResult>("azurerm:maps/v20200201preview:listAccountKeys", args ?? new ListAccountKeysArgs(), options.WithVersion());
    }


    public sealed class ListAccountKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Maps Account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListAccountKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListAccountKeysResult
    {
        /// <summary>
        /// The primary key for accessing the Maps REST APIs.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// The secondary key for accessing the Maps REST APIs.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListAccountKeysResult(
            string primaryKey,

            string secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}
