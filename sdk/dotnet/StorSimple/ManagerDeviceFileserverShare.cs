// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple
{
    /// <summary>
    /// The File Share.
    /// </summary>
    public partial class ManagerDeviceFileserverShare : Pulumi.CustomResource
    {
        /// <summary>
        /// The name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.FileSharePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagerDeviceFileserverShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagerDeviceFileserverShare(string name, ManagerDeviceFileserverShareArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storsimple:ManagerDeviceFileserverShare", name, args ?? new ManagerDeviceFileserverShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagerDeviceFileserverShare(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storsimple:ManagerDeviceFileserverShare", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ManagerDeviceFileserverShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagerDeviceFileserverShare Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagerDeviceFileserverShare(name, id, options);
        }
    }

    public sealed class ManagerDeviceFileserverShareArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The file server name.
        /// </summary>
        [Input("fileServerName", required: true)]
        public Input<string> FileServerName { get; set; } = null!;

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// The file share name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.FileSharePropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ManagerDeviceFileserverShareArgs()
        {
        }
    }
}
