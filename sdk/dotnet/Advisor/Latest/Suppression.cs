// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Advisor.Latest
{
    /// <summary>
    /// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
    /// </summary>
    public partial class Suppression : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The GUID of the suppression.
        /// </summary>
        [Output("suppressionId")]
        public Output<string?> SuppressionId { get; private set; } = null!;

        /// <summary>
        /// The duration for which the suppression is valid.
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Suppression resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Suppression(string name, SuppressionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:advisor/latest:Suppression", name, args ?? new SuppressionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Suppression(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:advisor/latest:Suppression", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:advisor/v20160712preview:Suppression"},
                    new Pulumi.Alias { Type = "azurerm:advisor/v20170331:Suppression"},
                    new Pulumi.Alias { Type = "azurerm:advisor/v20170419:Suppression"},
                    new Pulumi.Alias { Type = "azurerm:advisor/v20200101:Suppression"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Suppression resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Suppression Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Suppression(name, id, options);
        }
    }

    public sealed class SuppressionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the suppression.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The recommendation ID.
        /// </summary>
        [Input("recommendationId", required: true)]
        public Input<string> RecommendationId { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        /// <summary>
        /// The GUID of the suppression.
        /// </summary>
        [Input("suppressionId")]
        public Input<string>? SuppressionId { get; set; }

        /// <summary>
        /// The duration for which the suppression is valid.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public SuppressionArgs()
        {
        }
    }
}
