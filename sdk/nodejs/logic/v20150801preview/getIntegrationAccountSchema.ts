// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIntegrationAccountSchema(args: GetIntegrationAccountSchemaArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationAccountSchemaResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:logic/v20150801preview:getIntegrationAccountSchema", {
        "integrationAccountName": args.integrationAccountName,
        "resourceGroupName": args.resourceGroupName,
        "schemaName": args.schemaName,
    }, opts);
}

export interface GetIntegrationAccountSchemaArgs {
    /**
     * The integration account name.
     */
    readonly integrationAccountName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The integration account schema name.
     */
    readonly schemaName: string;
}

export interface GetIntegrationAccountSchemaResult {
    /**
     * The changed time.
     */
    readonly changedTime: string;
    /**
     * The content.
     */
    readonly content?: {[key: string]: any};
    /**
     * The content link.
     */
    readonly contentLink: outputs.logic.v20150801preview.IntegrationAccountContentLinkResponse;
    /**
     * The content type.
     */
    readonly contentType?: string;
    /**
     * The created time.
     */
    readonly createdTime: string;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * The metadata.
     */
    readonly metadata?: {[key: string]: any};
    /**
     * The resource name.
     */
    readonly name?: string;
    /**
     * The schema type.
     */
    readonly schemaType?: string;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The target namespace.
     */
    readonly targetNamespace?: string;
    /**
     * The resource type.
     */
    readonly type?: string;
}
