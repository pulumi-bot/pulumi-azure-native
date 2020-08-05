// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// Route Filter Resource.
    /// </summary>
    public sealed class RouteFilterResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag", required: true)]
        public string Etag { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Route Filter Resource
        /// </summary>
        [Input("properties")]
        public Inputs.RouteFilterPropertiesFormatResponseArgs? Properties { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public RouteFilterResponseArgs()
        {
        }
    }
}
