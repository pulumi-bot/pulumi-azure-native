// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview
{
    /// <summary>
    /// The Live Output.
    /// </summary>
    public partial class LiveOutput : Pulumi.CustomResource
    {
        /// <summary>
        /// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
        /// </summary>
        [Output("archiveWindowLength")]
        public Output<string> ArchiveWindowLength { get; private set; } = null!;

        /// <summary>
        /// The asset name.
        /// </summary>
        [Output("assetName")]
        public Output<string> AssetName { get; private set; } = null!;

        /// <summary>
        /// The exact time the Live Output was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The description of the Live Output.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The HLS configuration.
        /// </summary>
        [Output("hls")]
        public Output<Outputs.HlsResponseResult?> Hls { get; private set; } = null!;

        /// <summary>
        /// The exact time the Live Output was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The manifest file name.
        /// </summary>
        [Output("manifestName")]
        public Output<string?> ManifestName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The output snapshot time.
        /// </summary>
        [Output("outputSnapTime")]
        public Output<int?> OutputSnapTime { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the Live Output.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource state of the Live Output.
        /// </summary>
        [Output("resourceState")]
        public Output<string> ResourceState { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LiveOutput resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LiveOutput(string name, LiveOutputArgs args, CustomResourceOptions? options = null)
            : base("azurerm:media/v20180330preview:LiveOutput", name, args ?? new LiveOutputArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LiveOutput(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:media/v20180330preview:LiveOutput", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:media/latest:LiveOutput"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180601preview:LiveOutput"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180701:LiveOutput"},
                    new Pulumi.Alias { Type = "azurerm:media/v20190501preview:LiveOutput"},
                    new Pulumi.Alias { Type = "azurerm:media/v20200501:LiveOutput"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LiveOutput resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LiveOutput Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LiveOutput(name, id, options);
        }
    }

    public sealed class LiveOutputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
        /// </summary>
        [Input("archiveWindowLength", required: true)]
        public Input<string> ArchiveWindowLength { get; set; } = null!;

        /// <summary>
        /// The asset name.
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        /// <summary>
        /// The description of the Live Output.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The HLS configuration.
        /// </summary>
        [Input("hls")]
        public Input<Inputs.HlsArgs>? Hls { get; set; }

        /// <summary>
        /// The name of the Live Event.
        /// </summary>
        [Input("liveEventName", required: true)]
        public Input<string> LiveEventName { get; set; } = null!;

        /// <summary>
        /// The name of the Live Output.
        /// </summary>
        [Input("liveOutputName", required: true)]
        public Input<string> LiveOutputName { get; set; } = null!;

        /// <summary>
        /// The manifest file name.
        /// </summary>
        [Input("manifestName")]
        public Input<string>? ManifestName { get; set; }

        /// <summary>
        /// The output snapshot time.
        /// </summary>
        [Input("outputSnapTime")]
        public Input<int>? OutputSnapTime { get; set; }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public LiveOutputArgs()
        {
        }
    }
}
