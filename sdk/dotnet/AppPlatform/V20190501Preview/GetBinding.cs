// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.V20190501Preview
{
    public static class GetBinding
    {
        public static Task<GetBindingResult> InvokeAsync(GetBindingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBindingResult>("azurerm:appplatform/v20190501preview:getBinding", args ?? new GetBindingArgs(), options.WithVersion());
    }


    public sealed class GetBindingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the App resource.
        /// </summary>
        [Input("appName", required: true)]
        public string AppName { get; set; } = null!;

        /// <summary>
        /// The name of the Binding resource.
        /// </summary>
        [Input("bindingName", required: true)]
        public string BindingName { get; set; } = null!;

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

        public GetBindingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBindingResult
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the Binding resource
        /// </summary>
        public readonly Outputs.BindingResourcePropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBindingResult(
            string name,

            Outputs.BindingResourcePropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
