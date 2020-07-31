// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Inputs
{

    /// <summary>
    /// Route Table resource
    /// </summary>
    public sealed class RouteTablePropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
        /// </summary>
        [Input("disableBgpRoutePropagation")]
        public bool? DisableBgpRoutePropagation { get; set; }

        /// <summary>
        /// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        [Input("routes")]
        private List<Inputs.RouteResponseArgs>? _routes;

        /// <summary>
        /// Collection of routes contained within a route table.
        /// </summary>
        public List<Inputs.RouteResponseArgs> Routes
        {
            get => _routes ?? (_routes = new List<Inputs.RouteResponseArgs>());
            set => _routes = value;
        }

        [Input("subnets", required: true)]
        private List<Inputs.SubnetResponseArgs>? _subnets;

        /// <summary>
        /// A collection of references to subnets.
        /// </summary>
        public List<Inputs.SubnetResponseArgs> Subnets
        {
            get => _subnets ?? (_subnets = new List<Inputs.SubnetResponseArgs>());
            set => _subnets = value;
        }

        public RouteTablePropertiesFormatResponseArgs()
        {
        }
    }
}