// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Hybrid network device resource.
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hybridnetwork/v20200101preview:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * The reference to the azure stack edge device.
     */
    public readonly azureStackEdge!: pulumi.Output<outputs.hybridnetwork.v20200101preview.SubResourceResponse | undefined>;
    /**
     * The type of the hybrid network device.
     */
    public readonly deviceType!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the hybrid network device resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The current device status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The list of virtual network functions deployed on the hybrid network device.
     */
    public /*out*/ readonly virtualNetworkFunctions!: pulumi.Output<outputs.hybridnetwork.v20200101preview.SubResourceResponse[]>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.deviceType === undefined) {
                throw new Error("Missing required property 'deviceType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["azureStackEdge"] = args ? args.azureStackEdge : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["deviceType"] = args ? args.deviceType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkFunctions"] = undefined /*out*/;
        } else {
            inputs["azureStackEdge"] = undefined /*out*/;
            inputs["deviceType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkFunctions"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Device.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * The reference to the azure stack edge device.
     */
    readonly azureStackEdge?: pulumi.Input<inputs.hybridnetwork.v20200101preview.SubResource>;
    /**
     * Resource name for the device resource.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The type of the hybrid network device.
     */
    readonly deviceType: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
