// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getView(args: GetViewArgs, opts?: pulumi.InvokeOptions): Promise<GetViewResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:costmanagement/v20200601:getView", {
        "name": args.name,
    }, opts);
}

export interface GetViewArgs {
    /**
     * View name
     */
    readonly name: string;
}

/**
 * States and configurations of Cost Analysis.
 */
export interface GetViewResult {
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    readonly eTag?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The properties of the view.
     */
    readonly properties: outputs.costmanagement.v20200601.ViewPropertiesResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}