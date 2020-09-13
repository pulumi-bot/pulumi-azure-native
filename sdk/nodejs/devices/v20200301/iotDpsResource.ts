// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The description of the provisioning service.
 */
export class IotDpsResource extends pulumi.CustomResource {
    /**
     * Get an existing IotDpsResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IotDpsResource {
        return new IotDpsResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devices/v20200301:IotDpsResource';

    /**
     * Returns true if the given object is an instance of IotDpsResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotDpsResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotDpsResource.__pulumiType;
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
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Service specific properties for a provisioning service
     */
    public readonly properties!: pulumi.Output<outputs.devices.v20200301.IotDpsPropertiesDescriptionResponse>;
    /**
     * Sku info for a provisioning Service.
     */
    public readonly sku!: pulumi.Output<outputs.devices.v20200301.IotDpsSkuInfoResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IotDpsResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotDpsResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.provisioningServiceName === undefined) {
                throw new Error("Missing required property 'provisioningServiceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["provisioningServiceName"] = args ? args.provisioningServiceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devices/latest:IotDpsResource" }, { type: "azurerm:devices/v20170821preview:IotDpsResource" }, { type: "azurerm:devices/v20171115:IotDpsResource" }, { type: "azurerm:devices/v20180122:IotDpsResource" }, { type: "azurerm:devices/v20200101:IotDpsResource" }, { type: "azurerm:devices/v20200901preview:IotDpsResource" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IotDpsResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IotDpsResource resource.
 */
export interface IotDpsResourceArgs {
    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Service specific properties for a provisioning service
     */
    readonly properties: pulumi.Input<inputs.devices.v20200301.IotDpsPropertiesDescription>;
    /**
     * Name of provisioning service to create or update.
     */
    readonly provisioningServiceName: pulumi.Input<string>;
    /**
     * Resource group identifier.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sku info for a provisioning Service.
     */
    readonly sku: pulumi.Input<inputs.devices.v20200301.IotDpsSkuInfo>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
