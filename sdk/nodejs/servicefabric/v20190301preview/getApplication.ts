// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:servicefabric/v20190301preview:getApplication", {
        "applicationName": args.applicationName,
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The name of the application resource.
     */
    readonly applicationName: string;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The application resource.
 */
export interface GetApplicationResult {
    /**
     * Azure resource etag.
     */
    readonly etag: string;
    /**
     * Azure resource location.
     */
    readonly location?: string;
    /**
     * The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
     */
    readonly maximumNodes?: number;
    /**
     * List of application capacity metric description.
     */
    readonly metrics?: outputs.servicefabric.v20190301preview.ApplicationMetricDescriptionResponse[];
    /**
     * The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
     */
    readonly minimumNodes?: number;
    /**
     * Azure resource name.
     */
    readonly name: string;
    /**
     * List of application parameters with overridden values from their default values specified in the application manifest.
     */
    readonly parameters?: {[key: string]: string};
    /**
     * The current deployment or provisioning state, which only appears in the response
     */
    readonly provisioningState: string;
    /**
     * Remove the current application capacity settings.
     */
    readonly removeApplicationCapacity?: boolean;
    /**
     * Azure resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Azure resource type.
     */
    readonly type: string;
    /**
     * The application type name as defined in the application manifest.
     */
    readonly typeName?: string;
    /**
     * The version of the application type as defined in the application manifest.
     */
    readonly typeVersion?: string;
    /**
     * Describes the policy for a monitored application upgrade.
     */
    readonly upgradePolicy?: outputs.servicefabric.v20190301preview.ApplicationUpgradePolicyResponse;
}
