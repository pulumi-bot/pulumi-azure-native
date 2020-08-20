// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getTagAtScope(args: GetTagAtScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetTagAtScopeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:resources/v20191001:getTagAtScope", {
        "scope": args.scope,
    }, opts);
}

export interface GetTagAtScopeArgs {
    /**
     * The resource scope.
     */
    readonly scope: string;
}

/**
 * Wrapper resource for tags API requests and responses.
 */
export interface GetTagAtScopeResult {
    /**
     * The name of the tags wrapper resource.
     */
    readonly name: string;
    /**
     * The set of tags.
     */
    readonly properties: outputs.resources.v20191001.TagsResponse;
    /**
     * The type of the tags wrapper resource.
     */
    readonly type: string;
}
