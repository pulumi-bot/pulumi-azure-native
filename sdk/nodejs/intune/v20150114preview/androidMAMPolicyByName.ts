// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Android Policy entity for Intune MAM.
 */
export class AndroidMAMPolicyByName extends pulumi.CustomResource {
    /**
     * Get an existing AndroidMAMPolicyByName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AndroidMAMPolicyByName {
        return new AndroidMAMPolicyByName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:intune/v20150114preview:AndroidMAMPolicyByName';

    /**
     * Returns true if the given object is an instance of AndroidMAMPolicyByName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AndroidMAMPolicyByName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AndroidMAMPolicyByName.__pulumiType;
    }

    public readonly accessRecheckOfflineTimeout!: pulumi.Output<string | undefined>;
    public readonly accessRecheckOnlineTimeout!: pulumi.Output<string | undefined>;
    public readonly appSharingFromLevel!: pulumi.Output<string | undefined>;
    public readonly appSharingToLevel!: pulumi.Output<string | undefined>;
    public readonly authentication!: pulumi.Output<string | undefined>;
    public readonly clipboardSharingLevel!: pulumi.Output<string | undefined>;
    public readonly dataBackup!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly deviceCompliance!: pulumi.Output<string | undefined>;
    public readonly fileEncryption!: pulumi.Output<string | undefined>;
    public readonly fileSharingSaveAs!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string>;
    public /*out*/ readonly groupStatus!: pulumi.Output<string>;
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    public readonly managedBrowser!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public /*out*/ readonly numOfApps!: pulumi.Output<number>;
    public readonly offlineWipeTimeout!: pulumi.Output<string | undefined>;
    public readonly pin!: pulumi.Output<string | undefined>;
    public readonly pinNumRetry!: pulumi.Output<number | undefined>;
    public readonly screenCapture!: pulumi.Output<string | undefined>;
    /**
     * Resource Tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AndroidMAMPolicyByName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AndroidMAMPolicyByNameArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.friendlyName === undefined) {
                throw new Error("Missing required property 'friendlyName'");
            }
            if (!args || args.hostName === undefined) {
                throw new Error("Missing required property 'hostName'");
            }
            if (!args || args.policyName === undefined) {
                throw new Error("Missing required property 'policyName'");
            }
            inputs["accessRecheckOfflineTimeout"] = args ? args.accessRecheckOfflineTimeout : undefined;
            inputs["accessRecheckOnlineTimeout"] = args ? args.accessRecheckOnlineTimeout : undefined;
            inputs["appSharingFromLevel"] = args ? args.appSharingFromLevel : undefined;
            inputs["appSharingToLevel"] = args ? args.appSharingToLevel : undefined;
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["clipboardSharingLevel"] = args ? args.clipboardSharingLevel : undefined;
            inputs["dataBackup"] = args ? args.dataBackup : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceCompliance"] = args ? args.deviceCompliance : undefined;
            inputs["fileEncryption"] = args ? args.fileEncryption : undefined;
            inputs["fileSharingSaveAs"] = args ? args.fileSharingSaveAs : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["hostName"] = args ? args.hostName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedBrowser"] = args ? args.managedBrowser : undefined;
            inputs["offlineWipeTimeout"] = args ? args.offlineWipeTimeout : undefined;
            inputs["pin"] = args ? args.pin : undefined;
            inputs["pinNumRetry"] = args ? args.pinNumRetry : undefined;
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["screenCapture"] = args ? args.screenCapture : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["groupStatus"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["numOfApps"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["accessRecheckOfflineTimeout"] = undefined /*out*/;
            inputs["accessRecheckOnlineTimeout"] = undefined /*out*/;
            inputs["appSharingFromLevel"] = undefined /*out*/;
            inputs["appSharingToLevel"] = undefined /*out*/;
            inputs["authentication"] = undefined /*out*/;
            inputs["clipboardSharingLevel"] = undefined /*out*/;
            inputs["dataBackup"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["deviceCompliance"] = undefined /*out*/;
            inputs["fileEncryption"] = undefined /*out*/;
            inputs["fileSharingSaveAs"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["groupStatus"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedBrowser"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["numOfApps"] = undefined /*out*/;
            inputs["offlineWipeTimeout"] = undefined /*out*/;
            inputs["pin"] = undefined /*out*/;
            inputs["pinNumRetry"] = undefined /*out*/;
            inputs["screenCapture"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:intune/latest:AndroidMAMPolicyByName" }, { type: "azurerm:intune/v20150114privatepreview:AndroidMAMPolicyByName" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AndroidMAMPolicyByName.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AndroidMAMPolicyByName resource.
 */
export interface AndroidMAMPolicyByNameArgs {
    readonly accessRecheckOfflineTimeout?: pulumi.Input<string>;
    readonly accessRecheckOnlineTimeout?: pulumi.Input<string>;
    readonly appSharingFromLevel?: pulumi.Input<string>;
    readonly appSharingToLevel?: pulumi.Input<string>;
    readonly authentication?: pulumi.Input<string>;
    readonly clipboardSharingLevel?: pulumi.Input<string>;
    readonly dataBackup?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly deviceCompliance?: pulumi.Input<string>;
    readonly fileEncryption?: pulumi.Input<string>;
    readonly fileSharingSaveAs?: pulumi.Input<string>;
    readonly friendlyName: pulumi.Input<string>;
    /**
     * Location hostName for the tenant
     */
    readonly hostName: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location?: pulumi.Input<string>;
    readonly managedBrowser?: pulumi.Input<string>;
    readonly offlineWipeTimeout?: pulumi.Input<string>;
    readonly pin?: pulumi.Input<string>;
    readonly pinNumRetry?: pulumi.Input<number>;
    /**
     * Unique name for the policy
     */
    readonly policyName: pulumi.Input<string>;
    readonly screenCapture?: pulumi.Input<string>;
    /**
     * Resource Tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
