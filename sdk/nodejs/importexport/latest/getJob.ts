// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Contains the job information.
 * Latest API Version: 2020-08-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:importexport:getJob'. */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    pulumi.log.warn("getJob is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:importexport:getJob'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:importexport/latest:getJob", {
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetJobArgs {
    /**
     * The name of the import/export job.
     */
    readonly jobName: string;
    /**
     * The resource group name uniquely identifies the resource group within the user subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains the job information.
 */
export interface GetJobResult {
    /**
     * Specifies the resource identifier of the job.
     */
    readonly id: string;
    /**
     * Specifies the job identity details
     */
    readonly identity?: outputs.importexport.latest.IdentityDetailsResponse;
    /**
     * Specifies the Azure location where the job is created.
     */
    readonly location?: string;
    /**
     * Specifies the name of the job.
     */
    readonly name: string;
    /**
     * Specifies the job properties
     */
    readonly properties: outputs.importexport.latest.JobDetailsResponse;
    /**
     * SystemData of ImportExport Jobs.
     */
    readonly systemData: outputs.importexport.latest.SystemDataResponse;
    /**
     * Specifies the tags that are assigned to the job.
     */
    readonly tags?: any;
    /**
     * Specifies the type of the job resource.
     */
    readonly type: string;
}
