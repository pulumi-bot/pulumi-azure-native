// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights
{
    public static class GetHubPrediction
    {
        public static Task<GetHubPredictionResult> InvokeAsync(GetHubPredictionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHubPredictionResult>("azurerm:customerinsights:getHubPrediction", args ?? new GetHubPredictionArgs(), options.WithVersion());
    }


    public sealed class GetHubPredictionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public string HubName { get; set; } = null!;

        /// <summary>
        /// The name of the Prediction.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHubPredictionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHubPredictionResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The prediction definition.
        /// </summary>
        public readonly Outputs.PredictionResponseResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetHubPredictionResult(
            string name,

            Outputs.PredictionResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}