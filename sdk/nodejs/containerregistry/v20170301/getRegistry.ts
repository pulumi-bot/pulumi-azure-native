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
    return pulumi.runtime.invoke("azurerm:containerregistry/v20170301:getRegistry", {
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
     * The provisioning state of the container registry at the time the operation was called.
     */
    readonly provisioningState: ProvisioningState;
    /**
     * The SKU of the container registry.
     */
    readonly sku: outputs.containerregistry.v20170301.SkuResponse;
    /**
     * The properties of the storage account for the container registry.
     */
    readonly storageAccount?: outputs.containerregistry.v20170301.StorageAccountPropertiesResponse;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
