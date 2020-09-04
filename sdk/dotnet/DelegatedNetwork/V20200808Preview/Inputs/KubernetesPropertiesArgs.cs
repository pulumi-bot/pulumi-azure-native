// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DelegatedNetwork.V20200808Preview.Inputs
{

    /// <summary>
    /// Properties of kubernetes cluster
    /// </summary>
    public sealed class KubernetesPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIServer url
        /// </summary>
        [Input("apiServerEndpoint")]
        public Input<string>? ApiServerEndpoint { get; set; }

        /// <summary>
        /// RootCA certificate of kubernetes cluster
        /// </summary>
        [Input("clusterRootCA")]
        public Input<string>? ClusterRootCA { get; set; }

        /// <summary>
        /// AAD ID used with apiserver
        /// </summary>
        [Input("serverAppID")]
        public Input<string>? ServerAppID { get; set; }

        /// <summary>
        /// TenantID of server App ID
        /// </summary>
        [Input("serverTenantID")]
        public Input<string>? ServerTenantID { get; set; }

        public KubernetesPropertiesArgs()
        {
        }
    }
}
