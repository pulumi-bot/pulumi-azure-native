// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web
{
    /// <summary>
    /// A web app, a mobile app backend, or an API app.
    /// </summary>
    public partial class AppServiceSlot : Pulumi.CustomResource
    {
        /// <summary>
        /// Managed service identity.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Site resource specific properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SiteResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AppServiceSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppServiceSlot(string name, AppServiceSlotArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web:AppServiceSlot", name, args ?? new AppServiceSlotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppServiceSlot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web:AppServiceSlot", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AppServiceSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppServiceSlot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppServiceSlot(name, id, options);
        }
    }

    public sealed class AppServiceSlotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed service identity.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot to create or update. By default, this API attempts to create or modify the production slot.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Site resource specific properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.SitePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public AppServiceSlotArgs()
        {
        }
    }
}
