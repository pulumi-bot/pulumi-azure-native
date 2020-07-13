// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getBatchAccountCertificate(args: GetBatchAccountCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetBatchAccountCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batch:getBatchAccountCertificate", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetBatchAccountCertificateArgs {
    /**
     * The name of the Batch account.
     */
    readonly accountName: string;
    /**
     * The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
     */
    readonly name: string;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains information about a certificate.
 */
export interface GetBatchAccountCertificateResult {
    /**
     * The ETag of the resource, used for concurrency statements.
     */
    readonly etag: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The properties associated with the certificate.
     */
    readonly properties: outputs.batch.CertificatePropertiesResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
}