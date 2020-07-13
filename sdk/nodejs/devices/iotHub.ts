// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The description of the IoT hub.
 */
export class IotHub extends pulumi.CustomResource {
    /**
     * Get an existing IotHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IotHub {
        return new IotHub(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devices:IotHub';

    /**
     * Returns true if the given object is an instance of IotHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotHub.__pulumiType;
    }

    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IotHub properties
     */
    public readonly properties!: pulumi.Output<outputs.devices.IotHubPropertiesResponse>;
    /**
     * IotHub SKU info
     */
    public readonly sku!: pulumi.Output<outputs.devices.IotHubSkuInfoResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IotHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IotHubArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
        inputs["etag"] = args ? args.etag : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["sku"] = args ? args.sku : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IotHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IotHub resource.
 */
export interface IotHubArgs {
    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the IoT hub.
     */
    readonly name: pulumi.Input<string>;
    /**
     * IotHub properties
     */
    readonly properties?: pulumi.Input<inputs.devices.IotHubProperties>;
    /**
     * The name of the resource group that contains the IoT hub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * IotHub SKU info
     */
    readonly sku: pulumi.Input<inputs.devices.IotHubSkuInfo>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
