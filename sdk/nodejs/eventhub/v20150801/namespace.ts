// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Single Namespace item in List or Get Operation
 *
 * ## Example Usage
 * ### NamespaceCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const namespace = new azurerm.eventhub.v20150801.Namespace("namespace", {
 *     location: "South Central US",
 *     namespaceName: "sdk-Namespace8107",
 *     resourceGroupName: "Default-ServiceBus-WestUS",
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
        return new Namespace(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventhub/v20150801:Namespace';

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
     * The time the Namespace was created.
     */
    public readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether this instance is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Resource location
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
     * Provisioning state of the Namespace.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Endpoint you can use to perform Service Bus operations.
     */
    public readonly serviceBusEndpoint!: pulumi.Output<string | undefined>;
    /**
     * SKU parameters supplied to the create Namespace operation
     */
    public readonly sku!: pulumi.Output<outputs.eventhub.v20150801.SkuResponse | undefined>;
    /**
     * State of the Namespace.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The time the Namespace was updated.
     */
    public readonly updatedAt!: pulumi.Output<string | undefined>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as NamespaceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["createdAt"] = args ? args.createdAt : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceBusEndpoint"] = args ? args.serviceBusEndpoint : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["updatedAt"] = args ? args.updatedAt : undefined;
            inputs["metricId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:eventhub/latest:Namespace" }, { type: "azurerm:eventhub/v20140901:Namespace" }, { type: "azurerm:eventhub/v20170401:Namespace" }, { type: "azurerm:eventhub/v20180101preview:Namespace" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The time the Namespace was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Specifies whether this instance is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Namespace location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The Namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Provisioning state of the Namespace.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Name of the resource group within the azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Endpoint you can use to perform Service Bus operations.
     */
    readonly serviceBusEndpoint?: pulumi.Input<string>;
    /**
     * SKU parameters supplied to the create Namespace operation
     */
    readonly sku?: pulumi.Input<inputs.eventhub.v20150801.Sku>;
    /**
     * State of the Namespace.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Namespace tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time the Namespace was updated.
     */
    readonly updatedAt?: pulumi.Input<string>;
}
