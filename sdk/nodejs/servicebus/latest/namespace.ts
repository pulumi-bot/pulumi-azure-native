// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Description of a namespace resource.
 *
 * ## Example Usage
 * ### NameSpaceCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const namespace = new azurerm.servicebus.latest.Namespace("namespace", {
 *     location: "South Central US",
 *     namespaceName: "sdk-Namespace2924",
 *     resourceGroupName: "ArunMonocle",
 *     sku: {
 *         name: "Standard",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         tag1: "value1",
 *         tag2: "value2",
 *     },
 * });
 *
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicebus/latest:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * The time the namespace was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The Geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Identifier for Azure Insights metrics
     */
    public /*out*/ readonly metricId!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the namespace.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Endpoint you can use to perform Service Bus operations.
     */
    public /*out*/ readonly serviceBusEndpoint!: pulumi.Output<string>;
    /**
     * Properties of Sku
     */
    public readonly sku!: pulumi.Output<outputs.servicebus.latest.SBSkuResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The time the namespace was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["metricId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serviceBusEndpoint"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        } else {
            inputs["createdAt"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["metricId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serviceBusEndpoint"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:servicebus/v20140901:Namespace" }, { type: "azurerm:servicebus/v20150801:Namespace" }, { type: "azurerm:servicebus/v20170401:Namespace" }, { type: "azurerm:servicebus/v20180101preview:Namespace" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The Geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The namespace name.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Properties of Sku
     */
    readonly sku?: pulumi.Input<inputs.servicebus.latest.SBSku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
