// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Policy Contract details.
 */
export class ProductPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ProductPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProductPolicy {
        return new ProductPolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20170301:ProductPolicy';

    /**
     * Returns true if the given object is an instance of ProductPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductPolicy.__pulumiType;
    }

    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Json escaped Xml Encoded contents of the Policy.
     */
    public readonly policyContent!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProductPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ProductPolicyArgs | undefined;
            if (!args || args.policyContent === undefined) {
                throw new Error("Missing required property 'policyContent'");
            }
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            if (!args || args.productId === undefined) {
                throw new Error("Missing required property 'productId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["policyContent"] = args ? args.policyContent : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ProductPolicy" }, { type: "azurerm:apimanagement/v20180101:ProductPolicy" }, { type: "azurerm:apimanagement/v20180601preview:ProductPolicy" }, { type: "azurerm:apimanagement/v20190101:ProductPolicy" }, { type: "azurerm:apimanagement/v20191201:ProductPolicy" }, { type: "azurerm:apimanagement/v20191201preview:ProductPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ProductPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProductPolicy resource.
 */
export interface ProductPolicyArgs {
    /**
     * Json escaped Xml Encoded contents of the Policy.
     */
    readonly policyContent: pulumi.Input<string>;
    /**
     * The identifier of the Policy.
     */
    readonly policyId: pulumi.Input<string>;
    /**
     * Product identifier. Must be unique in the current API Management service instance.
     */
    readonly productId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
