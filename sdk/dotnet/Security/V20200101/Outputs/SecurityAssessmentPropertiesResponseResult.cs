// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20200101.Outputs
{

    [OutputType]
    public sealed class SecurityAssessmentPropertiesResponseResult
    {
        /// <summary>
        /// Additional data regarding the assessment
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AdditionalData;
        /// <summary>
        /// User friendly display name of the assessment
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Links relevant to the assessment
        /// </summary>
        public readonly Outputs.AssessmentLinksResponseResult? Links;
        /// <summary>
        /// Describes properties of an assessment metadata.
        /// </summary>
        public readonly Outputs.SecurityAssessmentMetadataPropertiesResponseResult? Metadata;
        /// <summary>
        /// Data regarding 3rd party partner integration
        /// </summary>
        public readonly Outputs.SecurityAssessmentPartnerDataResponseResult? PartnersData;
        /// <summary>
        /// Details of the resource that was assessed
        /// </summary>
        public readonly Outputs.ResourceDetailsResponseResult ResourceDetails;
        /// <summary>
        /// The result of the assessment
        /// </summary>
        public readonly Outputs.AssessmentStatusResponseResult Status;

        [OutputConstructor]
        private SecurityAssessmentPropertiesResponseResult(
            ImmutableDictionary<string, string>? additionalData,

            string displayName,

            Outputs.AssessmentLinksResponseResult? links,

            Outputs.SecurityAssessmentMetadataPropertiesResponseResult? metadata,

            Outputs.SecurityAssessmentPartnerDataResponseResult? partnersData,

            Outputs.ResourceDetailsResponseResult resourceDetails,

            Outputs.AssessmentStatusResponseResult status)
        {
            AdditionalData = additionalData;
            DisplayName = displayName;
            Links = links;
            Metadata = metadata;
            PartnersData = partnersData;
            ResourceDetails = resourceDetails;
            Status = status;
        }
    }
}