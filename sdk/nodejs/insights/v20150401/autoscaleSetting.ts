// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The autoscale setting resource.
 */
export class AutoscaleSetting extends pulumi.CustomResource {
    /**
     * Get an existing AutoscaleSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AutoscaleSetting {
        return new AutoscaleSetting(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:insights/v20150401:AutoscaleSetting';

    /**
     * Returns true if the given object is an instance of AutoscaleSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoscaleSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoscaleSetting.__pulumiType;
    }

    /**
     * the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'true'.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * the collection of notifications.
     */
    public readonly notifications!: pulumi.Output<outputs.insights.v20150401.AutoscaleNotificationResponse[] | undefined>;
    /**
     * the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
     */
    public readonly profiles!: pulumi.Output<outputs.insights.v20150401.AutoscaleProfileResponse[]>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * the resource identifier of the resource that the autoscale setting should be added to.
     */
    public readonly targetResourceUri!: pulumi.Output<string | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AutoscaleSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoscaleSettingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.autoscaleSettingName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'autoscaleSettingName'");
            }
            if ((!args || args.profiles === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'profiles'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoscaleSettingName"] = args ? args.autoscaleSettingName : undefined;
            inputs["enabled"] = (args ? args.enabled : undefined) || true;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["profiles"] = args ? args.profiles : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetResourceUri"] = args ? args.targetResourceUri : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["enabled"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notifications"] = undefined /*out*/;
            inputs["profiles"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["targetResourceUri"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:insights/latest:AutoscaleSetting" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AutoscaleSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AutoscaleSetting resource.
 */
export interface AutoscaleSettingArgs {
    /**
     * The autoscale setting name.
     */
    readonly autoscaleSettingName: pulumi.Input<string>;
    /**
     * the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'true'.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * the name of the autoscale setting.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * the collection of notifications.
     */
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.insights.v20150401.AutoscaleNotification>[]>;
    /**
     * the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
     */
    readonly profiles: pulumi.Input<pulumi.Input<inputs.insights.v20150401.AutoscaleProfile>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * the resource identifier of the resource that the autoscale setting should be added to.
     */
    readonly targetResourceUri?: pulumi.Input<string>;
}
