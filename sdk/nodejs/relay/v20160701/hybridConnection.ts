// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Description of HybridConnection Resource.
 */
export class HybridConnection extends pulumi.CustomResource {
    /**
     * Get an existing HybridConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HybridConnection {
        return new HybridConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:relay/v20160701:HybridConnection';

    /**
     * Returns true if the given object is an instance of HybridConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HybridConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HybridConnection.__pulumiType;
    }

    /**
     * The time the HybridConnection was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The number of listeners for this HybridConnection. min : 1 and max:25 supported
     */
    public /*out*/ readonly listenerCount!: pulumi.Output<number>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * true if client authorization is needed for this HybridConnection; otherwise, false.
     */
    public readonly requiresClientAuthorization!: pulumi.Output<boolean | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The time the namespace was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
     */
    public readonly userMetadata!: pulumi.Output<string | undefined>;

    /**
     * Create a HybridConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HybridConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hybridConnectionName === undefined) {
                throw new Error("Missing required property 'hybridConnectionName'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hybridConnectionName"] = args ? args.hybridConnectionName : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["requiresClientAuthorization"] = args ? args.requiresClientAuthorization : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["userMetadata"] = args ? args.userMetadata : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["listenerCount"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        } else {
            inputs["createdAt"] = undefined /*out*/;
            inputs["listenerCount"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["requiresClientAuthorization"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
            inputs["userMetadata"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:relay/latest:HybridConnection" }, { type: "azurerm:relay/v20170401:HybridConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(HybridConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HybridConnection resource.
 */
export interface HybridConnectionArgs {
    /**
     * The hybrid connection name.
     */
    readonly hybridConnectionName: pulumi.Input<string>;
    /**
     * The Namespace Name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * true if client authorization is needed for this HybridConnection; otherwise, false.
     */
    readonly requiresClientAuthorization?: pulumi.Input<boolean>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
     */
    readonly userMetadata?: pulumi.Input<string>;
}
