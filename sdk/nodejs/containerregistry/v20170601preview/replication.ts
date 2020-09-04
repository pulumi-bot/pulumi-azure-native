// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An object that represents a replication for a container registry.
 */
export class Replication extends pulumi.CustomResource {
    /**
     * Get an existing Replication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Replication {
        return new Replication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerregistry/v20170601preview:Replication';

    /**
     * Returns true if the given object is an instance of Replication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Replication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Replication.__pulumiType;
    }

    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the replication at the time the operation was called.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The status of the replication at the time the operation was called.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.containerregistry.v20170601preview.StatusResponse>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Replication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.replicationName === undefined) {
                throw new Error("Missing required property 'replicationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["replicationName"] = args ? args.replicationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerregistry/latest:Replication" }, { type: "azurerm:containerregistry/v20171001:Replication" }, { type: "azurerm:containerregistry/v20190501:Replication" }, { type: "azurerm:containerregistry/v20191201preview:Replication" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Replication.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Replication resource.
 */
export interface ReplicationArgs {
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the replication.
     */
    readonly replicationName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
