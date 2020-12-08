// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200710Preview.Inputs
{

    /// <summary>
    /// The current state of a private endpoint connection
    /// </summary>
    public sealed class PrivateLinkServiceConnectionStateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actions required for a private endpoint connection
        /// </summary>
        [Input("actionsRequired")]
        public Input<string>? ActionsRequired { get; set; }

        /// <summary>
        /// The description for the current state of a private endpoint connection
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The status of a private endpoint connection
        /// </summary>
        [Input("status", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Devices.V20200710Preview.PrivateLinkServiceConnectionStatus> Status { get; set; } = null!;

        public PrivateLinkServiceConnectionStateArgs()
        {
        }
    }
}
