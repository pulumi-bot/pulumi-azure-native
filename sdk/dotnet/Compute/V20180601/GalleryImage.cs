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
    /// Specifies information about the gallery Image Definition that you want to create or update.
    /// 
    /// ## Example Usage
    /// ### Create or update a simple gallery image.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var galleryImage = new AzureRM.Compute.V20180601.GalleryImage("galleryImage", new AzureRM.Compute.V20180601.GalleryImageArgs
    ///         {
    ///             GalleryImageName = "myGalleryImageName",
    ///             GalleryName = "myGalleryName",
    ///             Identifier = new AzureRM.Compute.V20180601.Inputs.GalleryImageIdentifierArgs
    ///             {
    ///                 Offer = "myOfferName",
    ///                 Publisher = "myPublisherName",
    ///                 Sku = "mySkuName",
    ///             },
    ///             Location = "West US",
    ///             OsState = "Generalized",
    ///             OsType = "Windows",
    ///             ResourceGroupName = "myResourceGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class GalleryImage : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this gallery Image Definition resource. This property is updatable.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Describes the disallowed disk types.
        /// </summary>
        [Output("disallowed")]
        public Output<Outputs.DisallowedResponseResult?> Disallowed { get; private set; } = null!;

        /// <summary>
        /// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
        /// </summary>
        [Output("endOfLifeDate")]
        public Output<string?> EndOfLifeDate { get; private set; } = null!;

        /// <summary>
        /// The Eula agreement for the gallery Image Definition.
        /// </summary>
        [Output("eula")]
        public Output<string?> Eula { get; private set; } = null!;

        /// <summary>
        /// This is the gallery Image Definition identifier.
        /// </summary>
        [Output("identifier")]
        public Output<Outputs.GalleryImageIdentifierResponseResult> Identifier { get; private set; } = null!;

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
        /// The allowed values for OS State are 'Generalized'.
        /// </summary>
        [Output("osState")]
        public Output<string> OsState { get; private set; } = null!;

        /// <summary>
        /// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **Windows** &lt;br&gt;&lt;br&gt; **Linux**
        /// </summary>
        [Output("osType")]
        public Output<string> OsType { get; private set; } = null!;

        /// <summary>
        /// The privacy statement uri.
        /// </summary>
        [Output("privacyStatementUri")]
        public Output<string?> PrivacyStatementUri { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
        /// </summary>
        [Output("purchasePlan")]
        public Output<Outputs.ImagePurchasePlanResponseResult?> PurchasePlan { get; private set; } = null!;

        /// <summary>
        /// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
        /// </summary>
        [Output("recommended")]
        public Output<Outputs.RecommendedMachineConfigurationResponseResult?> Recommended { get; private set; } = null!;

        /// <summary>
        /// The release note uri.
        /// </summary>
        [Output("releaseNoteUri")]
        public Output<string?> ReleaseNoteUri { get; private set; } = null!;

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
        /// Create a GalleryImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GalleryImage(string name, GalleryImageArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20180601:GalleryImage", name, args ?? new GalleryImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GalleryImage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20180601:GalleryImage", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/latest:GalleryImage"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190301:GalleryImage"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190701:GalleryImage"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20191201:GalleryImage"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GalleryImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GalleryImage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GalleryImage(name, id, options);
        }
    }

    public sealed class GalleryImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this gallery Image Definition resource. This property is updatable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Describes the disallowed disk types.
        /// </summary>
        [Input("disallowed")]
        public Input<Inputs.DisallowedArgs>? Disallowed { get; set; }

        /// <summary>
        /// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
        /// </summary>
        [Input("endOfLifeDate")]
        public Input<string>? EndOfLifeDate { get; set; }

        /// <summary>
        /// The Eula agreement for the gallery Image Definition.
        /// </summary>
        [Input("eula")]
        public Input<string>? Eula { get; set; }

        /// <summary>
        /// The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
        /// </summary>
        [Input("galleryImageName", required: true)]
        public Input<string> GalleryImageName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image Gallery in which the Image Definition is to be created.
        /// </summary>
        [Input("galleryName", required: true)]
        public Input<string> GalleryName { get; set; } = null!;

        /// <summary>
        /// This is the gallery Image Definition identifier.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<Inputs.GalleryImageIdentifierArgs> Identifier { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The allowed values for OS State are 'Generalized'.
        /// </summary>
        [Input("osState", required: true)]
        public Input<string> OsState { get; set; } = null!;

        /// <summary>
        /// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **Windows** &lt;br&gt;&lt;br&gt; **Linux**
        /// </summary>
        [Input("osType", required: true)]
        public Input<string> OsType { get; set; } = null!;

        /// <summary>
        /// The privacy statement uri.
        /// </summary>
        [Input("privacyStatementUri")]
        public Input<string>? PrivacyStatementUri { get; set; }

        /// <summary>
        /// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
        /// </summary>
        [Input("purchasePlan")]
        public Input<Inputs.ImagePurchasePlanArgs>? PurchasePlan { get; set; }

        /// <summary>
        /// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
        /// </summary>
        [Input("recommended")]
        public Input<Inputs.RecommendedMachineConfigurationArgs>? Recommended { get; set; }

        /// <summary>
        /// The release note uri.
        /// </summary>
        [Input("releaseNoteUri")]
        public Input<string>? ReleaseNoteUri { get; set; }

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

        public GalleryImageArgs()
        {
        }
    }
}
