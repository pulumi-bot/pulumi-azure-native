// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBoxEdge.V20200501Preview
{
    /// <summary>
    /// Represents a container on the  Data Box Edge/Gateway device.
    /// </summary>
    public partial class Container : Pulumi.CustomResource
    {
        /// <summary>
        /// Current status of the container.
        /// </summary>
        [Output("containerStatus")]
        public Output<string> ContainerStatus { get; private set; } = null!;

        /// <summary>
        /// The UTC time when container got created.
        /// </summary>
        [Output("createdDateTime")]
        public Output<string> CreatedDateTime { get; private set; } = null!;

        /// <summary>
        /// DataFormat for Container
        /// </summary>
        [Output("dataFormat")]
        public Output<string> DataFormat { get; private set; } = null!;

        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Details of the refresh job on this container.
        /// </summary>
        [Output("refreshDetails")]
        public Output<Outputs.RefreshDetailsResponse> RefreshDetails { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Container resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Container(string name, ContainerArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:databoxedge/v20200501preview:Container", name, args ?? new ContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Container(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:databoxedge/v20200501preview:Container", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:databoxedge/latest:Container"},
                    new Pulumi.Alias { Type = "azure-nextgen:databoxedge/v20190801:Container"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Container resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Container Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Container(name, id, options);
        }
    }

    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The container name.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// DataFormat for Container
        /// </summary>
        [Input("dataFormat", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.DataBoxEdge.V20200501Preview.AzureContainerDataFormat> DataFormat { get; set; } = null!;

        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Storage Account Name
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        public ContainerArgs()
        {
        }
    }
}
