// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getMoveCollection(args: GetMoveCollectionArgs, opts?: pulumi.InvokeOptions): Promise<GetMoveCollectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:migrate/v20191001preview:getMoveCollection", {
        "moveCollectionName": args.moveCollectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMoveCollectionArgs {
    /**
     * The Move Collection Name.
     */
    readonly moveCollectionName: string;
    /**
     * The Resource Group Name.
     */
    readonly resourceGroupName: string;
}

/**
 * Define the move collection.
 */
export interface GetMoveCollectionResult {
    /**
     * Defines the MSI properties of the Move Collection.
     */
    readonly identity?: outputs.migrate.v20191001preview.IdentityResponse;
    /**
     * The geo-location where the resource lives.
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Defines the move collection properties.
     */
    readonly properties: outputs.migrate.v20191001preview.MoveCollectionPropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
