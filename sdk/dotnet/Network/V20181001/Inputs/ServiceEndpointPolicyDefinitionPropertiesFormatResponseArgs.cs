// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181001.Inputs
{

    /// <summary>
    /// Service Endpoint policy definition resource.
    /// </summary>
    public sealed class ServiceEndpointPolicyDefinitionPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The provisioning state of the service end point policy definition. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        /// <summary>
        /// service endpoint name.
        /// </summary>
        [Input("service")]
        public string? Service { get; set; }

        [Input("serviceResources")]
        private List<string>? _serviceResources;

        /// <summary>
        /// A list of service resources.
        /// </summary>
        public List<string> ServiceResources
        {
            get => _serviceResources ?? (_serviceResources = new List<string>());
            set => _serviceResources = value;
        }

        public ServiceEndpointPolicyDefinitionPropertiesFormatResponseArgs()
        {
        }
    }
}
