// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class NetworkResourcePropertiesResponseResult
    {
        /// <summary>
        /// User readable description of the network.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The type of a Service Fabric container network.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// State of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Status of the network.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Gives additional information about the current status of the network.
        /// </summary>
        public readonly string StatusDetails;

        [OutputConstructor]
        private NetworkResourcePropertiesResponseResult(
            string? description,

            string kind,

            string provisioningState,

            string status,

            string statusDetails)
        {
            Description = description;
            Kind = kind;
            ProvisioningState = provisioningState;
            Status = status;
            StatusDetails = statusDetails;
        }
    }
}
