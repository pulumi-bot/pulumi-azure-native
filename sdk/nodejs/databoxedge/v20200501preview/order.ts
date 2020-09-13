// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The order details.
 */
export class Order extends pulumi.CustomResource {
    /**
     * Get an existing Order resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Order {
        return new Order(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:databoxedge/v20200501preview:Order';

    /**
     * Returns true if the given object is an instance of Order.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Order {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Order.__pulumiType;
    }

    /**
     * The contact details.
     */
    public readonly contactInformation!: pulumi.Output<outputs.databoxedge.v20200501preview.ContactDetailsResponse>;
    /**
     * Current status of the order.
     */
    public readonly currentStatus!: pulumi.Output<outputs.databoxedge.v20200501preview.OrderStatusResponse | undefined>;
    /**
     * Tracking information for the package delivered to the customer whether it has an original or a replacement device.
     */
    public /*out*/ readonly deliveryTrackingInfo!: pulumi.Output<outputs.databoxedge.v20200501preview.TrackingInfoResponse[]>;
    /**
     * The object name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of status changes in the order.
     */
    public /*out*/ readonly orderHistory!: pulumi.Output<outputs.databoxedge.v20200501preview.OrderStatusResponse[]>;
    /**
     * Tracking information for the package returned from the customer whether it has an original or a replacement device.
     */
    public /*out*/ readonly returnTrackingInfo!: pulumi.Output<outputs.databoxedge.v20200501preview.TrackingInfoResponse[]>;
    /**
     * Serial number of the device.
     */
    public /*out*/ readonly serialNumber!: pulumi.Output<string>;
    /**
     * The shipping address.
     */
    public readonly shippingAddress!: pulumi.Output<outputs.databoxedge.v20200501preview.AddressResponse>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Order resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.contactInformation === undefined) {
                throw new Error("Missing required property 'contactInformation'");
            }
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.shippingAddress === undefined) {
                throw new Error("Missing required property 'shippingAddress'");
            }
            inputs["contactInformation"] = args ? args.contactInformation : undefined;
            inputs["currentStatus"] = args ? args.currentStatus : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shippingAddress"] = args ? args.shippingAddress : undefined;
            inputs["deliveryTrackingInfo"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["orderHistory"] = undefined /*out*/;
            inputs["returnTrackingInfo"] = undefined /*out*/;
            inputs["serialNumber"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["contactInformation"] = undefined /*out*/;
            inputs["currentStatus"] = undefined /*out*/;
            inputs["deliveryTrackingInfo"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["orderHistory"] = undefined /*out*/;
            inputs["returnTrackingInfo"] = undefined /*out*/;
            inputs["serialNumber"] = undefined /*out*/;
            inputs["shippingAddress"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:databoxedge/latest:Order" }, { type: "azurerm:databoxedge/v20190301:Order" }, { type: "azurerm:databoxedge/v20190701:Order" }, { type: "azurerm:databoxedge/v20190801:Order" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Order.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Order resource.
 */
export interface OrderArgs {
    /**
     * The contact details.
     */
    readonly contactInformation: pulumi.Input<inputs.databoxedge.v20200501preview.ContactDetails>;
    /**
     * Current status of the order.
     */
    readonly currentStatus?: pulumi.Input<inputs.databoxedge.v20200501preview.OrderStatus>;
    /**
     * The order details of a device.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The shipping address.
     */
    readonly shippingAddress: pulumi.Input<inputs.databoxedge.v20200501preview.Address>;
}
