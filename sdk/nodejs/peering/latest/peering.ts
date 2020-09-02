// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
 *
 * ## Example Usage
 * ### Create a direct peering
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const peering = new azurerm.peering.latest.Peering("peering", {
 *     direct: {
 *         connections: [
 *             {
 *                 bandwidthInMbps: 10000,
 *                 bgpSession: {
 *                     maxPrefixesAdvertisedV4: 1000,
 *                     maxPrefixesAdvertisedV6: 100,
 *                     md5AuthenticationKey: "test-md5-auth-key",
 *                     sessionPrefixV4: "192.168.0.0/31",
 *                     sessionPrefixV6: "fd00::0/127",
 *                 },
 *                 connectionIdentifier: "5F4CB5C7-6B43-4444-9338-9ABC72606C16",
 *                 peeringDBFacilityId: 99999,
 *                 sessionAddressProvider: "Peer",
 *                 useForPeeringService: false,
 *             },
 *             {
 *                 bandwidthInMbps: 10000,
 *                 connectionIdentifier: "8AB00818-D533-4504-A25A-03A17F61201C",
 *                 peeringDBFacilityId: 99999,
 *                 sessionAddressProvider: "Microsoft",
 *                 useForPeeringService: true,
 *             },
 *         ],
 *         directPeeringType: "Edge",
 *         peerAsn: {
 *             id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1",
 *         },
 *     },
 *     kind: "Direct",
 *     location: "eastus",
 *     peeringLocation: "peeringLocation0",
 *     peeringName: "peeringName",
 *     resourceGroupName: "rgName",
 *     sku: {
 *         name: "Basic_Direct_Free",
 *     },
 * });
 *
 * ```
 * ### Create a peering with exchange route server
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const peering = new azurerm.peering.latest.Peering("peering", {
 *     direct: {
 *         connections: [{
 *             bandwidthInMbps: 10000,
 *             bgpSession: {
 *                 maxPrefixesAdvertisedV4: 1000,
 *                 maxPrefixesAdvertisedV6: 100,
 *                 microsoftSessionIPv4Address: "192.168.0.123",
 *                 peerSessionIPv4Address: "192.168.0.234",
 *                 sessionPrefixV4: "192.168.0.0/24",
 *             },
 *             connectionIdentifier: "5F4CB5C7-6B43-4444-9338-9ABC72606C16",
 *             peeringDBFacilityId: 99999,
 *             sessionAddressProvider: "Peer",
 *             useForPeeringService: true,
 *         }],
 *         directPeeringType: "IxRs",
 *         peerAsn: {
 *             id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1",
 *         },
 *     },
 *     kind: "Direct",
 *     location: "eastus",
 *     peeringLocation: "peeringLocation0",
 *     peeringName: "peeringName",
 *     resourceGroupName: "rgName",
 *     sku: {
 *         name: "Premium_Direct_Free",
 *     },
 * });
 *
 * ```
 * ### Create an exchange peering
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const peering = new azurerm.peering.latest.Peering("peering", {
 *     exchange: {
 *         connections: [
 *             {
 *                 bgpSession: {
 *                     maxPrefixesAdvertisedV4: 1000,
 *                     maxPrefixesAdvertisedV6: 100,
 *                     md5AuthenticationKey: "test-md5-auth-key",
 *                     peerSessionIPv4Address: "192.168.2.1",
 *                     peerSessionIPv6Address: "fd00::1",
 *                 },
 *                 connectionIdentifier: "CE495334-0E94-4E51-8164-8116D6CD284D",
 *                 peeringDBFacilityId: 99999,
 *             },
 *             {
 *                 bgpSession: {
 *                     maxPrefixesAdvertisedV4: 1000,
 *                     maxPrefixesAdvertisedV6: 100,
 *                     md5AuthenticationKey: "test-md5-auth-key",
 *                     peerSessionIPv4Address: "192.168.2.2",
 *                     peerSessionIPv6Address: "fd00::2",
 *                 },
 *                 connectionIdentifier: "CDD8E673-CB07-47E6-84DE-3739F778762B",
 *                 peeringDBFacilityId: 99999,
 *             },
 *         ],
 *         peerAsn: {
 *             id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1",
 *         },
 *     },
 *     kind: "Exchange",
 *     location: "eastus",
 *     peeringLocation: "peeringLocation0",
 *     peeringName: "peeringName",
 *     resourceGroupName: "rgName",
 *     sku: {
 *         name: "Basic_Exchange_Free",
 *     },
 * });
 *
 * ```
 */
export class Peering extends pulumi.CustomResource {
    /**
     * Get an existing Peering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Peering {
        return new Peering(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:peering/latest:Peering';

    /**
     * Returns true if the given object is an instance of Peering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Peering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Peering.__pulumiType;
    }

    /**
     * The properties that define a direct peering.
     */
    public readonly direct!: pulumi.Output<outputs.peering.latest.PeeringPropertiesDirectResponse | undefined>;
    /**
     * The properties that define an exchange peering.
     */
    public readonly exchange!: pulumi.Output<outputs.peering.latest.PeeringPropertiesExchangeResponse | undefined>;
    /**
     * The kind of the peering.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The location of the peering.
     */
    public readonly peeringLocation!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The SKU that defines the tier and kind of the peering.
     */
    public readonly sku!: pulumi.Output<outputs.peering.latest.PeeringSkuResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Peering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PeeringArgs | undefined;
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.peeringName === undefined) {
                throw new Error("Missing required property 'peeringName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["direct"] = args ? args.direct : undefined;
            inputs["exchange"] = args ? args.exchange : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["peeringLocation"] = args ? args.peeringLocation : undefined;
            inputs["peeringName"] = args ? args.peeringName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:peering/v20200401:Peering" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Peering.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Peering resource.
 */
export interface PeeringArgs {
    /**
     * The properties that define a direct peering.
     */
    readonly direct?: pulumi.Input<inputs.peering.latest.PeeringPropertiesDirect>;
    /**
     * The properties that define an exchange peering.
     */
    readonly exchange?: pulumi.Input<inputs.peering.latest.PeeringPropertiesExchange>;
    /**
     * The kind of the peering.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The location of the peering.
     */
    readonly peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the peering.
     */
    readonly peeringName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU that defines the tier and kind of the peering.
     */
    readonly sku: pulumi.Input<inputs.peering.latest.PeeringSku>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
