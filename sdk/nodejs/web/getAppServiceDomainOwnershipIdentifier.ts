// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAppServiceDomainOwnershipIdentifier(args: GetAppServiceDomainOwnershipIdentifierArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceDomainOwnershipIdentifierResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web:getAppServiceDomainOwnershipIdentifier", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppServiceDomainOwnershipIdentifierArgs {
    /**
     * Name of domain ownership identifier.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * A domain specific resource identifier.
 */
export interface GetAppServiceDomainOwnershipIdentifierResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Identifier resource specific properties
     */
    readonly properties: outputs.web.IdentifierResponseProperties;
    /**
     * Resource type.
     */
    readonly type: string;
}
