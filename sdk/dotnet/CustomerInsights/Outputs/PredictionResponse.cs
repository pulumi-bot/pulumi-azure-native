// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Outputs
{

    [OutputType]
    public sealed class PredictionResponse
    {
        /// <summary>
        /// Whether do auto analyze.
        /// </summary>
        public readonly bool AutoAnalyze;
        /// <summary>
        /// Description of the prediction.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Description;
        /// <summary>
        /// Display name of the prediction.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// The prediction grades.
        /// </summary>
        public readonly ImmutableArray<Outputs.PredictionResponseProperties> Grades;
        /// <summary>
        /// Interaction types involved in the prediction.
        /// </summary>
        public readonly ImmutableArray<string> InvolvedInteractionTypes;
        /// <summary>
        /// KPI types involved in the prediction.
        /// </summary>
        public readonly ImmutableArray<string> InvolvedKpiTypes;
        /// <summary>
        /// Relationships involved in the prediction.
        /// </summary>
        public readonly ImmutableArray<string> InvolvedRelationships;
        /// <summary>
        /// Definition of the link mapping of prediction.
        /// </summary>
        public readonly Outputs.PredictionResponseProperties Mappings;
        /// <summary>
        /// Negative outcome expression.
        /// </summary>
        public readonly string NegativeOutcomeExpression;
        /// <summary>
        /// Positive outcome expression.
        /// </summary>
        public readonly string PositiveOutcomeExpression;
        /// <summary>
        /// Name of the prediction.
        /// </summary>
        public readonly string? PredictionName;
        /// <summary>
        /// Primary profile type.
        /// </summary>
        public readonly string PrimaryProfileType;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Scope expression.
        /// </summary>
        public readonly string ScopeExpression;
        /// <summary>
        /// Score label.
        /// </summary>
        public readonly string ScoreLabel;
        /// <summary>
        /// System generated entities.
        /// </summary>
        public readonly Outputs.PredictionResponseProperties SystemGeneratedEntities;
        /// <summary>
        /// The hub name.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private PredictionResponse(
            bool autoAnalyze,

            ImmutableDictionary<string, string>? description,

            ImmutableDictionary<string, string>? displayName,

            ImmutableArray<Outputs.PredictionResponseProperties> grades,

            ImmutableArray<string> involvedInteractionTypes,

            ImmutableArray<string> involvedKpiTypes,

            ImmutableArray<string> involvedRelationships,

            Outputs.PredictionResponseProperties mappings,

            string negativeOutcomeExpression,

            string positiveOutcomeExpression,

            string? predictionName,

            string primaryProfileType,

            string provisioningState,

            string scopeExpression,

            string scoreLabel,

            Outputs.PredictionResponseProperties systemGeneratedEntities,

            string tenantId)
        {
            AutoAnalyze = autoAnalyze;
            Description = description;
            DisplayName = displayName;
            Grades = grades;
            InvolvedInteractionTypes = involvedInteractionTypes;
            InvolvedKpiTypes = involvedKpiTypes;
            InvolvedRelationships = involvedRelationships;
            Mappings = mappings;
            NegativeOutcomeExpression = negativeOutcomeExpression;
            PositiveOutcomeExpression = positiveOutcomeExpression;
            PredictionName = predictionName;
            PrimaryProfileType = primaryProfileType;
            ProvisioningState = provisioningState;
            ScopeExpression = scopeExpression;
            ScoreLabel = scoreLabel;
            SystemGeneratedEntities = systemGeneratedEntities;
            TenantId = tenantId;
        }
    }
}