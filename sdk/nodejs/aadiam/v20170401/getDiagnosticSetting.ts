// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDiagnosticSetting(args: GetDiagnosticSettingArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagnosticSettingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:aadiam/v20170401:getDiagnosticSetting", {
        "name": args.name,
    }, opts);
}

export interface GetDiagnosticSettingArgs {
    /**
     * The name of the diagnostic setting.
     */
    readonly name: string;
}

/**
 * The diagnostic setting resource.
 */
export interface GetDiagnosticSettingResult {
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * Properties of a Diagnostic Settings Resource.
     */
    readonly properties: outputs.aadiam.v20170401.DiagnosticSettingsResponse;
    /**
     * Azure resource type
     */
    readonly type: string;
}