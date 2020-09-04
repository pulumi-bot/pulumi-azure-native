// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview
{
    public static class GetLiveEvent
    {
        public static Task<GetLiveEventResult> InvokeAsync(GetLiveEventArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLiveEventResult>("azurerm:media/v20180330preview:getLiveEvent", args ?? new GetLiveEventArgs(), options.WithVersion());
    }


    public sealed class GetLiveEventArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Live Event.
        /// </summary>
        [Input("liveEventName", required: true)]
        public string LiveEventName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLiveEventArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLiveEventResult
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
        /// The Azure Region of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
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
        /// The stream options.
        /// </summary>
        public readonly ImmutableArray<string> StreamOptions;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The Live Event vanity URL flag.
        /// </summary>
        public readonly bool? VanityUrl;

        [OutputConstructor]
        private GetLiveEventResult(
            string created,

            Outputs.CrossSiteAccessPoliciesResponseResult? crossSiteAccessPolicies,

            string? description,

            Outputs.LiveEventEncodingResponseResult? encoding,

            Outputs.LiveEventInputResponseResult input,

            string lastModified,

            string? location,

            string name,

            Outputs.LiveEventPreviewResponseResult? preview,

            string provisioningState,

            string resourceState,

            ImmutableArray<string> streamOptions,

            ImmutableDictionary<string, string>? tags,

            string type,

            bool? vanityUrl)
        {
            Created = created;
            CrossSiteAccessPolicies = crossSiteAccessPolicies;
            Description = description;
            Encoding = encoding;
            Input = input;
            LastModified = lastModified;
            Location = location;
            Name = name;
            Preview = preview;
            ProvisioningState = provisioningState;
            ResourceState = resourceState;
            StreamOptions = streamOptions;
            Tags = tags;
            Type = type;
            VanityUrl = vanityUrl;
        }
    }
}
