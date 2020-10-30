// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:synapse/v20190601preview:getWorkspace", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace
     */
    readonly workspaceName: string;
}

/**
 * A workspace
 */
export interface GetWorkspaceResult {
    /**
     * Babylon Configuration
     */
    readonly babylonConfiguration?: outputs.synapse.v20190601preview.BabylonConfigurationResponse;
    /**
     * Connectivity endpoints
     */
    readonly connectivityEndpoints?: {[key: string]: string};
    /**
     * Workspace default data lake storage account details
     */
    readonly defaultDataLakeStorage?: outputs.synapse.v20190601preview.DataLakeStorageAccountDetailsResponse;
    /**
     * The encryption details of the workspace
     */
    readonly encryption?: outputs.synapse.v20190601preview.EncryptionDetailsResponse;
    /**
     * Workspace level configs and feature flags
     */
    readonly extraProperties: {[key: string]: any};
    /**
     * Identity of the workspace
     */
    readonly identity?: outputs.synapse.v20190601preview.ManagedIdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
     */
    readonly managedResourceGroupName?: string;
    /**
     * Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
     */
    readonly managedVirtualNetwork?: string;
    /**
     * Managed Virtual Network Settings
     */
    readonly managedVirtualNetworkSettings?: outputs.synapse.v20190601preview.ManagedVirtualNetworkSettingsResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Private endpoint connections to the workspace
     */
    readonly privateEndpointConnections?: outputs.synapse.v20190601preview.PrivateEndpointConnectionResponse[];
    /**
     * Resource provisioning state
     */
    readonly provisioningState: string;
    /**
     * Login for workspace SQL active directory administrator
     */
    readonly sqlAdministratorLogin?: string;
    /**
     * SQL administrator login password
     */
    readonly sqlAdministratorLoginPassword?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Virtual Network profile
     */
    readonly virtualNetworkProfile?: outputs.synapse.v20190601preview.VirtualNetworkProfileResponse;
    /**
     * The workspace unique identifier
     */
    readonly workspaceUID: string;
}
