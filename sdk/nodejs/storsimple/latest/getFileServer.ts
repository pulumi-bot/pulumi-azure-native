// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The file server.
 * Latest API Version: 2016-10-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:storsimple:getFileServer'. */
export function getFileServer(args: GetFileServerArgs, opts?: pulumi.InvokeOptions): Promise<GetFileServerResult> {
    pulumi.log.warn("getFileServer is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:storsimple:getFileServer'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storsimple/latest:getFileServer", {
        "deviceName": args.deviceName,
        "fileServerName": args.fileServerName,
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFileServerArgs {
    /**
     * The device name.
     */
    readonly deviceName: string;
    /**
     * The file server name.
     */
    readonly fileServerName: string;
    /**
     * The manager name
     */
    readonly managerName: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
}

/**
 * The file server.
 */
export interface GetFileServerResult {
    /**
     * The backup policy id.
     */
    readonly backupScheduleGroupId: string;
    /**
     * The description of the file server
     */
    readonly description?: string;
    /**
     * Domain of the file server
     */
    readonly domainName: string;
    /**
     * The identifier.
     */
    readonly id: string;
    /**
     * The name.
     */
    readonly name: string;
    /**
     * The storage domain id.
     */
    readonly storageDomainId: string;
    /**
     * The type.
     */
    readonly type: string;
}
