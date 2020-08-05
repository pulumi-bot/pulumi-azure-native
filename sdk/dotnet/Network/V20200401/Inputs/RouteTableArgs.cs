// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Inputs
{

    /// <summary>
    /// Route table resource.
    /// </summary>
    public sealed class RouteTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to disable the routes learned by BGP on that route table. True means disable.
        /// </summary>
        [Input("disableBgpRoutePropagation")]
        public Input<bool>? DisableBgpRoutePropagation { get; set; }

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

        [Input("routes")]
        private InputList<Inputs.RouteArgs>? _routes;

        /// <summary>
        /// Collection of routes contained within a route table.
        /// </summary>
        public InputList<Inputs.RouteArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.RouteArgs>());
            set => _routes = value;
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

        public RouteTableArgs()
        {
        }
    }
}
