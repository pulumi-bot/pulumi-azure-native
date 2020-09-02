// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getContainerGroup(args: GetContainerGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:containerinstance/v20170801preview:getContainerGroup", {
        "containerGroupName": args.containerGroupName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetContainerGroupArgs {
    /**
     * The name of the container group.
     */
    readonly containerGroupName: string;
    /**
     * The name of the resource group that contains the container group.
     */
    readonly resourceGroupName: string;
}

/**
 * A container group.
 */
export interface GetContainerGroupResult {
    /**
     * The containers within the container group.
     */
    readonly containers: outputs.containerinstance.v20170801preview.ContainerResponse[];
    /**
     * The image registry credentials by which the container group is created from.
     */
    readonly imageRegistryCredentials?: outputs.containerinstance.v20170801preview.ImageRegistryCredentialResponse[];
    /**
     * The IP address type of the container group.
     */
    readonly ipAddress?: outputs.containerinstance.v20170801preview.IpAddressResponse;
    /**
     * The resource location.
     */
    readonly location: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * The operating system type required by the containers in the container group.
     */
    readonly osType: string;
    /**
     * The provisioning state of the container group. This only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Restart policy for all containers within the container group. Currently the only available option is `always`.
     */
    readonly restartPolicy?: string;
    /**
     * The current state of the container group. This is only valid for the response.
     */
    readonly state: string;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type.
     */
    readonly type: string;
    /**
     * The list of volumes that can be mounted by containers in this container group.
     */
    readonly volumes?: outputs.containerinstance.v20170801preview.VolumeResponse[];
}
