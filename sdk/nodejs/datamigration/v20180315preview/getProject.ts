// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datamigration/v20180315preview:getProject", {
        "groupName": args.groupName,
        "projectName": args.projectName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetProjectArgs {
    /**
     * Name of the resource group
     */
    readonly groupName: string;
    /**
     * Name of the project
     */
    readonly projectName: string;
    /**
     * Name of the service
     */
    readonly serviceName: string;
}

/**
 * A project resource
 */
export interface GetProjectResult {
    /**
     * UTC Date and time when project was created
     */
    readonly creationTime: string;
    /**
     * List of DatabaseInfo
     */
    readonly databasesInfo?: outputs.datamigration.v20180315preview.DatabaseInfoResponse[];
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The project's provisioning state
     */
    readonly provisioningState: string;
    /**
     * Information for connecting to source
     */
    readonly sourceConnectionInfo?: outputs.datamigration.v20180315preview.ConnectionInfoResponse;
    /**
     * Source platform for the project
     */
    readonly sourcePlatform: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Information for connecting to target
     */
    readonly targetConnectionInfo?: outputs.datamigration.v20180315preview.ConnectionInfoResponse;
    /**
     * Target platform for the project
     */
    readonly targetPlatform: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
