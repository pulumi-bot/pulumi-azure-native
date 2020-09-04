// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getMachine(args: GetMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetMachineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:hybridcompute/v20190802preview:getMachine", {
        "expand": args.expand,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMachineArgs {
    /**
     * The expand expression to apply on the operation.
     */
    readonly expand?: string;
    /**
     * The name of the hybrid machine.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Describes a hybrid machine.
 */
export interface GetMachineResult {
    /**
     * The hybrid machine agent full version.
     */
    readonly agentVersion: string;
    /**
     * Public Key that the client provides to be used during initial resource onboarding
     */
    readonly clientPublicKey?: string;
    /**
     * Specifies the hybrid machine display name.
     */
    readonly displayName: string;
    /**
     * Details about the error state.
     */
    readonly errorDetails: outputs.hybridcompute.v20190802preview.ErrorDetailResponse[];
    /**
     * Machine Extensions information
     */
    readonly extensions?: outputs.hybridcompute.v20190802preview.MachineExtensionInstanceViewResponse[];
    /**
     * The time of the last status change.
     */
    readonly lastStatusChange: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Specifies the hybrid machine FQDN.
     */
    readonly machineFqdn: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The Operating System running on the hybrid machine.
     */
    readonly osName?: string;
    /**
     * Specifies the operating system settings for the hybrid machine.
     */
    readonly osProfile?: outputs.hybridcompute.v20190802preview.OSProfileResponse;
    /**
     * The version of Operating System running on the hybrid machine.
     */
    readonly osVersion?: string;
    /**
     * Resource's Physical Location
     */
    readonly physicalLocation?: string;
    /**
     * The identity's principal id.
     */
    readonly principalId: string;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * The status of the hybrid machine agent.
     */
    readonly status: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * The identity's tenant id.
     */
    readonly tenantId: string;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Specifies the hybrid machine unique ID.
     */
    readonly vmId: string;
}
