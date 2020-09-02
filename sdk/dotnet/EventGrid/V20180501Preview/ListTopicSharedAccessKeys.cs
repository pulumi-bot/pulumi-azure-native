// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20180501Preview
{
    public static class ListTopicSharedAccessKeys
    {
        public static Task<ListTopicSharedAccessKeysResult> InvokeAsync(ListTopicSharedAccessKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListTopicSharedAccessKeysResult>("azurerm:eventgrid/v20180501preview:listTopicSharedAccessKeys", args ?? new ListTopicSharedAccessKeysArgs(), options.WithVersion());
    }


    public sealed class ListTopicSharedAccessKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the topic
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public ListTopicSharedAccessKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListTopicSharedAccessKeysResult
    {
        /// <summary>
        /// Shared access key1 for the topic.
        /// </summary>
        public readonly string? Key1;
        /// <summary>
        /// Shared access key2 for the topic.
        /// </summary>
        public readonly string? Key2;

        [OutputConstructor]
        private ListTopicSharedAccessKeysResult(
            string? key1,

            string? key2)
        {
            Key1 = key1;
            Key2 = key2;
        }
    }
}
