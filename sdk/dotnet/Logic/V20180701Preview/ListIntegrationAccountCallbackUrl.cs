// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20180701Preview
{
    public static class ListIntegrationAccountCallbackUrl
    {
        public static Task<ListIntegrationAccountCallbackUrlResult> InvokeAsync(ListIntegrationAccountCallbackUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListIntegrationAccountCallbackUrlResult>("azure-nextgen:logic/v20180701preview:listIntegrationAccountCallbackUrl", args ?? new ListIntegrationAccountCallbackUrlArgs(), options.WithVersion());
    }


    public sealed class ListIntegrationAccountCallbackUrlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The key type.
        /// </summary>
        [Input("keyType")]
        public Union<string, Pulumi.AzureNextGen.Logic.V20180701Preview.KeyType>? KeyType { get; set; }

        /// <summary>
        /// The expiry time.
        /// </summary>
        [Input("notAfter")]
        public string? NotAfter { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListIntegrationAccountCallbackUrlArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListIntegrationAccountCallbackUrlResult
    {
        /// <summary>
        /// The URL value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ListIntegrationAccountCallbackUrlResult(string? value)
        {
            Value = value;
        }
    }
}
