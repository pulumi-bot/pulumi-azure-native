// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getBlueprint(args: GetBlueprintArgs, opts?: pulumi.InvokeOptions): Promise<GetBlueprintResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:blueprint/v20181101preview:getBlueprint", {
        "blueprintName": args.blueprintName,
        "resourceScope": args.resourceScope,
    }, opts);
}

export interface GetBlueprintArgs {
    /**
     * Name of the blueprint definition.
     */
    readonly blueprintName: string;
    /**
     * The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
     */
    readonly resourceScope: string;
}

/**
 * Represents a Blueprint definition.
 */
export interface GetBlueprintResult {
    /**
     * Multi-line explain this resource.
     */
    readonly description?: string;
    /**
     * One-liner string explain this resource.
     */
    readonly displayName?: string;
    /**
     * Layout view of the blueprint definition for UI reference.
     */
    readonly layout?: {[key: string]: any};
    /**
     * Name of this resource.
     */
    readonly name: string;
    /**
     * Parameters required by this blueprint definition.
     */
    readonly parameters?: {[key: string]: outputs.blueprint.v20181101preview.ParameterDefinitionResponse};
    /**
     * Resource group placeholders defined by this blueprint definition.
     */
    readonly resourceGroups?: {[key: string]: outputs.blueprint.v20181101preview.ResourceGroupDefinitionResponse};
    /**
     * Status of the blueprint. This field is readonly.
     */
    readonly status: outputs.blueprint.v20181101preview.BlueprintStatusResponse;
    /**
     * The scope where this blueprint definition can be assigned.
     */
    readonly targetScope: string;
    /**
     * Type of this resource.
     */
    readonly type: string;
    /**
     * Published versions of this blueprint definition.
     */
    readonly versions?: {[key: string]: any};
}
