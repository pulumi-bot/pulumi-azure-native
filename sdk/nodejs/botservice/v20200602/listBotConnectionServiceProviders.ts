// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listBotConnectionServiceProviders(args?: ListBotConnectionServiceProvidersArgs, opts?: pulumi.InvokeOptions): Promise<ListBotConnectionServiceProvidersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:botservice/v20200602:listBotConnectionServiceProviders", {
    }, opts);
}

export interface ListBotConnectionServiceProvidersArgs {
}

/**
 * The list of bot service providers response.
 */
export interface ListBotConnectionServiceProvidersResult {
    /**
     * The link used to get the next page of bot service providers.
     */
    readonly nextLink?: string;
    /**
     * Gets the list of bot service providers and their properties.
     */
    readonly value: outputs.botservice.v20200602.ServiceProviderResponse[];
}
