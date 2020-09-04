// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents settings of an environment, from which environment instances would be created
 */
export class EnvironmentSetting extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EnvironmentSetting {
        return new EnvironmentSetting(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:labservices/latest:EnvironmentSetting';

    /**
     * Returns true if the given object is an instance of EnvironmentSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentSetting.__pulumiType;
    }

    /**
     * Describes the user's progress in configuring their environment setting
     */
    public readonly configurationState!: pulumi.Output<string | undefined>;
    /**
     * Describes the environment and its resource settings
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Time when the template VM was last changed.
     */
    public /*out*/ readonly lastChanged!: pulumi.Output<string>;
    /**
     * Time when the template VM was last sent for publishing.
     */
    public /*out*/ readonly lastPublished!: pulumi.Output<string>;
    /**
     * The details of the latest operation. ex: status, error
     */
    public /*out*/ readonly latestOperationResult!: pulumi.Output<outputs.labservices.latest.LatestOperationResultResponse>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning status of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Describes the readiness of this environment setting
     */
    public /*out*/ readonly publishingState!: pulumi.Output<string>;
    /**
     * The resource specific settings
     */
    public readonly resourceSettings!: pulumi.Output<outputs.labservices.latest.ResourceSettingsResponse>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Brief title describing the environment and its resource settings
     */
    public readonly title!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public readonly uniqueIdentifier!: pulumi.Output<string | undefined>;

    /**
     * Create a EnvironmentSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentSettingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.environmentSettingName === undefined) {
                throw new Error("Missing required property 'environmentSettingName'");
            }
            if (!args || args.labAccountName === undefined) {
                throw new Error("Missing required property 'labAccountName'");
            }
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceSettings === undefined) {
                throw new Error("Missing required property 'resourceSettings'");
            }
            inputs["configurationState"] = args ? args.configurationState : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environmentSettingName"] = args ? args.environmentSettingName : undefined;
            inputs["labAccountName"] = args ? args.labAccountName : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceSettings"] = args ? args.resourceSettings : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["uniqueIdentifier"] = args ? args.uniqueIdentifier : undefined;
            inputs["lastChanged"] = undefined /*out*/;
            inputs["lastPublished"] = undefined /*out*/;
            inputs["latestOperationResult"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["publishingState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["configurationState"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["lastChanged"] = undefined /*out*/;
            inputs["lastPublished"] = undefined /*out*/;
            inputs["latestOperationResult"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publishingState"] = undefined /*out*/;
            inputs["resourceSettings"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:labservices/v20181015:EnvironmentSetting" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(EnvironmentSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EnvironmentSetting resource.
 */
export interface EnvironmentSettingArgs {
    /**
     * Describes the user's progress in configuring their environment setting
     */
    readonly configurationState?: pulumi.Input<string>;
    /**
     * Describes the environment and its resource settings
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the environment Setting.
     */
    readonly environmentSettingName: pulumi.Input<string>;
    /**
     * The name of the lab Account.
     */
    readonly labAccountName: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource specific settings
     */
    readonly resourceSettings: pulumi.Input<inputs.labservices.latest.ResourceSettings>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Brief title describing the environment and its resource settings
     */
    readonly title?: pulumi.Input<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier?: pulumi.Input<string>;
}
