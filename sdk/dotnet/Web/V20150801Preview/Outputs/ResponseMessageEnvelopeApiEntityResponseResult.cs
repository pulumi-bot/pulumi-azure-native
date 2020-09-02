// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class ResponseMessageEnvelopeApiEntityResponseResult
    {
        /// <summary>
        /// Resource Id. Typically id is populated only for responses to GET requests. Caller is responsible for passing in this
        ///             value for GET requests only.
        ///             For example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupId}/providers/Microsoft.Web/sites/{sitename}
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Geo region resource belongs to e.g. SouthCentralUS, SouthEastAsia
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Name of resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Azure resource manager plan
        /// </summary>
        public readonly Outputs.ArmPlanResponseResult? Plan;
        /// <summary>
        /// Resource specific properties
        /// </summary>
        public readonly Outputs.ApiEntityResponseResult? Properties;
        /// <summary>
        /// Sku description of the resource
        /// </summary>
        public readonly Outputs.SkuDescriptionResponseResult? Sku;
        /// <summary>
        /// Tags associated with resource
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Type of resource e.g Microsoft.Web/sites
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ResponseMessageEnvelopeApiEntityResponseResult(
            string? id,

            string? location,

            string? name,

            Outputs.ArmPlanResponseResult? plan,

            Outputs.ApiEntityResponseResult? properties,

            Outputs.SkuDescriptionResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            Id = id;
            Location = location;
            Name = name;
            Plan = plan;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
