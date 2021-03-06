// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The device security group resource
 * Latest API Version: 2019-08-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:security:getDeviceSecurityGroup'. */
export function getDeviceSecurityGroup(args: GetDeviceSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceSecurityGroupResult> {
    pulumi.log.warn("getDeviceSecurityGroup is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:security:getDeviceSecurityGroup'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:security/latest:getDeviceSecurityGroup", {
        "deviceSecurityGroupName": args.deviceSecurityGroupName,
        "resourceId": args.resourceId,
    }, opts);
}

export interface GetDeviceSecurityGroupArgs {
    /**
     * The name of the device security group. Note that the name of the device security group is case insensitive.
     */
    readonly deviceSecurityGroupName: string;
    /**
     * The identifier of the resource.
     */
    readonly resourceId: string;
}

/**
 * The device security group resource
 */
export interface GetDeviceSecurityGroupResult {
    /**
     * The allow-list custom alert rules.
     */
    readonly allowlistRules?: outputs.security.latest.AllowlistCustomAlertRuleResponse[];
    /**
     * The deny-list custom alert rules.
     */
    readonly denylistRules?: outputs.security.latest.DenylistCustomAlertRuleResponse[];
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The list of custom alert threshold rules.
     */
    readonly thresholdRules?: outputs.security.latest.ThresholdCustomAlertRuleResponse[];
    /**
     * The list of custom alert time-window rules.
     */
    readonly timeWindowRules?: outputs.security.latest.TimeWindowCustomAlertRuleResponse[];
    /**
     * Resource type
     */
    readonly type: string;
}
