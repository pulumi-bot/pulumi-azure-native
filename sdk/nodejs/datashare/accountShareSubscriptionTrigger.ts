// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Trigger data transfer object.
 */
export class AccountShareSubscriptionTrigger extends pulumi.CustomResource {
    /**
     * Get an existing AccountShareSubscriptionTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccountShareSubscriptionTrigger {
        return new AccountShareSubscriptionTrigger(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datashare:AccountShareSubscriptionTrigger';

    /**
     * Returns true if the given object is an instance of AccountShareSubscriptionTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountShareSubscriptionTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountShareSubscriptionTrigger.__pulumiType;
    }

    /**
     * Kind of synchronization
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Name of the azure resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Type of the azure resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AccountShareSubscriptionTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountShareSubscriptionTriggerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.shareSubscriptionName === undefined) {
                throw new Error("Missing required property 'shareSubscriptionName'");
            }
        inputs["accountName"] = args ? args.accountName : undefined;
        inputs["kind"] = args ? args.kind : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["shareSubscriptionName"] = args ? args.shareSubscriptionName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccountShareSubscriptionTrigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccountShareSubscriptionTrigger resource.
 */
export interface AccountShareSubscriptionTriggerArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Kind of synchronization
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The name of the trigger.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the share subscription which will hold the data set sink.
     */
    readonly shareSubscriptionName: pulumi.Input<string>;
}