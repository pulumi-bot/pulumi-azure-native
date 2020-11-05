// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HybridNetwork.V20200101Preview
{
    public static class GetVendorSkuPreview
    {
        public static Task<GetVendorSkuPreviewResult> InvokeAsync(GetVendorSkuPreviewArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVendorSkuPreviewResult>("azure-nextgen:hybridnetwork/v20200101preview:getVendorSkuPreview", args ?? new GetVendorSkuPreviewArgs(), options.WithVersion());
    }


    public sealed class GetVendorSkuPreviewArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Preview subscription ID.
        /// </summary>
        [Input("previewSubscription", required: true)]
        public string PreviewSubscription { get; set; } = null!;

        /// <summary>
        /// The name of the vendor sku.
        /// </summary>
        [Input("skuName", required: true)]
        public string SkuName { get; set; } = null!;

        /// <summary>
        /// The name of the vendor.
        /// </summary>
        [Input("vendorName", required: true)]
        public string VendorName { get; set; } = null!;

        public GetVendorSkuPreviewArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVendorSkuPreviewResult
    {
        /// <summary>
        /// The preview subscription ID.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVendorSkuPreviewResult(
            string name,

            string type)
        {
            Name = name;
            Type = type;
        }
    }
}
