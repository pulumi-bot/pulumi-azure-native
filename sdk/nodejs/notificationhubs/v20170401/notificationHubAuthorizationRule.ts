// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Description of a Namespace AuthorizationRules.
 */
export class NotificationHubAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing NotificationHubAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NotificationHubAuthorizationRule {
        return new NotificationHubAuthorizationRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:notificationhubs/v20170401:NotificationHubAuthorizationRule';

    /**
     * Returns true if the given object is an instance of NotificationHubAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationHubAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationHubAuthorizationRule.__pulumiType;
    }

    /**
     * A string that describes the claim type
     */
    public /*out*/ readonly claimType!: pulumi.Output<string>;
    /**
     * A string that describes the claim value
     */
    public /*out*/ readonly claimValue!: pulumi.Output<string>;
    /**
     * The created time for this rule
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * A string that describes the authorization rule.
     */
    public /*out*/ readonly keyName!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * The last modified time for this rule
     */
    public /*out*/ readonly modifiedTime!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * The revision number for the rule
     */
    public /*out*/ readonly revision!: pulumi.Output<number>;
    /**
     * The rights associated with the rule.
     */
    public /*out*/ readonly rights!: pulumi.Output<AccessRights[] | undefined>;
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * The sku of the created namespace
     */
    public /*out*/ readonly sku!: pulumi.Output<outputs.notificationhubs.v20170401.SkuResponse | undefined>;
    /**
     * Resource tags
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NotificationHubAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationHubAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationHubAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as NotificationHubAuthorizationRuleArgs | undefined;
            if (!args || args.authorizationRuleName === undefined) {
                throw new Error("Missing required property 'authorizationRuleName'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.notificationHubName === undefined) {
                throw new Error("Missing required property 'notificationHubName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["authorizationRuleName"] = args ? args.authorizationRuleName : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["notificationHubName"] = args ? args.notificationHubName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["claimType"] = undefined /*out*/;
            inputs["claimValue"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["keyName"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["modifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["primaryKey"] = undefined /*out*/;
            inputs["revision"] = undefined /*out*/;
            inputs["rights"] = undefined /*out*/;
            inputs["secondaryKey"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:notificationhubs/v20140901:NotificationHubAuthorizationRule" }, { type: "azurerm:notificationhubs/v20160301:NotificationHubAuthorizationRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NotificationHubAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NotificationHubAuthorizationRule resource.
 */
export interface NotificationHubAuthorizationRuleArgs {
    /**
     * Authorization Rule Name.
     */
    readonly authorizationRuleName: pulumi.Input<string>;
    /**
     * The namespace name.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The notification hub name.
     */
    readonly notificationHubName: pulumi.Input<string>;
    /**
     * Properties of the Namespace AuthorizationRules.
     */
    readonly properties: pulumi.Input<inputs.notificationhubs.v20170401.SharedAccessAuthorizationRuleProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
