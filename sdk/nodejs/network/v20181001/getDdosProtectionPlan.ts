// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getDdosProtectionPlan(args: GetDdosProtectionPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetDdosProtectionPlanResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20181001:getDdosProtectionPlan", {
        "ddosProtectionPlanName": args.ddosProtectionPlanName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDdosProtectionPlanArgs {
    /**
     * The name of the DDoS protection plan.
     */
    readonly ddosProtectionPlanName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * A DDoS protection plan in a resource group.
 */
export interface GetDdosProtectionPlanResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the DDoS protection plan resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
     */
    readonly resourceGuid: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
     */
    readonly virtualNetworks: outputs.network.v20181001.SubResourceResponse[];
}
