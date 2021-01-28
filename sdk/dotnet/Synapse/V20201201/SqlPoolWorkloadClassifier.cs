// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20201201
{
    /// <summary>
    /// Workload classifier operations for a data warehouse
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:synapse/v20201201:SqlPoolWorkloadClassifier")]
    public partial class SqlPoolWorkloadClassifier : Pulumi.CustomResource
    {
        /// <summary>
        /// The workload classifier context.
        /// </summary>
        [Output("context")]
        public Output<string?> Context { get; private set; } = null!;

        /// <summary>
        /// The workload classifier end time for classification.
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// The workload classifier importance.
        /// </summary>
        [Output("importance")]
        public Output<string?> Importance { get; private set; } = null!;

        /// <summary>
        /// The workload classifier label.
        /// </summary>
        [Output("label")]
        public Output<string?> Label { get; private set; } = null!;

        /// <summary>
        /// The workload classifier member name.
        /// </summary>
        [Output("memberName")]
        public Output<string> MemberName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The workload classifier start time for classification.
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SqlPoolWorkloadClassifier resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlPoolWorkloadClassifier(string name, SqlPoolWorkloadClassifierArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20201201:SqlPoolWorkloadClassifier", name, args ?? new SqlPoolWorkloadClassifierArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlPoolWorkloadClassifier(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20201201:SqlPoolWorkloadClassifier", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:synapse/v20190601preview:SqlPoolWorkloadClassifier"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlPoolWorkloadClassifier resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlPoolWorkloadClassifier Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlPoolWorkloadClassifier(name, id, options);
        }
    }

    public sealed class SqlPoolWorkloadClassifierArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The workload classifier context.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// The workload classifier end time for classification.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The workload classifier importance.
        /// </summary>
        [Input("importance")]
        public Input<string>? Importance { get; set; }

        /// <summary>
        /// The workload classifier label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The workload classifier member name.
        /// </summary>
        [Input("memberName", required: true)]
        public Input<string> MemberName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SQL pool name
        /// </summary>
        [Input("sqlPoolName", required: true)]
        public Input<string> SqlPoolName { get; set; } = null!;

        /// <summary>
        /// The workload classifier start time for classification.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The name of the workload classifier.
        /// </summary>
        [Input("workloadClassifierName", required: true)]
        public Input<string> WorkloadClassifierName { get; set; } = null!;

        /// <summary>
        /// The name of the workload group.
        /// </summary>
        [Input("workloadGroupName", required: true)]
        public Input<string> WorkloadGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public SqlPoolWorkloadClassifierArgs()
        {
        }
    }
}
