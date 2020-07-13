// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions
{
    /// <summary>
    /// Information about managed application definition.
    /// </summary>
    public partial class ApplicationDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// ID of the resource that manages this resource.
        /// </summary>
        [Output("managedBy")]
        public Output<string?> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The managed application definition properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ApplicationDefinitionPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The SKU of the resource.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

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
        /// Create a ApplicationDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationDefinition(string name, ApplicationDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:solutions:ApplicationDefinition", name, args ?? new ApplicationDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:solutions:ApplicationDefinition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApplicationDefinition(name, id, options);
        }
    }

    public sealed class ApplicationDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// ID of the resource that manages this resource.
        /// </summary>
        [Input("managedBy")]
        public Input<string>? ManagedBy { get; set; }

        /// <summary>
        /// The name of the managed application definition.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The managed application definition properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.ApplicationDefinitionPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the resource.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

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

        public ApplicationDefinitionArgs()
        {
        }
    }
}
