// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The HDInsight cluster application
 */
export class ClusterApplication extends pulumi.CustomResource {
    /**
     * Get an existing ClusterApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClusterApplication {
        return new ClusterApplication(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hdinsight:ClusterApplication';

    /**
     * Returns true if the given object is an instance of ClusterApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterApplication.__pulumiType;
    }

    /**
     * The ETag for the application
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the application.
     */
    public readonly properties!: pulumi.Output<outputs.hdinsight.ApplicationPropertiesResponse>;
    /**
     * The tags for the application.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ClusterApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["clusterName"] = args ? args.clusterName : undefined;
        inputs["etag"] = args ? args.etag : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterApplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClusterApplication resource.
 */
export interface ClusterApplicationArgs {
    /**
     * The name of the cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The ETag for the application
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The constant value for the application name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The properties of the application.
     */
    readonly properties?: pulumi.Input<inputs.hdinsight.ApplicationProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tags for the application.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
