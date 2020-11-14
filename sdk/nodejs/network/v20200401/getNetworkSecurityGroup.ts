// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getNetworkSecurityGroup(args: GetNetworkSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkSecurityGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20200401:getNetworkSecurityGroup", {
        "expand": args.expand,
        "networkSecurityGroupName": args.networkSecurityGroupName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetworkSecurityGroupArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the network security group.
     */
    readonly networkSecurityGroupName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * NetworkSecurityGroup resource.
 */
export interface GetNetworkSecurityGroupResult {
    /**
     * The default security rules of network security group.
     */
    readonly defaultSecurityRules: outputs.network.v20200401.SecurityRuleResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * A collection of references to flow log resources.
     */
    readonly flowLogs: outputs.network.v20200401.FlowLogResponse[];
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * A collection of references to network interfaces.
     */
    readonly networkInterfaces: outputs.network.v20200401.NetworkInterfaceResponse[];
    /**
     * The provisioning state of the network security group resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the network security group resource.
     */
    readonly resourceGuid: string;
    /**
     * A collection of security rules of the network security group.
     */
    readonly securityRules?: outputs.network.v20200401.SecurityRuleResponse[];
    /**
     * A collection of references to subnets.
     */
    readonly subnets: outputs.network.v20200401.SubnetResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
