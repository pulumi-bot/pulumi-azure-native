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
 * const managedCluster = new azurerm.containerservice.v20170831.ManagedCluster("managedCluster", {
 *     agentPoolProfiles: [{
 *         count: 1,
 *         name: "agentpool1",
 *         vmSize: "Standard_D2_v2",
 *     }],
 *     dnsPrefix: "dnsprefix1",
 *     kubernetesVersion: "1.7.7",
 *     location: "location1",
 *     resourceGroupName: "rg1",
 *     resourceName: "clustername1",
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
    public static readonly __pulumiType = 'azurerm:containerservice/v20170831:ManagedCluster';

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
     * Properties of the agent pool.
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.v20170831.ContainerServiceAgentPoolProfileResponse[] | undefined>;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    public readonly dnsPrefix!: pulumi.Output<string | undefined>;
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
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.v20170831.ContainerServiceLinuxProfileResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
     */
    public readonly servicePrincipalProfile!: pulumi.Output<outputs.containerservice.v20170831.ContainerServiceServicePrincipalProfileResponse | undefined>;
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
            inputs["agentPoolProfiles"] = args ? args.agentPoolProfiles : undefined;
            inputs["dnsPrefix"] = args ? args.dnsPrefix : undefined;
            inputs["kubernetesVersion"] = args ? args.kubernetesVersion : undefined;
            inputs["linuxProfile"] = args ? args.linuxProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["servicePrincipalProfile"] = args ? args.servicePrincipalProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["fqdn"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:containerservice/latest:ManagedCluster" }, { type: "azurerm:containerservice/v20180331:ManagedCluster" }, { type: "azurerm:containerservice/v20180801preview:ManagedCluster" }, { type: "azurerm:containerservice/v20190201:ManagedCluster" }, { type: "azurerm:containerservice/v20190401:ManagedCluster" }, { type: "azurerm:containerservice/v20190601:ManagedCluster" }, { type: "azurerm:containerservice/v20190801:ManagedCluster" }, { type: "azurerm:containerservice/v20191001:ManagedCluster" }, { type: "azurerm:containerservice/v20191101:ManagedCluster" }, { type: "azurerm:containerservice/v20200101:ManagedCluster" }, { type: "azurerm:containerservice/v20200201:ManagedCluster" }, { type: "azurerm:containerservice/v20200301:ManagedCluster" }, { type: "azurerm:containerservice/v20200401:ManagedCluster" }, { type: "azurerm:containerservice/v20200601:ManagedCluster" }, { type: "azurerm:containerservice/v20200701:ManagedCluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagedCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedCluster resource.
 */
export interface ManagedClusterArgs {
    /**
     * Properties of the agent pool.
     */
    readonly agentPoolProfiles?: pulumi.Input<pulumi.Input<inputs.containerservice.v20170831.ContainerServiceAgentPoolProfile>[]>;
    /**
     * DNS prefix specified when creating the managed cluster.
     */
    readonly dnsPrefix?: pulumi.Input<string>;
    /**
     * Version of Kubernetes specified when creating the managed cluster.
     */
    readonly kubernetesVersion?: pulumi.Input<string>;
    /**
     * Profile for Linux VMs in the container service cluster.
     */
    readonly linuxProfile?: pulumi.Input<inputs.containerservice.v20170831.ContainerServiceLinuxProfile>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the managed cluster resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
     */
    readonly servicePrincipalProfile?: pulumi.Input<inputs.containerservice.v20170831.ContainerServiceServicePrincipalProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
