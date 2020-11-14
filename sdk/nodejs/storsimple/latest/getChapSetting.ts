// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getChapSetting(args: GetChapSettingArgs, opts?: pulumi.InvokeOptions): Promise<GetChapSettingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:storsimple/latest:getChapSetting", {
        "chapUserName": args.chapUserName,
        "deviceName": args.deviceName,
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetChapSettingArgs {
    /**
     * The user name of chap to be fetched.
     */
    readonly chapUserName: string;
    /**
     * The device name.
     */
    readonly deviceName: string;
    /**
     * The manager name
     */
    readonly managerName: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
}

/**
 * Challenge-Handshake Authentication Protocol (CHAP) setting
 */
export interface GetChapSettingResult {
    /**
     * The name.
     */
    readonly name: string;
    /**
     * The chap password.
     */
    readonly password: outputs.storsimple.latest.AsymmetricEncryptedSecretResponse;
    /**
     * The type.
     */
    readonly type: string;
}
