// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Site REST Resource.
 *
 * ## Example Usage
 * ### Create Hyper-V site
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const hyperVSite = new azurerm.offazure.latest.HyperVSite("hyperVSite", {
 *     location: "eastus",
 *     resourceGroupName: "pajindTest",
 *     siteName: "appliance1e39site",
 * });
 *
 * ```
 */
export class HyperVSite extends pulumi.CustomResource {
    /**
     * Get an existing HyperVSite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HyperVSite {
        return new HyperVSite(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:offazure/latest:HyperVSite';

    /**
     * Returns true if the given object is an instance of HyperVSite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HyperVSite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HyperVSite.__pulumiType;
    }

    /**
     * eTag for concurrency control.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * Azure location in which Sites is created.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of the Hyper-V site.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Nested properties of Hyper-V site.
     */
    public readonly properties!: pulumi.Output<outputs.offazure.latest.SitePropertiesResponse>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of resource. Type = Microsoft.OffAzure/HyperVSites.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a HyperVSite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HyperVSiteArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.siteName === undefined) {
                throw new Error("Missing required property 'siteName'");
            }
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["siteName"] = args ? args.siteName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["eTag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:offazure/v20200101:HyperVSite" }, { type: "azurerm:offazure/v20200707:HyperVSite" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(HyperVSite.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HyperVSite resource.
 */
export interface HyperVSiteArgs {
    /**
     * eTag for concurrency control.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * Azure location in which Sites is created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Name of the Hyper-V site.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Nested properties of Hyper-V site.
     */
    readonly properties?: pulumi.Input<inputs.offazure.latest.SiteProperties>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Site name.
     */
    readonly siteName: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
