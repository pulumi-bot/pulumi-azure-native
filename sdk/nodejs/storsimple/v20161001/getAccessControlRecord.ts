// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getAccessControlRecord(args: GetAccessControlRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessControlRecordResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:storsimple/v20161001:getAccessControlRecord", {
        "accessControlRecordName": args.accessControlRecordName,
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAccessControlRecordArgs {
    /**
     * Name of access control record to be fetched.
     */
    readonly accessControlRecordName: string;
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
 * The access control record
 */
export interface GetAccessControlRecordResult {
    /**
     * The Iscsi initiator name (IQN)
     */
    readonly initiatorName: string;
    /**
     * The name.
     */
    readonly name: string;
    /**
     * The type.
     */
    readonly type: string;
}
