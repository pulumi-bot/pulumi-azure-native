// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Class representing a Kusto cluster.
 *
 * ## Example Usage
 * ### KustoClustersCreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.kusto.v20190121.Cluster("cluster", {
 *     clusterName: "KustoClusterRPTest4",
 *     location: "westus",
 *     resourceGroupName: "kustorptest",
 *     sku: {
 *         capacity: 2,
 *         name: "L8",
 *         tier: "Standard",
 *     },
 * });
 *
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:kusto/v20190121:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The cluster data ingestion URI.
     */
    public /*out*/ readonly dataIngestionUri!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioned state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The SKU of the cluster.
     */
    public readonly sku!: pulumi.Output<outputs.kusto.v20190121.AzureSkuResponse>;
    /**
     * The state of the resource.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The cluster's external tenants.
     */
    public readonly trustedExternalTenants!: pulumi.Output<outputs.kusto.v20190121.TrustedExternalTenantResponse[] | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The cluster URI.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trustedExternalTenants"] = args ? args.trustedExternalTenants : undefined;
            inputs["dataIngestionUri"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:kusto/latest:Cluster" }, { type: "azurerm:kusto/v20190515:Cluster" }, { type: "azurerm:kusto/v20190907:Cluster" }, { type: "azurerm:kusto/v20191109:Cluster" }, { type: "azurerm:kusto/v20200215:Cluster" }, { type: "azurerm:kusto/v20200614:Cluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The name of the Kusto cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group containing the Kusto cluster.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the cluster.
     */
    readonly sku: pulumi.Input<inputs.kusto.v20190121.AzureSku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The cluster's external tenants.
     */
    readonly trustedExternalTenants?: pulumi.Input<pulumi.Input<inputs.kusto.v20190121.TrustedExternalTenant>[]>;
}
