// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The essential information related to the peer's ASN.
 */
export class PeerAsn extends pulumi.CustomResource {
    /**
     * Get an existing PeerAsn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PeerAsn {
        return new PeerAsn(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:peering/v20200401:PeerAsn';

    /**
     * Returns true if the given object is an instance of PeerAsn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PeerAsn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PeerAsn.__pulumiType;
    }

    /**
     * The error message for the validation state
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Autonomous System Number (ASN) of the peer.
     */
    public readonly peerAsn!: pulumi.Output<number | undefined>;
    /**
     * The contact details of the peer.
     */
    public readonly peerContactDetail!: pulumi.Output<outputs.peering.v20200401.ContactDetailResponse[] | undefined>;
    /**
     * The name of the peer.
     */
    public readonly peerName!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The validation state of the ASN associated with the peer.
     */
    public readonly validationState!: pulumi.Output<string | undefined>;

    /**
     * Create a PeerAsn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PeerAsnArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.peerAsnName === undefined) {
                throw new Error("Missing required property 'peerAsnName'");
            }
            inputs["peerAsn"] = args ? args.peerAsn : undefined;
            inputs["peerAsnName"] = args ? args.peerAsnName : undefined;
            inputs["peerContactDetail"] = args ? args.peerContactDetail : undefined;
            inputs["peerName"] = args ? args.peerName : undefined;
            inputs["validationState"] = args ? args.validationState : undefined;
            inputs["errorMessage"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["errorMessage"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peerAsn"] = undefined /*out*/;
            inputs["peerContactDetail"] = undefined /*out*/;
            inputs["peerName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["validationState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:peering/latest:PeerAsn" }, { type: "azurerm:peering/v20190801preview:PeerAsn" }, { type: "azurerm:peering/v20190901preview:PeerAsn" }, { type: "azurerm:peering/v20200101preview:PeerAsn" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PeerAsn.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PeerAsn resource.
 */
export interface PeerAsnArgs {
    /**
     * The Autonomous System Number (ASN) of the peer.
     */
    readonly peerAsn?: pulumi.Input<number>;
    /**
     * The peer ASN name.
     */
    readonly peerAsnName: pulumi.Input<string>;
    /**
     * The contact details of the peer.
     */
    readonly peerContactDetail?: pulumi.Input<pulumi.Input<inputs.peering.v20200401.ContactDetail>[]>;
    /**
     * The name of the peer.
     */
    readonly peerName?: pulumi.Input<string>;
    /**
     * The validation state of the ASN associated with the peer.
     */
    readonly validationState?: pulumi.Input<string>;
}
