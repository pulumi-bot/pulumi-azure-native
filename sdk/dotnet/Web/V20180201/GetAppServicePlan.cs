// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201
{
    public static class GetAppServicePlan
    {
        public static Task<GetAppServicePlanResult> InvokeAsync(GetAppServicePlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServicePlanResult>("azurerm:web/v20180201:getAppServicePlan", args ?? new GetAppServicePlanArgs(), options.WithVersion());
    }


    public sealed class GetAppServicePlanArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the App Service plan.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppServicePlanArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServicePlanResult
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
        /// AppServicePlan resource specific properties
        /// </summary>
        public readonly Outputs.AppServicePlanResponsePropertiesResult Properties;
        /// <summary>
        /// Description of a SKU for a scalable resource.
        /// </summary>
        public readonly Outputs.SkuDescriptionResponseResult? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppServicePlanResult(
            string? kind,

            string location,

            string name,

            Outputs.AppServicePlanResponsePropertiesResult properties,

            Outputs.SkuDescriptionResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}