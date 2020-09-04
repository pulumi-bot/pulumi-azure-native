// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Hybrid Connection for an App Service app.
 */
export class WebAppRelayServiceConnectionSlot extends pulumi.CustomResource {
    /**
     * Get an existing WebAppRelayServiceConnectionSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppRelayServiceConnectionSlot {
        return new WebAppRelayServiceConnectionSlot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/latest:WebAppRelayServiceConnectionSlot';

    /**
     * Returns true if the given object is an instance of WebAppRelayServiceConnectionSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppRelayServiceConnectionSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppRelayServiceConnectionSlot.__pulumiType;
    }

    public readonly biztalkUri!: pulumi.Output<string | undefined>;
    public readonly entityConnectionString!: pulumi.Output<string | undefined>;
    public readonly entityName!: pulumi.Output<string | undefined>;
    public readonly hostname!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly resourceConnectionString!: pulumi.Output<string | undefined>;
    public readonly resourceType!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebAppRelayServiceConnectionSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppRelayServiceConnectionSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.entityName === undefined) {
                throw new Error("Missing required property 'entityName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.slot === undefined) {
                throw new Error("Missing required property 'slot'");
            }
            inputs["biztalkUri"] = args ? args.biztalkUri : undefined;
            inputs["entityConnectionString"] = args ? args.entityConnectionString : undefined;
            inputs["entityName"] = args ? args.entityName : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["resourceConnectionString"] = args ? args.resourceConnectionString : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["biztalkUri"] = undefined /*out*/;
            inputs["entityConnectionString"] = undefined /*out*/;
            inputs["entityName"] = undefined /*out*/;
            inputs["hostname"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["resourceConnectionString"] = undefined /*out*/;
            inputs["resourceType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/v20150801:WebAppRelayServiceConnectionSlot" }, { type: "azurerm:web/v20160801:WebAppRelayServiceConnectionSlot" }, { type: "azurerm:web/v20180201:WebAppRelayServiceConnectionSlot" }, { type: "azurerm:web/v20181101:WebAppRelayServiceConnectionSlot" }, { type: "azurerm:web/v20190801:WebAppRelayServiceConnectionSlot" }, { type: "azurerm:web/v20200601:WebAppRelayServiceConnectionSlot" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebAppRelayServiceConnectionSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppRelayServiceConnectionSlot resource.
 */
export interface WebAppRelayServiceConnectionSlotArgs {
    readonly biztalkUri?: pulumi.Input<string>;
    readonly entityConnectionString?: pulumi.Input<string>;
    readonly entityName: pulumi.Input<string>;
    readonly hostname?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the app.
     */
    readonly name: pulumi.Input<string>;
    readonly port?: pulumi.Input<number>;
    readonly resourceConnectionString?: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    readonly resourceType?: pulumi.Input<string>;
    /**
     * Name of the deployment slot. If a slot is not specified, the API will create or update a hybrid connection for the production slot.
     */
    readonly slot: pulumi.Input<string>;
}
