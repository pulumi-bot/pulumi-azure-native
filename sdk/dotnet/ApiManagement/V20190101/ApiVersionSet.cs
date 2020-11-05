// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20190101
{
    /// <summary>
    /// Api Version Set Contract details.
    /// </summary>
    public partial class ApiVersionSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of API Version Set.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of API Version Set
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
        /// </summary>
        [Output("versionHeaderName")]
        public Output<string?> VersionHeaderName { get; private set; } = null!;

        /// <summary>
        /// Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
        /// </summary>
        [Output("versionQueryName")]
        public Output<string?> VersionQueryName { get; private set; } = null!;

        /// <summary>
        /// An value that determines where the API Version identifer will be located in a HTTP request.
        /// </summary>
        [Output("versioningScheme")]
        public Output<string> VersioningScheme { get; private set; } = null!;


        /// <summary>
        /// Create a ApiVersionSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiVersionSet(string name, ApiVersionSetArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:apimanagement/v20190101:ApiVersionSet", name, args ?? new ApiVersionSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiVersionSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:apimanagement/v20190101:ApiVersionSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/latest:ApiVersionSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20170301:ApiVersionSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20180101:ApiVersionSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20180601preview:ApiVersionSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20191201:ApiVersionSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20191201preview:ApiVersionSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20200601preview:ApiVersionSet"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiVersionSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiVersionSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiVersionSet(name, id, options);
        }
    }

    public sealed class ApiVersionSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of API Version Set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of API Version Set
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
        /// </summary>
        [Input("versionHeaderName")]
        public Input<string>? VersionHeaderName { get; set; }

        /// <summary>
        /// Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
        /// </summary>
        [Input("versionQueryName")]
        public Input<string>? VersionQueryName { get; set; }

        /// <summary>
        /// Api Version Set identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("versionSetId", required: true)]
        public Input<string> VersionSetId { get; set; } = null!;

        /// <summary>
        /// An value that determines where the API Version identifer will be located in a HTTP request.
        /// </summary>
        [Input("versioningScheme", required: true)]
        public Input<string> VersioningScheme { get; set; } = null!;

        public ApiVersionSetArgs()
        {
        }
    }
}
