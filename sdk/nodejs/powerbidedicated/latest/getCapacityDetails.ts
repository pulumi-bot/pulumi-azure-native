// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getCapacityDetails(args: GetCapacityDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetCapacityDetailsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:powerbidedicated/latest:getCapacityDetails", {
        "dedicatedCapacityName": args.dedicatedCapacityName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCapacityDetailsArgs {
    /**
     * The name of the dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
     */
    readonly dedicatedCapacityName: string;
    /**
     * The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
     */
    readonly resourceGroupName: string;
}

/**
 * Represents an instance of a Dedicated Capacity resource.
 */
export interface GetCapacityDetailsResult {
    /**
     * A collection of Dedicated capacity administrators
     */
    readonly administration?: outputs.powerbidedicated.latest.DedicatedCapacityAdministratorsResponse;
    /**
     * Location of the PowerBI Dedicated resource.
     */
    readonly location: string;
    /**
     * The name of the PowerBI Dedicated resource.
     */
    readonly name: string;
    /**
     * The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
     */
    readonly provisioningState: string;
    /**
     * The SKU of the PowerBI Dedicated resource.
     */
    readonly sku: outputs.powerbidedicated.latest.ResourceSkuResponse;
    /**
     * The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
     */
    readonly state: string;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the PowerBI Dedicated resource.
     */
    readonly type: string;
}
