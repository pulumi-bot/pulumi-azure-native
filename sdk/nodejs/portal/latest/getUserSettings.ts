// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Response to get user settings
 * Latest API Version: 2018-10-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:portal:getUserSettings'. */
export function getUserSettings(args: GetUserSettingsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserSettingsResult> {
    pulumi.log.warn("getUserSettings is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:portal:getUserSettings'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:portal/latest:getUserSettings", {
        "userSettingsName": args.userSettingsName,
    }, opts);
}

export interface GetUserSettingsArgs {
    /**
     * The name of the user settings
     */
    readonly userSettingsName: string;
}

/**
 * Response to get user settings
 */
export interface GetUserSettingsResult {
    /**
     * The cloud shell user settings properties.
     */
    readonly properties: outputs.portal.latest.UserPropertiesResponse;
}
