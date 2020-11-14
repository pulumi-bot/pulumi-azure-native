// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearningServices.V20200801
{
    /// <summary>
    /// Workspace connection.
    /// </summary>
    public partial class WorkspaceConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Authorization type of the workspace connection.
        /// </summary>
        [Output("authType")]
        public Output<string?> AuthType { get; private set; } = null!;

        /// <summary>
        /// Category of the workspace connection.
        /// </summary>
        [Output("category")]
        public Output<string?> Category { get; private set; } = null!;

        /// <summary>
        /// Friendly name of the workspace connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Target of the workspace connection.
        /// </summary>
        [Output("target")]
        public Output<string?> Target { get; private set; } = null!;

        /// <summary>
        /// Resource type of workspace connection.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Value details of the workspace connection.
        /// </summary>
        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspaceConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspaceConnection(string name, WorkspaceConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:machinelearningservices/v20200801:WorkspaceConnection", name, args ?? new WorkspaceConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspaceConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:machinelearningservices/v20200801:WorkspaceConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:machinelearningservices/latest:WorkspaceConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:machinelearningservices/v20200601:WorkspaceConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:machinelearningservices/v20200901preview:WorkspaceConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkspaceConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspaceConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkspaceConnection(name, id, options);
        }
    }

    public sealed class WorkspaceConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authorization type of the workspace connection.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// Category of the workspace connection.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Friendly name of the workspace connection
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Friendly name of the workspace connection
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Target of the workspace connection.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// Value details of the workspace connection.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public WorkspaceConnectionArgs()
        {
        }
    }
}
