// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getEventSource(args: GetEventSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetEventSourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:timeseriesinsights/v20170228preview:getEventSource", {
        "environmentName": args.environmentName,
        "eventSourceName": args.eventSourceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEventSourceArgs {
    /**
     * The name of the Time Series Insights environment associated with the specified resource group.
     */
    readonly environmentName: string;
    /**
     * The name of the Time Series Insights event source associated with the specified environment.
     */
    readonly eventSourceName: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * An environment receives data from one or more event sources. Each event source has associated connection info that allows the Time Series Insights ingress pipeline to connect to and pull data from the event source
 */
export interface GetEventSourceResult {
    /**
     * The kind of the event source.
     */
    readonly kind: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
