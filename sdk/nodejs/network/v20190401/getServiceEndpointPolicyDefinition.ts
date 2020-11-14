// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getServiceEndpointPolicyDefinition(args: GetServiceEndpointPolicyDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEndpointPolicyDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20190401:getServiceEndpointPolicyDefinition", {
        "resourceGroupName": args.resourceGroupName,
        "serviceEndpointPolicyDefinitionName": args.serviceEndpointPolicyDefinitionName,
        "serviceEndpointPolicyName": args.serviceEndpointPolicyName,
    }, opts);
}

export interface GetServiceEndpointPolicyDefinitionArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the service endpoint policy definition name.
     */
    readonly serviceEndpointPolicyDefinitionName: string;
    /**
     * The name of the service endpoint policy name.
     */
    readonly serviceEndpointPolicyName: string;
}

/**
 * Service Endpoint policy definitions.
 */
export interface GetServiceEndpointPolicyDefinitionResult {
    /**
     * A description for this rule. Restricted to 140 chars.
     */
    readonly description?: string;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * The provisioning state of the service end point policy definition. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState: string;
    /**
     * Service endpoint name.
     */
    readonly service?: string;
    /**
     * A list of service resources.
     */
    readonly serviceResources?: string[];
}
