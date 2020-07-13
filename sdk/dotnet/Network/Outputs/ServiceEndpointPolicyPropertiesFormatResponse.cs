// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class ServiceEndpointPolicyPropertiesFormatResponse
    {
        /// <summary>
        /// The provisioning state of the service endpoint policy resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource GUID property of the service endpoint policy resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// A collection of service endpoint policy definitions of the service endpoint policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceEndpointPolicyDefinitionResponse> ServiceEndpointPolicyDefinitions;
        /// <summary>
        /// A collection of references to subnets.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponse> Subnets;

        [OutputConstructor]
        private ServiceEndpointPolicyPropertiesFormatResponse(
            string provisioningState,

            string resourceGuid,

            ImmutableArray<Outputs.ServiceEndpointPolicyDefinitionResponse> serviceEndpointPolicyDefinitions,

            ImmutableArray<Outputs.SubnetResponse> subnets)
        {
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            ServiceEndpointPolicyDefinitions = serviceEndpointPolicyDefinitions;
            Subnets = subnets;
        }
    }
}