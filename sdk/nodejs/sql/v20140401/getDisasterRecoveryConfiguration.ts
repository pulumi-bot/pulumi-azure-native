// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getDisasterRecoveryConfiguration(args: GetDisasterRecoveryConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetDisasterRecoveryConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:sql/v20140401:getDisasterRecoveryConfiguration", {
        "disasterRecoveryConfigurationName": args.disasterRecoveryConfigurationName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetDisasterRecoveryConfigurationArgs {
    /**
     * The name of the disaster recovery configuration.
     */
    readonly disasterRecoveryConfigurationName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * Represents a disaster recovery configuration.
 */
export interface GetDisasterRecoveryConfigurationResult {
    /**
     * Whether or not failover can be done automatically.
     */
    readonly autoFailover: string;
    /**
     * How aggressive the automatic failover should be.
     */
    readonly failoverPolicy: string;
    /**
     * Location of the server that contains this disaster recovery configuration.
     */
    readonly location: string;
    /**
     * Logical name of the server.
     */
    readonly logicalServerName: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Logical name of the partner server.
     */
    readonly partnerLogicalServerName: string;
    /**
     * Id of the partner server.
     */
    readonly partnerServerId: string;
    /**
     * The role of the current server in the disaster recovery configuration.
     */
    readonly role: string;
    /**
     * The status of the disaster recovery configuration.
     */
    readonly status: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
