// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20190501
{
    /// <summary>
    /// The integration account map.
    /// </summary>
    public partial class IntegrationAccountMap : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Gets the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The integration account map properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.IntegrationAccountMapPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Gets the resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationAccountMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationAccountMap(string name, IntegrationAccountMapArgs args, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20190501:IntegrationAccountMap", name, args ?? new IntegrationAccountMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationAccountMap(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20190501:IntegrationAccountMap", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationAccountMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationAccountMap Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IntegrationAccountMap(name, id, options);
        }
    }

    public sealed class IntegrationAccountMapArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The content type.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public Input<string> IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The map type.
        /// </summary>
        [Input("mapType", required: true)]
        public Input<string> MapType { get; set; } = null!;

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The integration account map name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The parameters schema of integration account map.
        /// </summary>
        [Input("parametersSchema")]
        public Input<Inputs.IntegrationAccountMapPropertiesPropertiesArgs>? ParametersSchema { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IntegrationAccountMapArgs()
        {
        }
    }
}
