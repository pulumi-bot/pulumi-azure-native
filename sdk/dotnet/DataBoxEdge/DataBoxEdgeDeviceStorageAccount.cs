// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge
{
    /// <summary>
    /// Represents a Storage Account on the  Data Box Edge/Gateway device.
    /// </summary>
    public partial class DataBoxEdgeDeviceStorageAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Storage Account properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.StorageAccountPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataBoxEdgeDeviceStorageAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataBoxEdgeDeviceStorageAccount(string name, DataBoxEdgeDeviceStorageAccountArgs args, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge:DataBoxEdgeDeviceStorageAccount", name, args ?? new DataBoxEdgeDeviceStorageAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataBoxEdgeDeviceStorageAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge:DataBoxEdgeDeviceStorageAccount", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DataBoxEdgeDeviceStorageAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataBoxEdgeDeviceStorageAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataBoxEdgeDeviceStorageAccount(name, id, options);
        }
    }

    public sealed class DataBoxEdgeDeviceStorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The StorageAccount name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Storage Account properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.StorageAccountPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DataBoxEdgeDeviceStorageAccountArgs()
        {
        }
    }
}
