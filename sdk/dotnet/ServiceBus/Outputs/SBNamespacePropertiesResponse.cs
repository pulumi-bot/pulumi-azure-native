// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.Outputs
{

    [OutputType]
    public sealed class SBNamespacePropertiesResponse
    {
        /// <summary>
        /// The time the namespace was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Identifier for Azure Insights metrics
        /// </summary>
        public readonly string MetricId;
        /// <summary>
        /// Provisioning state of the namespace.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Endpoint you can use to perform Service Bus operations.
        /// </summary>
        public readonly string ServiceBusEndpoint;
        /// <summary>
        /// The time the namespace was updated.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private SBNamespacePropertiesResponse(
            string createdAt,

            string metricId,

            string provisioningState,

            string serviceBusEndpoint,

            string updatedAt)
        {
            CreatedAt = createdAt;
            MetricId = metricId;
            ProvisioningState = provisioningState;
            ServiceBusEndpoint = serviceBusEndpoint;
            UpdatedAt = updatedAt;
        }
    }
}