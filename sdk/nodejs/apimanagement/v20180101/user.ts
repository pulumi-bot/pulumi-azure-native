// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * User details.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20180101:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Email address.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * First name.
     */
    public readonly firstName!: pulumi.Output<string | undefined>;
    /**
     * Collection of groups user is part of.
     */
    public /*out*/ readonly groups!: pulumi.Output<outputs.apimanagement.v20180101.GroupContractPropertiesResponse[]>;
    /**
     * Collection of user identities.
     */
    public readonly identities!: pulumi.Output<outputs.apimanagement.v20180101.UserIdentityContractResponse[] | undefined>;
    /**
     * Last name.
     */
    public readonly lastName!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional note about a user set by the administrator.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * Date of user registration. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    public /*out*/ readonly registrationDate!: pulumi.Output<string | undefined>;
    /**
     * Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.email === undefined) {
                throw new Error("Missing required property 'email'");
            }
            if (!args || args.firstName === undefined) {
                throw new Error("Missing required property 'firstName'");
            }
            if (!args || args.lastName === undefined) {
                throw new Error("Missing required property 'lastName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.uid === undefined) {
                throw new Error("Missing required property 'uid'");
            }
            inputs["confirmation"] = args ? args.confirmation : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["firstName"] = args ? args.firstName : undefined;
            inputs["identities"] = args ? args.identities : undefined;
            inputs["lastName"] = args ? args.lastName : undefined;
            inputs["note"] = args ? args.note : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["uid"] = args ? args.uid : undefined;
            inputs["groups"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["registrationDate"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["email"] = undefined /*out*/;
            inputs["firstName"] = undefined /*out*/;
            inputs["groups"] = undefined /*out*/;
            inputs["identities"] = undefined /*out*/;
            inputs["lastName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["note"] = undefined /*out*/;
            inputs["registrationDate"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:User" }, { type: "azurerm:apimanagement/v20160707:User" }, { type: "azurerm:apimanagement/v20161010:User" }, { type: "azurerm:apimanagement/v20170301:User" }, { type: "azurerm:apimanagement/v20180601preview:User" }, { type: "azurerm:apimanagement/v20190101:User" }, { type: "azurerm:apimanagement/v20191201:User" }, { type: "azurerm:apimanagement/v20191201preview:User" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Determines the type of confirmation e-mail that will be sent to the newly created user.
     */
    readonly confirmation?: pulumi.Input<string>;
    /**
     * Email address. Must not be empty and must be unique within the service instance.
     */
    readonly email: pulumi.Input<string>;
    /**
     * First name.
     */
    readonly firstName: pulumi.Input<string>;
    /**
     * Collection of user identities.
     */
    readonly identities?: pulumi.Input<pulumi.Input<inputs.apimanagement.v20180101.UserIdentityContract>[]>;
    /**
     * Last name.
     */
    readonly lastName: pulumi.Input<string>;
    /**
     * Optional note about a user set by the administrator.
     */
    readonly note?: pulumi.Input<string>;
    /**
     * User Password. If no value is provided, a default password is generated.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * User identifier. Must be unique in the current API Management service instance.
     */
    readonly uid: pulumi.Input<string>;
}
