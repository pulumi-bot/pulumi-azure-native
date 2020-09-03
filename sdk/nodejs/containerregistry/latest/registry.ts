// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An object that represents a container registry.
 *
 * ## Example Usage
 * ### RegistryCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const registry = new azurerm.containerregistry.latest.Registry("registry", {
 *     adminUserEnabled: true,
 *     location: "westus",
 *     registryName: "myRegistry",
 *     resourceGroupName: "myResourceGroup",
 *     sku: {
 *         name: "Standard",
 *     },
 *     tags: {
 *         key: "value",
 *     },
 * });
 *
 * ```
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerregistry/latest:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    /**
     * The value that indicates whether the admin user is enabled.
     */
    public readonly adminUserEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The creation date of the container registry in ISO8601 format.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The URL that can be used to log into the container registry.
     */
    public /*out*/ readonly loginServer!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The network rule set for a container registry.
     */
    public readonly networkRuleSet!: pulumi.Output<outputs.containerregistry.latest.NetworkRuleSetResponse | undefined>;
    /**
     * The policies for a container registry.
     */
    public readonly policies!: pulumi.Output<outputs.containerregistry.latest.PoliciesResponse | undefined>;
    /**
     * The provisioning state of the container registry at the time the operation was called.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The SKU of the container registry.
     */
    public readonly sku!: pulumi.Output<outputs.containerregistry.latest.SkuResponse>;
    /**
     * The status of the container registry at the time the operation was called.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.containerregistry.latest.StatusResponse>;
    /**
     * The properties of the storage account for the container registry. Only applicable to Classic SKU.
     */
    public readonly storageAccount!: pulumi.Output<outputs.containerregistry.latest.StorageAccountPropertiesResponse | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RegistryArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["adminUserEnabled"] = args ? args.adminUserEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkRuleSet"] = args ? args.networkRuleSet : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["storageAccount"] = args ? args.storageAccount : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["loginServer"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerregistry/v20160627preview:Registry" }, { type: "azurerm:containerregistry/v20170301:Registry" }, { type: "azurerm:containerregistry/v20170601preview:Registry" }, { type: "azurerm:containerregistry/v20171001:Registry" }, { type: "azurerm:containerregistry/v20190501:Registry" }, { type: "azurerm:containerregistry/v20191201preview:Registry" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Registry.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    /**
     * The value that indicates whether the admin user is enabled.
     */
    readonly adminUserEnabled?: pulumi.Input<boolean>;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The network rule set for a container registry.
     */
    readonly networkRuleSet?: pulumi.Input<inputs.containerregistry.latest.NetworkRuleSet>;
    /**
     * The policies for a container registry.
     */
    readonly policies?: pulumi.Input<inputs.containerregistry.latest.Policies>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the container registry.
     */
    readonly sku: pulumi.Input<inputs.containerregistry.latest.Sku>;
    /**
     * The properties of the storage account for the container registry. Only applicable to Classic SKU.
     */
    readonly storageAccount?: pulumi.Input<inputs.containerregistry.latest.StorageAccountProperties>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
