// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CognitiveServices.V20160201Preview
{
    public static class ListCognitiveServicesAccountKeys
    {
        public static Task<ListCognitiveServicesAccountKeysResult> InvokeAsync(ListCognitiveServicesAccountKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListCognitiveServicesAccountKeysResult>("azurerm:cognitiveservices/v20160201preview:listCognitiveServicesAccountKeys", args ?? new ListCognitiveServicesAccountKeysArgs(), options.WithVersion());
    }


    public sealed class ListCognitiveServicesAccountKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cognitive services account within the specified resource group. Cognitive Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.  
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListCognitiveServicesAccountKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListCognitiveServicesAccountKeysResult
    {
        /// <summary>
        /// Gets the value of key 1.
        /// </summary>
        public readonly string? Key1;
        /// <summary>
        /// Gets the value of key 2.
        /// </summary>
        public readonly string? Key2;

        [OutputConstructor]
        private ListCognitiveServicesAccountKeysResult(
            string? key1,

            string? key2)
        {
            Key1 = key1;
            Key2 = key2;
        }
    }
}
