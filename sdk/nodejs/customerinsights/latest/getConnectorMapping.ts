// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getConnectorMapping(args: GetConnectorMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectorMappingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:customerinsights/latest:getConnectorMapping", {
        "connectorName": args.connectorName,
        "hubName": args.hubName,
        "mappingName": args.mappingName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectorMappingArgs {
    /**
     * The name of the connector.
     */
    readonly connectorName: string;
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * The name of the connector mapping.
     */
    readonly mappingName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The connector mapping resource format.
 */
export interface GetConnectorMappingResult {
    /**
     * The connector mapping name
     */
    readonly connectorMappingName: string;
    /**
     * The connector name.
     */
    readonly connectorName: string;
    /**
     * Type of connector.
     */
    readonly connectorType?: string;
    /**
     * The created time.
     */
    readonly created: string;
    /**
     * The DataFormat ID.
     */
    readonly dataFormatId: string;
    /**
     * The description of the connector mapping.
     */
    readonly description?: string;
    /**
     * Display name for the connector mapping.
     */
    readonly displayName?: string;
    /**
     * Defines which entity type the file should map to.
     */
    readonly entityType: string;
    /**
     * The mapping entity name.
     */
    readonly entityTypeName: string;
    /**
     * The last modified time.
     */
    readonly lastModified: string;
    /**
     * The properties of the mapping.
     */
    readonly mappingProperties: outputs.customerinsights.latest.ConnectorMappingPropertiesResponse;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The next run time based on customer's settings.
     */
    readonly nextRunTime: string;
    /**
     * The RunId.
     */
    readonly runId: string;
    /**
     * State of connector mapping.
     */
    readonly state: string;
    /**
     * The hub name.
     */
    readonly tenantId: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
