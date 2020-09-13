// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Event Channel.
 */
export class EventChannel extends pulumi.CustomResource {
    /**
     * Get an existing EventChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventChannel {
        return new EventChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventgrid/v20200401preview:EventChannel';

    /**
     * Returns true if the given object is an instance of EventChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventChannel.__pulumiType;
    }

    /**
     * Represents the destination of an event channel.
     */
    public readonly destination!: pulumi.Output<outputs.eventgrid.v20200401preview.EventChannelDestinationResponse | undefined>;
    /**
     * Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
     * the event channel and corresponding partner topic are deleted.
     */
    public readonly expirationTimeIfNotActivatedUtc!: pulumi.Output<string | undefined>;
    /**
     * Information about the filter for the event channel.
     */
    public readonly filter!: pulumi.Output<outputs.eventgrid.v20200401preview.EventChannelFilterResponse | undefined>;
    /**
     * Name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
     * This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
     */
    public readonly partnerTopicFriendlyDescription!: pulumi.Output<string | undefined>;
    /**
     * The readiness state of the corresponding partner topic.
     */
    public /*out*/ readonly partnerTopicReadinessState!: pulumi.Output<string>;
    /**
     * Provisioning state of the event channel.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Source of the event channel. This represents a unique resource in the partner's resource model.
     */
    public readonly source!: pulumi.Output<outputs.eventgrid.v20200401preview.EventChannelSourceResponse | undefined>;
    /**
     * Type of the resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EventChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.eventChannelName === undefined) {
                throw new Error("Missing required property 'eventChannelName'");
            }
            if (!args || args.partnerNamespaceName === undefined) {
                throw new Error("Missing required property 'partnerNamespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["destination"] = args ? args.destination : undefined;
            inputs["eventChannelName"] = args ? args.eventChannelName : undefined;
            inputs["expirationTimeIfNotActivatedUtc"] = args ? args.expirationTimeIfNotActivatedUtc : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["partnerNamespaceName"] = args ? args.partnerNamespaceName : undefined;
            inputs["partnerTopicFriendlyDescription"] = args ? args.partnerTopicFriendlyDescription : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["partnerTopicReadinessState"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["destination"] = undefined /*out*/;
            inputs["expirationTimeIfNotActivatedUtc"] = undefined /*out*/;
            inputs["filter"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnerTopicFriendlyDescription"] = undefined /*out*/;
            inputs["partnerTopicReadinessState"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["source"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventChannel resource.
 */
export interface EventChannelArgs {
    /**
     * Represents the destination of an event channel.
     */
    readonly destination?: pulumi.Input<inputs.eventgrid.v20200401preview.EventChannelDestination>;
    /**
     * Name of the event channel.
     */
    readonly eventChannelName: pulumi.Input<string>;
    /**
     * Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
     * the event channel and corresponding partner topic are deleted.
     */
    readonly expirationTimeIfNotActivatedUtc?: pulumi.Input<string>;
    /**
     * Information about the filter for the event channel.
     */
    readonly filter?: pulumi.Input<inputs.eventgrid.v20200401preview.EventChannelFilter>;
    /**
     * Name of the partner namespace.
     */
    readonly partnerNamespaceName: pulumi.Input<string>;
    /**
     * Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
     * This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
     */
    readonly partnerTopicFriendlyDescription?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Source of the event channel. This represents a unique resource in the partner's resource model.
     */
    readonly source?: pulumi.Input<inputs.eventgrid.v20200401preview.EventChannelSource>;
}
