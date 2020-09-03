// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Event Subscription
 *
 * ## Example Usage
 * ### EventSubscriptions_CreateOrUpdateForCustomTopic_EventHubDestination
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     deadLetterDestination: {
 *         endpointType: "StorageBlob",
 *     },
 *     destination: {
 *         endpointType: "EventHub",
 *     },
 *     eventSubscriptionName: "examplesubscription1",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *         subjectBeginsWith: "ExamplePrefix",
 *         subjectEndsWith: "ExampleSuffix",
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
 * });
 *
 * ```
 * ### EventSubscriptions_CreateOrUpdateForCustomTopic_HybridConnectionDestination
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     deadLetterDestination: {
 *         endpointType: "StorageBlob",
 *     },
 *     destination: {
 *         endpointType: "HybridConnection",
 *     },
 *     eventSubscriptionName: "examplesubscription1",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *         subjectBeginsWith: "ExamplePrefix",
 *         subjectEndsWith: "ExampleSuffix",
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
 * });
 *
 * ```
 * ### EventSubscriptions_CreateOrUpdateForCustomTopic_StorageQueueDestination
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     deadLetterDestination: {
 *         endpointType: "StorageBlob",
 *     },
 *     destination: {
 *         endpointType: "StorageQueue",
 *     },
 *     eventSubscriptionName: "examplesubscription1",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *         subjectBeginsWith: "ExamplePrefix",
 *         subjectEndsWith: "ExampleSuffix",
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
 * });
 *
 * ```
 * ### EventSubscriptions_CreateOrUpdateForCustomTopic_WebhookDestination
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     destination: {
 *         endpointType: "EventHub",
 *     },
 *     eventSubscriptionName: "examplesubscription1",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *         subjectBeginsWith: "ExamplePrefix",
 *         subjectEndsWith: "ExampleSuffix",
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
 * });
 *
 * ```
 * ### EventSubscriptions_CreateOrUpdateForResource
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     destination: {
 *         endpointType: "WebHook",
 *     },
 *     eventSubscriptionName: "examplesubscription10",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *         subjectBeginsWith: "ExamplePrefix",
 *         subjectEndsWith: "ExampleSuffix",
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1",
 * });
 *
 * ```
 * ### EventSubscriptions_CreateOrUpdateForResourceGroup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     destination: {
 *         endpointType: "WebHook",
 *     },
 *     eventSubscriptionName: "examplesubscription2",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *         subjectBeginsWith: "ExamplePrefix",
 *         subjectEndsWith: "ExampleSuffix",
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg",
 * });
 *
 * ```
 * ### EventSubscriptions_CreateOrUpdateForSubscription
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const eventSubscription = new azurerm.eventgrid.v20200601.EventSubscription("eventSubscription", {
 *     destination: {
 *         endpointType: "WebHook",
 *     },
 *     eventSubscriptionName: "examplesubscription3",
 *     filter: {
 *         isSubjectCaseSensitive: false,
 *     },
 *     scope: "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4",
 * });
 *
 * ```
 */
export class EventSubscription extends pulumi.CustomResource {
    /**
     * Get an existing EventSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventSubscription {
        return new EventSubscription(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventgrid/v20200601:EventSubscription';

    /**
     * Returns true if the given object is an instance of EventSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSubscription.__pulumiType;
    }

    /**
     * The DeadLetter destination of the event subscription.
     */
    public readonly deadLetterDestination!: pulumi.Output<outputs.eventgrid.v20200601.DeadLetterDestinationResponse | undefined>;
    /**
     * Information about the destination where events have to be delivered for the event subscription.
     */
    public readonly destination!: pulumi.Output<outputs.eventgrid.v20200601.EventSubscriptionDestinationResponse | undefined>;
    /**
     * The event delivery schema for the event subscription.
     */
    public readonly eventDeliverySchema!: pulumi.Output<string | undefined>;
    /**
     * Expiration time of the event subscription.
     */
    public readonly expirationTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * Information about the filter for the event subscription.
     */
    public readonly filter!: pulumi.Output<outputs.eventgrid.v20200601.EventSubscriptionFilterResponse | undefined>;
    /**
     * List of user defined labels.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the event subscription.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
     */
    public readonly retryPolicy!: pulumi.Output<outputs.eventgrid.v20200601.RetryPolicyResponse | undefined>;
    /**
     * Name of the topic of the event subscription.
     */
    public /*out*/ readonly topic!: pulumi.Output<string>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EventSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventSubscriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as EventSubscriptionArgs | undefined;
            if (!args || args.eventSubscriptionName === undefined) {
                throw new Error("Missing required property 'eventSubscriptionName'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["deadLetterDestination"] = args ? args.deadLetterDestination : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["eventDeliverySchema"] = args ? args.eventDeliverySchema : undefined;
            inputs["eventSubscriptionName"] = args ? args.eventSubscriptionName : undefined;
            inputs["expirationTimeUtc"] = args ? args.expirationTimeUtc : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["retryPolicy"] = args ? args.retryPolicy : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["topic"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:eventgrid/latest:EventSubscription" }, { type: "azurerm:eventgrid/v20170615preview:EventSubscription" }, { type: "azurerm:eventgrid/v20170915preview:EventSubscription" }, { type: "azurerm:eventgrid/v20180101:EventSubscription" }, { type: "azurerm:eventgrid/v20180501preview:EventSubscription" }, { type: "azurerm:eventgrid/v20180915preview:EventSubscription" }, { type: "azurerm:eventgrid/v20190101:EventSubscription" }, { type: "azurerm:eventgrid/v20190201preview:EventSubscription" }, { type: "azurerm:eventgrid/v20190601:EventSubscription" }, { type: "azurerm:eventgrid/v20200101preview:EventSubscription" }, { type: "azurerm:eventgrid/v20200401preview:EventSubscription" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(EventSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    /**
     * The DeadLetter destination of the event subscription.
     */
    readonly deadLetterDestination?: pulumi.Input<inputs.eventgrid.v20200601.DeadLetterDestination>;
    /**
     * Information about the destination where events have to be delivered for the event subscription.
     */
    readonly destination?: pulumi.Input<inputs.eventgrid.v20200601.EventSubscriptionDestination>;
    /**
     * The event delivery schema for the event subscription.
     */
    readonly eventDeliverySchema?: pulumi.Input<string>;
    /**
     * Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
     */
    readonly eventSubscriptionName: pulumi.Input<string>;
    /**
     * Expiration time of the event subscription.
     */
    readonly expirationTimeUtc?: pulumi.Input<string>;
    /**
     * Information about the filter for the event subscription.
     */
    readonly filter?: pulumi.Input<inputs.eventgrid.v20200601.EventSubscriptionFilter>;
    /**
     * List of user defined labels.
     */
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
     */
    readonly retryPolicy?: pulumi.Input<inputs.eventgrid.v20200601.RetryPolicy>;
    /**
     * The identifier of the resource to which the event subscription needs to be created or updated. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
     */
    readonly scope: pulumi.Input<string>;
}
