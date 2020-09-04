// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview
{
    /// <summary>
    /// Hybrid network virtual network function resource response.
    /// </summary>
    public partial class VirtualNetworkFunction : Pulumi.CustomResource
    {
        /// <summary>
        /// The reference to the hybrid network device.
        /// </summary>
        [Output("device")]
        public Output<Outputs.SubResourceResponseResult?> Device { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource URI of the managed application.
        /// </summary>
        [Output("managedApplication")]
        public Output<Outputs.SubResourceResponseResult> ManagedApplication { get; private set; } = null!;

        /// <summary>
        /// The parameters for the managed application.
        /// </summary>
        [Output("managedApplicationParameters")]
        public Output<ImmutableDictionary<string, object>?> ManagedApplicationParameters { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the hybrid network virtual network function resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The service key for the virtual network function resource.
        /// </summary>
        [Output("serviceKey")]
        public Output<string> ServiceKey { get; private set; } = null!;

        /// <summary>
        /// The sku name for the hybrid network virtual network function.
        /// </summary>
        [Output("skuName")]
        public Output<string?> SkuName { get; private set; } = null!;

        /// <summary>
        /// The sku type for the hybrid network virtual network function.
        /// </summary>
        [Output("skuType")]
        public Output<string> SkuType { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The vendor name for the hybrid network virtual network function.
        /// </summary>
        [Output("vendorName")]
        public Output<string?> VendorName { get; private set; } = null!;

        /// <summary>
        /// The vendor provisioning state for the virtual network function resource.
        /// </summary>
        [Output("vendorProvisioningState")]
        public Output<string> VendorProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The virtual network function configurations from the user.
        /// </summary>
        [Output("virtualNetworkFunctionUserConfigurations")]
        public Output<ImmutableArray<Outputs.VirtualNetworkFunctionUserConfigurationResponseResult>> VirtualNetworkFunctionUserConfigurations { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetworkFunction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkFunction(string name, VirtualNetworkFunctionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:hybridnetwork/v20200101preview:VirtualNetworkFunction", name, args ?? new VirtualNetworkFunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkFunction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:hybridnetwork/v20200101preview:VirtualNetworkFunction", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualNetworkFunction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkFunction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkFunction(name, id, options);
        }
    }

    public sealed class VirtualNetworkFunctionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference to the hybrid network device.
        /// </summary>
        [Input("device")]
        public Input<Inputs.SubResourceArgs>? Device { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("managedApplicationParameters")]
        private InputMap<object>? _managedApplicationParameters;

        /// <summary>
        /// The parameters for the managed application.
        /// </summary>
        public InputMap<object> ManagedApplicationParameters
        {
            get => _managedApplicationParameters ?? (_managedApplicationParameters = new InputMap<object>());
            set => _managedApplicationParameters = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku name for the hybrid network virtual network function.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The vendor name for the hybrid network virtual network function.
        /// </summary>
        [Input("vendorName")]
        public Input<string>? VendorName { get; set; }

        /// <summary>
        /// Resource name for the hybrid network virtual network function resource.
        /// </summary>
        [Input("virtualNetworkFunctionName", required: true)]
        public Input<string> VirtualNetworkFunctionName { get; set; } = null!;

        [Input("virtualNetworkFunctionUserConfigurations")]
        private InputList<Inputs.VirtualNetworkFunctionUserConfigurationArgs>? _virtualNetworkFunctionUserConfigurations;

        /// <summary>
        /// The virtual network function configurations from the user.
        /// </summary>
        public InputList<Inputs.VirtualNetworkFunctionUserConfigurationArgs> VirtualNetworkFunctionUserConfigurations
        {
            get => _virtualNetworkFunctionUserConfigurations ?? (_virtualNetworkFunctionUserConfigurations = new InputList<Inputs.VirtualNetworkFunctionUserConfigurationArgs>());
            set => _virtualNetworkFunctionUserConfigurations = value;
        }

        public VirtualNetworkFunctionArgs()
        {
        }
    }
}
