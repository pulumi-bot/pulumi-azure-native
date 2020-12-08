// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20180701.Inputs
{

    /// <summary>
    /// SKU of a public IP address
    /// </summary>
    public sealed class PublicIPAddressSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of a public IP address SKU.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20180701.PublicIPAddressSkuName>? Name { get; set; }

        public PublicIPAddressSkuArgs()
        {
        }
    }
}
