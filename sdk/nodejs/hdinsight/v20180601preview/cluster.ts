// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The HDInsight cluster.
 *
 * ## Example Usage
 * ### Create HDInsight cluster with Autoscale configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 * ### Create Hadoop cluster with Azure Data Lake Storage Gen 2
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "val1",
 *     },
 * });
 *
 * ```
 * ### Create Hadoop on Linux cluster with SSH password
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "val1",
 *     },
 * });
 *
 * ```
 * ### Create Hadoop on Linux cluster with SSH public key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "val1",
 *     },
 * });
 *
 * ```
 * ### Create Kafka cluster with Kafka Rest Proxy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 * ### Create Secure Hadoop cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "val1",
 *     },
 * });
 *
 * ```
 * ### Create Spark on Linux Cluster with SSH password
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "val1",
 *     },
 * });
 *
 * ```
 * ### Create cluster with TLS 1.2
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 * ### Create cluster with custom network settings
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 * ### Create cluster with encryption at host
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 * ### Create cluster with encryption in transit
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cluster = new azurerm.hdinsight.v20180601preview.Cluster("cluster", {
 *     clusterName: "cluster1",
 *     resourceGroupName: "rg1",
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
    public static readonly __pulumiType = 'azurerm:hdinsight/v20180601preview:Cluster';

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
     * The ETag for the resource
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The identity of the cluster, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.hdinsight.v20180601preview.ClusterIdentityResponse | undefined>;
    /**
     * The Azure Region where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The properties of the cluster.
     */
    public readonly properties!: pulumi.Output<outputs.hdinsight.v20180601preview.ClusterGetPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

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
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:hdinsight/latest:Cluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The name of the cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The identity of the cluster, if configured.
     */
    readonly identity?: pulumi.Input<inputs.hdinsight.v20180601preview.ClusterIdentity>;
    /**
     * The location of the cluster.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The cluster create parameters.
     */
    readonly properties?: pulumi.Input<inputs.hdinsight.v20180601preview.ClusterCreateProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
