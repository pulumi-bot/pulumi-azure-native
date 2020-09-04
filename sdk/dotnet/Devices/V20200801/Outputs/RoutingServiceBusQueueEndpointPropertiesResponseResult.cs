// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200801.Outputs
{

    [OutputType]
    public sealed class RoutingServiceBusQueueEndpointPropertiesResponseResult
    {
        /// <summary>
        /// Method used to authenticate against the service bus queue endpoint
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// The connection string of the service bus queue endpoint.
        /// </summary>
        public readonly string? ConnectionString;
        /// <summary>
        /// The url of the service bus queue endpoint. It must include the protocol sb://
        /// </summary>
        public readonly string? EndpointUri;
        /// <summary>
        /// Queue name on the service bus namespace
        /// </summary>
        public readonly string? EntityPath;
        /// <summary>
        /// Id of the service bus queue endpoint
        /// </summary>
        public readonly string? Id;
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
            string? authenticationType,

            string? connectionString,

            string? endpointUri,

            string? entityPath,

            string? id,

            string name,

            string? resourceGroup,

            string? subscriptionId)
        {
            AuthenticationType = authenticationType;
            ConnectionString = connectionString;
            EndpointUri = endpointUri;
            EntityPath = entityPath;
            Id = id;
            Name = name;
            ResourceGroup = resourceGroup;
            SubscriptionId = subscriptionId;
        }
    }
}
