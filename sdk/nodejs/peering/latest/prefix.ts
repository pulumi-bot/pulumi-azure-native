// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The peering service prefix class.
 *
 * ## Example Usage
 * ### Create or update a prefix for the peering service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const prefix = new azurerm.peering.latest.Prefix("prefix", {
 *     peeringServiceName: "peeringServiceName",
 *     peeringServicePrefixKey: "00000000-0000-0000-0000-000000000000",
 *     prefix: "192.168.1.0/24",
 *     prefixName: "peeringServicePrefixName",
 *     resourceGroupName: "rgName",
 * });
 *
 * ```
 */
export class Prefix extends pulumi.CustomResource {
    /**
     * Get an existing Prefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Prefix {
        return new Prefix(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:peering/latest:Prefix';

    /**
     * Returns true if the given object is an instance of Prefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Prefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Prefix.__pulumiType;
    }

    /**
     * The error message for validation state
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * The list of events for peering service prefix
     */
    public /*out*/ readonly events!: pulumi.Output<outputs.peering.latest.PeeringServicePrefixEventResponse[]>;
    /**
     * The prefix learned type
     */
    public /*out*/ readonly learnedType!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The peering service prefix key
     */
    public readonly peeringServicePrefixKey!: pulumi.Output<string | undefined>;
    /**
     * The prefix from which your traffic originates.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;
    /**
     * The prefix validation state
     */
    public /*out*/ readonly prefixValidationState!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Prefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrefixArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrefixArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PrefixArgs | undefined;
            if (!args || args.peeringServiceName === undefined) {
                throw new Error("Missing required property 'peeringServiceName'");
            }
            if (!args || args.prefixName === undefined) {
                throw new Error("Missing required property 'prefixName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["peeringServiceName"] = args ? args.peeringServiceName : undefined;
            inputs["peeringServicePrefixKey"] = args ? args.peeringServicePrefixKey : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["prefixName"] = args ? args.prefixName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["errorMessage"] = undefined /*out*/;
            inputs["events"] = undefined /*out*/;
            inputs["learnedType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["prefixValidationState"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:peering/v20200401:Prefix" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Prefix.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Prefix resource.
 */
export interface PrefixArgs {
    /**
     * The name of the peering service.
     */
    readonly peeringServiceName: pulumi.Input<string>;
    /**
     * The peering service prefix key
     */
    readonly peeringServicePrefixKey?: pulumi.Input<string>;
    /**
     * The prefix from which your traffic originates.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * The name of the prefix.
     */
    readonly prefixName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
