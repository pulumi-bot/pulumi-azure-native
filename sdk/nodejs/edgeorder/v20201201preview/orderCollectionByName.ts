// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Specifies the properties or parameters for an order collection. Order collection is a grouping of one or more orders.
 */
export class OrderCollectionByName extends pulumi.CustomResource {
    /**
     * Get an existing OrderCollectionByName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrderCollectionByName {
        return new OrderCollectionByName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:edgeorder/v20201201preview:OrderCollectionByName';

    /**
     * Returns true if the given object is an instance of OrderCollectionByName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrderCollectionByName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrderCollectionByName.__pulumiType;
    }

    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of order ARM Ids which are part of an order collection.
     */
    public readonly orderIds!: pulumi.Output<string[]>;
    /**
     * Represents resource creation and update time
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.edgeorder.v20201201preview.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a OrderCollectionByName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrderCollectionByNameArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.orderCollectionName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'orderCollectionName'");
            }
            if ((!args || args.orderIds === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'orderIds'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["orderCollectionName"] = args ? args.orderCollectionName : undefined;
            inputs["orderIds"] = args ? args.orderIds : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["orderIds"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrderCollectionByName.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrderCollectionByName resource.
 */
export interface OrderCollectionByNameArgs {
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the order collection
     */
    readonly orderCollectionName: pulumi.Input<string>;
    /**
     * List of order ARM Ids which are part of an order collection.
     */
    readonly orderIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
