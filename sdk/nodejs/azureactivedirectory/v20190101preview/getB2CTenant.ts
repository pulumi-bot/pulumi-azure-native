// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getB2CTenant(args: GetB2CTenantArgs, opts?: pulumi.InvokeOptions): Promise<GetB2CTenantResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:azureactivedirectory/v20190101preview:getB2CTenant", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetB2CTenantArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The initial domain name of the B2C tenant.
     */
    readonly resourceName: string;
}

export interface GetB2CTenantResult {
    /**
     * The billing configuration for the tenant.
     */
    readonly billingConfig?: outputs.azureactivedirectory.v20190101preview.B2CTenantResourcePropertiesResponseBillingConfig;
    /**
     * An identifier that represents the B2C tenant resource.
     */
    readonly id: string;
    /**
     * The location in which the resource is hosted and data resides. Refer to [this documentation](https://aka.ms/B2CDataResidency) to see valid data residency locations. Please choose one of 'United States', 'Europe', and 'Asia Pacific'.
     */
    readonly location: string;
    /**
     * The name of the B2C tenant resource.
     */
    readonly name: string;
    /**
     * SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
     */
    readonly sku: outputs.azureactivedirectory.v20190101preview.B2CResourceSKUResponse;
    /**
     * Resource Tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * An identifier of the B2C tenant.
     */
    readonly tenantId?: string;
    /**
     * The type of the B2C tenant resource.
     */
    readonly type: string;
}
