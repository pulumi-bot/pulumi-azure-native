// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Response on GET of a hybrid use benefit
 */
export class HybridUseBenefit extends pulumi.CustomResource {
    /**
     * Get an existing HybridUseBenefit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HybridUseBenefit {
        return new HybridUseBenefit(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:softwareplan/v20191201:HybridUseBenefit';

    /**
     * Returns true if the given object is an instance of HybridUseBenefit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HybridUseBenefit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HybridUseBenefit.__pulumiType;
    }

    /**
     * Indicates the revision of the hybrid use benefit
     */
    public /*out*/ readonly etag!: pulumi.Output<number>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Property bag for a hybrid use benefit response
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.softwareplan.v20191201.HybridUseBenefitPropertiesResponse>;
    /**
     * Hybrid use benefit SKU
     */
    public readonly sku!: pulumi.Output<outputs.softwareplan.v20191201.SkuResponse>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a HybridUseBenefit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HybridUseBenefitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HybridUseBenefitArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as HybridUseBenefitArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HybridUseBenefit.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HybridUseBenefit resource.
 */
export interface HybridUseBenefitArgs {
    /**
     * This is a unique identifier for a plan. Should be a guid.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
     */
    readonly scope: pulumi.Input<string>;
    /**
     * Hybrid use benefit SKU
     */
    readonly sku: pulumi.Input<inputs.softwareplan.v20191201.Sku>;
}
