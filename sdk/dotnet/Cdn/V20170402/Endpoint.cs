// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20170402
{
    /// <summary>
    /// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format &lt;endpointname&gt;.azureedge.net.
    /// </summary>
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The JSON object that contains the properties required to create an endpoint.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.EndpointPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("azurerm:cdn/v20170402:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:cdn/v20170402:Endpoint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        [Input("contentTypesToCompress")]
        private InputList<string>? _contentTypesToCompress;

        /// <summary>
        /// List of content types on which compression applies. The value should be a valid MIME type.
        /// </summary>
        public InputList<string> ContentTypesToCompress
        {
            get => _contentTypesToCompress ?? (_contentTypesToCompress = new InputList<string>());
            set => _contentTypesToCompress = value;
        }

        [Input("geoFilters")]
        private InputList<Inputs.GeoFilterArgs>? _geoFilters;

        /// <summary>
        /// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
        /// </summary>
        public InputList<Inputs.GeoFilterArgs> GeoFilters
        {
            get => _geoFilters ?? (_geoFilters = new InputList<Inputs.GeoFilterArgs>());
            set => _geoFilters = value;
        }

        /// <summary>
        /// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
        /// </summary>
        [Input("isCompressionEnabled")]
        public Input<bool>? IsCompressionEnabled { get; set; }

        /// <summary>
        /// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        /// </summary>
        [Input("isHttpAllowed")]
        public Input<bool>? IsHttpAllowed { get; set; }

        /// <summary>
        /// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        /// </summary>
        [Input("isHttpsAllowed")]
        public Input<bool>? IsHttpsAllowed { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
        /// </summary>
        [Input("optimizationType")]
        public Input<string>? OptimizationType { get; set; }

        /// <summary>
        /// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
        /// </summary>
        [Input("originHostHeader")]
        public Input<string>? OriginHostHeader { get; set; }

        /// <summary>
        /// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
        /// </summary>
        [Input("originPath")]
        public Input<string>? OriginPath { get; set; }

        [Input("origins", required: true)]
        private InputList<Inputs.DeepCreatedOriginArgs>? _origins;

        /// <summary>
        /// The source of the content being delivered via CDN.
        /// </summary>
        public InputList<Inputs.DeepCreatedOriginArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.DeepCreatedOriginArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
        /// </summary>
        [Input("probePath")]
        public Input<string>? ProbePath { get; set; }

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
        /// </summary>
        [Input("queryStringCachingBehavior")]
        public Input<string>? QueryStringCachingBehavior { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EndpointArgs()
        {
        }
    }
}
