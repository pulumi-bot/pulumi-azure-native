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
    /// Tap configuration in a Network Interface
    /// </summary>
    public sealed class NetworkInterfaceTapConfigurationResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public string? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Properties of the Virtual Network Tap configuration
        /// </summary>
        [Input("properties")]
        public Inputs.NetworkInterfaceTapConfigurationPropertiesFormatResponseArgs? Properties { get; set; }

        /// <summary>
        /// Sub Resource type.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public NetworkInterfaceTapConfigurationResponseArgs()
        {
        }
    }
}
