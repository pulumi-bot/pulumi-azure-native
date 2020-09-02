// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Contact details and configurations for notifications coming from Azure Security Center.
 */
export class SecurityContact extends pulumi.CustomResource {
    /**
     * Get an existing SecurityContact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecurityContact {
        return new SecurityContact(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20200101preview:SecurityContact';

    /**
     * Returns true if the given object is an instance of SecurityContact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityContact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityContact.__pulumiType;
    }

    /**
     * Defines whether to send email notifications about new security alerts
     */
    public /*out*/ readonly alertNotifications!: pulumi.Output<outputs.security.v20200101preview.SecurityContactPropertiesResponseAlertNotifications | undefined>;
    /**
     * List of email addresses which will get notifications from Azure Security Center by the configurations defined in this security contact.
     */
    public /*out*/ readonly emails!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Defines whether to send email notifications from Azure Security Center to persons with specific RBAC roles on the subscription.
     */
    public /*out*/ readonly notificationsByRole!: pulumi.Output<outputs.security.v20200101preview.SecurityContactPropertiesResponseNotificationsByRole | undefined>;
    /**
     * The security contact's phone number
     */
    public /*out*/ readonly phone!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SecurityContact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityContactArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityContactArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SecurityContactArgs | undefined;
            if (!args || args.securityContactName === undefined) {
                throw new Error("Missing required property 'securityContactName'");
            }
            inputs["securityContactName"] = args ? args.securityContactName : undefined;
            inputs["alertNotifications"] = undefined /*out*/;
            inputs["emails"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notificationsByRole"] = undefined /*out*/;
            inputs["phone"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:security/v20170801preview:SecurityContact" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SecurityContact.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecurityContact resource.
 */
export interface SecurityContactArgs {
    /**
     * Name of the security contact object
     */
    readonly securityContactName: pulumi.Input<string>;
}
