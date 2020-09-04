// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20181201Preview.Outputs
{

    [OutputType]
    public sealed class RoutingServiceBusQueueEndpointPropertiesResponseResult
    {
        /// <summary>
        /// The connection string of the service bus queue endpoint.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types. The name need not be the same as the actual queue name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the resource group of the service bus queue endpoint.
        /// </summary>
        public readonly string? ResourceGroup;
        /// <summary>
        /// The subscription identifier of the service bus queue endpoint.
        /// </summary>
        public readonly string? SubscriptionId;

        [OutputConstructor]
        private RoutingServiceBusQueueEndpointPropertiesResponseResult(
            string connectionString,

            string name,

            string? resourceGroup,

            string? subscriptionId)
        {
            ConnectionString = connectionString;
            Name = name;
            ResourceGroup = resourceGroup;
            SubscriptionId = subscriptionId;
        }
    }
}
