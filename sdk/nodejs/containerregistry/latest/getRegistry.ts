// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRegistry(args: GetRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:containerregistry/latest:getRegistry", {
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRegistryArgs {
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * An object that represents a container registry.
 */
export interface GetRegistryResult {
    /**
     * The value that indicates whether the admin user is enabled.
     */
    readonly adminUserEnabled?: boolean;
    /**
     * The creation date of the container registry in ISO8601 format.
     */
    readonly creationDate: string;
    /**
     * Enable a single data endpoint per region for serving data.
     */
    readonly dataEndpointEnabled?: boolean;
    /**
     * List of host names that will serve data when dataEndpointEnabled is true.
     */
    readonly dataEndpointHostNames: string[];
    /**
     * The encryption settings of container registry.
     */
    readonly encryption?: outputs.containerregistry.latest.EncryptionPropertyResponse;
    /**
     * The identity of the container registry.
     */
    readonly identity?: outputs.containerregistry.latest.IdentityPropertiesResponse;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: string;
    /**
     * The URL that can be used to log into the container registry.
     */
    readonly loginServer: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The network rule set for a container registry.
     */
    readonly networkRuleSet?: outputs.containerregistry.latest.NetworkRuleSetResponse;
    /**
     * The policies for a container registry.
     */
    readonly policies?: outputs.containerregistry.latest.PoliciesResponse;
    /**
     * List of private endpoint connections for a container registry.
     */
    readonly privateEndpointConnections: outputs.containerregistry.latest.PrivateEndpointConnectionResponse[];
    /**
     * The provisioning state of the container registry at the time the operation was called.
     */
    readonly provisioningState: string;
    /**
     * Whether or not public network access is allowed for the container registry.
     */
    readonly publicNetworkAccess?: string;
    /**
     * The SKU of the container registry.
     */
    readonly sku: outputs.containerregistry.latest.SkuResponse;
    /**
     * The status of the container registry at the time the operation was called.
     */
    readonly status: outputs.containerregistry.latest.StatusResponse;
    /**
     * The properties of the storage account for the container registry. Only applicable to Classic SKU.
     */
    readonly storageAccount?: outputs.containerregistry.latest.StorageAccountPropertiesResponse;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
