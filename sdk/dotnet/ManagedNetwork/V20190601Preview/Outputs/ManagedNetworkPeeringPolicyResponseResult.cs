// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedNetwork.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class ManagedNetworkPeeringPolicyResponseResult
    {
        /// <summary>
        /// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the properties of a Managed Network Policy
        /// </summary>
        public readonly Outputs.ManagedNetworkPeeringPolicyPropertiesResponseResult? Properties;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ManagedNetworkPeeringPolicyResponseResult(
            string id,

            string? location,

            string name,

            Outputs.ManagedNetworkPeeringPolicyPropertiesResponseResult? properties,

            string type)
        {
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
