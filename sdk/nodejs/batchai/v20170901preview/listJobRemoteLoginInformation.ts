// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Values returned by the List operation.
 */
export function listJobRemoteLoginInformation(args: ListJobRemoteLoginInformationArgs, opts?: pulumi.InvokeOptions): Promise<ListJobRemoteLoginInformationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:batchai/v20170901preview:listJobRemoteLoginInformation", {
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListJobRemoteLoginInformationArgs {
    /**
     * The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly jobName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Values returned by the List operation.
 */
export interface ListJobRemoteLoginInformationResult {
    /**
     * The continuation token.
     */
    readonly nextLink?: string;
    /**
     * The collection of returned remote login details.
     */
    readonly value?: outputs.batchai.v20170901preview.RemoteLoginInformationResponse[];
}
