// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Single item in List or Get Consumer group operation
 */
export class NamespaceEventhubConsumergroup extends pulumi.CustomResource {
    /**
     * Get an existing NamespaceEventhubConsumergroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NamespaceEventhubConsumergroup {
        return new NamespaceEventhubConsumergroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventhub:NamespaceEventhubConsumergroup';

    /**
     * Returns true if the given object is an instance of NamespaceEventhubConsumergroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamespaceEventhubConsumergroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamespaceEventhubConsumergroup.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Single item in List or Get Consumer group operation
     */
    public readonly properties!: pulumi.Output<outputs.eventhub.ConsumerGroupResponseProperties>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NamespaceEventhubConsumergroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NamespaceEventhubConsumergroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.eventHubName === undefined) {
                throw new Error("Missing required property 'eventHubName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["eventHubName"] = args ? args.eventHubName : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["namespaceName"] = args ? args.namespaceName : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NamespaceEventhubConsumergroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NamespaceEventhubConsumergroup resource.
 */
export interface NamespaceEventhubConsumergroupArgs {
    /**
     * The Event Hub name
     */
    readonly eventHubName: pulumi.Input<string>;
    /**
     * The consumer group name
     */
    readonly name: pulumi.Input<string>;
    /**
     * The Namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Single item in List or Get Consumer group operation
     */
    readonly properties?: pulumi.Input<inputs.eventhub.ConsumerGroupProperties>;
    /**
     * Name of the resource group within the azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
