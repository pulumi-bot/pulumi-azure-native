// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listJobOutputFiles(args: ListJobOutputFilesArgs, opts?: pulumi.InvokeOptions): Promise<ListJobOutputFilesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batchai/v20170901preview:listJobOutputFiles", {
        "jobName": args.jobName,
        "linkexpiryinminutes": args.linkexpiryinminutes,
        "maxResults": args.maxResults,
        "outputdirectoryid": args.outputdirectoryid,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListJobOutputFilesArgs {
    /**
     * The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly jobName: string;
    /**
     * The number of minutes after which the download link will expire.
     */
    readonly linkexpiryinminutes?: number;
    /**
     * The maximum number of items to return in the response. A maximum of 1000 files can be returned.
     */
    readonly maxResults?: number;
    /**
     * Id of the job output directory. This is the OutputDirectory-->id parameter that is given by the user during Create Job.
     */
    readonly outputdirectoryid: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Values returned by the List operation.
 */
export interface ListJobOutputFilesResult {
    /**
     * The continuation token.
     */
    readonly nextLink?: string;
    /**
     * The collection of returned job files.
     */
    readonly value: outputs.batchai.v20170901preview.FileResponse[];
}
