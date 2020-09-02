// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The description of the Windows IoT Device Service.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:windowsiot/v20180216preview:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * Windows IoT Device Service OEM AAD domain
     */
    public readonly adminDomainName!: pulumi.Output<string | undefined>;
    /**
     * Windows IoT Device Service ODM AAD domain
     */
    public /*out*/ readonly billingDomainName!: pulumi.Output<string>;
    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The Azure Region where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Windows IoT Device Service notes.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * Windows IoT Device Service device allocation,
     */
    public readonly quantity!: pulumi.Output<number | undefined>;
    /**
     * Windows IoT Device Service start date,
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["adminDomainName"] = args ? args.adminDomainName : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["quantity"] = args ? args.quantity : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["billingDomainName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startDate"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:windowsiot/latest:Service" }, { type: "azurerm:windowsiot/v20190601:Service" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Windows IoT Device Service OEM AAD domain
     */
    readonly adminDomainName?: pulumi.Input<string>;
    /**
     * The name of the Windows IoT Device Service.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The Azure Region where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Windows IoT Device Service notes.
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * Windows IoT Device Service device allocation,
     */
    readonly quantity?: pulumi.Input<number>;
    /**
     * The name of the resource group that contains the Windows IoT Device Service.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
