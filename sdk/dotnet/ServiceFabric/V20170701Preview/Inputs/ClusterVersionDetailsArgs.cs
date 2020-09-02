// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20170701Preview.Inputs
{

    /// <summary>
    /// The detail of the Service Fabric runtime version result
    /// </summary>
    public sealed class ClusterVersionDetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Service Fabric runtime version of the cluster.
        /// </summary>
        [Input("codeVersion")]
        public Input<string>? CodeVersion { get; set; }

        /// <summary>
        /// Indicates if this version is for Windows or Linux operating system.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The date of expiry of support of the version.
        /// </summary>
        [Input("supportExpiryUtc")]
        public Input<string>? SupportExpiryUtc { get; set; }

        public ClusterVersionDetailsArgs()
        {
        }
    }
}
