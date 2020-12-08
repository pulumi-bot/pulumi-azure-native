// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20200901.Inputs
{

    /// <summary>
    /// Describes the properties of an existing Shared Private Link Resource to use when connecting to a private origin.
    /// </summary>
    public sealed class SharedPrivateLinkResourcePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group id from the provider of resource the shared private link resource is for.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The resource id of the resource the shared private link resource is for.
        /// </summary>
        [Input("privateLink")]
        public Input<Inputs.ResourceReferenceArgs>? PrivateLink { get; set; }

        /// <summary>
        /// The location of the shared private link resource
        /// </summary>
        [Input("privateLinkLocation")]
        public Input<string>? PrivateLinkLocation { get; set; }

        /// <summary>
        /// The request message for requesting approval of the shared private link resource.
        /// </summary>
        [Input("requestMessage")]
        public Input<string>? RequestMessage { get; set; }

        /// <summary>
        /// Status of the shared private link resource. Can be Pending, Approved, Rejected, Disconnected, or Timeout.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AzureNextGen.Cdn.V20200901.SharedPrivateLinkResourceStatus>? Status { get; set; }

        public SharedPrivateLinkResourcePropertiesArgs()
        {
        }
    }
}
