// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Quantum.V20191104Preview.Inputs
{

    /// <summary>
    /// Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
    /// </summary>
    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The provider's marketplace application display name.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// A Uri identifying the specific instance of this provider.
        /// </summary>
        [Input("instanceUri")]
        public Input<string>? InstanceUri { get; set; }

        /// <summary>
        /// Unique id of this provider.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The sku associated with pricing information for this provider.
        /// </summary>
        [Input("providerSku")]
        public Input<string>? ProviderSku { get; set; }

        /// <summary>
        /// Provisioning status field
        /// </summary>
        [Input("provisioningState")]
        public InputUnion<string, Pulumi.AzureNextGen.Quantum.V20191104Preview.Status>? ProvisioningState { get; set; }

        /// <summary>
        /// Id to track resource usage for the provider.
        /// </summary>
        [Input("resourceUsageId")]
        public Input<string>? ResourceUsageId { get; set; }

        public ProviderArgs()
        {
        }
    }
}
