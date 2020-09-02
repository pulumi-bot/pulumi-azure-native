// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * External OAuth authorization server settings.
 */
export class AuthorizationServer extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizationServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AuthorizationServer {
        return new AuthorizationServer(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20191201preview:AuthorizationServer';

    /**
     * Returns true if the given object is an instance of AuthorizationServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizationServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizationServer.__pulumiType;
    }

    /**
     * OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
     */
    public readonly authorizationEndpoint!: pulumi.Output<string>;
    /**
     * HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
     */
    public readonly authorizationMethods!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the mechanism by which access token is passed to the API. 
     */
    public readonly bearerTokenSendingMethods!: pulumi.Output<string[] | undefined>;
    /**
     * Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
     */
    public readonly clientAuthenticationMethod!: pulumi.Output<string[] | undefined>;
    /**
     * Client or app id registered with this authorization server.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
     */
    public readonly clientRegistrationEndpoint!: pulumi.Output<string>;
    /**
     * Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
     */
    public readonly defaultScope!: pulumi.Output<string | undefined>;
    /**
     * Description of the authorization server. Can contain HTML formatting tags.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User-friendly authorization server name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Form of an authorization grant, which the client uses to request the access token.
     */
    public readonly grantTypes!: pulumi.Output<string[]>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
     */
    public readonly resourceOwnerPassword!: pulumi.Output<string | undefined>;
    /**
     * Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
     */
    public readonly resourceOwnerUsername!: pulumi.Output<string | undefined>;
    /**
     * If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
     */
    public readonly supportState!: pulumi.Output<boolean | undefined>;
    /**
     * Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
     */
    public readonly tokenBodyParameters!: pulumi.Output<outputs.apimanagement.v20191201preview.TokenBodyParameterContractResponse[] | undefined>;
    /**
     * OAuth token endpoint. Contains absolute URI to entity being referenced.
     */
    public readonly tokenEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AuthorizationServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizationServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AuthorizationServerArgs | undefined;
            if (!args || args.authorizationEndpoint === undefined) {
                throw new Error("Missing required property 'authorizationEndpoint'");
            }
            if (!args || args.authsid === undefined) {
                throw new Error("Missing required property 'authsid'");
            }
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientRegistrationEndpoint === undefined) {
                throw new Error("Missing required property 'clientRegistrationEndpoint'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.grantTypes === undefined) {
                throw new Error("Missing required property 'grantTypes'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["authorizationEndpoint"] = args ? args.authorizationEndpoint : undefined;
            inputs["authorizationMethods"] = args ? args.authorizationMethods : undefined;
            inputs["authsid"] = args ? args.authsid : undefined;
            inputs["bearerTokenSendingMethods"] = args ? args.bearerTokenSendingMethods : undefined;
            inputs["clientAuthenticationMethod"] = args ? args.clientAuthenticationMethod : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientRegistrationEndpoint"] = args ? args.clientRegistrationEndpoint : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["defaultScope"] = args ? args.defaultScope : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["grantTypes"] = args ? args.grantTypes : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceOwnerPassword"] = args ? args.resourceOwnerPassword : undefined;
            inputs["resourceOwnerUsername"] = args ? args.resourceOwnerUsername : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["supportState"] = args ? args.supportState : undefined;
            inputs["tokenBodyParameters"] = args ? args.tokenBodyParameters : undefined;
            inputs["tokenEndpoint"] = args ? args.tokenEndpoint : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:AuthorizationServer" }, { type: "azurerm:apimanagement/v20160707:AuthorizationServer" }, { type: "azurerm:apimanagement/v20161010:AuthorizationServer" }, { type: "azurerm:apimanagement/v20170301:AuthorizationServer" }, { type: "azurerm:apimanagement/v20180101:AuthorizationServer" }, { type: "azurerm:apimanagement/v20180601preview:AuthorizationServer" }, { type: "azurerm:apimanagement/v20190101:AuthorizationServer" }, { type: "azurerm:apimanagement/v20191201:AuthorizationServer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AuthorizationServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AuthorizationServer resource.
 */
export interface AuthorizationServerArgs {
    /**
     * OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
     */
    readonly authorizationEndpoint: pulumi.Input<string>;
    /**
     * HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
     */
    readonly authorizationMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier of the authorization server.
     */
    readonly authsid: pulumi.Input<string>;
    /**
     * Specifies the mechanism by which access token is passed to the API. 
     */
    readonly bearerTokenSendingMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
     */
    readonly clientAuthenticationMethod?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Client or app id registered with this authorization server.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
     */
    readonly clientRegistrationEndpoint: pulumi.Input<string>;
    /**
     * Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
     */
    readonly defaultScope?: pulumi.Input<string>;
    /**
     * Description of the authorization server. Can contain HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * User-friendly authorization server name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Form of an authorization grant, which the client uses to request the access token.
     */
    readonly grantTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
     */
    readonly resourceOwnerPassword?: pulumi.Input<string>;
    /**
     * Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
     */
    readonly resourceOwnerUsername?: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
     */
    readonly supportState?: pulumi.Input<boolean>;
    /**
     * Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
     */
    readonly tokenBodyParameters?: pulumi.Input<pulumi.Input<inputs.apimanagement.v20191201preview.TokenBodyParameterContract>[]>;
    /**
     * OAuth token endpoint. Contains absolute URI to entity being referenced.
     */
    readonly tokenEndpoint?: pulumi.Input<string>;
}
