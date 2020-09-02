// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20190322Preview.Inputs
{

    /// <summary>
    /// The properties related to service bus queue endpoint types.
    /// </summary>
    public sealed class RoutingServiceBusQueueEndpointPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string of the service bus queue endpoint.
        /// </summary>
        [Input("connectionString", required: true)]
        public Input<string> ConnectionString { get; set; } = null!;

        /// <summary>
        /// The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types. The name need not be the same as the actual queue name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group of the service bus queue endpoint.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// The subscription identifier of the service bus queue endpoint.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        public RoutingServiceBusQueueEndpointPropertiesArgs()
        {
        }
    }
}
