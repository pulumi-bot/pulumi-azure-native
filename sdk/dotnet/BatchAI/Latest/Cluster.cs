// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.BatchAI.Latest
{
    /// <summary>
    /// Information about a Cluster.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Allocation state of the cluster. Possible values are: steady - Indicates that the cluster is not resizing. There are no changes to the number of compute nodes in the cluster in progress. A cluster enters this state when it is created and when no operations are being performed on the cluster to change the number of compute nodes. resizing - Indicates that the cluster is resizing; that is, compute nodes are being added to or removed from the cluster.
        /// </summary>
        [Output("allocationState")]
        public Output<string> AllocationState { get; private set; } = null!;

        /// <summary>
        /// The time at which the cluster entered its current allocation state.
        /// </summary>
        [Output("allocationStateTransitionTime")]
        public Output<string> AllocationStateTransitionTime { get; private set; } = null!;

        /// <summary>
        /// The time when the cluster was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The number of compute nodes currently assigned to the cluster.
        /// </summary>
        [Output("currentNodeCount")]
        public Output<int> CurrentNodeCount { get; private set; } = null!;

        /// <summary>
        /// Collection of errors encountered by various compute nodes during node setup.
        /// </summary>
        [Output("errors")]
        public Output<ImmutableArray<Outputs.BatchAIErrorResponse>> Errors { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Setup (mount file systems, performance counters settings and custom setup task) to be performed on each compute node in the cluster.
        /// </summary>
        [Output("nodeSetup")]
        public Output<Outputs.NodeSetupResponse?> NodeSetup { get; private set; } = null!;

        /// <summary>
        /// Counts of various node states on the cluster.
        /// </summary>
        [Output("nodeStateCounts")]
        public Output<Outputs.NodeStateCountsResponse> NodeStateCounts { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the cluster. Possible value are: creating - Specifies that the cluster is being created. succeeded - Specifies that the cluster has been created successfully. failed - Specifies that the cluster creation has failed. deleting - Specifies that the cluster is being deleted.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Time when the provisioning state was changed.
        /// </summary>
        [Output("provisioningStateTransitionTime")]
        public Output<string> ProvisioningStateTransitionTime { get; private set; } = null!;

        /// <summary>
        /// Scale settings of the cluster.
        /// </summary>
        [Output("scaleSettings")]
        public Output<Outputs.ScaleSettingsResponse?> ScaleSettings { get; private set; } = null!;

        /// <summary>
        /// Virtual network subnet resource ID the cluster nodes belong to.
        /// </summary>
        [Output("subnet")]
        public Output<Outputs.ResourceIdResponse?> Subnet { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Administrator user account settings which can be used to SSH to compute nodes.
        /// </summary>
        [Output("userAccountSettings")]
        public Output<Outputs.UserAccountSettingsResponse?> UserAccountSettings { get; private set; } = null!;

        /// <summary>
        /// Virtual machine configuration (OS image) of the compute nodes. All nodes in a cluster have the same OS image configuration.
        /// </summary>
        [Output("virtualMachineConfiguration")]
        public Output<Outputs.VirtualMachineConfigurationResponse?> VirtualMachineConfiguration { get; private set; } = null!;

        /// <summary>
        /// VM priority of cluster nodes.
        /// </summary>
        [Output("vmPriority")]
        public Output<string?> VmPriority { get; private set; } = null!;

        /// <summary>
        /// The size of the virtual machines in the cluster. All nodes in a cluster have the same VM size.
        /// </summary>
        [Output("vmSize")]
        public Output<string?> VmSize { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:batchai/latest:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:batchai/latest:Cluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:batchai/v20180501:Cluster"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Setup to be performed on each compute node in the cluster.
        /// </summary>
        [Input("nodeSetup")]
        public Input<Inputs.NodeSetupArgs>? NodeSetup { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Scale settings for the cluster. Batch AI service supports manual and auto scale clusters.
        /// </summary>
        [Input("scaleSettings")]
        public Input<Inputs.ScaleSettingsArgs>? ScaleSettings { get; set; }

        /// <summary>
        /// Existing virtual network subnet to put the cluster nodes in. Note, if a File Server mount configured in node setup, the File Server's subnet will be used automatically.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.ResourceIdArgs>? Subnet { get; set; }

        /// <summary>
        /// Settings for an administrator user account that will be created on each compute node in the cluster.
        /// </summary>
        [Input("userAccountSettings", required: true)]
        public Input<Inputs.UserAccountSettingsArgs> UserAccountSettings { get; set; } = null!;

        /// <summary>
        /// OS image configuration for cluster nodes. All nodes in a cluster have the same OS image.
        /// </summary>
        [Input("virtualMachineConfiguration")]
        public Input<Inputs.VirtualMachineConfigurationArgs>? VirtualMachineConfiguration { get; set; }

        /// <summary>
        /// VM priority. Allowed values are: dedicated (default) and lowpriority.
        /// </summary>
        [Input("vmPriority")]
        public Input<Pulumi.AzureNextGen.BatchAI.Latest.VmPriority>? VmPriority { get; set; }

        /// <summary>
        /// The size of the virtual machines in the cluster. All nodes in a cluster have the same VM size. For information about available VM sizes for clusters using images from the Virtual Machines Marketplace see Sizes for Virtual Machines (Linux). Batch AI service supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        /// <summary>
        /// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ClusterArgs()
        {
        }
    }
}
