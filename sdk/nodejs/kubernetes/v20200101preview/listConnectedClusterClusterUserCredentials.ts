// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listConnectedClusterClusterUserCredentials(args: ListConnectedClusterClusterUserCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<ListConnectedClusterClusterUserCredentialsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:kubernetes/v20200101preview:listConnectedClusterClusterUserCredentials", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListConnectedClusterClusterUserCredentialsArgs {
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
 * The list of credential result response.
 */
export interface ListConnectedClusterClusterUserCredentialsResult {
    /**
     * Base64-encoded Kubernetes configuration file.
     */
    readonly kubeconfigs: outputs.kubernetes.v20200101preview.CredentialResultResponse[];
}
