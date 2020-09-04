// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Hub resource.
 *
 * ## Example Usage
 * ### Hubs_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const hub = new azurerm.customerinsights.latest.Hub("hub", {
 *     hubBillingInfo: {
 *         maxUnits: 5,
 *         minUnits: 1,
 *         skuName: "B0",
 *     },
 *     hubName: "sdkTestHub",
 *     location: "West US",
 *     resourceGroupName: "TestHubRG",
 * });
 *
 * ```
 */
export class Hub extends pulumi.CustomResource {
    /**
     * Get an existing Hub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Hub {
        return new Hub(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/latest:Hub';

    /**
     * Returns true if the given object is an instance of Hub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hub.__pulumiType;
    }

    /**
     * API endpoint URL of the hub.
     */
    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    /**
     * Billing settings of the hub.
     */
    public readonly hubBillingInfo!: pulumi.Output<outputs.customerinsights.latest.HubBillingInfoFormatResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the hub.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
     */
    public readonly tenantFeatures!: pulumi.Output<number | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Web endpoint URL of the hub.
     */
    public /*out*/ readonly webEndpoint!: pulumi.Output<string>;

    /**
     * Create a Hub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HubArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hubBillingInfo"] = args ? args.hubBillingInfo : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantFeatures"] = args ? args.tenantFeatures : undefined;
            inputs["apiEndpoint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["webEndpoint"] = undefined /*out*/;
        } else {
            inputs["apiEndpoint"] = undefined /*out*/;
            inputs["hubBillingInfo"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantFeatures"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["webEndpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/v20170101:Hub" }, { type: "azurerm:customerinsights/v20170426:Hub" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Hub.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Hub resource.
 */
export interface HubArgs {
    /**
     * Billing settings of the hub.
     */
    readonly hubBillingInfo?: pulumi.Input<inputs.customerinsights.latest.HubBillingInfoFormat>;
    /**
     * The name of the Hub.
     */
    readonly hubName: pulumi.Input<string>;
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
    /**
     * The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
     */
    readonly tenantFeatures?: pulumi.Input<number>;
}
