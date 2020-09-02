// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20170615Preview.Inputs
{

    /// <summary>
    /// Information about the destination for an event subscription
    /// </summary>
    public sealed class EventSubscriptionDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the endpoint for the event subscription destination
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The URL that represents the endpoint of the destination of an event subscription.
        /// </summary>
        [Input("endpointUrl")]
        public Input<string>? EndpointUrl { get; set; }

        public EventSubscriptionDestinationArgs()
        {
        }
    }
}
