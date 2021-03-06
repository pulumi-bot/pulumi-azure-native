// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents a scaling plan definition.
 */
export function getScalingPlan(args: GetScalingPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetScalingPlanResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:desktopvirtualization/v20210201preview:getScalingPlan", {
        "resourceGroupName": args.resourceGroupName,
        "scalingPlanName": args.scalingPlanName,
    }, opts);
}

export interface GetScalingPlanArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the scaling plan.
     */
    readonly scalingPlanName: string;
}

/**
 * Represents a scaling plan definition.
 */
export interface GetScalingPlanResult {
    /**
     * Description of scaling plan.
     */
    readonly description?: string;
    /**
     * Exclusion tag for scaling plan.
     */
    readonly exclusionTag?: string;
    /**
     * User friendly name of scaling plan.
     */
    readonly friendlyName?: string;
    /**
     * List of ScalingHostPoolReference definitions.
     */
    readonly hostPoolReferences?: outputs.desktopvirtualization.v20210201preview.ScalingHostPoolReferenceResponse[];
    /**
     * HostPool type for desktop.
     */
    readonly hostPoolType?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The ring number of scaling plan.
     */
    readonly ring?: number;
    /**
     * List of ScalingSchedule definitions.
     */
    readonly schedules?: outputs.desktopvirtualization.v20210201preview.ScalingScheduleResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Timezone of the scaling plan.
     */
    readonly timeZone?: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
