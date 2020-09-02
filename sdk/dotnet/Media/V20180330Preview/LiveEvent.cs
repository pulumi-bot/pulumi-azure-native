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
    /// The Live Event.
    /// </summary>
    public partial class LiveEvent : Pulumi.CustomResource
    {
        /// <summary>
        /// The exact time the Live Event was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The Live Event access policies.
        /// </summary>
        [Output("crossSiteAccessPolicies")]
        public Output<Outputs.CrossSiteAccessPoliciesResponseResult?> CrossSiteAccessPolicies { get; private set; } = null!;

        /// <summary>
        /// The Live Event description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Live Event encoding.
        /// </summary>
        [Output("encoding")]
        public Output<Outputs.LiveEventEncodingResponseResult?> Encoding { get; private set; } = null!;

        /// <summary>
        /// The Live Event input.
        /// </summary>
        [Output("input")]
        public Output<Outputs.LiveEventInputResponseResult> Input { get; private set; } = null!;

        /// <summary>
        /// The exact time the Live Event was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The Azure Region of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Live Event preview.
        /// </summary>
        [Output("preview")]
        public Output<Outputs.LiveEventPreviewResponseResult?> Preview { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the Live Event.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource state of the Live Event.
        /// </summary>
        [Output("resourceState")]
        public Output<string> ResourceState { get; private set; } = null!;

        /// <summary>
        /// The stream options.
        /// </summary>
        [Output("streamOptions")]
        public Output<ImmutableArray<string>> StreamOptions { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The Live Event vanity URL flag.
        /// </summary>
        [Output("vanityUrl")]
        public Output<bool?> VanityUrl { get; private set; } = null!;


        /// <summary>
        /// Create a LiveEvent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LiveEvent(string name, LiveEventArgs args, CustomResourceOptions? options = null)
            : base("azurerm:media/v20180330preview:LiveEvent", name, args ?? new LiveEventArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LiveEvent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:media/v20180330preview:LiveEvent", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:media/latest:LiveEvent"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180601preview:LiveEvent"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180701:LiveEvent"},
                    new Pulumi.Alias { Type = "azurerm:media/v20190501preview:LiveEvent"},
                    new Pulumi.Alias { Type = "azurerm:media/v20200501:LiveEvent"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LiveEvent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LiveEvent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LiveEvent(name, id, options);
        }
    }

    public sealed class LiveEventArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The flag indicates if auto start the Live Event.
        /// </summary>
        [Input("autoStart")]
        public Input<bool>? AutoStart { get; set; }

        /// <summary>
        /// The Live Event access policies.
        /// </summary>
        [Input("crossSiteAccessPolicies")]
        public Input<Inputs.CrossSiteAccessPoliciesArgs>? CrossSiteAccessPolicies { get; set; }

        /// <summary>
        /// The Live Event description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Live Event encoding.
        /// </summary>
        [Input("encoding")]
        public Input<Inputs.LiveEventEncodingArgs>? Encoding { get; set; }

        /// <summary>
        /// The Live Event input.
        /// </summary>
        [Input("input", required: true)]
        public Input<Inputs.LiveEventInputArgs> Input { get; set; } = null!;

        /// <summary>
        /// The name of the Live Event.
        /// </summary>
        [Input("liveEventName", required: true)]
        public Input<string> LiveEventName { get; set; } = null!;

        /// <summary>
        /// The Azure Region of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Live Event preview.
        /// </summary>
        [Input("preview")]
        public Input<Inputs.LiveEventPreviewArgs>? Preview { get; set; }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("streamOptions")]
        private InputList<string>? _streamOptions;

        /// <summary>
        /// The stream options.
        /// </summary>
        public InputList<string> StreamOptions
        {
            get => _streamOptions ?? (_streamOptions = new InputList<string>());
            set => _streamOptions = value;
        }

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
        /// The Live Event vanity URL flag.
        /// </summary>
        [Input("vanityUrl")]
        public Input<bool>? VanityUrl { get; set; }

        public LiveEventArgs()
        {
        }
    }
}
