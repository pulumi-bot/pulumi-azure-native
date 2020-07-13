// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Outputs
{

    [OutputType]
    public sealed class GalleryImageIdentifierResponse
    {
        /// <summary>
        /// The name of the gallery Image Definition offer.
        /// </summary>
        public readonly string Offer;
        /// <summary>
        /// The name of the gallery Image Definition publisher.
        /// </summary>
        public readonly string Publisher;
        /// <summary>
        /// The name of the gallery Image Definition SKU.
        /// </summary>
        public readonly string Sku;

        [OutputConstructor]
        private GalleryImageIdentifierResponse(
            string offer,

            string publisher,

            string sku)
        {
            Offer = offer;
            Publisher = publisher;
            Sku = sku;
        }
    }
}