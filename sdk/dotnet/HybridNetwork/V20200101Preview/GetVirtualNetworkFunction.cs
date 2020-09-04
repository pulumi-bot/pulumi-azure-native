// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview
{
    public static class GetVirtualNetworkFunction
    {
        public static Task<GetVirtualNetworkFunctionResult> InvokeAsync(GetVirtualNetworkFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkFunctionResult>("azurerm:hybridnetwork/v20200101preview:getVirtualNetworkFunction", args ?? new GetVirtualNetworkFunctionArgs(), options.WithVersion());
    }


    public sealed class GetVirtualNetworkFunctionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of hybrid network virtual network function resource.
        /// </summary>
        [Input("virtualNetworkFunctionName", required: true)]
        public string VirtualNetworkFunctionName { get; set; } = null!;

        public GetVirtualNetworkFunctionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualNetworkFunctionResult
    {
        /// <summary>
        /// The reference to the hybrid network device.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? Device;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource URI of the managed application.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult ManagedApplication;
        /// <summary>
        /// The parameters for the managed application.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ManagedApplicationParameters;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the hybrid network virtual network function resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The service key for the virtual network function resource.
        /// </summary>
        public readonly string ServiceKey;
        /// <summary>
        /// The sku name for the hybrid network virtual network function.
        /// </summary>
        public readonly string? SkuName;
        /// <summary>
        /// The sku type for the hybrid network virtual network function.
        /// </summary>
        public readonly string SkuType;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The vendor name for the hybrid network virtual network function.
        /// </summary>
        public readonly string? VendorName;
        /// <summary>
        /// The vendor provisioning state for the virtual network function resource.
        /// </summary>
        public readonly string VendorProvisioningState;
        /// <summary>
        /// The virtual network function configurations from the user.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualNetworkFunctionUserConfigurationResponseResult> VirtualNetworkFunctionUserConfigurations;

        [OutputConstructor]
        private GetVirtualNetworkFunctionResult(
            Outputs.SubResourceResponseResult? device,

            string? etag,

            string? location,

            Outputs.SubResourceResponseResult managedApplication,

            ImmutableDictionary<string, object>? managedApplicationParameters,

            string name,

            string provisioningState,

            string serviceKey,

            string? skuName,

            string skuType,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? vendorName,

            string vendorProvisioningState,

            ImmutableArray<Outputs.VirtualNetworkFunctionUserConfigurationResponseResult> virtualNetworkFunctionUserConfigurations)
        {
            Device = device;
            Etag = etag;
            Location = location;
            ManagedApplication = managedApplication;
            ManagedApplicationParameters = managedApplicationParameters;
            Name = name;
            ProvisioningState = provisioningState;
            ServiceKey = serviceKey;
            SkuName = skuName;
            SkuType = skuType;
            Tags = tags;
            Type = type;
            VendorName = vendorName;
            VendorProvisioningState = vendorProvisioningState;
            VirtualNetworkFunctionUserConfigurations = virtualNetworkFunctionUserConfigurations;
        }
    }
}
