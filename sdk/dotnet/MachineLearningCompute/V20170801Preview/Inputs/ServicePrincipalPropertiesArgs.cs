// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170801Preview.Inputs
{

    /// <summary>
    /// The Azure service principal used by Kubernetes for configuring load balancers
    /// </summary>
    public sealed class ServicePrincipalPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service principal client ID
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The service principal secret. This is not returned in response of GET/PUT on the resource. To see this please call listKeys.
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        public ServicePrincipalPropertiesArgs()
        {
        }
    }
}
