// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20200101
{
    /// <summary>
    /// Security assessment metadata
    /// </summary>
    public partial class AssessmentMetadataInSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes properties of an assessment metadata.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SecurityAssessmentMetadataPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AssessmentMetadataInSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssessmentMetadataInSubscription(string name, AssessmentMetadataInSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:security/v20200101:AssessmentMetadataInSubscription", name, args ?? new AssessmentMetadataInSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssessmentMetadataInSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:security/v20200101:AssessmentMetadataInSubscription", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AssessmentMetadataInSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssessmentMetadataInSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AssessmentMetadataInSubscription(name, id, options);
        }
    }

    public sealed class AssessmentMetadataInSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Assessment Key - Unique key for the assessment type
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Describes properties of an assessment metadata.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.SecurityAssessmentMetadataPropertiesArgs>? Properties { get; set; }

        public AssessmentMetadataInSubscriptionArgs()
        {
        }
    }
}