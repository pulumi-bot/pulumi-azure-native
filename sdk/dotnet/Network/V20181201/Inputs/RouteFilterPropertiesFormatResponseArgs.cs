// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181201.Inputs
{

    /// <summary>
    /// Route Filter Resource
    /// </summary>
    public sealed class RouteFilterPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        [Input("peerings")]
        private List<Inputs.ExpressRouteCircuitPeeringResponseArgs>? _peerings;

        /// <summary>
        /// A collection of references to express route circuit peerings.
        /// </summary>
        public List<Inputs.ExpressRouteCircuitPeeringResponseArgs> Peerings
        {
            get => _peerings ?? (_peerings = new List<Inputs.ExpressRouteCircuitPeeringResponseArgs>());
            set => _peerings = value;
        }

        /// <summary>
        /// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        [Input("rules")]
        private List<Inputs.RouteFilterRuleResponseArgs>? _rules;

        /// <summary>
        /// Collection of RouteFilterRules contained within a route filter.
        /// </summary>
        public List<Inputs.RouteFilterRuleResponseArgs> Rules
        {
            get => _rules ?? (_rules = new List<Inputs.RouteFilterRuleResponseArgs>());
            set => _rules = value;
        }

        public RouteFilterPropertiesFormatResponseArgs()
        {
        }
    }
}
