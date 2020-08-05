// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201
{
    /// <summary>
    /// Issue Contract details.
    /// </summary>
    public partial class ApiIssue : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the Issue.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.IssueContractPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApiIssue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiIssue(string name, ApiIssueArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20191201:ApiIssue", name, args ?? new ApiIssueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiIssue(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20191201:ApiIssue", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ApiIssue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiIssue Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiIssue(name, id, options);
        }
    }

    public sealed class ApiIssueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A resource identifier for the API the issue was created for.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Date and time when the issue was created.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// Text describing the issue.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Issue identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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
        /// Status of the issue.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The issue title.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// A resource identifier for the user created the issue.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public ApiIssueArgs()
        {
        }
    }
}
