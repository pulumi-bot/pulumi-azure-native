// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200710Preview
{
    public static class ListIotHubResourceKeys
    {
        public static Task<ListIotHubResourceKeysResult> InvokeAsync(ListIotHubResourceKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListIotHubResourceKeysResult>("azurerm:devices/v20200710preview:listIotHubResourceKeys", args ?? new ListIotHubResourceKeysArgs(), options.WithVersion());
    }


    public sealed class ListIotHubResourceKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public ListIotHubResourceKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListIotHubResourceKeysResult
    {
        /// <summary>
        /// The next link.
        /// </summary>
        public readonly string NextLink;
        /// <summary>
        /// The list of shared access policies.
        /// </summary>
        public readonly ImmutableArray<Outputs.SharedAccessSignatureAuthorizationRuleResponseResult> Value;

        [OutputConstructor]
        private ListIotHubResourceKeysResult(
            string nextLink,

            ImmutableArray<Outputs.SharedAccessSignatureAuthorizationRuleResponseResult> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
