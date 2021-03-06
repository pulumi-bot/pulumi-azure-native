// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Deployment information.
 * Latest API Version: 2020-10-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:resources:getDeployment'. */
export function getDeployment(args: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    pulumi.log.warn("getDeployment is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:resources:getDeployment'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:resources/latest:getDeployment", {
        "deploymentName": args.deploymentName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeploymentArgs {
    /**
     * The name of the deployment.
     */
    readonly deploymentName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Deployment information.
 */
export interface GetDeploymentResult {
    /**
     * The ID of the deployment.
     */
    readonly id: string;
    /**
     * the location of the deployment.
     */
    readonly location?: string;
    /**
     * The name of the deployment.
     */
    readonly name: string;
    /**
     * Deployment properties.
     */
    readonly properties: outputs.resources.latest.DeploymentPropertiesExtendedResponse;
    /**
     * Deployment tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the deployment.
     */
    readonly type: string;
}
