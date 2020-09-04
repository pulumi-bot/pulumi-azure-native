// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20190930Preview.Outputs
{

    [OutputType]
    public sealed class PurchasePlanResponseResult
    {
        /// <summary>
        /// The plan ID.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Specifies the product of the image from the marketplace. This is the same value as Offer under the imageReference element.
        /// </summary>
        public readonly string? Product;
        /// <summary>
        /// The promotion code.
        /// </summary>
        public readonly string? PromotionCode;
        /// <summary>
        /// The plan ID.
        /// </summary>
        public readonly string? Publisher;

        [OutputConstructor]
        private PurchasePlanResponseResult(
            string? name,

            string? product,

            string? promotionCode,

            string? publisher)
        {
            Name = name;
            Product = product;
            PromotionCode = promotionCode;
            Publisher = publisher;
        }
    }
}
