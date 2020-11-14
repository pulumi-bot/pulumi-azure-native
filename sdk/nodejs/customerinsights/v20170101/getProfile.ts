// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getProfile(args: GetProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:customerinsights/v20170101:getProfile", {
        "hubName": args.hubName,
        "localeCode": args.localeCode,
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProfileArgs {
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * Locale of profile to retrieve, default is en-us.
     */
    readonly localeCode?: string;
    /**
     * The name of the profile.
     */
    readonly profileName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The profile resource format.
 */
export interface GetProfileResult {
    /**
     * The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
     */
    readonly apiEntitySetName?: string;
    /**
     * The attributes for the Type.
     */
    readonly attributes?: {[key: string]: string[]};
    /**
     * Localized descriptions for the property.
     */
    readonly description?: {[key: string]: string};
    /**
     * Localized display names for the property.
     */
    readonly displayName?: {[key: string]: string};
    /**
     * Type of entity.
     */
    readonly entityType?: string;
    /**
     * The properties of the Profile.
     */
    readonly fields?: outputs.customerinsights.v20170101.PropertyDefinitionResponse[];
    /**
     * The instance count.
     */
    readonly instancesCount?: number;
    /**
     * Large Image associated with the Property or EntityType.
     */
    readonly largeImage?: string;
    /**
     * The last changed time for the type definition.
     */
    readonly lastChangedUtc: string;
    /**
     * Any custom localized attributes for the Type.
     */
    readonly localizedAttributes?: {[key: string]: {[key: string]: string}};
    /**
     * Medium Image associated with the Property or EntityType.
     */
    readonly mediumImage?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Provisioning state.
     */
    readonly provisioningState: string;
    /**
     * The schema org link. This helps ACI identify and suggest semantic models.
     */
    readonly schemaItemTypeLink?: string;
    /**
     * Small Image associated with the Property or EntityType.
     */
    readonly smallImage?: string;
    /**
     * The strong IDs.
     */
    readonly strongIds?: outputs.customerinsights.v20170101.StrongIdResponse[];
    /**
     * The hub name.
     */
    readonly tenantId: string;
    /**
     * The timestamp property name. Represents the time when the interaction or profile update happened.
     */
    readonly timestampFieldName?: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The name of the entity.
     */
    readonly typeName?: string;
}
