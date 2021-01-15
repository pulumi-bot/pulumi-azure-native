// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getSoftwareUpdateConfigurationByName(args: GetSoftwareUpdateConfigurationByNameArgs, opts?: pulumi.InvokeOptions): Promise<GetSoftwareUpdateConfigurationByNameResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:automation/latest:getSoftwareUpdateConfigurationByName", {
        "automationAccountName": args.automationAccountName,
        "resourceGroupName": args.resourceGroupName,
        "softwareUpdateConfigurationName": args.softwareUpdateConfigurationName,
    }, opts);
}

export interface GetSoftwareUpdateConfigurationByNameArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the software update configuration to be created.
     */
    readonly softwareUpdateConfigurationName: string;
}

/**
 * Software update configuration properties.
 */
export interface GetSoftwareUpdateConfigurationByNameResult {
    /**
     * CreatedBy property, which only appears in the response.
     */
    readonly createdBy: string;
    /**
     * Creation time of the resource, which only appears in the response.
     */
    readonly creationTime: string;
    /**
     * Details of provisioning error
     */
    readonly error?: outputs.automation.latest.ErrorResponseResponse;
    /**
     * Resource Id.
     */
    readonly id: string;
    /**
     * LastModifiedBy property, which only appears in the response.
     */
    readonly lastModifiedBy: string;
    /**
     * Last time resource was modified, which only appears in the response.
     */
    readonly lastModifiedTime: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Provisioning state for the software update configuration, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Schedule information for the Software update configuration
     */
    readonly scheduleInfo: outputs.automation.latest.SUCSchedulePropertiesResponse;
    /**
     * Tasks information for the Software update configuration.
     */
    readonly tasks?: outputs.automation.latest.SoftwareUpdateConfigurationTasksResponse;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * update specific properties for the Software update configuration
     */
    readonly updateConfiguration: outputs.automation.latest.UpdateConfigurationResponse;
}
