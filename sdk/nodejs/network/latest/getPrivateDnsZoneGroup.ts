// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Private dns zone group resource.
 * Latest API Version: 2020-08-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:network:getPrivateDnsZoneGroup'. */
export function getPrivateDnsZoneGroup(args: GetPrivateDnsZoneGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateDnsZoneGroupResult> {
    pulumi.log.warn("getPrivateDnsZoneGroup is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:network:getPrivateDnsZoneGroup'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:network/latest:getPrivateDnsZoneGroup", {
        "privateDnsZoneGroupName": args.privateDnsZoneGroupName,
        "privateEndpointName": args.privateEndpointName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateDnsZoneGroupArgs {
    /**
     * The name of the private dns zone group.
     */
    readonly privateDnsZoneGroupName: string;
    /**
     * The name of the private endpoint.
     */
    readonly privateEndpointName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Private dns zone group resource.
 */
export interface GetPrivateDnsZoneGroupResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * Name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * A collection of private dns zone configurations of the private dns zone group.
     */
    readonly privateDnsZoneConfigs?: outputs.network.latest.PrivateDnsZoneConfigResponse[];
    /**
     * The provisioning state of the private dns zone group resource.
     */
    readonly provisioningState: string;
}
