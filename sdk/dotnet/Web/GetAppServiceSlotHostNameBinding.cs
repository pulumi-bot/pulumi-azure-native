// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    public static class GetAppServiceSlotHostNameBinding
    {
        public static Task<GetAppServiceSlotHostNameBindingResult> InvokeAsync(GetAppServiceSlotHostNameBindingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServiceSlotHostNameBindingResult>("azurerm:web:getAppServiceSlotHostNameBinding", args ?? new GetAppServiceSlotHostNameBindingArgs(), options.WithVersion());
    }


    public sealed class GetAppServiceSlotHostNameBindingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Hostname in the hostname binding.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API the named binding for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetAppServiceSlotHostNameBindingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServiceSlotHostNameBindingResult
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
        /// HostNameBinding resource specific properties
        /// </summary>
        public readonly Outputs.HostNameBindingResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppServiceSlotHostNameBindingResult(
            string? kind,

            string name,

            Outputs.HostNameBindingResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
