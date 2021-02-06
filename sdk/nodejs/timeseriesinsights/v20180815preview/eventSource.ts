// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An environment receives data from one or more event sources. Each event source has associated connection info that allows the Time Series Insights ingress pipeline to connect to and pull data from the event source
 */
export class EventSource extends pulumi.CustomResource {
    /**
     * Get an existing EventSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventSource {
        return new EventSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:timeseriesinsights/v20180815preview:EventSource';

    /**
     * Returns true if the given object is an instance of EventSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSource.__pulumiType;
    }

    /**
     * The kind of the event source.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EventSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.environmentName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'environmentName'");
            }
            if ((!args || args.eventSourceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'eventSourceName'");
            }
            if ((!args || args.kind === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["environmentName"] = args ? args.environmentName : undefined;
            inputs["eventSourceName"] = args ? args.eventSourceName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["localTimestamp"] = args ? args.localTimestamp : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:timeseriesinsights/latest:EventSource" }, { type: "azure-nextgen:timeseriesinsights/v20170228preview:EventSource" }, { type: "azure-nextgen:timeseriesinsights/v20171115:EventSource" }, { type: "azure-nextgen:timeseriesinsights/v20200515:EventSource" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(EventSource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventSource resource.
 */
export interface EventSourceArgs {
    /**
     * The name of the Time Series Insights environment associated with the specified resource group.
     */
    readonly environmentName: pulumi.Input<string>;
    /**
     * Name of the event source.
     */
    readonly eventSourceName: pulumi.Input<string>;
    /**
     * The kind of the event source.
     */
    readonly kind: pulumi.Input<string | enums.timeseriesinsights.v20180815preview.Kind>;
    /**
     * An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
     */
    readonly localTimestamp?: pulumi.Input<inputs.timeseriesinsights.v20180815preview.LocalTimestamp>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Key-value pairs of additional properties for the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
