// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.V20190501Preview
{
    public static class GetCustomDomain
    {
        public static Task<GetCustomDomainResult> InvokeAsync(GetCustomDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomDomainResult>("azurerm:appplatform/v20190501preview:getCustomDomain", args ?? new GetCustomDomainArgs(), options.WithVersion());
    }


    public sealed class GetCustomDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the App resource.
        /// </summary>
        [Input("appName", required: true)]
        public string AppName { get; set; } = null!;

        /// <summary>
        /// The name of the custom domain resource.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCustomDomainArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomDomainResult
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the custom domain resource.
        /// </summary>
        public readonly Outputs.CustomDomainPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.
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
