// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getNamespace(args: GetNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:relay/v20170401:getNamespace", {
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNamespaceArgs {
    /**
     * The namespace name
     */
    readonly namespaceName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Description of a namespace resource.
 */
export interface GetNamespaceResult {
    /**
     * The time the namespace was created.
     */
    readonly createdAt: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Identifier for Azure Insights metrics.
     */
    readonly metricId: string;
    /**
     * Resource name.
     */
    readonly name: string;
    readonly provisioningState: provisioningStateEnum;
    /**
     * Endpoint you can use to perform Service Bus operations.
     */
    readonly serviceBusEndpoint: string;
    /**
     * SKU of the namespace.
     */
    readonly sku?: outputs.relay.v20170401.SkuResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The time the namespace was updated.
     */
    readonly updatedAt: string;
}
