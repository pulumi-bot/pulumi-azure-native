// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class LiveEventPreviewResponseResult
    {
        /// <summary>
        /// The access control for LiveEvent preview.
        /// </summary>
        public readonly Outputs.LiveEventPreviewAccessControlResponseResult? AccessControl;
        /// <summary>
        /// An Alternative Media Identifier associated with the StreamingLocator created for the preview.  This value is specified at creation time and cannot be updated.  The identifier can be used in the CustomLicenseAcquisitionUrlTemplate or the CustomKeyAcquisitionUrlTemplate of the StreamingPolicy specified in the StreamingPolicyName field.
        /// </summary>
        public readonly string? AlternativeMediaId;
        /// <summary>
        /// The endpoints for preview.
        /// </summary>
        public readonly ImmutableArray<Outputs.LiveEventEndpointResponseResult> Endpoints;
        /// <summary>
        /// The identifier of the preview locator in Guid format.  Specifying this at creation time allows the caller to know the preview locator url before the event is created.  If omitted, the service will generate a random identifier.  This value cannot be updated once the live event is created.
        /// </summary>
        public readonly string? PreviewLocator;
        /// <summary>
        /// The name of streaming policy used for the LiveEvent preview.  This value is specified at creation time and cannot be updated.
        /// </summary>
        public readonly string? StreamingPolicyName;

        [OutputConstructor]
        private LiveEventPreviewResponseResult(
            Outputs.LiveEventPreviewAccessControlResponseResult? accessControl,

            string? alternativeMediaId,

            ImmutableArray<Outputs.LiveEventEndpointResponseResult> endpoints,

            string? previewLocator,

            string? streamingPolicyName)
        {
            AccessControl = accessControl;
            AlternativeMediaId = alternativeMediaId;
            Endpoints = endpoints;
            PreviewLocator = previewLocator;
            StreamingPolicyName = streamingPolicyName;
        }
    }
}
