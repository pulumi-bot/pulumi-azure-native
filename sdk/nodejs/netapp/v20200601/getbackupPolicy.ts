// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getbackupPolicy(args: GetbackupPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetbackupPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:netapp/v20200601:getbackupPolicy", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetbackupPolicyArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: string;
    /**
     * Backup policy Name which uniquely identify backup policy.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Backup policy information
 */
export interface GetbackupPolicyResult {
    /**
     * Daily backups count to keep
     */
    readonly dailyBackupsToKeep?: number;
    /**
     * The property to decide policy is enabled or not
     */
    readonly enabled?: boolean;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Monthly backups count to keep
     */
    readonly monthlyBackupsToKeep?: number;
    /**
     * Name of backup policy
     */
    readonly name: string;
    /**
     * Azure lifecycle management
     */
    readonly provisioningState: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * A list of volumes assigned to this policy
     */
    readonly volumeBackups?: outputs.netapp.v20200601.VolumeBackupsResponse[];
    /**
     * Volumes using current backup policy
     */
    readonly volumesAssigned?: number;
    /**
     * Weekly backups count to keep
     */
    readonly weeklyBackupsToKeep?: number;
    /**
     * Yearly backups count to keep
     */
    readonly yearlyBackupsToKeep?: number;
}
