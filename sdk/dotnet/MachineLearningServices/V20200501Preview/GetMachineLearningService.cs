// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200501Preview
{
    public static class GetMachineLearningService
    {
        public static Task<GetMachineLearningServiceResult> InvokeAsync(GetMachineLearningServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMachineLearningServiceResult>("azurerm:machinelearningservices/v20200501preview:getMachineLearningService", args ?? new GetMachineLearningServiceArgs(), options.WithVersion());
    }


    public sealed class GetMachineLearningServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Set to True to include Model details.
        /// </summary>
        [Input("expand")]
        public bool? Expand { get; set; }

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Machine Learning service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetMachineLearningServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMachineLearningServiceResult
    {
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Specifies the name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Service properties
        /// </summary>
        public readonly Outputs.ServiceResponseBaseResponseResult Properties;
        /// <summary>
        /// The sku of the workspace.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies the type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMachineLearningServiceResult(
            Outputs.IdentityResponseResult? identity,

            string? location,

            string name,

            Outputs.ServiceResponseBaseResponseResult properties,

            Outputs.SkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
