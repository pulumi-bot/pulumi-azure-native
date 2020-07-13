// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getMachineProviderGuestConfigurationAssignment(args: GetMachineProviderGuestConfigurationAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetMachineProviderGuestConfigurationAssignmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:hybridcompute:getMachineProviderGuestConfigurationAssignment", {
        "machineName": args.machineName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMachineProviderGuestConfigurationAssignmentArgs {
    /**
     * The name of the ARC machine.
     */
    readonly machineName: string;
    /**
     * The guest configuration assignment name.
     */
    readonly name: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * Guest configuration assignment is an association between a machine and guest configuration.
 */
export interface GetMachineProviderGuestConfigurationAssignmentResult {
    /**
     * Region where the VM is located.
     */
    readonly location?: string;
    /**
     * Name of the guest configuration assignment.
     */
    readonly name?: string;
    /**
     * Properties of the Guest configuration assignment.
     */
    readonly properties: outputs.hybridcompute.GuestConfigurationAssignmentPropertiesResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
}