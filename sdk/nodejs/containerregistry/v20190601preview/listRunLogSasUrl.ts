// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listRunLogSasUrl(args: ListRunLogSasUrlArgs, opts?: pulumi.InvokeOptions): Promise<ListRunLogSasUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:containerregistry/v20190601preview:listRunLogSasUrl", {
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
        "runId": args.runId,
    }, opts);
}

export interface ListRunLogSasUrlArgs {
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The run ID.
     */
    readonly runId: string;
}

/**
 * The result of get log link operation.
 */
export interface ListRunLogSasUrlResult {
    /**
     * The link to logs in registry for a run on a azure container registry.
     */
    readonly logArtifactLink?: string;
    /**
     * The link to logs for a run on a azure container registry.
     */
    readonly logLink?: string;
}
