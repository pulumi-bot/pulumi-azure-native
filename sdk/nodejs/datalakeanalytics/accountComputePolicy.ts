// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Data Lake Analytics compute policy information.
 */
export class AccountComputePolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccountComputePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccountComputePolicy {
        return new AccountComputePolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datalakeanalytics:AccountComputePolicy';

    /**
     * Returns true if the given object is an instance of AccountComputePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountComputePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountComputePolicy.__pulumiType;
    }

    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The compute policy properties.
     */
    public readonly properties!: pulumi.Output<outputs.datalakeanalytics.ComputePolicyPropertiesResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AccountComputePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountComputePolicyArgs, opts?: pulumi.CustomResourceOptions) {
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
        super(AccountComputePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccountComputePolicy resource.
 */
export interface AccountComputePolicyArgs {
    /**
     * The name of the Data Lake Analytics account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the compute policy to create or update.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The compute policy properties to use when creating a new compute policy.
     */
    readonly properties: pulumi.Input<inputs.datalakeanalytics.CreateOrUpdateComputePolicyProperties>;
    /**
     * The name of the Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
