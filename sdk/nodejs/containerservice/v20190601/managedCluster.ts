// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Managed cluster.
 *
 * ## Example Usage
 * ### Create/Update Managed Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const managedCluster = new azurerm.containerservice.v20190601.ManagedCluster("managedCluster", {
 *     addonProfiles: {},
 *     agentPoolProfiles: [{
 *         availabilityZones: [
 *             "1",
 *             "2",
 *             "3",
 *         ],
 *         count: 3,
 *         enableNodePublicIP: true,
 *         name: "nodepool1",
 *         osType: "Linux",
 *         type: "VirtualMachineScaleSets",
 *         vmSize: "Standard_DS1_v2",
 *     }],
 *     dnsPrefix: "dnsprefix1",
 *     enablePodSecurityPolicy: true,
 *     enableRBAC: true,
 *     kubernetesVersion: "",
 *     linuxProfile: {
 *         adminUsername: "azureuser",
 *         ssh: {
 *             publicKeys: [{
 *                 keyData: "keydata",
 *             }],
 *         },
 *     },
 *     location: "location1",
 *     networkProfile: {
 *         loadBalancerSku: "basic",
 *     },
 *     resourceGroupName: "rg1",
 *     resourceName: "clustername1",
 *     servicePrincipalProfile: {
 *         clientId: "clientid",
 *         secret: "secret",
 *     },
 *     tags: {
 *         archv2: "",
 *         tier: "production",
 *     },
 *     windowsProfile: {
 *         adminPassword: `replacePassword1234$`,
 *         adminUsername: "azureuser",
 *     },
 * });
 *
 * ```
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
        return new ManagedCluster(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerservice/v20190601:ManagedCluster';

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
    public readonly aadProfile!: pulumi.Output<outputs.containerservice.v20190601.ManagedClusterAADProfileResponse | undefined>;
    /**
     * Profile of managed cluster add-on.
     */
    public readonly addonProfiles!: pulumi.Output<{[key: string]: outputs.containerservice.v20190601.ManagedClusterAddonProfileResponse} | undefined>;
    /**
     * Properties of the agent pool.
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.v20190601.ManagedClusterAgentPoolProfileResponse[] | undefined>;
    /**
     * (PREVIEW) Authorized IP Ranges to kubernetes API server.
     */
    public readonly apiServerAuthorizedIPRanges!: pulumi.Output<string[] | undefined>;
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
    public readonly identity!: pulumi.Output<outputs.containerservice.v20190601.ManagedClusterIdentityResponse | undefined>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    public readonly kubernetesVersion!: pulumi.Output<string | undefined>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.v20190601.ContainerServiceLinuxProfileResponse | undefined>;
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
    public readonly networkProfile!: pulumi.Output<outputs.containerservice.v20190601.ContainerServiceNetworkProfileResponse | undefined>;
    /**
     * Name of the resource group containing agent pool nodes.
     */
    public readonly nodeResourceGroup!: pulumi.Output<string | undefined>;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs.
     */
    public readonly servicePrincipalProfile!: pulumi.Output<outputs.containerservice.v20190601.ManagedClusterServicePrincipalProfileResponse | undefined>;
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
    public readonly windowsProfile!: pulumi.Output<outputs.containerservice.v20190601.ManagedClusterWindowsProfileResponse | undefined>;

    /**
     * Create a ManagedCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ManagedClusterArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["aadProfile"] = args ? args.aadProfile : undefined;
            inputs["addonProfiles"] = args ? args.addonProfiles : undefined;
            inputs["agentPoolProfiles"] = args ? args.agentPoolProfiles : undefined;
            inputs["apiServerAuthorizedIPRanges"] = args ? args.apiServerAuthorizedIPRanges : undefined;
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
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerservice/latest:ManagedCluster" }, { type: "azurerm:containerservice/v20170831:ManagedCluster" }, { type: "azurerm:containerservice/v20180331:ManagedCluster" }, { type: "azurerm:containerservice/v20180801preview:ManagedCluster" }, { type: "azurerm:containerservice/v20190201:ManagedCluster" }, { type: "azurerm:containerservice/v20190401:ManagedCluster" }, { type: "azurerm:containerservice/v20190801:ManagedCluster" }, { type: "azurerm:containerservice/v20191001:ManagedCluster" }, { type: "azurerm:containerservice/v20191101:ManagedCluster" }, { type: "azurerm:containerservice/v20200101:ManagedCluster" }, { type: "azurerm:containerservice/v20200201:ManagedCluster" }, { type: "azurerm:containerservice/v20200301:ManagedCluster" }, { type: "azurerm:containerservice/v20200401:ManagedCluster" }, { type: "azurerm:containerservice/v20200601:ManagedCluster" }, { type: "azurerm:containerservice/v20200701:ManagedCluster" }] };
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
    readonly aadProfile?: pulumi.Input<inputs.containerservice.v20190601.ManagedClusterAADProfile>;
    /**
     * Profile of managed cluster add-on.
     */
    readonly addonProfiles?: pulumi.Input<{[key: string]: pulumi.Input<inputs.containerservice.v20190601.ManagedClusterAddonProfile>}>;
    /**
     * Properties of the agent pool.
     */
    readonly agentPoolProfiles?: pulumi.Input<pulumi.Input<inputs.containerservice.v20190601.ManagedClusterAgentPoolProfile>[]>;
    /**
     * (PREVIEW) Authorized IP Ranges to kubernetes API server.
     */
    readonly apiServerAuthorizedIPRanges?: pulumi.Input<pulumi.Input<string>[]>;
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
    readonly identity?: pulumi.Input<inputs.containerservice.v20190601.ManagedClusterIdentity>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    readonly kubernetesVersion?: pulumi.Input<string>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    readonly linuxProfile?: pulumi.Input<inputs.containerservice.v20190601.ContainerServiceLinuxProfile>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Profile of network configuration.
     */
    readonly networkProfile?: pulumi.Input<inputs.containerservice.v20190601.ContainerServiceNetworkProfile>;
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
    readonly servicePrincipalProfile?: pulumi.Input<inputs.containerservice.v20190601.ManagedClusterServicePrincipalProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Profile for Windows VMs in the container service cluster.
     */
    readonly windowsProfile?: pulumi.Input<inputs.containerservice.v20190601.ManagedClusterWindowsProfile>;
}
