// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getServerDetails(args: GetServerDetailsArgs, opts?: pulumi.InvokeOptions): Promise<GetServerDetailsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:analysisservices/v20170714:getServerDetails", {
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerDetailsArgs {
    /**
     * The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
     */
    readonly serverName: string;
}

/**
 * Represents an instance of an Analysis Services resource.
 */
export interface GetServerDetailsResult {
    /**
     * A collection of AS server administrators
     */
    readonly asAdministrators?: outputs.analysisservices.v20170714.ServerAdministratorsResponse;
    /**
     * The SAS container URI to the backup container.
     */
    readonly backupBlobContainerUri?: string;
    /**
     * The gateway details configured for the AS server.
     */
    readonly gatewayDetails?: outputs.analysisservices.v20170714.GatewayDetailsResponse;
    /**
     * Location of the Analysis Services resource.
     */
    readonly location: string;
    /**
     * The name of the Analysis Services resource.
     */
    readonly name: string;
    /**
     * The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
     */
    readonly provisioningState: string;
    /**
     * The full name of the Analysis Services resource.
     */
    readonly serverFullName: string;
    /**
     * The SKU of the Analysis Services resource.
     */
    readonly sku: outputs.analysisservices.v20170714.ResourceSkuResponse;
    /**
     * The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
     */
    readonly state: string;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the Analysis Services resource.
     */
    readonly type: string;
}
