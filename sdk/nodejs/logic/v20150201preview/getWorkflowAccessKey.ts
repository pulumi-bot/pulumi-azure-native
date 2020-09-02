// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWorkflowAccessKey(args: GetWorkflowAccessKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkflowAccessKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:logic/v20150201preview:getWorkflowAccessKey", {
        "accessKeyName": args.accessKeyName,
        "resourceGroupName": args.resourceGroupName,
        "workflowName": args.workflowName,
    }, opts);
}

export interface GetWorkflowAccessKeyArgs {
    /**
     * The workflow access key name.
     */
    readonly accessKeyName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The workflow name.
     */
    readonly workflowName: string;
}

export interface GetWorkflowAccessKeyResult {
    /**
     * Gets the workflow access key name.
     */
    readonly name: string;
    /**
     * Gets or sets the not-after time.
     */
    readonly notAfter?: string;
    /**
     * Gets or sets the not-before time.
     */
    readonly notBefore?: string;
    /**
     * Gets the workflow access key type.
     */
    readonly type: string;
}
