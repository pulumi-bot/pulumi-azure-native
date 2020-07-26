// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Inputs
{

    /// <summary>
    /// Private Link Service Connection State
    /// </summary>
    public sealed class PrivateLinkServiceConnectionStateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets actions required
        /// </summary>
        [Input("actionRequired")]
        public Input<string>? ActionRequired { get; set; }

        /// <summary>
        /// Gets or sets description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Gets or sets the status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateLinkServiceConnectionStateArgs()
        {
        }
    }
}
