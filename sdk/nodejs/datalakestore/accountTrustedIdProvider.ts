// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Data Lake Store trusted identity provider information.
 */
export class AccountTrustedIdProvider extends pulumi.CustomResource {
    /**
     * Get an existing AccountTrustedIdProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccountTrustedIdProvider {
        return new AccountTrustedIdProvider(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datalakestore:AccountTrustedIdProvider';

    /**
     * Returns true if the given object is an instance of AccountTrustedIdProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountTrustedIdProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountTrustedIdProvider.__pulumiType;
    }

    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The trusted identity provider properties.
     */
    public readonly properties!: pulumi.Output<outputs.datalakestore.TrustedIdProviderPropertiesResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AccountTrustedIdProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountTrustedIdProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["accountName"] = args ? args.accountName : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccountTrustedIdProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccountTrustedIdProvider resource.
 */
export interface AccountTrustedIdProviderArgs {
    /**
     * The name of the Data Lake Store account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the trusted identity provider. This is used for differentiation of providers in the account.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The trusted identity provider properties to use when creating a new trusted identity provider.
     */
    readonly properties: pulumi.Input<inputs.datalakestore.CreateOrUpdateTrustedIdProviderProperties>;
    /**
     * The name of the Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
