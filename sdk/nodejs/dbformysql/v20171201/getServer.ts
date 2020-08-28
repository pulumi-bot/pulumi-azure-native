// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:dbformysql/v20171201:getServer", {
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * Represents a server.
 */
export interface GetServerResult {
    /**
     * The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
     */
    readonly administratorLogin?: string;
    /**
     * Status showing whether the server data encryption is enabled with customer-managed keys.
     */
    readonly byokEnforcement: string;
    /**
     * Earliest restore point creation time (ISO8601 format)
     */
    readonly earliestRestoreDate?: string;
    /**
     * The fully qualified domain name of a server.
     */
    readonly fullyQualifiedDomainName?: string;
    /**
     * The Azure Active Directory identity of the server.
     */
    readonly identity?: outputs.dbformysql.v20171201.ResourceIdentityResponse;
    /**
     * Status showing whether the server enabled infrastructure encryption.
     */
    readonly infrastructureEncryption?: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The master server id of a replica server.
     */
    readonly masterServerId?: string;
    /**
     * Enforce a minimal Tls version for the server.
     */
    readonly minimalTlsVersion?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * List of private endpoint connections on a server
     */
    readonly privateEndpointConnections: outputs.dbformysql.v20171201.ServerPrivateEndpointConnectionResponse[];
    /**
     * Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
     */
    readonly publicNetworkAccess?: string;
    /**
     * The maximum number of replicas that a master server can have.
     */
    readonly replicaCapacity?: number;
    /**
     * The replication role of the server.
     */
    readonly replicationRole?: string;
    /**
     * The SKU (pricing tier) of the server.
     */
    readonly sku?: outputs.dbformysql.v20171201.SkuResponse;
    /**
     * Enable ssl enforcement or not when connect to server.
     */
    readonly sslEnforcement?: string;
    /**
     * Storage profile of a server.
     */
    readonly storageProfile?: outputs.dbformysql.v20171201.StorageProfileResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
    /**
     * A state of a server that is visible to user.
     */
    readonly userVisibleState?: string;
    /**
     * Server version.
     */
    readonly version?: string;
}
