// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kubernetes.V20200101Preview.Inputs
{

    public sealed class ConnectedClusterAADProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client app id configured on target K8 cluster 
        /// </summary>
        [Input("clientAppId", required: true)]
        public Input<string> ClientAppId { get; set; } = null!;

        /// <summary>
        /// The server app id to access AD server
        /// </summary>
        [Input("serverAppId", required: true)]
        public Input<string> ServerAppId { get; set; } = null!;

        /// <summary>
        /// The aad tenant id which is configured on target K8s cluster
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public ConnectedClusterAADProfileArgs()
        {
        }
    }
}
