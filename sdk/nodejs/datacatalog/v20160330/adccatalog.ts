// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Azure Data Catalog.
 *
 * ## Example Usage
 * ### Create Azure Data Catalog Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const adcCatalog = new azurerm.datacatalog.v20160330.ADCCatalog("adcCatalog", {
 *     admins: [{
 *         objectId: "99999999-9999-9999-999999999999",
 *         upn: "myupn@microsoft.com",
 *     }],
 *     catalogName: "exampleCatalog",
 *     enableAutomaticUnitAdjustment: false,
 *     location: "North US",
 *     resourceGroupName: "exampleResourceGroup",
 *     sku: "Standard",
 *     tags: {
 *         mykey: "myvalue",
 *         mykey2: "myvalue2",
 *     },
 *     units: 1,
 *     users: [{
 *         objectId: "99999999-9999-9999-999999999999",
 *         upn: "myupn@microsoft.com",
 *     }],
 * });
 *
 * ```
 */
export class ADCCatalog extends pulumi.CustomResource {
    /**
     * Get an existing ADCCatalog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ADCCatalog {
        return new ADCCatalog(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datacatalog/v20160330:ADCCatalog';

    /**
     * Returns true if the given object is an instance of ADCCatalog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ADCCatalog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ADCCatalog.__pulumiType;
    }

    /**
     * Azure data catalog admin list.
     */
    public readonly admins!: pulumi.Output<outputs.datacatalog.v20160330.PrincipalsResponse[] | undefined>;
    /**
     * Automatic unit adjustment enabled or not.
     */
    public readonly enableAutomaticUnitAdjustment!: pulumi.Output<boolean | undefined>;
    /**
     * Resource etag
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure data catalog SKU.
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * Azure data catalog provision status.
     */
    public readonly successfullyProvisioned!: pulumi.Output<boolean | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Azure data catalog units.
     */
    public readonly units!: pulumi.Output<number | undefined>;
    /**
     * Azure data catalog user list.
     */
    public readonly users!: pulumi.Output<outputs.datacatalog.v20160330.PrincipalsResponse[] | undefined>;

    /**
     * Create a ADCCatalog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ADCCatalogArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.catalogName === undefined) {
                throw new Error("Missing required property 'catalogName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["admins"] = args ? args.admins : undefined;
            inputs["catalogName"] = args ? args.catalogName : undefined;
            inputs["enableAutomaticUnitAdjustment"] = args ? args.enableAutomaticUnitAdjustment : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["successfullyProvisioned"] = args ? args.successfullyProvisioned : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["units"] = args ? args.units : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["admins"] = undefined /*out*/;
            inputs["enableAutomaticUnitAdjustment"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["successfullyProvisioned"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["units"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:datacatalog/latest:ADCCatalog" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ADCCatalog.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ADCCatalog resource.
 */
export interface ADCCatalogArgs {
    /**
     * Azure data catalog admin list.
     */
    readonly admins?: pulumi.Input<pulumi.Input<inputs.datacatalog.v20160330.Principals>[]>;
    /**
     * The name of the data catalog in the specified subscription and resource group.
     */
    readonly catalogName: pulumi.Input<string>;
    /**
     * Automatic unit adjustment enabled or not.
     */
    readonly enableAutomaticUnitAdjustment?: pulumi.Input<boolean>;
    /**
     * Resource etag
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Azure data catalog SKU.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * Azure data catalog provision status.
     */
    readonly successfullyProvisioned?: pulumi.Input<boolean>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Azure data catalog units.
     */
    readonly units?: pulumi.Input<number>;
    /**
     * Azure data catalog user list.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.datacatalog.v20160330.Principals>[]>;
}
