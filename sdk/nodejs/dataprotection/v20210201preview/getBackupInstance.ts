// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * BackupInstance Resource
 */
export function getBackupInstance(args: GetBackupInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:dataprotection/v20210201preview:getBackupInstance", {
        "backupInstanceName": args.backupInstanceName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetBackupInstanceArgs {
    /**
     * The name of the backup instance
     */
    readonly backupInstanceName: string;
    /**
     * The name of the resource group where the backup vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the backup vault.
     */
    readonly vaultName: string;
}

/**
 * BackupInstance Resource
 */
export interface GetBackupInstanceResult {
    /**
     * Gets or sets the data source information.
     */
    readonly dataSourceInfo: outputs.dataprotection.v20210201preview.DatasourceResponse;
    /**
     * Gets or sets the data source set information.
     */
    readonly dataSourceSetInfo?: outputs.dataprotection.v20210201preview.DatasourceSetResponse;
    /**
     * Resource Id represents the complete path to the resource.
     */
    readonly id: string;
    /**
     * Resource name associated with the resource.
     */
    readonly name: string;
    readonly objectType?: string;
    /**
     * Gets or sets the policy information.
     */
    readonly policyInfo: outputs.dataprotection.v20210201preview.PolicyInfoResponse;
    /**
     * Specifies the protection status of the resource
     */
    readonly protectionStatus: outputs.dataprotection.v20210201preview.ProtectionStatusDetailsResponse;
    /**
     * Specifies the provisioning state of the resource i.e. provisioning/updating/Succeeded/Failed
     */
    readonly provisioningState: string;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.dataprotection.v20210201preview.SystemDataResponse;
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
