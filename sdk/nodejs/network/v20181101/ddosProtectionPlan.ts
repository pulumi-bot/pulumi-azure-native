// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A DDoS protection plan in a resource group.
 *
 * ## Example Usage
 * ### Create DDoS protection plan
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const ddosProtectionPlan = new azurerm.network.v20181101.DdosProtectionPlan("ddosProtectionPlan", {
 *     ddosProtectionPlanName: "test-plan",
 *     location: "westus",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class DdosProtectionPlan extends pulumi.CustomResource {
    /**
     * Get an existing DdosProtectionPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DdosProtectionPlan {
        return new DdosProtectionPlan(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20181101:DdosProtectionPlan';

    /**
     * Returns true if the given object is an instance of DdosProtectionPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DdosProtectionPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DdosProtectionPlan.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the DDoS protection plan resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
     */
    public /*out*/ readonly virtualNetworks!: pulumi.Output<outputs.network.v20181101.SubResourceResponse[]>;

    /**
     * Create a DdosProtectionPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DdosProtectionPlanArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.ddosProtectionPlanName === undefined) {
                throw new Error("Missing required property 'ddosProtectionPlanName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["ddosProtectionPlanName"] = args ? args.ddosProtectionPlanName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworks"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworks"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:DdosProtectionPlan" }, { type: "azurerm:network/v20180201:DdosProtectionPlan" }, { type: "azurerm:network/v20180401:DdosProtectionPlan" }, { type: "azurerm:network/v20180601:DdosProtectionPlan" }, { type: "azurerm:network/v20180701:DdosProtectionPlan" }, { type: "azurerm:network/v20180801:DdosProtectionPlan" }, { type: "azurerm:network/v20181001:DdosProtectionPlan" }, { type: "azurerm:network/v20181201:DdosProtectionPlan" }, { type: "azurerm:network/v20190201:DdosProtectionPlan" }, { type: "azurerm:network/v20190401:DdosProtectionPlan" }, { type: "azurerm:network/v20190601:DdosProtectionPlan" }, { type: "azurerm:network/v20190701:DdosProtectionPlan" }, { type: "azurerm:network/v20190801:DdosProtectionPlan" }, { type: "azurerm:network/v20190901:DdosProtectionPlan" }, { type: "azurerm:network/v20191101:DdosProtectionPlan" }, { type: "azurerm:network/v20191201:DdosProtectionPlan" }, { type: "azurerm:network/v20200301:DdosProtectionPlan" }, { type: "azurerm:network/v20200401:DdosProtectionPlan" }, { type: "azurerm:network/v20200501:DdosProtectionPlan" }, { type: "azurerm:network/v20200601:DdosProtectionPlan" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DdosProtectionPlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DdosProtectionPlan resource.
 */
export interface DdosProtectionPlanArgs {
    /**
     * The name of the DDoS protection plan.
     */
    readonly ddosProtectionPlanName: pulumi.Input<string>;
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
