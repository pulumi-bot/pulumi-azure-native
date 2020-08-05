// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20160801
{
    /// <summary>
    /// User credentials used for publishing activity.
    /// </summary>
    public partial class WebAppDeployment : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Deployment resource specific properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DeploymentResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppDeployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppDeployment(string name, WebAppDeploymentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20160801:WebAppDeployment", name, args ?? new WebAppDeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppDeployment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20160801:WebAppDeployment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WebAppDeployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppDeployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppDeployment(name, id, options);
        }
    }

    public sealed class WebAppDeploymentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if deployment is currently active, false if completed and null if not started.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Who authored the deployment.
        /// </summary>
        [Input("author")]
        public Input<string>? Author { get; set; }

        /// <summary>
        /// Author email.
        /// </summary>
        [Input("authorEmail")]
        public Input<string>? AuthorEmail { get; set; }

        /// <summary>
        /// Who performed the deployment.
        /// </summary>
        [Input("deployer")]
        public Input<string>? Deployer { get; set; }

        /// <summary>
        /// Details on deployment.
        /// </summary>
        [Input("details")]
        public Input<string>? Details { get; set; }

        /// <summary>
        /// End time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Details about deployment status.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Identifier for deployment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Deployment status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public WebAppDeploymentArgs()
        {
        }
    }
}
