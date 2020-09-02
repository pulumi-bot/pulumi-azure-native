// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.Latest.Outputs
{

    [OutputType]
    public sealed class ImageReferenceResponseResult
    {
        /// <summary>
        /// Offer of the image.
        /// </summary>
        public readonly string Offer;
        /// <summary>
        /// Publisher of the image.
        /// </summary>
        public readonly string Publisher;
        /// <summary>
        /// SKU of the image.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// Version of the image.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// The ARM resource identifier of the virtual machine image for the compute nodes. This is of the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}. The virtual machine image must be in the same region and subscription as the cluster. For information about the firewall settings for the Batch node agent to communicate with the Batch service see https://docs.microsoft.com/en-us/azure/batch/batch-api-basics#virtual-network-vnet-and-firewall-configuration. Note, you need to provide publisher, offer and sku of the base OS image of which the custom image has been derived from.
        /// </summary>
        public readonly string? VirtualMachineImageId;

        [OutputConstructor]
        private ImageReferenceResponseResult(
            string offer,

            string publisher,

            string sku,

            string? version,

            string? virtualMachineImageId)
        {
            Offer = offer;
            Publisher = publisher;
            Sku = sku;
            Version = version;
            VirtualMachineImageId = virtualMachineImageId;
        }
    }
}