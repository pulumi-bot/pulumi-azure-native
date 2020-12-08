// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.SecurityInsights.V20190101Preview
{
    /// <summary>
    /// Threat intelligence information object.
    /// </summary>
    public partial class ThreatIntelligenceIndicator : Pulumi.CustomResource
    {
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The kind of the entity.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Azure resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ThreatIntelligenceIndicator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThreatIntelligenceIndicator(string name, ThreatIntelligenceIndicatorArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:securityinsights/v20190101preview:ThreatIntelligenceIndicator", name, args ?? new ThreatIntelligenceIndicatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ThreatIntelligenceIndicator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:securityinsights/v20190101preview:ThreatIntelligenceIndicator", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ThreatIntelligenceIndicator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThreatIntelligenceIndicator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ThreatIntelligenceIndicator(name, id, options);
        }
    }

    public sealed class ThreatIntelligenceIndicatorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Confidence of threat intelligence entity
        /// </summary>
        [Input("confidence")]
        public Input<int>? Confidence { get; set; }

        /// <summary>
        /// Created by
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Created by reference of threat intelligence entity
        /// </summary>
        [Input("createdByRef")]
        public Input<string>? CreatedByRef { get; set; }

        /// <summary>
        /// Description of a threat intelligence entity
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name of a threat intelligence entity
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// External ID of threat intelligence entity
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("externalReferences")]
        private InputList<string>? _externalReferences;

        /// <summary>
        /// External References
        /// </summary>
        public InputList<string> ExternalReferences
        {
            get => _externalReferences ?? (_externalReferences = new InputList<string>());
            set => _externalReferences = value;
        }

        [Input("granularMarkings")]
        private InputList<Inputs.ThreatIntelligenceGranularMarkingModelArgs>? _granularMarkings;

        /// <summary>
        /// Granular Markings
        /// </summary>
        public InputList<Inputs.ThreatIntelligenceGranularMarkingModelArgs> GranularMarkings
        {
            get => _granularMarkings ?? (_granularMarkings = new InputList<Inputs.ThreatIntelligenceGranularMarkingModelArgs>());
            set => _granularMarkings = value;
        }

        [Input("indicatorTypes")]
        private InputList<string>? _indicatorTypes;

        /// <summary>
        /// Indicator types of threat intelligence entities
        /// </summary>
        public InputList<string> IndicatorTypes
        {
            get => _indicatorTypes ?? (_indicatorTypes = new InputList<string>());
            set => _indicatorTypes = value;
        }

        [Input("killChainPhases")]
        private InputList<Inputs.ThreatIntelligenceKillChainPhaseArgs>? _killChainPhases;

        /// <summary>
        /// Kill chain phases
        /// </summary>
        public InputList<Inputs.ThreatIntelligenceKillChainPhaseArgs> KillChainPhases
        {
            get => _killChainPhases ?? (_killChainPhases = new InputList<Inputs.ThreatIntelligenceKillChainPhaseArgs>());
            set => _killChainPhases = value;
        }

        /// <summary>
        /// The kind of the entity.
        /// </summary>
        [Input("kind", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.SecurityInsights.V20190101Preview.ThreatIntelligenceResourceKind> Kind { get; set; } = null!;

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// Labels  of threat intelligence entity
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Last updated time in UTC
        /// </summary>
        [Input("lastUpdatedTimeUtc")]
        public Input<string>? LastUpdatedTimeUtc { get; set; }

        /// <summary>
        /// Modified by
        /// </summary>
        [Input("modified")]
        public Input<string>? Modified { get; set; }

        /// <summary>
        /// Threat intelligence indicator name field.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        /// </summary>
        [Input("operationalInsightsResourceProvider", required: true)]
        public Input<string> OperationalInsightsResourceProvider { get; set; } = null!;

        /// <summary>
        /// Pattern of a threat intelligence entity
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Pattern type of a threat intelligence entity
        /// </summary>
        [Input("patternType")]
        public Input<string>? PatternType { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Is threat intelligence entity revoked
        /// </summary>
        [Input("revoked")]
        public Input<bool>? Revoked { get; set; }

        /// <summary>
        /// Source of a threat intelligence entity
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("threatIntelligenceTags")]
        private InputList<string>? _threatIntelligenceTags;

        /// <summary>
        /// List of tags
        /// </summary>
        public InputList<string> ThreatIntelligenceTags
        {
            get => _threatIntelligenceTags ?? (_threatIntelligenceTags = new InputList<string>());
            set => _threatIntelligenceTags = value;
        }

        [Input("threatTypes")]
        private InputList<string>? _threatTypes;

        /// <summary>
        /// Threat types
        /// </summary>
        public InputList<string> ThreatTypes
        {
            get => _threatTypes ?? (_threatTypes = new InputList<string>());
            set => _threatTypes = value;
        }

        /// <summary>
        /// Valid from
        /// </summary>
        [Input("validFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// Valid until
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ThreatIntelligenceIndicatorArgs()
        {
        }
    }
}
