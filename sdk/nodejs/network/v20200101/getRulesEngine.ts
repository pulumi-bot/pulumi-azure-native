// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRulesEngine(args: GetRulesEngineArgs, opts?: pulumi.InvokeOptions): Promise<GetRulesEngineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20200101:getRulesEngine", {
        "frontDoorName": args.frontDoorName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRulesEngineArgs {
    /**
     * Name of the Front Door which is globally unique.
     */
    readonly frontDoorName: string;
    /**
     * Name of the Rules Engine which is unique within the Front Door.
     */
    readonly name: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * A rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.
 */
export interface GetRulesEngineResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties of the Rules Engine Configuration.
     */
    readonly properties: outputs.network.v20200101.RulesEnginePropertiesResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}