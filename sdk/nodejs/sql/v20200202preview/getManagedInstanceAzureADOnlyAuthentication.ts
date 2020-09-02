// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getManagedInstanceAzureADOnlyAuthentication(args: GetManagedInstanceAzureADOnlyAuthenticationArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedInstanceAzureADOnlyAuthenticationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:sql/v20200202preview:getManagedInstanceAzureADOnlyAuthentication", {
        "authenticationName": args.authenticationName,
        "managedInstanceName": args.managedInstanceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedInstanceAzureADOnlyAuthenticationArgs {
    /**
     * The name of server azure active directory only authentication.
     */
    readonly authenticationName: string;
    /**
     * The name of the managed instance.
     */
    readonly managedInstanceName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
}

/**
 * Azure Active Directory only authentication.
 */
export interface GetManagedInstanceAzureADOnlyAuthenticationResult {
    /**
     * Azure Active Directory only Authentication enabled.
     */
    readonly azureADOnlyAuthentication: boolean;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
