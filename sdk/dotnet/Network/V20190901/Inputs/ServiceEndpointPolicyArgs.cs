// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901.Inputs
{

    /// <summary>
    /// Service End point policy resource.
    /// </summary>
    public sealed class ServiceEndpointPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("serviceEndpointPolicyDefinitions")]
        private InputList<Inputs.ServiceEndpointPolicyDefinitionArgs>? _serviceEndpointPolicyDefinitions;

        /// <summary>
        /// A collection of service endpoint policy definitions of the service endpoint policy.
        /// </summary>
        public InputList<Inputs.ServiceEndpointPolicyDefinitionArgs> ServiceEndpointPolicyDefinitions
        {
            get => _serviceEndpointPolicyDefinitions ?? (_serviceEndpointPolicyDefinitions = new InputList<Inputs.ServiceEndpointPolicyDefinitionArgs>());
            set => _serviceEndpointPolicyDefinitions = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceEndpointPolicyArgs()
        {
        }
    }
}
