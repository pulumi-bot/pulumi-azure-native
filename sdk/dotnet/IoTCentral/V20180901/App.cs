// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.IoTCentral.V20180901
{
    /// <summary>
    /// The IoT Central application.
    /// </summary>
    public partial class App : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ARM resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The common properties of an IoT Central application.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.AppPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// A valid instance SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.AppSkuInfoResponseResult> Sku { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs args, CustomResourceOptions? options = null)
            : base("azurerm:iotcentral/v20180901:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:iotcentral/v20180901:App", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new App(name, id, options);
        }
    }

    public sealed class AppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the application.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The ARM resource name of the IoT Central application.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the IoT Central application.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A valid instance SKU.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.AppSkuInfoArgs> Sku { get; set; } = null!;

        /// <summary>
        /// The subdomain of the application.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

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

        /// <summary>
        /// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public AppArgs()
        {
        }
    }
}
