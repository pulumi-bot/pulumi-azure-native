// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The customer's prefix that is registered by the peering service provider.
 */
export class PeeringRegisteredPrefix extends pulumi.CustomResource {
    /**
     * Get an existing PeeringRegisteredPrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PeeringRegisteredPrefix {
        return new PeeringRegisteredPrefix(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:peering:PeeringRegisteredPrefix';

    /**
     * Returns true if the given object is an instance of PeeringRegisteredPrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PeeringRegisteredPrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PeeringRegisteredPrefix.__pulumiType;
    }

    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties that define a registered prefix.
     */
    public readonly properties!: pulumi.Output<outputs.peering.PeeringRegisteredPrefixPropertiesResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PeeringRegisteredPrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PeeringRegisteredPrefixArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.peeringName === undefined) {
                throw new Error("Missing required property 'peeringName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["name"] = args ? args.name : undefined;
        inputs["peeringName"] = args ? args.peeringName : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PeeringRegisteredPrefix.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PeeringRegisteredPrefix resource.
 */
export interface PeeringRegisteredPrefixArgs {
    /**
     * The name of the registered prefix.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the peering.
     */
    readonly peeringName: pulumi.Input<string>;
    /**
     * The properties that define a registered prefix.
     */
    readonly properties?: pulumi.Input<inputs.peering.PeeringRegisteredPrefixProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
