// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20190801
{
    /// <summary>
    /// Hybrid Connection contract. This is used to configure a Hybrid Connection.
    /// </summary>
    public partial class WebAppHybridConnectionSlot : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// HybridConnection resource specific properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.HybridConnectionResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppHybridConnectionSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppHybridConnectionSlot(string name, WebAppHybridConnectionSlotArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20190801:WebAppHybridConnectionSlot", name, args ?? new WebAppHybridConnectionSlotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppHybridConnectionSlot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20190801:WebAppHybridConnectionSlot", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WebAppHybridConnectionSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppHybridConnectionSlot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppHybridConnectionSlot(name, id, options);
        }
    }

    public sealed class WebAppHybridConnectionSlotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The relay name for this hybrid connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace for this hybrid connection.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// HybridConnection resource specific properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.HybridConnectionPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the slot for the web app.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        public WebAppHybridConnectionSlotArgs()
        {
        }
    }
}