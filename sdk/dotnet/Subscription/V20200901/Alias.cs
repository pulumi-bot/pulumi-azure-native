// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Subscription.V20200901
{
    /// <summary>
    /// Subscription Information with the alias.
    /// </summary>
    public partial class Alias : Pulumi.CustomResource
    {
        /// <summary>
        /// Alias ID.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Put Alias response properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PutAliasResponsePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type, Microsoft.Subscription/aliases.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Alias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alias(string name, AliasArgs args, CustomResourceOptions? options = null)
            : base("azurerm:subscription/v20200901:Alias", name, args ?? new AliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alias(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:subscription/v20200901:Alias", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:subscription/latest:Alias"},
                    new Pulumi.Alias { Type = "azurerm:subscription/v20191001preview:Alias"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Alias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alias Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Alias(name, id, options);
        }
    }

    public sealed class AliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Input("aliasName", required: true)]
        public Input<string> AliasName { get; set; } = null!;

        /// <summary>
        /// Put alias request properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.PutAliasRequestPropertiesArgs> Properties { get; set; } = null!;

        public AliasArgs()
        {
        }
    }
}
