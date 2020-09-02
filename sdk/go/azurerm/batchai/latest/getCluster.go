// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azurerm:batchai/latest:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	ClusterName string `pulumi:"clusterName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Information about a Cluster.
type LookupClusterResult struct {
	// Allocation state of the cluster. Possible values are: steady - Indicates that the cluster is not resizing. There are no changes to the number of compute nodes in the cluster in progress. A cluster enters this state when it is created and when no operations are being performed on the cluster to change the number of compute nodes. resizing - Indicates that the cluster is resizing; that is, compute nodes are being added to or removed from the cluster.
	AllocationState string `pulumi:"allocationState"`
	// The time at which the cluster entered its current allocation state.
	AllocationStateTransitionTime string `pulumi:"allocationStateTransitionTime"`
	// The time when the cluster was created.
	CreationTime string `pulumi:"creationTime"`
	// The number of compute nodes currently assigned to the cluster.
	CurrentNodeCount int `pulumi:"currentNodeCount"`
	// Collection of errors encountered by various compute nodes during node setup.
	Errors []BatchAIErrorResponse `pulumi:"errors"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Setup (mount file systems, performance counters settings and custom setup task) to be performed on each compute node in the cluster.
	NodeSetup *NodeSetupResponse `pulumi:"nodeSetup"`
	// Counts of various node states on the cluster.
	NodeStateCounts NodeStateCountsResponse `pulumi:"nodeStateCounts"`
	// Provisioning state of the cluster. Possible value are: creating - Specifies that the cluster is being created. succeeded - Specifies that the cluster has been created successfully. failed - Specifies that the cluster creation has failed. deleting - Specifies that the cluster is being deleted.
	ProvisioningState string `pulumi:"provisioningState"`
	// Time when the provisioning state was changed.
	ProvisioningStateTransitionTime string `pulumi:"provisioningStateTransitionTime"`
	// Scale settings of the cluster.
	ScaleSettings *ScaleSettingsResponse `pulumi:"scaleSettings"`
	// Virtual network subnet resource ID the cluster nodes belong to.
	Subnet *ResourceIdResponse `pulumi:"subnet"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Administrator user account settings which can be used to SSH to compute nodes.
	UserAccountSettings *UserAccountSettingsResponse `pulumi:"userAccountSettings"`
	// Virtual machine configuration (OS image) of the compute nodes. All nodes in a cluster have the same OS image configuration.
	VirtualMachineConfiguration *VirtualMachineConfigurationResponse `pulumi:"virtualMachineConfiguration"`
	// VM priority of cluster nodes.
	VmPriority *string `pulumi:"vmPriority"`
	// The size of the virtual machines in the cluster. All nodes in a cluster have the same VM size.
	VmSize *string `pulumi:"vmSize"`
}