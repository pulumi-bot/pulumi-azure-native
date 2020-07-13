// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement
{
    public static class GetServiceNamedValue
    {
        public static Task<GetServiceNamedValueResult> InvokeAsync(GetServiceNamedValueArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceNamedValueResult>("azurerm:apimanagement:getServiceNamedValue", args ?? new GetServiceNamedValueArgs(), options.WithVersion());
    }


    public sealed class GetServiceNamedValueArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the NamedValue.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

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

        public GetServiceNamedValueArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceNamedValueResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NamedValue entity contract properties.
        /// </summary>
        public readonly Outputs.NamedValueContractPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServiceNamedValueResult(
            string name,

            Outputs.NamedValueContractPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}