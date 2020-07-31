// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.V20150819
{
    public static class ListAdminKey
    {
        public static Task<ListAdminKeyResult> InvokeAsync(ListAdminKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListAdminKeyResult>("azurerm:search/v20150819:listAdminKey", args ?? new ListAdminKeyArgs(), options.WithVersion());
    }


    public sealed class ListAdminKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Cognitive Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public string SearchServiceName { get; set; } = null!;

        public ListAdminKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListAdminKeyResult
    {
        /// <summary>
        /// The primary admin API key of the Search service.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// The secondary admin API key of the Search service.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListAdminKeyResult(
            string primaryKey,

            string secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}