// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Response for all the Bastion Shareable Link endpoints.
 * Latest API Version: 2020-08-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:network:getBastionShareableLink'. */
export function getBastionShareableLink(args: GetBastionShareableLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetBastionShareableLinkResult> {
    pulumi.log.warn("getBastionShareableLink is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:network:getBastionShareableLink'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:network/latest:getBastionShareableLink", {
        "bastionHostName": args.bastionHostName,
        "resourceGroupName": args.resourceGroupName,
        "vms": args.vms,
    }, opts);
}

export interface GetBastionShareableLinkArgs {
    /**
     * The name of the Bastion Host.
     */
    readonly bastionHostName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * List of VM references.
     */
    readonly vms?: inputs.network.latest.BastionShareableLink[];
}

/**
 * Response for all the Bastion Shareable Link endpoints.
 */
export interface GetBastionShareableLinkResult {
    /**
     * The URL to get the next set of results.
     */
    readonly nextLink?: string;
    /**
     * List of Bastion Shareable Links for the request.
     */
    readonly value?: outputs.network.latest.BastionShareableLinkResponse[];
}
