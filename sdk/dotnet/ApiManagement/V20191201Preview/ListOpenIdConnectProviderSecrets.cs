// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    public static class ListOpenIdConnectProviderSecrets
    {
        public static Task<ListOpenIdConnectProviderSecretsResult> InvokeAsync(ListOpenIdConnectProviderSecretsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListOpenIdConnectProviderSecretsResult>("azurerm:apimanagement/v20191201preview:listOpenIdConnectProviderSecrets", args ?? new ListOpenIdConnectProviderSecretsArgs(), options.WithVersion());
    }


    public sealed class ListOpenIdConnectProviderSecretsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the OpenID Connect Provider.
        /// </summary>
        [Input("opid", required: true)]
        public string Opid { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public ListOpenIdConnectProviderSecretsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListOpenIdConnectProviderSecretsResult
    {
        /// <summary>
        /// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
        /// </summary>
        public readonly string? ClientSecret;

        [OutputConstructor]
        private ListOpenIdConnectProviderSecretsResult(string? clientSecret)
        {
            ClientSecret = clientSecret;
        }
    }
}
