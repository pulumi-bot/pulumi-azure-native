// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getConnectedCluster(args: GetConnectedClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectedClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:kubernetes/v20200101preview:getConnectedCluster", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectedClusterArgs {
    /**
     * The name of the Kubernetes cluster on which get is called.
     */
    readonly clusterName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Represents a connected cluster.
 */
export interface GetConnectedClusterResult {
    readonly aadProfile: outputs.kubernetes.v20200101preview.ConnectedClusterAADProfileResponse;
    /**
     * Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
     */
    readonly agentPublicKeyCertificate: string;
    /**
     * Version of the agent running on the connected cluster resource
     */
    readonly agentVersion: string;
    /**
     * The identity of the connected cluster.
     */
    readonly identity: outputs.kubernetes.v20200101preview.ConnectedClusterIdentityResponse;
    /**
     * The Kubernetes version of the connected cluster resource
     */
    readonly kubernetesVersion: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The current deployment state of connectedClusters.
     */
    readonly provisioningState?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Number of nodes present in the connected cluster resource
     */
    readonly totalNodeCount: number;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
