// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getFileServer(args: GetFileServerArgs, opts?: pulumi.InvokeOptions): Promise<GetFileServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batchai/v20180301:getFileServer", {
        "fileServerName": args.fileServerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFileServerArgs {
    /**
     * The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly fileServerName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains information about the File Server.
 */
export interface GetFileServerResult {
    readonly creationTime: string;
    /**
     * Settings for the data disk which would be created for the File Server.
     */
    readonly dataDisks?: outputs.batchai.v20180301.DataDisksResponse;
    /**
     * The location of the resource
     */
    readonly location: string;
    /**
     * Details of the File Server.
     */
    readonly mountSettings: outputs.batchai.v20180301.MountSettingsResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Possible values: creating - The File Server is getting created. updating - The File Server creation has been accepted and it is getting updated. deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted. failed - The File Server creation has failed with the specified errorCode. Details about the error code are specified in the message field. succeeded - The File Server creation has succeeded.
     */
    readonly provisioningState: FileServerProvisioningState;
    readonly provisioningStateTransitionTime: string;
    /**
     * SSH configuration settings for the VM
     */
    readonly sshConfiguration?: outputs.batchai.v20180301.SshConfigurationResponse;
    /**
     * Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
     */
    readonly subnet?: outputs.batchai.v20180301.ResourceIdResponse;
    /**
     * The tags of the resource
     */
    readonly tags: {[key: string]: string};
    /**
     * The type of the resource
     */
    readonly type: string;
    /**
     * For information about available VM sizes for File Server from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
     */
    readonly vmSize?: string;
}
