// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191101.Inputs
{

    /// <summary>
    /// Route Filter Rule Resource.
    /// </summary>
    public sealed class RouteFilterRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access type of the rule.
        /// </summary>
        [Input("access", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20191101.Access> Access { get; set; } = null!;

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
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The rule type of the rule.
        /// </summary>
        [Input("routeFilterRuleType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20191101.RouteFilterRuleType> RouteFilterRuleType { get; set; } = null!;

        public RouteFilterRuleArgs()
        {
        }
    }
}
