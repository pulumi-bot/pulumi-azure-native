// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The description of the IoTSpaces service.
 */
export class IoTSpace extends pulumi.CustomResource {
    /**
     * Get an existing IoTSpace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IoTSpace {
        return new IoTSpace(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:iotspaces/v20171001preview:IoTSpace';

    /**
     * Returns true if the given object is an instance of IoTSpace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IoTSpace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IoTSpace.__pulumiType;
    }

    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The common properties of a IoTSpaces service.
     */
    public readonly properties!: pulumi.Output<outputs.iotspaces.v20171001preview.IoTSpacesPropertiesResponse>;
    /**
     * A valid instance SKU.
     */
    public readonly sku!: pulumi.Output<outputs.iotspaces.v20171001preview.IoTSpacesSkuInfoResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IoTSpace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IoTSpaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IoTSpaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IoTSpaceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IoTSpace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IoTSpace resource.
 */
export interface IoTSpaceArgs {
    /**
     * The resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The common properties of a IoTSpaces service.
     */
    readonly properties?: pulumi.Input<inputs.iotspaces.v20171001preview.IoTSpacesProperties>;
    /**
     * The name of the resource group that contains the IoTSpaces instance.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the IoTSpaces instance.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * A valid instance SKU.
     */
    readonly sku: pulumi.Input<inputs.iotspaces.v20171001preview.IoTSpacesSkuInfo>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
