// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getDataExport(args: GetDataExportArgs, opts?: pulumi.InvokeOptions): Promise<GetDataExportResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:operationalinsights/v20190801preview:getDataExport", {
        "dataExportName": args.dataExportName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetDataExportArgs {
    /**
     * The data export rule name.
     */
    readonly dataExportName: string;
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The Log Analytics workspace name.
     */
    readonly workspaceName: string;
}

/**
 * The top level data export resource container.
 */
export interface GetDataExportResult {
    /**
     * The latest data export rule modification time.
     */
    readonly createdDate?: string;
    /**
     * The data export rule ID.
     */
    readonly dataExportId?: string;
    /**
     * Active when enabled.
     */
    readonly enable?: boolean;
    /**
     * Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
     */
    readonly eventHubName?: string;
    /**
     * Date and time when the export was last modified.
     */
    readonly lastModifiedDate?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
     */
    readonly resourceId: string;
    /**
     * An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
     */
    readonly tableNames?: string[];
    /**
     * Resource type.
     */
    readonly type: string;
}
