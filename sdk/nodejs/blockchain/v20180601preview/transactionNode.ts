// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Payload of the transaction node which is the request/response of the resource provider.
 */
export class TransactionNode extends pulumi.CustomResource {
    /**
     * Get an existing TransactionNode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransactionNode {
        return new TransactionNode(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:blockchain/v20180601preview:TransactionNode';

    /**
     * Returns true if the given object is an instance of TransactionNode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransactionNode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransactionNode.__pulumiType;
    }

    /**
     * Gets or sets the transaction node dns endpoint.
     */
    public /*out*/ readonly dns!: pulumi.Output<string>;
    /**
     * Gets or sets the firewall rules.
     */
    public readonly firewallRules!: pulumi.Output<outputs.blockchain.v20180601preview.FirewallRuleResponse[] | undefined>;
    /**
     * Gets or sets the transaction node location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Sets the transaction node dns endpoint basic auth password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the blockchain member provision state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Gets or sets the transaction node public key.
     */
    public /*out*/ readonly publicKey!: pulumi.Output<string>;
    /**
     * The type of the service - e.g. "Microsoft.Blockchain"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Gets or sets the transaction node dns endpoint basic auth user name.
     */
    public /*out*/ readonly userName!: pulumi.Output<string>;

    /**
     * Create a TransactionNode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransactionNodeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.blockchainMemberName === undefined) {
                throw new Error("Missing required property 'blockchainMemberName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.transactionNodeName === undefined) {
                throw new Error("Missing required property 'transactionNodeName'");
            }
            inputs["blockchainMemberName"] = args ? args.blockchainMemberName : undefined;
            inputs["firewallRules"] = args ? args.firewallRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["transactionNodeName"] = args ? args.transactionNodeName : undefined;
            inputs["dns"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicKey"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userName"] = undefined /*out*/;
        } else {
            inputs["dns"] = undefined /*out*/;
            inputs["firewallRules"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicKey"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TransactionNode.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransactionNode resource.
 */
export interface TransactionNodeArgs {
    /**
     * Blockchain member name.
     */
    readonly blockchainMemberName: pulumi.Input<string>;
    /**
     * Gets or sets the firewall rules.
     */
    readonly firewallRules?: pulumi.Input<pulumi.Input<inputs.blockchain.v20180601preview.FirewallRule>[]>;
    /**
     * Gets or sets the transaction node location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Sets the transaction node dns endpoint basic auth password.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Transaction node name.
     */
    readonly transactionNodeName: pulumi.Input<string>;
}
