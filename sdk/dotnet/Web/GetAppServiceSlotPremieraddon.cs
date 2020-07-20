// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    public static class GetAppServiceSlotPremieraddon
    {
        public static Task<GetAppServiceSlotPremieraddonResult> InvokeAsync(GetAppServiceSlotPremieraddonArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServiceSlotPremieraddonResult>("azurerm:web:getAppServiceSlotPremieraddon", args ?? new GetAppServiceSlotPremieraddonArgs(), options.WithVersion());
    }


    public sealed class GetAppServiceSlotPremieraddonArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Add-on name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will get the named add-on for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetAppServiceSlotPremieraddonArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServiceSlotPremieraddonResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// PremierAddOn resource specific properties
        /// </summary>
        public readonly Outputs.PremierAddOnResponsePropertiesResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppServiceSlotPremieraddonResult(
            string? kind,

            string location,

            string name,

            Outputs.PremierAddOnResponsePropertiesResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
