// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:devices/v20190322preview:getCertificate", {
        "certificateName": args.certificateName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetCertificateArgs {
    /**
     * The name of the certificate
     */
    readonly certificateName: string;
    /**
     * The name of the resource group that contains the IoT hub.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the IoT hub.
     */
    readonly resourceName: string;
}

/**
 * The X509 Certificate.
 */
export interface GetCertificateResult {
    /**
     * The entity tag.
     */
    readonly etag: string;
    /**
     * The name of the certificate.
     */
    readonly name: string;
    /**
     * The description of an X509 CA Certificate.
     */
    readonly properties: outputs.devices.v20190322preview.CertificatePropertiesResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
