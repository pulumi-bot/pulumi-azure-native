// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media
{
    public static class ListMediaServiceStreamingLocatorPaths
    {
        public static Task<ListMediaServiceStreamingLocatorPathsResult> InvokeAsync(ListMediaServiceStreamingLocatorPathsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListMediaServiceStreamingLocatorPathsResult>("azurerm:media:listMediaServiceStreamingLocatorPaths", args ?? new ListMediaServiceStreamingLocatorPathsArgs(), options.WithVersion());
    }


    public sealed class ListMediaServiceStreamingLocatorPathsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Streaming Locator name.
        /// </summary>
        [Input("streamingLocatorName", required: true)]
        public string StreamingLocatorName { get; set; } = null!;

        public ListMediaServiceStreamingLocatorPathsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListMediaServiceStreamingLocatorPathsResult
    {
        /// <summary>
        /// Download Paths supported by current Streaming Locator
        /// </summary>
        public readonly ImmutableArray<string> DownloadPaths;
        /// <summary>
        /// Streaming Paths supported by current Streaming Locator
        /// </summary>
        public readonly ImmutableArray<Outputs.StreamingPathResponseResult> StreamingPaths;

        [OutputConstructor]
        private ListMediaServiceStreamingLocatorPathsResult(
            ImmutableArray<string> downloadPaths,

            ImmutableArray<Outputs.StreamingPathResponseResult> streamingPaths)
        {
            DownloadPaths = downloadPaths;
            StreamingPaths = streamingPaths;
        }
    }
}