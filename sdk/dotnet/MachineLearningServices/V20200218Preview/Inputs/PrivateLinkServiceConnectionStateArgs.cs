// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200218Preview.Inputs
{

    /// <summary>
    /// A collection of information about the state of the connection between service consumer and provider.
    /// </summary>
    public sealed class PrivateLinkServiceConnectionStateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A message indicating if changes on the service provider require any updates on the consumer.
        /// </summary>
        [Input("actionRequired")]
        public Input<string>? ActionRequired { get; set; }

        /// <summary>
        /// The reason for approval/rejection of the connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateLinkServiceConnectionStateArgs()
        {
        }
    }
}
