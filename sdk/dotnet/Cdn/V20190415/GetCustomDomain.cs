// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20190415
{
    public static class GetCustomDomain
    {
        public static Task<GetCustomDomainResult> InvokeAsync(GetCustomDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomDomainResult>("azurerm:cdn/v20190415:getCustomDomain", args ?? new GetCustomDomainArgs(), options.WithVersion());
    }


    public sealed class GetCustomDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// Name of the custom domain within an endpoint.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCustomDomainArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomDomainResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The JSON object that contains the properties of the custom domain to create.
        /// </summary>
        public readonly Outputs.CustomDomainPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCustomDomainResult(
            string name,

            Outputs.CustomDomainPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}