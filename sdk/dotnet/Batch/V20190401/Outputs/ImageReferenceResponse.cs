// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20190401.Outputs
{

    [OutputType]
    public sealed class ImageReferenceResponse
    {
        /// <summary>
        /// This property is mutually exclusive with other properties. The virtual machine image must be in the same region and subscription as the Azure Batch account. For information about the firewall settings for Batch node agent to communicate with Batch service see https://docs.microsoft.com/en-us/azure/batch/batch-api-basics#virtual-network-vnet-and-firewall-configuration .
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// For example, UbuntuServer or WindowsServer.
        /// </summary>
        public readonly string? Offer;
        /// <summary>
        /// For example, Canonical or MicrosoftWindowsServer.
        /// </summary>
        public readonly string? Publisher;
        /// <summary>
        /// For example, 14.04.0-LTS or 2012-R2-Datacenter.
        /// </summary>
        public readonly string? Sku;
        /// <summary>
        /// A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ImageReferenceResponse(
            string? id,

            string? offer,

            string? publisher,

            string? sku,

            string? version)
        {
            Id = id;
            Offer = offer;
            Publisher = publisher;
            Sku = sku;
            Version = version;
        }
    }
}
