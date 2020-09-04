// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Single item in List or Get Event Hub operation
 */
export class EventHub extends pulumi.CustomResource {
    /**
     * Get an existing EventHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventHub {
        return new EventHub(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventhub/v20140901:EventHub';

    /**
     * Returns true if the given object is an instance of EventHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventHub.__pulumiType;
    }

    /**
     * Exact time the Event Hub was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Number of days to retain the events for this Event Hub.
     */
    public readonly messageRetentionInDays!: pulumi.Output<number | undefined>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of partitions created for the Event Hub.
     */
    public readonly partitionCount!: pulumi.Output<number | undefined>;
    /**
     * Current number of shards on the Event Hub.
     */
    public /*out*/ readonly partitionIds!: pulumi.Output<string[]>;
    /**
     * Enumerates the possible values for the status of the Event Hub.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The exact time the message was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a EventHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventHubArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.eventHubName === undefined) {
                throw new Error("Missing required property 'eventHubName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["eventHubName"] = args ? args.eventHubName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["messageRetentionInDays"] = args ? args.messageRetentionInDays : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["partitionCount"] = args ? args.partitionCount : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["partitionIds"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        } else {
            inputs["createdAt"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["messageRetentionInDays"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partitionCount"] = undefined /*out*/;
            inputs["partitionIds"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:eventhub/latest:EventHub" }, { type: "azurerm:eventhub/v20150801:EventHub" }, { type: "azurerm:eventhub/v20170401:EventHub" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(EventHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventHub resource.
 */
export interface EventHubArgs {
    /**
     * The Event Hub name
     */
    readonly eventHubName: pulumi.Input<string>;
    /**
     * Location of the resource.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Number of days to retain the events for this Event Hub.
     */
    readonly messageRetentionInDays?: pulumi.Input<number>;
    /**
     * Name of the Event Hub.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Number of partitions created for the Event Hub.
     */
    readonly partitionCount?: pulumi.Input<number>;
    /**
     * Name of the resource group within the azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Enumerates the possible values for the status of the Event Hub.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * ARM type of the Namespace.
     */
    readonly type?: pulumi.Input<string>;
}
