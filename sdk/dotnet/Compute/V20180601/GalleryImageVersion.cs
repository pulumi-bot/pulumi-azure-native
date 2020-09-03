// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180601
{
    /// <summary>
    /// Specifies information about the gallery Image Version that you want to create or update.
    /// 
    /// ## Example Usage
    /// ### Create or update a simple gallery Image Version.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var galleryImageVersion = new AzureRM.Compute.V20180601.GalleryImageVersion("galleryImageVersion", new AzureRM.Compute.V20180601.GalleryImageVersionArgs
    ///         {
    ///             GalleryImageName = "myGalleryImageName",
    ///             GalleryImageVersionName = "1.0.0",
    ///             GalleryName = "myGalleryName",
    ///             Location = "West US",
    ///             PublishingProfile = new AzureRM.Compute.V20180601.Inputs.GalleryImageVersionPublishingProfileArgs
    ///             {
    ///                 Source = new AzureRM.Compute.V20180601.Inputs.GalleryArtifactSourceArgs
    ///                 {
    ///                     ManagedImage = new AzureRM.Compute.V20180601.Inputs.ManagedArtifactArgs
    ///                     {
    ///                         Id = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}",
    ///                     },
    ///                 },
    ///                 TargetRegions = 
    ///                 {
    ///                     new AzureRM.Compute.V20180601.Inputs.TargetRegionArgs
    ///                     {
    ///                         Name = "West US",
    ///                         RegionalReplicaCount = 1,
    ///                     },
    ///                     new AzureRM.Compute.V20180601.Inputs.TargetRegionArgs
    ///                     {
    ///                         Name = "East US",
    ///                         RegionalReplicaCount = 2,
    ///                     },
    ///                 },
    ///             },
    ///             ResourceGroupName = "myResourceGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class GalleryImageVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The publishing profile of a gallery Image Version.
        /// </summary>
        [Output("publishingProfile")]
        public Output<Outputs.GalleryImageVersionPublishingProfileResponseResult> PublishingProfile { get; private set; } = null!;

        /// <summary>
        /// This is the replication status of the gallery Image Version.
        /// </summary>
        [Output("replicationStatus")]
        public Output<Outputs.ReplicationStatusResponseResult> ReplicationStatus { get; private set; } = null!;

        /// <summary>
        /// This is the storage profile of a gallery Image Version.
        /// </summary>
        [Output("storageProfile")]
        public Output<Outputs.GalleryImageVersionStorageProfileResponseResult> StorageProfile { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a GalleryImageVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GalleryImageVersion(string name, GalleryImageVersionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20180601:GalleryImageVersion", name, args ?? new GalleryImageVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GalleryImageVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20180601:GalleryImageVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/latest:GalleryImageVersion"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190301:GalleryImageVersion"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190701:GalleryImageVersion"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20191201:GalleryImageVersion"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GalleryImageVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GalleryImageVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GalleryImageVersion(name, id, options);
        }
    }

    public sealed class GalleryImageVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the gallery Image Definition in which the Image Version is to be created.
        /// </summary>
        [Input("galleryImageName", required: true)]
        public Input<string> GalleryImageName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: &lt;MajorVersion&gt;.&lt;MinorVersion&gt;.&lt;Patch&gt;
        /// </summary>
        [Input("galleryImageVersionName", required: true)]
        public Input<string> GalleryImageVersionName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image Gallery in which the Image Definition resides.
        /// </summary>
        [Input("galleryName", required: true)]
        public Input<string> GalleryName { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The publishing profile of a gallery Image Version.
        /// </summary>
        [Input("publishingProfile", required: true)]
        public Input<Inputs.GalleryImageVersionPublishingProfileArgs> PublishingProfile { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GalleryImageVersionArgs()
        {
        }
    }
}
