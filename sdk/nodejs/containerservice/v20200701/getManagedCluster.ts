// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getManagedCluster(args: GetManagedClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:containerservice/v20200701:getManagedCluster", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetManagedClusterArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the managed cluster resource.
     */
    readonly resourceName: string;
}

/**
 * Managed cluster.
 */
export interface GetManagedClusterResult {
    /**
     * Profile of Azure Active Directory configuration.
     */
    readonly aadProfile?: outputs.containerservice.v20200701.ManagedClusterAADProfileResponse;
    /**
     * Profile of managed cluster add-on.
     */
    readonly addonProfiles?: {[key: string]: outputs.containerservice.v20200701.ManagedClusterAddonProfileResponse};
    /**
     * Properties of the agent pool.
     */
    readonly agentPoolProfiles?: outputs.containerservice.v20200701.ManagedClusterAgentPoolProfileResponse[];
    /**
     * Access profile for managed cluster API server.
     */
    readonly apiServerAccessProfile?: outputs.containerservice.v20200701.ManagedClusterAPIServerAccessProfileResponse;
    /**
     * Parameters to be applied to the cluster-autoscaler when enabled
     */
    readonly autoScalerProfile?: outputs.containerservice.v20200701.ManagedClusterPropertiesResponseAutoScalerProfile;
    /**
     * ResourceId of the disk encryption set to use for enabling encryption at rest.
     */
    readonly diskEncryptionSetID?: string;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    readonly dnsPrefix?: string;
    /**
     * (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
     */
    readonly enablePodSecurityPolicy?: boolean;
    /**
     * Whether to enable Kubernetes Role-Based Access Control.
     */
    readonly enableRBAC?: boolean;
    /**
     * FQDN for the master pool.
     */
    readonly fqdn: string;
    /**
     * The identity of the managed cluster, if configured.
     */
    readonly identity?: outputs.containerservice.v20200701.ManagedClusterIdentityResponse;
    /**
     * Identities associated with the cluster.
     */
    readonly identityProfile?: {[key: string]: outputs.containerservice.v20200701.ManagedClusterPropertiesResponseIdentityProfile};
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    readonly kubernetesVersion?: string;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    readonly linuxProfile?: outputs.containerservice.v20200701.ContainerServiceLinuxProfileResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * The max number of agent pools for the managed cluster.
     */
    readonly maxAgentPools: number;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Profile of network configuration.
     */
    readonly networkProfile?: outputs.containerservice.v20200701.ContainerServiceNetworkProfileResponse;
    /**
     * Name of the resource group containing agent pool nodes.
     */
    readonly nodeResourceGroup?: string;
    /**
     * FQDN of private cluster.
     */
    readonly privateFQDN: string;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs.
     */
    readonly servicePrincipalProfile?: outputs.containerservice.v20200701.ManagedClusterServicePrincipalProfileResponse;
    /**
     * The managed cluster SKU.
     */
    readonly sku?: outputs.containerservice.v20200701.ManagedClusterSKUResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Profile for Windows VMs in the container service cluster.
     */
    readonly windowsProfile?: outputs.containerservice.v20200701.ManagedClusterWindowsProfileResponse;
}
