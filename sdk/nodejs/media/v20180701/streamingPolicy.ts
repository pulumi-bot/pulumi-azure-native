// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A Streaming Policy resource
 */
export class StreamingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing StreamingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StreamingPolicy {
        return new StreamingPolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/v20180701:StreamingPolicy';

    /**
     * Returns true if the given object is an instance of StreamingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamingPolicy.__pulumiType;
    }

    /**
     * Configuration of CommonEncryptionCbcs
     */
    public readonly commonEncryptionCbcs!: pulumi.Output<outputs.media.v20180701.CommonEncryptionCbcsResponse | undefined>;
    /**
     * Configuration of CommonEncryptionCenc
     */
    public readonly commonEncryptionCenc!: pulumi.Output<outputs.media.v20180701.CommonEncryptionCencResponse | undefined>;
    /**
     * Creation time of Streaming Policy
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Default ContentKey used by current Streaming Policy
     */
    public readonly defaultContentKeyPolicyName!: pulumi.Output<string | undefined>;
    /**
     * Configuration of EnvelopeEncryption
     */
    public readonly envelopeEncryption!: pulumi.Output<outputs.media.v20180701.EnvelopeEncryptionResponse | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Configurations of NoEncryption
     */
    public readonly noEncryption!: pulumi.Output<outputs.media.v20180701.NoEncryptionResponse | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StreamingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamingPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as StreamingPolicyArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.streamingPolicyName === undefined) {
                throw new Error("Missing required property 'streamingPolicyName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["commonEncryptionCbcs"] = args ? args.commonEncryptionCbcs : undefined;
            inputs["commonEncryptionCenc"] = args ? args.commonEncryptionCenc : undefined;
            inputs["defaultContentKeyPolicyName"] = args ? args.defaultContentKeyPolicyName : undefined;
            inputs["envelopeEncryption"] = args ? args.envelopeEncryption : undefined;
            inputs["noEncryption"] = args ? args.noEncryption : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["streamingPolicyName"] = args ? args.streamingPolicyName : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:media/latest:StreamingPolicy" }, { type: "azurerm:media/v20180330preview:StreamingPolicy" }, { type: "azurerm:media/v20180601preview:StreamingPolicy" }, { type: "azurerm:media/v20200501:StreamingPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(StreamingPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StreamingPolicy resource.
 */
export interface StreamingPolicyArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Configuration of CommonEncryptionCbcs
     */
    readonly commonEncryptionCbcs?: pulumi.Input<inputs.media.v20180701.CommonEncryptionCbcs>;
    /**
     * Configuration of CommonEncryptionCenc
     */
    readonly commonEncryptionCenc?: pulumi.Input<inputs.media.v20180701.CommonEncryptionCenc>;
    /**
     * Default ContentKey used by current Streaming Policy
     */
    readonly defaultContentKeyPolicyName?: pulumi.Input<string>;
    /**
     * Configuration of EnvelopeEncryption
     */
    readonly envelopeEncryption?: pulumi.Input<inputs.media.v20180701.EnvelopeEncryption>;
    /**
     * Configurations of NoEncryption
     */
    readonly noEncryption?: pulumi.Input<inputs.media.v20180701.NoEncryption>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Streaming Policy name.
     */
    readonly streamingPolicyName: pulumi.Input<string>;
}
