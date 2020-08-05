// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901
{
    /// <summary>
    /// Route Filter Rule Resource.
    /// </summary>
    public partial class RouteFilterRule : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the route filter rule.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.RouteFilterRulePropertiesFormatResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a RouteFilterRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteFilterRule(string name, RouteFilterRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190901:RouteFilterRule", name, args ?? new RouteFilterRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteFilterRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190901:RouteFilterRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteFilterRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteFilterRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RouteFilterRule(name, id, options);
        }
    }

    public sealed class RouteFilterRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access type of the rule.
        /// </summary>
        [Input("access", required: true)]
        public Input<string> Access { get; set; } = null!;

        [Input("communities", required: true)]
        private InputList<string>? _communities;

        /// <summary>
        /// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
        /// </summary>
        public InputList<string> Communities
        {
            get => _communities ?? (_communities = new InputList<string>());
            set => _communities = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the route filter rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the route filter.
        /// </summary>
        [Input("routeFilterName", required: true)]
        public Input<string> RouteFilterName { get; set; } = null!;

        /// <summary>
        /// The rule type of the rule.
        /// </summary>
        [Input("routeFilterRuleType", required: true)]
        public Input<string> RouteFilterRuleType { get; set; } = null!;

        public RouteFilterRuleArgs()
        {
        }
    }
}
