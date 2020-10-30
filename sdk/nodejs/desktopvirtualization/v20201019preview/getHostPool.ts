// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getHostPool(args: GetHostPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetHostPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:desktopvirtualization/v20201019preview:getHostPool", {
        "hostPoolName": args.hostPoolName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostPoolArgs {
    /**
     * The name of the host pool within the specified resource group
     */
    readonly hostPoolName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Represents a HostPool definition.
 */
export interface GetHostPoolResult {
    /**
     * List of applicationGroup links.
     */
    readonly applicationGroupReferences: string[];
    /**
     * Custom rdp property of HostPool.
     */
    readonly customRdpProperty?: string;
    /**
     * Description of HostPool.
     */
    readonly description?: string;
    /**
     * Friendly name of HostPool.
     */
    readonly friendlyName?: string;
    /**
     * HostPool type for desktop.
     */
    readonly hostPoolType: string;
    /**
     * The type of the load balancer.
     */
    readonly loadBalancerType: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The max session limit of HostPool.
     */
    readonly maxSessionLimit?: number;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * PersonalDesktopAssignment type for HostPool.
     */
    readonly personalDesktopAssignmentType?: string;
    /**
     * The type of preferred application group type, default to Desktop Application Group
     */
    readonly preferredAppGroupType: string;
    /**
     * The registration info of HostPool.
     */
    readonly registrationInfo?: outputs.desktopvirtualization.v20201019preview.RegistrationInfoResponse;
    /**
     * The ring number of HostPool.
     */
    readonly ring?: number;
    /**
     * ClientId for the registered Relying Party used to issue WVD SSO certificates.
     */
    readonly ssoClientId?: string;
    /**
     * Path to Azure KeyVault storing the secret used for communication to ADFS.
     */
    readonly ssoClientSecretKeyVaultPath?: string;
    /**
     * Path to keyvault containing ssoContext secret.
     */
    readonly ssoContext?: string;
    /**
     * The type of single sign on Secret Type.
     */
    readonly ssoSecretType?: string;
    /**
     * URL to customer ADFS server for signing WVD SSO certificates.
     */
    readonly ssoadfsAuthority?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Is validation environment.
     */
    readonly validationEnvironment?: boolean;
    /**
     * VM template for sessionhosts configuration within hostpool.
     */
    readonly vmTemplate?: string;
}
