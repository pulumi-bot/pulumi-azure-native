// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.Outputs
{

    [OutputType]
    public sealed class LiveEventPropertiesResponseResult
    {
        /// <summary>
        /// The exact time the Live Event was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The Live Event access policies.
        /// </summary>
        public readonly Outputs.CrossSiteAccessPoliciesResponseResult? CrossSiteAccessPolicies;
        /// <summary>
        /// The Live Event description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The Live Event encoding.
        /// </summary>
        public readonly Outputs.LiveEventEncodingResponseResult? Encoding;
        /// <summary>
        /// The Live Event input.
        /// </summary>
        public readonly Outputs.LiveEventInputResponseResult Input;
        /// <summary>
        /// The exact time the Live Event was last modified.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The Live Event preview.
        /// </summary>
        public readonly Outputs.LiveEventPreviewResponseResult? Preview;
        /// <summary>
        /// The provisioning state of the Live Event.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource state of the Live Event.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
        /// </summary>
        public readonly ImmutableArray<string> StreamOptions;
        /// <summary>
        /// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
        /// </summary>
        public readonly bool? UseStaticHostname;

        [OutputConstructor]
        private LiveEventPropertiesResponseResult(
            string created,

            Outputs.CrossSiteAccessPoliciesResponseResult? crossSiteAccessPolicies,

            string? description,

            Outputs.LiveEventEncodingResponseResult? encoding,

            Outputs.LiveEventInputResponseResult input,

            string lastModified,

            Outputs.LiveEventPreviewResponseResult? preview,

            string provisioningState,

            string resourceState,

            ImmutableArray<string> streamOptions,

            bool? useStaticHostname)
        {
            Created = created;
            CrossSiteAccessPolicies = crossSiteAccessPolicies;
            Description = description;
            Encoding = encoding;
            Input = input;
            LastModified = lastModified;
            Preview = preview;
            ProvisioningState = provisioningState;
            ResourceState = resourceState;
            StreamOptions = streamOptions;
            UseStaticHostname = useStaticHostname;
        }
    }
}