// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Portal.V20181001Preview
{
    /// <summary>
    /// The shared dashboard resource definition.
    /// </summary>
    public partial class Dashboard : Pulumi.CustomResource
    {
        /// <summary>
        /// The dashboard lenses.
        /// </summary>
        [Output("lenses")]
        public Output<ImmutableDictionary<string, Outputs.DashboardLensResponseResult>?> Lenses { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The dashboard metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, ImmutableDictionary<string, object>>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("azurerm:portal/v20181001preview:Dashboard", name, args ?? new DashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:portal/v20181001preview:Dashboard", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:portal/v20150801preview:Dashboard"},
                    new Pulumi.Alias { Type = "azurerm:portal/v20190101preview:Dashboard"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, options);
        }
    }

    public sealed class DashboardArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the dashboard.
        /// </summary>
        [Input("dashboardName", required: true)]
        public Input<string> DashboardName { get; set; } = null!;

        [Input("lenses")]
        private InputMap<Inputs.DashboardLensArgs>? _lenses;

        /// <summary>
        /// The dashboard lenses.
        /// </summary>
        public InputMap<Inputs.DashboardLensArgs> Lenses
        {
            get => _lenses ?? (_lenses = new InputMap<Inputs.DashboardLensArgs>());
            set => _lenses = value;
        }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("metadata")]
        private InputMap<ImmutableDictionary<string, object>>? _metadata;

        /// <summary>
        /// The dashboard metadata.
        /// </summary>
        public InputMap<ImmutableDictionary<string, object>> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<ImmutableDictionary<string, object>>());
            set => _metadata = value;
        }

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

        public DashboardArgs()
        {
        }
    }
}
