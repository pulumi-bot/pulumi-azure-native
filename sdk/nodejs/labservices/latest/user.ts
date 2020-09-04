// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The User registered to a lab
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
    public static readonly __pulumiType = 'azurerm:labservices/latest:User';

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
     * The user email address, as it was specified during registration.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The user family name, as it was specified during registration.
     */
    public /*out*/ readonly familyName!: pulumi.Output<string>;
    /**
     * The user given name, as it was specified during registration.
     */
    public /*out*/ readonly givenName!: pulumi.Output<string>;
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
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The user tenant ID, as it was specified during registration.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * How long the user has used his VMs in this lab
     */
    public /*out*/ readonly totalUsage!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public readonly uniqueIdentifier!: pulumi.Output<string | undefined>;

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
            if (!args || args.labAccountName === undefined) {
                throw new Error("Missing required property 'labAccountName'");
            }
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.userName === undefined) {
                throw new Error("Missing required property 'userName'");
            }
            inputs["labAccountName"] = args ? args.labAccountName : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["uniqueIdentifier"] = args ? args.uniqueIdentifier : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["email"] = undefined /*out*/;
            inputs["familyName"] = undefined /*out*/;
            inputs["givenName"] = undefined /*out*/;
            inputs["latestOperationResult"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["totalUsage"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["email"] = undefined /*out*/;
            inputs["familyName"] = undefined /*out*/;
            inputs["givenName"] = undefined /*out*/;
            inputs["latestOperationResult"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["totalUsage"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:labservices/v20181015:User" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
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
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    readonly userName: pulumi.Input<string>;
}
