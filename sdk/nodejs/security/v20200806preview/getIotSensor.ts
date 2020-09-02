// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIotSensor(args: GetIotSensorArgs, opts?: pulumi.InvokeOptions): Promise<GetIotSensorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:security/v20200806preview:getIotSensor", {
        "iotSensorName": args.iotSensorName,
        "scope": args.scope,
    }, opts);
}

export interface GetIotSensorArgs {
    /**
     * Name of the IoT sensor
     */
    readonly iotSensorName: string;
    /**
     * Scope of the query, can be subscription (/subscriptions/326b1ffa-8ac7-4034-8437-69bef733dede) or IoT Hub (/providers/Microsoft.Devices/iotHubs/myHub)
     */
    readonly scope: string;
}

/**
 * IoT sensor
 */
export interface GetIotSensorResult {
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Resource type
     */
    readonly type: string;
}
