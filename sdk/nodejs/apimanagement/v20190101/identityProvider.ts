// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Identity Provider details.
 *
 * ## Example Usage
 * ### ApiManagementCreateIdentityProvider
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const identityProvider = new azurerm.apimanagement.v20190101.IdentityProvider("identityProvider", {
 *     clientId: "facebookid",
 *     clientSecret: "facebookapplicationsecret",
 *     identityProviderName: "facebook",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 */
export class IdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IdentityProvider {
        return new IdentityProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20190101:IdentityProvider';

    /**
     * Returns true if the given object is an instance of IdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProvider.__pulumiType;
    }

    /**
     * List of Allowed Tenants when configuring Azure Active Directory login.
     */
    public readonly allowedTenants!: pulumi.Output<string[] | undefined>;
    /**
     * OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
     */
    public readonly authority!: pulumi.Output<string | undefined>;
    /**
     * Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
     */
    public readonly passwordResetPolicyName!: pulumi.Output<string | undefined>;
    /**
     * Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
     */
    public readonly profileEditingPolicyName!: pulumi.Output<string | undefined>;
    /**
     * Signin Policy Name. Only applies to AAD B2C Identity Provider.
     */
    public readonly signinPolicyName!: pulumi.Output<string | undefined>;
    /**
     * The TenantId to use instead of Common when logging into Active Directory
     */
    public readonly signinTenant!: pulumi.Output<string | undefined>;
    /**
     * Signup Policy Name. Only applies to AAD B2C Identity Provider.
     */
    public readonly signupPolicyName!: pulumi.Output<string | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a IdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientSecret === undefined) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if (!args || args.identityProviderName === undefined) {
                throw new Error("Missing required property 'identityProviderName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["allowedTenants"] = args ? args.allowedTenants : undefined;
            inputs["authority"] = args ? args.authority : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["identityProviderName"] = args ? args.identityProviderName : undefined;
            inputs["passwordResetPolicyName"] = args ? args.passwordResetPolicyName : undefined;
            inputs["profileEditingPolicyName"] = args ? args.profileEditingPolicyName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["signinPolicyName"] = args ? args.signinPolicyName : undefined;
            inputs["signinTenant"] = args ? args.signinTenant : undefined;
            inputs["signupPolicyName"] = args ? args.signupPolicyName : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["name"] = undefined /*out*/;
        } else {
            inputs["allowedTenants"] = undefined /*out*/;
            inputs["authority"] = undefined /*out*/;
            inputs["clientId"] = undefined /*out*/;
            inputs["clientSecret"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["passwordResetPolicyName"] = undefined /*out*/;
            inputs["profileEditingPolicyName"] = undefined /*out*/;
            inputs["signinPolicyName"] = undefined /*out*/;
            inputs["signinTenant"] = undefined /*out*/;
            inputs["signupPolicyName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:IdentityProvider" }, { type: "azurerm:apimanagement/v20160707:IdentityProvider" }, { type: "azurerm:apimanagement/v20161010:IdentityProvider" }, { type: "azurerm:apimanagement/v20170301:IdentityProvider" }, { type: "azurerm:apimanagement/v20180101:IdentityProvider" }, { type: "azurerm:apimanagement/v20180601preview:IdentityProvider" }, { type: "azurerm:apimanagement/v20191201:IdentityProvider" }, { type: "azurerm:apimanagement/v20191201preview:IdentityProvider" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IdentityProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IdentityProvider resource.
 */
export interface IdentityProviderArgs {
    /**
     * List of Allowed Tenants when configuring Azure Active Directory login.
     */
    readonly allowedTenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
     */
    readonly authority?: pulumi.Input<string>;
    /**
     * Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft.
     */
    readonly clientSecret: pulumi.Input<string>;
    /**
     * Identity Provider Type identifier.
     */
    readonly identityProviderName: pulumi.Input<string>;
    /**
     * Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
     */
    readonly passwordResetPolicyName?: pulumi.Input<string>;
    /**
     * Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
     */
    readonly profileEditingPolicyName?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Signin Policy Name. Only applies to AAD B2C Identity Provider.
     */
    readonly signinPolicyName?: pulumi.Input<string>;
    /**
     * The TenantId to use instead of Common when logging into Active Directory
     */
    readonly signinTenant?: pulumi.Input<string>;
    /**
     * Signup Policy Name. Only applies to AAD B2C Identity Provider.
     */
    readonly signupPolicyName?: pulumi.Input<string>;
    /**
     * Identity Provider Type identifier.
     */
    readonly type?: pulumi.Input<string>;
}
