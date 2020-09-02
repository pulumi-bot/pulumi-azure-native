// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.Latest
{
    /// <summary>
    /// A Streaming Locator resource
    /// </summary>
    public partial class StreamingLocator : Pulumi.CustomResource
    {
        /// <summary>
        /// Alternative Media ID of this Streaming Locator
        /// </summary>
        [Output("alternativeMediaId")]
        public Output<string?> AlternativeMediaId { get; private set; } = null!;

        /// <summary>
        /// Asset Name
        /// </summary>
        [Output("assetName")]
        public Output<string> AssetName { get; private set; } = null!;

        /// <summary>
        /// The ContentKeys used by this Streaming Locator.
        /// </summary>
        [Output("contentKeys")]
        public Output<ImmutableArray<Outputs.StreamingLocatorContentKeyResponseResult>> ContentKeys { get; private set; } = null!;

        /// <summary>
        /// The creation time of the Streaming Locator.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Name of the default ContentKeyPolicy used by this Streaming Locator.
        /// </summary>
        [Output("defaultContentKeyPolicyName")]
        public Output<string?> DefaultContentKeyPolicyName { get; private set; } = null!;

        /// <summary>
        /// The end time of the Streaming Locator.
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// A list of asset or account filters which apply to this streaming locator
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<string>> Filters { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The start time of the Streaming Locator.
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;

        /// <summary>
        /// The StreamingLocatorId of the Streaming Locator.
        /// </summary>
        [Output("streamingLocatorId")]
        public Output<string?> StreamingLocatorId { get; private set; } = null!;

        /// <summary>
        /// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        /// </summary>
        [Output("streamingPolicyName")]
        public Output<string> StreamingPolicyName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StreamingLocator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamingLocator(string name, StreamingLocatorArgs args, CustomResourceOptions? options = null)
            : base("azurerm:media/latest:StreamingLocator", name, args ?? new StreamingLocatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamingLocator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:media/latest:StreamingLocator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:media/v20180330preview:StreamingLocator"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180601preview:StreamingLocator"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180701:StreamingLocator"},
                    new Pulumi.Alias { Type = "azurerm:media/v20200501:StreamingLocator"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamingLocator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamingLocator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StreamingLocator(name, id, options);
        }
    }

    public sealed class StreamingLocatorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Alternative Media ID of this Streaming Locator
        /// </summary>
        [Input("alternativeMediaId")]
        public Input<string>? AlternativeMediaId { get; set; }

        /// <summary>
        /// Asset Name
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        [Input("contentKeys")]
        private InputList<Inputs.StreamingLocatorContentKeyArgs>? _contentKeys;

        /// <summary>
        /// The ContentKeys used by this Streaming Locator.
        /// </summary>
        public InputList<Inputs.StreamingLocatorContentKeyArgs> ContentKeys
        {
            get => _contentKeys ?? (_contentKeys = new InputList<Inputs.StreamingLocatorContentKeyArgs>());
            set => _contentKeys = value;
        }

        /// <summary>
        /// Name of the default ContentKeyPolicy used by this Streaming Locator.
        /// </summary>
        [Input("defaultContentKeyPolicyName")]
        public Input<string>? DefaultContentKeyPolicyName { get; set; }

        /// <summary>
        /// The end time of the Streaming Locator.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("filters")]
        private InputList<string>? _filters;

        /// <summary>
        /// A list of asset or account filters which apply to this streaming locator
        /// </summary>
        public InputList<string> Filters
        {
            get => _filters ?? (_filters = new InputList<string>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The start time of the Streaming Locator.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The StreamingLocatorId of the Streaming Locator.
        /// </summary>
        [Input("streamingLocatorId")]
        public Input<string>? StreamingLocatorId { get; set; }

        /// <summary>
        /// The Streaming Locator name.
        /// </summary>
        [Input("streamingLocatorName", required: true)]
        public Input<string> StreamingLocatorName { get; set; } = null!;

        /// <summary>
        /// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        /// </summary>
        [Input("streamingPolicyName", required: true)]
        public Input<string> StreamingPolicyName { get; set; } = null!;

        public StreamingLocatorArgs()
        {
        }
    }
}
