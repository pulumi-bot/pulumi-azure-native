// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HybridNetwork.V20200101Preview
{
    /// <summary>
    /// Device resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:hybridnetwork/v20200101preview:Device")]
    public partial class Device : Pulumi.CustomResource
    {
        /// <summary>
        /// The reference to the Azure stack edge device.
        /// </summary>
        [Output("azureStackEdge")]
        public Output<Outputs.SubResourceResponse?> AzureStackEdge { get; private set; } = null!;

        /// <summary>
        /// The type of the device.
        /// </summary>
        [Output("deviceType")]
        public Output<string> DeviceType { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of network functions deployed on the device.
        /// </summary>
        [Output("networkFunctions")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> NetworkFunctions { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the device resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The current device status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:hybridnetwork/v20200101preview:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:hybridnetwork/v20200101preview:Device", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Device(name, id, options);
        }
    }

    public sealed class DeviceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference to the Azure stack edge device.
        /// </summary>
        [Input("azureStackEdge")]
        public Input<Inputs.SubResourceArgs>? AzureStackEdge { get; set; }

        /// <summary>
        /// Resource name for the device resource.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The type of the device.
        /// </summary>
        [Input("deviceType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.HybridNetwork.V20200101Preview.DeviceType> DeviceType { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public DeviceArgs()
        {
        }
    }
}
