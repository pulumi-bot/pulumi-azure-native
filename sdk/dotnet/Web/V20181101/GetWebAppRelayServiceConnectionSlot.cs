// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20181101
{
    public static class GetWebAppRelayServiceConnectionSlot
    {
        public static Task<GetWebAppRelayServiceConnectionSlotResult> InvokeAsync(GetWebAppRelayServiceConnectionSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebAppRelayServiceConnectionSlotResult>("azurerm:web/v20181101:getWebAppRelayServiceConnectionSlot", args ?? new GetWebAppRelayServiceConnectionSlotArgs(), options.WithVersion());
    }


    public sealed class GetWebAppRelayServiceConnectionSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the hybrid connection.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will get a hybrid connection for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetWebAppRelayServiceConnectionSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebAppRelayServiceConnectionSlotResult
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
        /// RelayServiceConnectionEntity resource specific properties
        /// </summary>
        public readonly Outputs.RelayServiceConnectionEntityResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebAppRelayServiceConnectionSlotResult(
            string? kind,

            string name,

            Outputs.RelayServiceConnectionEntityResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}