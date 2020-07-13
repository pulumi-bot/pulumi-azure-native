// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn
{
    /// <summary>
    /// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
    /// </summary>
    public partial class ProfileEndpointOriginGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The JSON object that contains the properties of the origin group.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.OriginGroupPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileEndpointOriginGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileEndpointOriginGroup(string name, ProfileEndpointOriginGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:cdn:ProfileEndpointOriginGroup", name, args ?? new ProfileEndpointOriginGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileEndpointOriginGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:cdn:ProfileEndpointOriginGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileEndpointOriginGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileEndpointOriginGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProfileEndpointOriginGroup(name, id, options);
        }
    }

    public sealed class ProfileEndpointOriginGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// Name of the origin group which is unique within the endpoint.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// The JSON object that contains the properties of the origin group.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.OriginGroupPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ProfileEndpointOriginGroupArgs()
        {
        }
    }
}
