// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getCloudConnector(args: GetCloudConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudConnectorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:costmanagement/v20190301preview:getCloudConnector", {
        "connectorName": args.connectorName,
        "expand": args.expand,
    }, opts);
}

export interface GetCloudConnectorArgs {
    /**
     * Connector Name.
     */
    readonly connectorName: string;
    /**
     * May be used to expand the collectionInfo property. By default, collectionInfo is not included.
     */
    readonly expand?: string;
}

/**
 * The Connector model definition
 */
export interface GetCloudConnectorResult {
    /**
     * Connector billing model
     */
    readonly billingModel?: string;
    /**
     * Collection information
     */
    readonly collectionInfo: outputs.costmanagement.v20190301preview.ConnectorCollectionInfoResponse;
    /**
     * Connector definition creation datetime
     */
    readonly createdOn: string;
    /**
     * Credentials authentication key (eg AWS ARN)
     */
    readonly credentialsKey?: string;
    /**
     * Credentials secret (eg AWS ExternalId)
     */
    readonly credentialsSecret?: string;
    /**
     * Number of days remaining of trial
     */
    readonly daysTrialRemaining: number;
    /**
     * Default ManagementGroupId
     */
    readonly defaultManagementGroupId?: string;
    /**
     * Connector DisplayName
     */
    readonly displayName?: string;
    /**
     * Associated ExternalBillingAccountId
     */
    readonly externalBillingAccountId: string;
    /**
     * Connector kind (eg aws)
     */
    readonly kind?: string;
    /**
     * Connector last modified datetime
     */
    readonly modifiedOn: string;
    /**
     * Connector name
     */
    readonly name: string;
    /**
     * The display name of the providerBillingAccountId as defined on the external provider
     */
    readonly providerBillingAccountDisplayName: string;
    /**
     * Connector providerBillingAccountId, determined from credentials (eg AWS Consolidated account number)
     */
    readonly providerBillingAccountId: string;
    /**
     * Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
     */
    readonly reportId?: string;
    /**
     * Connector status
     */
    readonly status: string;
    /**
     * Billing SubscriptionId
     */
    readonly subscriptionId?: string;
    /**
     * Connector type
     */
    readonly type: string;
}
