// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20190101Preview
{
    /// <summary>
    /// Security assessment metadata
    /// </summary>
    public partial class AssessmentsMetadataSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
        /// </summary>
        [Output("assessmentType")]
        public Output<string> AssessmentType { get; private set; } = null!;

        [Output("category")]
        public Output<ImmutableArray<string>> Category { get; private set; } = null!;

        /// <summary>
        /// Human readable description of the assessment
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User friendly display name of the assessment
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The implementation effort required to remediate this assessment
        /// </summary>
        [Output("implementationEffort")]
        public Output<string?> ImplementationEffort { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure resource ID of the policy definition that turns this assessment calculation on
        /// </summary>
        [Output("policyDefinitionId")]
        public Output<string> PolicyDefinitionId { get; private set; } = null!;

        /// <summary>
        /// True if this assessment is in preview release status
        /// </summary>
        [Output("preview")]
        public Output<bool?> Preview { get; private set; } = null!;

        /// <summary>
        /// Human readable description of what you should do to mitigate this security issue
        /// </summary>
        [Output("remediationDescription")]
        public Output<string?> RemediationDescription { get; private set; } = null!;

        /// <summary>
        /// The severity level of the assessment
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        [Output("threats")]
        public Output<ImmutableArray<string>> Threats { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The user impact of the assessment
        /// </summary>
        [Output("userImpact")]
        public Output<string?> UserImpact { get; private set; } = null!;


        /// <summary>
        /// Create a AssessmentsMetadataSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssessmentsMetadataSubscription(string name, AssessmentsMetadataSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:security/v20190101preview:AssessmentsMetadataSubscription", name, args ?? new AssessmentsMetadataSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssessmentsMetadataSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:security/v20190101preview:AssessmentsMetadataSubscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:security/latest:AssessmentsMetadataSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:security/v20200101:AssessmentsMetadataSubscription"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AssessmentsMetadataSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssessmentsMetadataSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AssessmentsMetadataSubscription(name, id, options);
        }
    }

    public sealed class AssessmentsMetadataSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Assessment Key - Unique key for the assessment type
        /// </summary>
        [Input("assessmentMetadataName", required: true)]
        public Input<string> AssessmentMetadataName { get; set; } = null!;

        /// <summary>
        /// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
        /// </summary>
        [Input("assessmentType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20190101Preview.AssessmentType> AssessmentType { get; set; } = null!;

        [Input("category")]
        private InputList<Union<string, Pulumi.AzureNextGen.Security.V20190101Preview.Category>>? _category;
        public InputList<Union<string, Pulumi.AzureNextGen.Security.V20190101Preview.Category>> Category
        {
            get => _category ?? (_category = new InputList<Union<string, Pulumi.AzureNextGen.Security.V20190101Preview.Category>>());
            set => _category = value;
        }

        /// <summary>
        /// Human readable description of the assessment
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User friendly display name of the assessment
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The implementation effort required to remediate this assessment
        /// </summary>
        [Input("implementationEffort")]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20190101Preview.ImplementationEffort>? ImplementationEffort { get; set; }

        /// <summary>
        /// True if this assessment is in preview release status
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// Human readable description of what you should do to mitigate this security issue
        /// </summary>
        [Input("remediationDescription")]
        public Input<string>? RemediationDescription { get; set; }

        /// <summary>
        /// The severity level of the assessment
        /// </summary>
        [Input("severity", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20190101Preview.Severity> Severity { get; set; } = null!;

        [Input("threats")]
        private InputList<Union<string, Pulumi.AzureNextGen.Security.V20190101Preview.Threats>>? _threats;
        public InputList<Union<string, Pulumi.AzureNextGen.Security.V20190101Preview.Threats>> Threats
        {
            get => _threats ?? (_threats = new InputList<Union<string, Pulumi.AzureNextGen.Security.V20190101Preview.Threats>>());
            set => _threats = value;
        }

        /// <summary>
        /// The user impact of the assessment
        /// </summary>
        [Input("userImpact")]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20190101Preview.UserImpact>? UserImpact { get; set; }

        public AssessmentsMetadataSubscriptionArgs()
        {
        }
    }
}
