// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRegisteredAsn(args: GetRegisteredAsnArgs, opts?: pulumi.InvokeOptions): Promise<GetRegisteredAsnResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:peering/v20200101preview:getRegisteredAsn", {
        "peeringName": args.peeringName,
        "registeredAsnName": args.registeredAsnName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRegisteredAsnArgs {
    /**
     * The name of the peering.
     */
    readonly peeringName: string;
    /**
     * The name of the registered ASN.
     */
    readonly registeredAsnName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The customer's ASN that is registered by the peering service provider.
 */
export interface GetRegisteredAsnResult {
    /**
     * The customer's ASN from which traffic originates.
     */
    readonly asn?: number;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The peering service prefix key that is to be shared with the customer.
     */
    readonly peeringServicePrefixKey: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
