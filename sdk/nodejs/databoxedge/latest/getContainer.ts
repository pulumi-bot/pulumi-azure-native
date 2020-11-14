// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getContainer(args: GetContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:databoxedge/latest:getContainer", {
        "containerName": args.containerName,
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
        "storageAccountName": args.storageAccountName,
    }, opts);
}

export interface GetContainerArgs {
    /**
     * The container Name
     */
    readonly containerName: string;
    /**
     * The device name.
     */
    readonly deviceName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The Storage Account Name
     */
    readonly storageAccountName: string;
}

/**
 * Represents a container on the  Data Box Edge/Gateway device.
 */
export interface GetContainerResult {
    /**
     * Current status of the container.
     */
    readonly containerStatus: string;
    /**
     * The UTC time when container got created.
     */
    readonly createdDateTime: string;
    /**
     * DataFormat for Container
     */
    readonly dataFormat: string;
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * Details of the refresh job on this container.
     */
    readonly refreshDetails: outputs.databoxedge.latest.RefreshDetailsResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}
