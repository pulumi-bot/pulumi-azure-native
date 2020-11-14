// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getAgentPool(args: GetAgentPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetAgentPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:containerservice/v20190201:getAgentPool", {
        "agentPoolName": args.agentPoolName,
        "managedClusterName": args.managedClusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAgentPoolArgs {
    /**
     * The name of the agent pool.
     */
    readonly agentPoolName: string;
    /**
     * The name of the managed cluster resource.
     */
    readonly managedClusterName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Agent Pool.
 */
export interface GetAgentPoolResult {
    /**
     * (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
     */
    readonly availabilityZones?: string[];
    /**
     * Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
     */
    readonly count: number;
    /**
     * Whether to enable auto-scaler
     */
    readonly enableAutoScaling?: boolean;
    /**
     * Maximum number of nodes for auto-scaling
     */
    readonly maxCount?: number;
    /**
     * Maximum number of pods that can run on a node.
     */
    readonly maxPods?: number;
    /**
     * Minimum number of nodes for auto-scaling
     */
    readonly minCount?: number;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name: string;
    /**
     * Version of orchestrator specified when creating the managed cluster.
     */
    readonly orchestratorVersion?: string;
    /**
     * OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
     */
    readonly osDiskSizeGB?: number;
    /**
     * OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
     */
    readonly osType?: string;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * AgentPoolType represents types of an agent pool
     */
    readonly type: string;
    /**
     * Size of agent VMs.
     */
    readonly vmSize: string;
    /**
     * VNet SubnetID specifies the VNet's subnet identifier.
     */
    readonly vnetSubnetID?: string;
}
