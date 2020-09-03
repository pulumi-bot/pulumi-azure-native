// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The customer's ASN that is registered by the peering service provider.
 *
 * ## Example Usage
 * ### Create or update a registered ASN for the peering
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const registeredAsn = new azurerm.peering.latest.RegisteredAsn("registeredAsn", {
 *     asn: 65000,
 *     peeringName: "peeringName",
 *     registeredAsnName: "registeredAsnName",
 *     resourceGroupName: "rgName",
 * });
 *
 * ```
 */
export class RegisteredAsn extends pulumi.CustomResource {
    /**
     * Get an existing RegisteredAsn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegisteredAsn {
        return new RegisteredAsn(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:peering/latest:RegisteredAsn';

    /**
     * Returns true if the given object is an instance of RegisteredAsn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegisteredAsn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegisteredAsn.__pulumiType;
    }

    /**
     * The customer's ASN from which traffic originates.
     */
    public readonly asn!: pulumi.Output<number | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The peering service prefix key that is to be shared with the customer.
     */
    public /*out*/ readonly peeringServicePrefixKey!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RegisteredAsn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegisteredAsnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegisteredAsnArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RegisteredAsnArgs | undefined;
            if (!args || args.peeringName === undefined) {
                throw new Error("Missing required property 'peeringName'");
            }
            if (!args || args.registeredAsnName === undefined) {
                throw new Error("Missing required property 'registeredAsnName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["asn"] = args ? args.asn : undefined;
            inputs["peeringName"] = args ? args.peeringName : undefined;
            inputs["registeredAsnName"] = args ? args.registeredAsnName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["peeringServicePrefixKey"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:peering/v20200101preview:RegisteredAsn" }, { type: "azurerm:peering/v20200401:RegisteredAsn" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RegisteredAsn.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegisteredAsn resource.
 */
export interface RegisteredAsnArgs {
    /**
     * The customer's ASN from which traffic originates.
     */
    readonly asn?: pulumi.Input<number>;
    /**
     * The name of the peering.
     */
    readonly peeringName: pulumi.Input<string>;
    /**
     * The name of the ASN.
     */
    readonly registeredAsnName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
