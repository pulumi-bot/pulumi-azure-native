// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getLogger(args: GetLoggerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoggerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20191201preview:getLogger", {
        "loggerId": args.loggerId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetLoggerArgs {
    /**
     * Logger identifier. Must be unique in the API Management service instance.
     */
    readonly loggerId: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Logger details.
 */
export interface GetLoggerResult {
    /**
     * The name and SendRule connection string of the event hub for azureEventHub logger.
     * Instrumentation key for applicationInsights logger.
     */
    readonly credentials: {[key: string]: string};
    /**
     * Logger description.
     */
    readonly description?: string;
    /**
     * Whether records are buffered in the logger before publishing. Default is assumed to be true.
     */
    readonly isBuffered?: boolean;
    /**
     * Logger type.
     */
    readonly loggerType: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
     */
    readonly resourceId?: string;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}
