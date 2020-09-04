// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
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
    public static readonly __pulumiType = 'azurerm:containerservice/v20180801preview:ManagedCluster';

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
    public readonly aadProfile!: pulumi.Output<outputs.containerservice.v20180801preview.ManagedClusterAADProfileResponse | undefined>;
    /**
     * Profile of managed cluster add-on.
     */
    public readonly addonProfiles!: pulumi.Output<{[key: string]: outputs.containerservice.v20180801preview.ManagedClusterAddonProfileResponse} | undefined>;
    /**
     * Properties of the agent pool.
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.v20180801preview.ManagedClusterAgentPoolProfileResponse[] | undefined>;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    public readonly dnsPrefix!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable Kubernetes Role-Based Access Control.
     */
    public readonly enableRBAC!: pulumi.Output<boolean | undefined>;
    /**
     * FQDN for the master pool.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    public readonly kubernetesVersion!: pulumi.Output<string | undefined>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.v20180801preview.ContainerServiceLinuxProfileResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Profile of network configuration.
     */
    public readonly networkProfile!: pulumi.Output<outputs.containerservice.v20180801preview.ContainerServiceNetworkProfileResponse | undefined>;
    /**
     * Name of the resource group containing agent pool nodes.
     */
    public /*out*/ readonly nodeResourceGroup!: pulumi.Output<string>;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs.
     */
    public readonly servicePrincipalProfile!: pulumi.Output<outputs.containerservice.v20180801preview.ManagedClusterServicePrincipalProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

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
            inputs["dnsPrefix"] = args ? args.dnsPrefix : undefined;
            inputs["enableRBAC"] = args ? args.enableRBAC : undefined;
            inputs["kubernetesVersion"] = args ? args.kubernetesVersion : undefined;
            inputs["linuxProfile"] = args ? args.linuxProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["servicePrincipalProfile"] = args ? args.servicePrincipalProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["fqdn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeResourceGroup"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["aadProfile"] = undefined /*out*/;
            inputs["addonProfiles"] = undefined /*out*/;
            inputs["agentPoolProfiles"] = undefined /*out*/;
            inputs["dnsPrefix"] = undefined /*out*/;
            inputs["enableRBAC"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["kubernetesVersion"] = undefined /*out*/;
            inputs["linuxProfile"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkProfile"] = undefined /*out*/;
            inputs["nodeResourceGroup"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["servicePrincipalProfile"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerservice/latest:ManagedCluster" }, { type: "azurerm:containerservice/v20170831:ManagedCluster" }, { type: "azurerm:containerservice/v20180331:ManagedCluster" }, { type: "azurerm:containerservice/v20190201:ManagedCluster" }, { type: "azurerm:containerservice/v20190401:ManagedCluster" }, { type: "azurerm:containerservice/v20190601:ManagedCluster" }, { type: "azurerm:containerservice/v20190801:ManagedCluster" }, { type: "azurerm:containerservice/v20191001:ManagedCluster" }, { type: "azurerm:containerservice/v20191101:ManagedCluster" }, { type: "azurerm:containerservice/v20200101:ManagedCluster" }, { type: "azurerm:containerservice/v20200201:ManagedCluster" }, { type: "azurerm:containerservice/v20200301:ManagedCluster" }, { type: "azurerm:containerservice/v20200401:ManagedCluster" }, { type: "azurerm:containerservice/v20200601:ManagedCluster" }, { type: "azurerm:containerservice/v20200701:ManagedCluster" }] };
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
    readonly aadProfile?: pulumi.Input<inputs.containerservice.v20180801preview.ManagedClusterAADProfile>;
    /**
     * Profile of managed cluster add-on.
     */
    readonly addonProfiles?: pulumi.Input<{[key: string]: pulumi.Input<inputs.containerservice.v20180801preview.ManagedClusterAddonProfile>}>;
    /**
     * Properties of the agent pool.
     */
    readonly agentPoolProfiles?: pulumi.Input<pulumi.Input<inputs.containerservice.v20180801preview.ManagedClusterAgentPoolProfile>[]>;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    readonly dnsPrefix?: pulumi.Input<string>;
    /**
     * Whether to enable Kubernetes Role-Based Access Control.
     */
    readonly enableRBAC?: pulumi.Input<boolean>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    readonly kubernetesVersion?: pulumi.Input<string>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    readonly linuxProfile?: pulumi.Input<inputs.containerservice.v20180801preview.ContainerServiceLinuxProfile>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Profile of network configuration.
     */
    readonly networkProfile?: pulumi.Input<inputs.containerservice.v20180801preview.ContainerServiceNetworkProfile>;
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
    readonly servicePrincipalProfile?: pulumi.Input<inputs.containerservice.v20180801preview.ManagedClusterServicePrincipalProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
