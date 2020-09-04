// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Describes a Virtual Machine Scale Set Extension.
 */
export class VirtualMachineScaleSetExtension extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachineScaleSetExtension {
        return new VirtualMachineScaleSetExtension(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20200601:VirtualMachineScaleSetExtension';

    /**
     * Returns true if the given object is an instance of VirtualMachineScaleSetExtension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachineScaleSetExtension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineScaleSetExtension.__pulumiType;
    }

    /**
     * Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
     */
    public readonly autoUpgradeMinorVersion!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
     */
    public readonly enableAutomaticUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
     */
    public readonly forceUpdateTag!: pulumi.Output<string | undefined>;
    /**
     * The name of the extension.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
     */
    public readonly protectedSettings!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Collection of extension names after which this extension needs to be provisioned.
     */
    public readonly provisionAfterExtensions!: pulumi.Output<string[] | undefined>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The name of the extension handler publisher.
     */
    public readonly publisher!: pulumi.Output<string | undefined>;
    /**
     * Json formatted public settings for the extension.
     */
    public readonly settings!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the version of the script handler.
     */
    public readonly typeHandlerVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a VirtualMachineScaleSetExtension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineScaleSetExtensionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vmScaleSetName === undefined) {
                throw new Error("Missing required property 'vmScaleSetName'");
            }
            if (!args || args.vmssExtensionName === undefined) {
                throw new Error("Missing required property 'vmssExtensionName'");
            }
            inputs["autoUpgradeMinorVersion"] = args ? args.autoUpgradeMinorVersion : undefined;
            inputs["enableAutomaticUpgrade"] = args ? args.enableAutomaticUpgrade : undefined;
            inputs["forceUpdateTag"] = args ? args.forceUpdateTag : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protectedSettings"] = args ? args.protectedSettings : undefined;
            inputs["provisionAfterExtensions"] = args ? args.provisionAfterExtensions : undefined;
            inputs["publisher"] = args ? args.publisher : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["typeHandlerVersion"] = args ? args.typeHandlerVersion : undefined;
            inputs["vmScaleSetName"] = args ? args.vmScaleSetName : undefined;
            inputs["vmssExtensionName"] = args ? args.vmssExtensionName : undefined;
            inputs["provisioningState"] = undefined /*out*/;
        } else {
            inputs["autoUpgradeMinorVersion"] = undefined /*out*/;
            inputs["enableAutomaticUpgrade"] = undefined /*out*/;
            inputs["forceUpdateTag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["protectedSettings"] = undefined /*out*/;
            inputs["provisionAfterExtensions"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publisher"] = undefined /*out*/;
            inputs["settings"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["typeHandlerVersion"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20170330:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20171201:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20180401:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20180601:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20181001:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20190301:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20190701:VirtualMachineScaleSetExtension" }, { type: "azurerm:compute/v20191201:VirtualMachineScaleSetExtension" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualMachineScaleSetExtension.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachineScaleSetExtension resource.
 */
export interface VirtualMachineScaleSetExtensionArgs {
    /**
     * Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
     */
    readonly autoUpgradeMinorVersion?: pulumi.Input<boolean>;
    /**
     * Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
     */
    readonly enableAutomaticUpgrade?: pulumi.Input<boolean>;
    /**
     * If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
     */
    readonly forceUpdateTag?: pulumi.Input<string>;
    /**
     * The name of the extension.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
     */
    readonly protectedSettings?: pulumi.Input<{[key: string]: any}>;
    /**
     * Collection of extension names after which this extension needs to be provisioned.
     */
    readonly provisionAfterExtensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the extension handler publisher.
     */
    readonly publisher?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Json formatted public settings for the extension.
     */
    readonly settings?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the type of the extension; an example is "CustomScriptExtension".
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Specifies the version of the script handler.
     */
    readonly typeHandlerVersion?: pulumi.Input<string>;
    /**
     * The name of the VM scale set where the extension should be create or updated.
     */
    readonly vmScaleSetName: pulumi.Input<string>;
    /**
     * The name of the VM scale set extension.
     */
    readonly vmssExtensionName: pulumi.Input<string>;
}
