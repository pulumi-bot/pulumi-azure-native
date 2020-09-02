// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A DDoS custom policy in a resource group.
 *
 * ## Example Usage
 * ### Create DDoS custom policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const ddosCustomPolicy = new azurerm.network.v20200601.DdosCustomPolicy("ddosCustomPolicy", {
 *     ddosCustomPolicyName: "test-ddos-custom-policy",
 *     location: "centraluseuap",
 *     protocolCustomSettings: [{
 *         protocol: "Tcp",
 *     }],
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class DdosCustomPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DdosCustomPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DdosCustomPolicy {
        return new DdosCustomPolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:DdosCustomPolicy';

    /**
     * Returns true if the given object is an instance of DdosCustomPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DdosCustomPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DdosCustomPolicy.__pulumiType;
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
     * The protocol-specific DDoS policy customization parameters.
     */
    public readonly protocolCustomSettings!: pulumi.Output<outputs.network.v20200601.ProtocolCustomSettingsFormatResponse[] | undefined>;
    /**
     * The provisioning state of the DDoS custom policy resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The list of public IPs associated with the DDoS custom policy resource. This list is read-only.
     */
    public /*out*/ readonly publicIPAddresses!: pulumi.Output<outputs.network.v20200601.SubResourceResponse[]>;
    /**
     * The resource GUID property of the DDoS custom policy resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
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
     * Create a DdosCustomPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DdosCustomPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DdosCustomPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DdosCustomPolicyArgs | undefined;
            if (!args || args.ddosCustomPolicyName === undefined) {
                throw new Error("Missing required property 'ddosCustomPolicyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["ddosCustomPolicyName"] = args ? args.ddosCustomPolicyName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["protocolCustomSettings"] = args ? args.protocolCustomSettings : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicIPAddresses"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:DdosCustomPolicy" }, { type: "azurerm:network/v20181101:DdosCustomPolicy" }, { type: "azurerm:network/v20181201:DdosCustomPolicy" }, { type: "azurerm:network/v20190201:DdosCustomPolicy" }, { type: "azurerm:network/v20190401:DdosCustomPolicy" }, { type: "azurerm:network/v20190601:DdosCustomPolicy" }, { type: "azurerm:network/v20190701:DdosCustomPolicy" }, { type: "azurerm:network/v20190801:DdosCustomPolicy" }, { type: "azurerm:network/v20190901:DdosCustomPolicy" }, { type: "azurerm:network/v20191101:DdosCustomPolicy" }, { type: "azurerm:network/v20191201:DdosCustomPolicy" }, { type: "azurerm:network/v20200301:DdosCustomPolicy" }, { type: "azurerm:network/v20200401:DdosCustomPolicy" }, { type: "azurerm:network/v20200501:DdosCustomPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DdosCustomPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DdosCustomPolicy resource.
 */
export interface DdosCustomPolicyArgs {
    /**
     * The name of the DDoS custom policy.
     */
    readonly ddosCustomPolicyName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The protocol-specific DDoS policy customization parameters.
     */
    readonly protocolCustomSettings?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ProtocolCustomSettingsFormat>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
