// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20171101.Inputs
{

    /// <summary>
    /// Route Filter Resource.
    /// </summary>
    public sealed class RouteFilterDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("peerings")]
        private InputList<Inputs.ExpressRouteCircuitPeeringArgs>? _peerings;

        /// <summary>
        /// A collection of references to express route circuit peerings.
        /// </summary>
        public InputList<Inputs.ExpressRouteCircuitPeeringArgs> Peerings
        {
            get => _peerings ?? (_peerings = new InputList<Inputs.ExpressRouteCircuitPeeringArgs>());
            set => _peerings = value;
        }

        [Input("rules")]
        private InputList<Inputs.RouteFilterRuleArgs>? _rules;

        /// <summary>
        /// Collection of RouteFilterRules contained within a route filter.
        /// </summary>
        public InputList<Inputs.RouteFilterRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RouteFilterRuleArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RouteFilterDefinitionArgs()
        {
        }
    }
}
