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
    /// A domain specific resource identifier.
    /// </summary>
    public partial class AppServiceSlotDomainOwnershipIdentifier : Pulumi.CustomResource
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
        /// Identifier resource specific properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.IdentifierResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AppServiceSlotDomainOwnershipIdentifier resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppServiceSlotDomainOwnershipIdentifier(string name, AppServiceSlotDomainOwnershipIdentifierArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web:AppServiceSlotDomainOwnershipIdentifier", name, args ?? new AppServiceSlotDomainOwnershipIdentifierArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppServiceSlotDomainOwnershipIdentifier(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web:AppServiceSlotDomainOwnershipIdentifier", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AppServiceSlotDomainOwnershipIdentifier resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppServiceSlotDomainOwnershipIdentifier Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppServiceSlotDomainOwnershipIdentifier(name, id, options);
        }
    }

    public sealed class AppServiceSlotDomainOwnershipIdentifierArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of domain ownership identifier.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Identifier resource specific properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.IdentifierPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        public AppServiceSlotDomainOwnershipIdentifierArgs()
        {
        }
    }
}
