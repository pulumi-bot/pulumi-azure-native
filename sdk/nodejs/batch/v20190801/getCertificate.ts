// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batch/v20190801:getCertificate", {
        "accountName": args.accountName,
        "certificateName": args.certificateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCertificateArgs {
    /**
     * The name of the Batch account.
     */
    readonly accountName: string;
    /**
     * The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
     */
    readonly certificateName: string;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains information about a certificate.
 */
export interface GetCertificateResult {
    /**
     * This is only returned when the certificate provisioningState is 'Failed'.
     */
    readonly deleteCertificateError: outputs.batch.v20190801.DeleteCertificateErrorResponse;
    /**
     * The ETag of the resource, used for concurrency statements.
     */
    readonly etag: string;
    /**
     * The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
     */
    readonly format?: CertificateFormat;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The previous provisioned state of the resource
     */
    readonly previousProvisioningState: CertificateProvisioningState;
    readonly previousProvisioningStateTransitionTime: string;
    readonly provisioningState: CertificateProvisioningState;
    readonly provisioningStateTransitionTime: string;
    /**
     * The public key of the certificate.
     */
    readonly publicData: string;
    /**
     * This must match the thumbprint from the name.
     */
    readonly thumbprint?: string;
    /**
     * This must match the first portion of the certificate name. Currently required to be 'SHA1'.
     */
    readonly thumbprintAlgorithm?: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
