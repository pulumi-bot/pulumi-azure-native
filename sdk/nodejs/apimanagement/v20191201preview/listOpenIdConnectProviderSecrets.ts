// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listOpenIdConnectProviderSecrets(args: ListOpenIdConnectProviderSecretsArgs, opts?: pulumi.InvokeOptions): Promise<ListOpenIdConnectProviderSecretsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20191201preview:listOpenIdConnectProviderSecrets", {
        "opid": args.opid,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface ListOpenIdConnectProviderSecretsArgs {
    /**
     * Identifier of the OpenID Connect Provider.
     */
    readonly opid: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
 */
export interface ListOpenIdConnectProviderSecretsResult {
    /**
     * Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
     */
    readonly clientSecret?: string;
}
