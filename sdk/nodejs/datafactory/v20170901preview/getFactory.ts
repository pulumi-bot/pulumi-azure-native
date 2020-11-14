// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getFactory(args: GetFactoryArgs, opts?: pulumi.InvokeOptions): Promise<GetFactoryResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:datafactory/v20170901preview:getFactory", {
        "factoryName": args.factoryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFactoryArgs {
    /**
     * The factory name.
     */
    readonly factoryName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * Factory resource type.
 */
export interface GetFactoryResult {
    /**
     * Time the factory was created in ISO8601 format.
     */
    readonly createTime: string;
    /**
     * Managed service identity of the factory.
     */
    readonly identity?: outputs.datafactory.v20170901preview.FactoryIdentityResponse;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * Factory provisioning state, example Succeeded.
     */
    readonly provisioningState: string;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type.
     */
    readonly type: string;
    /**
     * Version of the factory.
     */
    readonly version: string;
    /**
     * VSTS repo information of the factory.
     */
    readonly vstsConfiguration?: outputs.datafactory.v20170901preview.FactoryVSTSConfigurationResponse;
}
