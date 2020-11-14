// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * The properties of the EventHubConsumerGroupInfo object.
 */
export class IotHubResourceEventHubConsumerGroup extends pulumi.CustomResource {
    /**
     * Get an existing IotHubResourceEventHubConsumerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IotHubResourceEventHubConsumerGroup {
        return new IotHubResourceEventHubConsumerGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:devices/latest:IotHubResourceEventHubConsumerGroup';

    /**
     * Returns true if the given object is an instance of IotHubResourceEventHubConsumerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotHubResourceEventHubConsumerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotHubResourceEventHubConsumerGroup.__pulumiType;
    }

    /**
     * The etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The Event Hub-compatible consumer group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The tags.
     */
    public readonly properties!: pulumi.Output<{[key: string]: string}>;
    /**
     * the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IotHubResourceEventHubConsumerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotHubResourceEventHubConsumerGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.eventHubEndpointName === undefined) {
                throw new Error("Missing required property 'eventHubEndpointName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["eventHubEndpointName"] = args ? args.eventHubEndpointName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:devices/v20160203:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20170119:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20170701:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20180122:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20180401:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20181201preview:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20190322:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20190322preview:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20190701preview:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20191104:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20200301:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20200401:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20200615:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20200710preview:IotHubResourceEventHubConsumerGroup" }, { type: "azure-nextgen:devices/v20200801:IotHubResourceEventHubConsumerGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IotHubResourceEventHubConsumerGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IotHubResourceEventHubConsumerGroup resource.
 */
export interface IotHubResourceEventHubConsumerGroupArgs {
    /**
     * The name of the Event Hub-compatible endpoint in the IoT hub.
     */
    readonly eventHubEndpointName: pulumi.Input<string>;
    /**
     * The name of the consumer group to add.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The EventHub consumer group name.
     */
    readonly properties?: pulumi.Input<inputs.devices.latest.EventHubConsumerGroupName>;
    /**
     * The name of the resource group that contains the IoT hub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the IoT hub.
     */
    readonly resourceName: pulumi.Input<string>;
}
