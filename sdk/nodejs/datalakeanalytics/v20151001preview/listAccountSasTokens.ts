// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listAccountSasTokens(args: ListAccountSasTokensArgs, opts?: pulumi.InvokeOptions): Promise<ListAccountSasTokensResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datalakeanalytics/v20151001preview:listAccountSasTokens", {
        "accountName": args.accountName,
        "containerName": args.containerName,
        "resourceGroupName": args.resourceGroupName,
        "storageAccountName": args.storageAccountName,
    }, opts);
}

export interface ListAccountSasTokensArgs {
    /**
     * The name of the Data Lake Analytics account from which an Azure Storage account's SAS token is being requested.
     */
    readonly accountName: string;
    /**
     * The name of the Azure storage container for which the SAS token is being requested.
     */
    readonly containerName: string;
    /**
     * The name of the Azure resource group that contains the Data Lake Analytics account.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Azure storage account for which the SAS token is being requested.
     */
    readonly storageAccountName: string;
}

/**
 * The SAS response that contains the storage account, container and associated SAS token for connection use.
 */
export interface ListAccountSasTokensResult {
    /**
     * the link (url) to the next page of results.
     */
    readonly nextLink: string;
    readonly value: outputs.datalakeanalytics.v20151001preview.SasTokenInfoResponse[];
}
