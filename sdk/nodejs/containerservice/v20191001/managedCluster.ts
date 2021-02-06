// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Managed cluster.
 */
export class ManagedCluster extends pulumi.CustomResource {
    /**
     * Get an existing ManagedCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedCluster {
        return new ManagedCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:containerservice/v20191001:ManagedCluster';

    /**
     * Returns true if the given object is an instance of ManagedCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedCluster.__pulumiType;
    }

    /**
     * Profile of Azure Active Directory configuration.
     */
    public readonly aadProfile!: pulumi.Output<outputs.containerservice.v20191001.ManagedClusterAADProfileResponse | undefined>;
    /**
     * Profile of managed cluster add-on.
     */
    public readonly addonProfiles!: pulumi.Output<{[key: string]: outputs.containerservice.v20191001.ManagedClusterAddonProfileResponse} | undefined>;
    /**
     * Properties of the agent pool.
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.v20191001.ManagedClusterAgentPoolProfileResponse[] | undefined>;
    /**
     * Access profile for managed cluster API server.
     */
    public readonly apiServerAccessProfile!: pulumi.Output<outputs.containerservice.v20191001.ManagedClusterAPIServerAccessProfileResponse | undefined>;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    public readonly dnsPrefix!: pulumi.Output<string | undefined>;
    /**
     * (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
     */
    public readonly enablePodSecurityPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable Kubernetes Role-Based Access Control.
     */
    public readonly enableRBAC!: pulumi.Output<boolean | undefined>;
    /**
     * FQDN for the master pool.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The identity of the managed cluster, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.containerservice.v20191001.ManagedClusterIdentityResponse | undefined>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    public readonly kubernetesVersion!: pulumi.Output<string | undefined>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.v20191001.ContainerServiceLinuxProfileResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The max number of agent pools for the managed cluster.
     */
    public /*out*/ readonly maxAgentPools!: pulumi.Output<number>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Profile of network configuration.
     */
    public readonly networkProfile!: pulumi.Output<outputs.containerservice.v20191001.ContainerServiceNetworkProfileResponse | undefined>;
    /**
     * Name of the resource group containing agent pool nodes.
     */
    public readonly nodeResourceGroup!: pulumi.Output<string | undefined>;
    /**
     * FQDN of private cluster.
     */
    public /*out*/ readonly privateFQDN!: pulumi.Output<string>;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs.
     */
    public readonly servicePrincipalProfile!: pulumi.Output<outputs.containerservice.v20191001.ManagedClusterServicePrincipalProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Profile for Windows VMs in the container service cluster.
     */
    public readonly windowsProfile!: pulumi.Output<outputs.containerservice.v20191001.ManagedClusterWindowsProfileResponse | undefined>;

    /**
     * Create a ManagedCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.resourceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["aadProfile"] = args ? args.aadProfile : undefined;
            inputs["addonProfiles"] = args ? args.addonProfiles : undefined;
            inputs["agentPoolProfiles"] = args ? args.agentPoolProfiles : undefined;
            inputs["apiServerAccessProfile"] = args ? args.apiServerAccessProfile : undefined;
            inputs["dnsPrefix"] = args ? args.dnsPrefix : undefined;
            inputs["enablePodSecurityPolicy"] = args ? args.enablePodSecurityPolicy : undefined;
            inputs["enableRBAC"] = args ? args.enableRBAC : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["kubernetesVersion"] = args ? args.kubernetesVersion : undefined;
            inputs["linuxProfile"] = args ? args.linuxProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["nodeResourceGroup"] = args ? args.nodeResourceGroup : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["servicePrincipalProfile"] = args ? args.servicePrincipalProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["windowsProfile"] = args ? args.windowsProfile : undefined;
            inputs["fqdn"] = undefined /*out*/;
            inputs["maxAgentPools"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateFQDN"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["aadProfile"] = undefined /*out*/;
            inputs["addonProfiles"] = undefined /*out*/;
            inputs["agentPoolProfiles"] = undefined /*out*/;
            inputs["apiServerAccessProfile"] = undefined /*out*/;
            inputs["dnsPrefix"] = undefined /*out*/;
            inputs["enablePodSecurityPolicy"] = undefined /*out*/;
            inputs["enableRBAC"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["kubernetesVersion"] = undefined /*out*/;
            inputs["linuxProfile"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maxAgentPools"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkProfile"] = undefined /*out*/;
            inputs["nodeResourceGroup"] = undefined /*out*/;
            inputs["privateFQDN"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["servicePrincipalProfile"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["windowsProfile"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:containerservice/latest:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20170831:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20180331:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20180801preview:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20190201:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20190401:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20190601:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20190801:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20191101:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200101:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200201:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200301:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200401:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200601:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200701:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20200901:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20201101:ManagedCluster" }, { type: "azure-nextgen:containerservice/v20201201:ManagedCluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagedCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedCluster resource.
 */
export interface ManagedClusterArgs {
    /**
     * Profile of Azure Active Directory configuration.
     */
    readonly aadProfile?: pulumi.Input<inputs.containerservice.v20191001.ManagedClusterAADProfile>;
    /**
     * Profile of managed cluster add-on.
     */
    readonly addonProfiles?: pulumi.Input<{[key: string]: pulumi.Input<inputs.containerservice.v20191001.ManagedClusterAddonProfile>}>;
    /**
     * Properties of the agent pool.
     */
    readonly agentPoolProfiles?: pulumi.Input<pulumi.Input<inputs.containerservice.v20191001.ManagedClusterAgentPoolProfile>[]>;
    /**
     * Access profile for managed cluster API server.
     */
    readonly apiServerAccessProfile?: pulumi.Input<inputs.containerservice.v20191001.ManagedClusterAPIServerAccessProfile>;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    readonly dnsPrefix?: pulumi.Input<string>;
    /**
     * (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
     */
    readonly enablePodSecurityPolicy?: pulumi.Input<boolean>;
    /**
     * Whether to enable Kubernetes Role-Based Access Control.
     */
    readonly enableRBAC?: pulumi.Input<boolean>;
    /**
     * The identity of the managed cluster, if configured.
     */
    readonly identity?: pulumi.Input<inputs.containerservice.v20191001.ManagedClusterIdentity>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    readonly kubernetesVersion?: pulumi.Input<string>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    readonly linuxProfile?: pulumi.Input<inputs.containerservice.v20191001.ContainerServiceLinuxProfile>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Profile of network configuration.
     */
    readonly networkProfile?: pulumi.Input<inputs.containerservice.v20191001.ContainerServiceNetworkProfile>;
    /**
     * Name of the resource group containing agent pool nodes.
     */
    readonly nodeResourceGroup?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the managed cluster resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs.
     */
    readonly servicePrincipalProfile?: pulumi.Input<inputs.containerservice.v20191001.ManagedClusterServicePrincipalProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Profile for Windows VMs in the container service cluster.
     */
    readonly windowsProfile?: pulumi.Input<inputs.containerservice.v20191001.ManagedClusterWindowsProfile>;
}
