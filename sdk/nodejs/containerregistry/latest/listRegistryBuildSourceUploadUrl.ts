// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listRegistryBuildSourceUploadUrl(args: ListRegistryBuildSourceUploadUrlArgs, opts?: pulumi.InvokeOptions): Promise<ListRegistryBuildSourceUploadUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:containerregistry/latest:listRegistryBuildSourceUploadUrl", {
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListRegistryBuildSourceUploadUrlArgs {
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * The properties of a response to source upload request.
 */
export interface ListRegistryBuildSourceUploadUrlResult {
    /**
     * The relative path to the source. This is used to submit the subsequent queue build request.
     */
    readonly relativePath?: string;
    /**
     * The URL where the client can upload the source.
     */
    readonly uploadUrl?: string;
}
