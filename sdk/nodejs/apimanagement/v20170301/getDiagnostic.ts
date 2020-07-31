// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDiagnostic(args: GetDiagnosticArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagnosticResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20170301:getDiagnostic", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetDiagnosticArgs {
    /**
     * Diagnostic identifier. Must be unique in the current API Management service instance.
     */
    readonly name: string;
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
 * Diagnostic details.
 */
export interface GetDiagnosticResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Diagnostic entity contract properties.
     */
    readonly properties: outputs.apimanagement.v20170301.DiagnosticContractPropertiesResponse;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}