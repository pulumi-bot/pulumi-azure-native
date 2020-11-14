// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getServiceUnit(args: GetServiceUnitArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceUnitResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:deploymentmanager/v20191101preview:getServiceUnit", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "serviceTopologyName": args.serviceTopologyName,
        "serviceUnitName": args.serviceUnitName,
    }, opts);
}

export interface GetServiceUnitArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the service resource.
     */
    readonly serviceName: string;
    /**
     * The name of the service topology .
     */
    readonly serviceTopologyName: string;
    /**
     * The name of the service unit resource.
     */
    readonly serviceUnitName: string;
}

/**
 * Represents the response of a service unit resource.
 */
export interface GetServiceUnitResult {
    /**
     * The artifacts for the service unit.
     */
    readonly artifacts?: outputs.deploymentmanager.v20191101preview.ServiceUnitArtifactsResponse;
    /**
     * Describes the type of ARM deployment to be performed on the resource.
     */
    readonly deploymentMode: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
     */
    readonly targetResourceGroup: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
