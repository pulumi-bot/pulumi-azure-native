// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * DigitalTwinsInstance endpoint resource.
 */
export class DigitalTwinsEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing DigitalTwinsEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DigitalTwinsEndpoint {
        return new DigitalTwinsEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:digitaltwins/latest:DigitalTwinsEndpoint';

    /**
     * Returns true if the given object is an instance of DigitalTwinsEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DigitalTwinsEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DigitalTwinsEndpoint.__pulumiType;
    }

    /**
     * Extension resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * DigitalTwinsInstance endpoint resource properties.
     */
    public readonly properties!: pulumi.Output<outputs.digitaltwins.latest.EventGridResponse | outputs.digitaltwins.latest.EventHubResponse | outputs.digitaltwins.latest.ServiceBusResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DigitalTwinsEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DigitalTwinsEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.endpointName === undefined) {
                throw new Error("Missing required property 'endpointName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["endpointName"] = args ? args.endpointName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:digitaltwins/v20200301preview:DigitalTwinsEndpoint" }, { type: "azure-nextgen:digitaltwins/v20201031:DigitalTwinsEndpoint" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DigitalTwinsEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DigitalTwinsEndpoint resource.
 */
export interface DigitalTwinsEndpointArgs {
    /**
     * Name of Endpoint Resource.
     */
    readonly endpointName: pulumi.Input<string>;
    /**
     * DigitalTwinsInstance endpoint resource properties.
     */
    readonly properties: pulumi.Input<inputs.digitaltwins.latest.EventGrid | inputs.digitaltwins.latest.EventHub | inputs.digitaltwins.latest.ServiceBus>;
    /**
     * The name of the resource group that contains the DigitalTwinsInstance.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the DigitalTwinsInstance.
     */
    readonly resourceName: pulumi.Input<string>;
}
