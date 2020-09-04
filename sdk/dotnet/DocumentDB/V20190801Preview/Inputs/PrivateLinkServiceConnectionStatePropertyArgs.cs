// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20190801Preview.Inputs
{

    /// <summary>
    /// Connection State of the Private Endpoint Connection.
    /// </summary>
    public sealed class PrivateLinkServiceConnectionStatePropertyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The private link service connection description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The private link service connection status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateLinkServiceConnectionStatePropertyArgs()
        {
        }
    }
}
