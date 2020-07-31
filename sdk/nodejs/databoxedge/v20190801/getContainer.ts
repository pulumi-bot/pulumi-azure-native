// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getContainer(args: GetContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:databoxedge/v20190801:getContainer", {
        "deviceName": args.deviceName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "storageAccountName": args.storageAccountName,
    }, opts);
}

export interface GetContainerArgs {
    /**
     * The device name.
     */
    readonly deviceName: string;
    /**
     * The container Name
     */
    readonly name: string;
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
     * The object name.
     */
    readonly name: string;
    /**
     * The container properties.
     */
    readonly properties: outputs.databoxedge.v20190801.ContainerPropertiesResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}