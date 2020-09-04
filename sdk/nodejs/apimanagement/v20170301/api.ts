// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * API details.
 *
 * ## Example Usage
 * ### ApiManagementCreateApi
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const api = new azurerm.apimanagement.v20170301.Api("api", {
 *     apiId: "tempgroup",
 *     authenticationSettings: {
 *         oAuth2: {
 *             authorizationServerId: "authorizationServerId2283",
 *             scope: "oauth2scope2580",
 *         },
 *     },
 *     description: "apidescription5200",
 *     displayName: "apiname1463",
 *     path: "newapiPath",
 *     protocols: [
 *         "https",
 *         "http",
 *     ],
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 *     serviceUrl: "http://newechoapi.cloudapp.net/api",
 *     subscriptionKeyParameterNames: {
 *         header: "header4520",
 *         query: "query3037",
 *     },
 * });
 *
 * ```
 * ### ApiManagementCreateApiRevision
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const api = new azurerm.apimanagement.v20170301.Api("api", {
 *     apiId: "5a838fd48f33670ed070d77c;rev=4",
 *     description: "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.",
 *     displayName: "Swagger Petstore V2",
 *     path: "petstore2",
 *     protocols: ["https"],
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 *     serviceUrl: "http://petstore.swagger.io/v4",
 *     subscriptionKeyParameterNames: {
 *         header: "Ocp-Apim-Subscription-Key",
 *         query: "subscription-key",
 *     },
 * });
 *
 * ```
 * ### ApiManagementCreateApiUsingSwaggerImport
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const api = new azurerm.apimanagement.v20170301.Api("api", {
 *     apiId: "petstore",
 *     contentFormat: "swagger-link-json",
 *     contentValue: "http://petstore.swagger.io/v2/swagger.json",
 *     path: "petstore",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 * ### ApiManagementCreateApiUsingWadlImport
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const api = new azurerm.apimanagement.v20170301.Api("api", {
 *     apiId: "petstore",
 *     contentFormat: "wadl-link-json",
 *     contentValue: "https://developer.cisco.com/media/wae-release-6-2-api-reference/wae-collector-rest-api/application.wadl",
 *     path: "collector",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 * ### ApiManagementCreateApiUsingWsdlImport
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const api = new azurerm.apimanagement.v20170301.Api("api", {
 *     apiId: "soapApi",
 *     contentFormat: "wsdl-link",
 *     contentValue: "http://www.webservicex.net/CurrencyConvertor.asmx?WSDL",
 *     path: "currency",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 *     wsdlSelector: {
 *         wsdlEndpointName: "CurrencyConvertorSoap",
 *         wsdlServiceName: "CurrencyConvertor",
 *     },
 * });
 *
 * ```
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20170301:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * Describes the Revision of the Api. If no value is provided, default revision 1 is created
     */
    public readonly apiRevision!: pulumi.Output<string | undefined>;
    /**
     * Type of API.
     */
    public readonly apiType!: pulumi.Output<string | undefined>;
    /**
     * Indicates the Version identifier of the API if the API is versioned
     */
    public readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * Api Version Set Contract details.
     */
    public readonly apiVersionSet!: pulumi.Output<outputs.apimanagement.v20170301.ApiVersionSetContractResponse | undefined>;
    /**
     * A resource identifier for the related ApiVersionSet.
     */
    public readonly apiVersionSetId!: pulumi.Output<string | undefined>;
    /**
     * Collection of authentication settings included into this API.
     */
    public readonly authenticationSettings!: pulumi.Output<outputs.apimanagement.v20170301.AuthenticationSettingsContractResponse | undefined>;
    /**
     * Description of the API. May include HTML formatting tags.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * API name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Indicates if API revision is current api revision.
     */
    public /*out*/ readonly isCurrent!: pulumi.Output<boolean>;
    /**
     * Indicates if API revision is accessible via the gateway.
     */
    public /*out*/ readonly isOnline!: pulumi.Output<boolean>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Describes on which protocols the operations in this API can be invoked.
     */
    public readonly protocols!: pulumi.Output<string[] | undefined>;
    /**
     * Absolute URL of the backend service implementing this API.
     */
    public readonly serviceUrl!: pulumi.Output<string | undefined>;
    /**
     * Protocols over which API is made available.
     */
    public readonly subscriptionKeyParameterNames!: pulumi.Output<outputs.apimanagement.v20170301.SubscriptionKeyParameterNamesContractResponse | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["apiRevision"] = args ? args.apiRevision : undefined;
            inputs["apiType"] = args ? args.apiType : undefined;
            inputs["apiVersion"] = args ? args.apiVersion : undefined;
            inputs["apiVersionSet"] = args ? args.apiVersionSet : undefined;
            inputs["apiVersionSetId"] = args ? args.apiVersionSetId : undefined;
            inputs["authenticationSettings"] = args ? args.authenticationSettings : undefined;
            inputs["contentFormat"] = args ? args.contentFormat : undefined;
            inputs["contentValue"] = args ? args.contentValue : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["protocols"] = args ? args.protocols : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["serviceUrl"] = args ? args.serviceUrl : undefined;
            inputs["subscriptionKeyParameterNames"] = args ? args.subscriptionKeyParameterNames : undefined;
            inputs["wsdlSelector"] = args ? args.wsdlSelector : undefined;
            inputs["isCurrent"] = undefined /*out*/;
            inputs["isOnline"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["apiRevision"] = undefined /*out*/;
            inputs["apiType"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["apiVersionSet"] = undefined /*out*/;
            inputs["apiVersionSetId"] = undefined /*out*/;
            inputs["authenticationSettings"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["isCurrent"] = undefined /*out*/;
            inputs["isOnline"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["path"] = undefined /*out*/;
            inputs["protocols"] = undefined /*out*/;
            inputs["serviceUrl"] = undefined /*out*/;
            inputs["subscriptionKeyParameterNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:Api" }, { type: "azurerm:apimanagement/v20160707:Api" }, { type: "azurerm:apimanagement/v20161010:Api" }, { type: "azurerm:apimanagement/v20180101:Api" }, { type: "azurerm:apimanagement/v20180601preview:Api" }, { type: "azurerm:apimanagement/v20190101:Api" }, { type: "azurerm:apimanagement/v20191201:Api" }, { type: "azurerm:apimanagement/v20191201preview:Api" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Api.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Describes the Revision of the Api. If no value is provided, default revision 1 is created
     */
    readonly apiRevision?: pulumi.Input<string>;
    /**
     * Type of API.
     */
    readonly apiType?: pulumi.Input<string>;
    /**
     * Indicates the Version identifier of the API if the API is versioned
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Api Version Set Contract details.
     */
    readonly apiVersionSet?: pulumi.Input<inputs.apimanagement.v20170301.ApiVersionSetContract>;
    /**
     * A resource identifier for the related ApiVersionSet.
     */
    readonly apiVersionSetId?: pulumi.Input<string>;
    /**
     * Collection of authentication settings included into this API.
     */
    readonly authenticationSettings?: pulumi.Input<inputs.apimanagement.v20170301.AuthenticationSettingsContract>;
    /**
     * Format of the Content in which the API is getting imported.
     */
    readonly contentFormat?: pulumi.Input<string>;
    /**
     * Content value when Importing an API.
     */
    readonly contentValue?: pulumi.Input<string>;
    /**
     * Description of the API. May include HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * API name.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
     */
    readonly path: pulumi.Input<string>;
    /**
     * Describes on which protocols the operations in this API can be invoked.
     */
    readonly protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Absolute URL of the backend service implementing this API.
     */
    readonly serviceUrl?: pulumi.Input<string>;
    /**
     * Protocols over which API is made available.
     */
    readonly subscriptionKeyParameterNames?: pulumi.Input<inputs.apimanagement.v20170301.SubscriptionKeyParameterNamesContract>;
    /**
     * Criteria to limit import of WSDL to a subset of the document.
     */
    readonly wsdlSelector?: pulumi.Input<inputs.apimanagement.v20170301.ApiCreateOrUpdatePropertiesWsdlSelector>;
}
