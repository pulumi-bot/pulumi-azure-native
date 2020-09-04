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
    return pulumi.runtime.invoke("azurerm:eventhub/v20180101preview:getNamespace", {
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNamespaceArgs {
    /**
     * The Namespace name
     */
    readonly namespaceName: string;
    /**
     * Name of the resource group within the azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Single Namespace item in List or Get Operation
 */
export interface GetNamespaceResult {
    /**
     * Cluster ARM ID of the Namespace.
     */
    readonly clusterArmId?: string;
    /**
     * The time the Namespace was created.
     */
    readonly createdAt: string;
    /**
     * Properties of BYOK Encryption description
     */
    readonly encryption?: outputs.eventhub.v20180101preview.EncryptionResponse;
    /**
     * Properties of BYOK Identity description
     */
    readonly identity?: outputs.eventhub.v20180101preview.IdentityResponse;
    /**
     * Value that indicates whether AutoInflate is enabled for eventhub namespace.
     */
    readonly isAutoInflateEnabled?: boolean;
    /**
     * Value that indicates whether Kafka is enabled for eventhub namespace.
     */
    readonly kafkaEnabled?: boolean;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
     */
    readonly maximumThroughputUnits?: number;
    /**
     * Identifier for Azure Insights metrics.
     */
    readonly metricId: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Provisioning state of the Namespace.
     */
    readonly provisioningState: string;
    /**
     * Endpoint you can use to perform Service Bus operations.
     */
    readonly serviceBusEndpoint: string;
    /**
     * Properties of sku resource
     */
    readonly sku?: outputs.eventhub.v20180101preview.SkuResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The time the Namespace was updated.
     */
    readonly updatedAt: string;
    /**
     * Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
     */
    readonly zoneRedundant?: boolean;
}
