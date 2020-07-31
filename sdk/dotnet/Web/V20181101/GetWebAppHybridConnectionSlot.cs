// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20181101
{
    public static class GetWebAppHybridConnectionSlot
    {
        public static Task<GetWebAppHybridConnectionSlotResult> InvokeAsync(GetWebAppHybridConnectionSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebAppHybridConnectionSlotResult>("azurerm:web/v20181101:getWebAppHybridConnectionSlot", args ?? new GetWebAppHybridConnectionSlotArgs(), options.WithVersion());
    }


    public sealed class GetWebAppHybridConnectionSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The relay name for this hybrid connection.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace for this hybrid connection.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the slot for the web app.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetWebAppHybridConnectionSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebAppHybridConnectionSlotResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// HybridConnection resource specific properties
        /// </summary>
        public readonly Outputs.HybridConnectionResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebAppHybridConnectionSlotResult(
            string? kind,

            string name,

            Outputs.HybridConnectionResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}