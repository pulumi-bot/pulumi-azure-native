// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getWebAppPublicCertificate(args: GetWebAppPublicCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppPublicCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/v20181101:getWebAppPublicCertificate", {
        "name": args.name,
        "publicCertificateName": args.publicCertificateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetWebAppPublicCertificateArgs {
    /**
     * Name of the app.
     */
    readonly name: string;
    /**
     * Public certificate name.
     */
    readonly publicCertificateName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Public certificate object
 */
export interface GetWebAppPublicCertificateResult {
    /**
     * Public Certificate byte array
     */
    readonly blob?: string;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Public Certificate Location
     */
    readonly publicCertificateLocation?: string;
    /**
     * Certificate Thumbprint
     */
    readonly thumbprint: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
