// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Data Lake Store trusted identity provider information.
 */
export class TrustedIdProvider extends pulumi.CustomResource {
    /**
     * Get an existing TrustedIdProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TrustedIdProvider {
        return new TrustedIdProvider(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datalakestore/v20161101:TrustedIdProvider';

    /**
     * Returns true if the given object is an instance of TrustedIdProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrustedIdProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrustedIdProvider.__pulumiType;
    }

    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The trusted identity provider properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.datalakestore.v20161101.TrustedIdProviderPropertiesResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a TrustedIdProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrustedIdProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrustedIdProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as TrustedIdProviderArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.idProvider === undefined) {
                throw new Error("Missing required property 'idProvider'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["idProvider"] = args ? args.idProvider : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TrustedIdProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TrustedIdProvider resource.
 */
export interface TrustedIdProviderArgs {
    /**
     * The name of the Data Lake Store account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The URL of this trusted identity provider.
     */
    readonly idProvider: pulumi.Input<string>;
    /**
     * The name of the trusted identity provider. This is used for differentiation of providers in the account.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
