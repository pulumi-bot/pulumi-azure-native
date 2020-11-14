// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listAppServicePlanHybridConnectionKeys(args: ListAppServicePlanHybridConnectionKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListAppServicePlanHybridConnectionKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/latest:listAppServicePlanHybridConnectionKeys", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "relayName": args.relayName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListAppServicePlanHybridConnectionKeysArgs {
    /**
     * Name of the App Service plan.
     */
    readonly name: string;
    /**
     * The name of the Service Bus namespace.
     */
    readonly namespaceName: string;
    /**
     * The name of the Service Bus relay.
     */
    readonly relayName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Hybrid Connection key contract. This has the send key name and value for a Hybrid Connection.
 */
export interface ListAppServicePlanHybridConnectionKeysResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * The name of the send key.
     */
    readonly sendKeyName: string;
    /**
     * The value of the send key.
     */
    readonly sendKeyValue: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
