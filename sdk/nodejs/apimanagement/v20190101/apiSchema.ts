// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Schema Contract details.
 *
 * ## Example Usage
 * ### ApiManagementCreateApiSchema
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const apiSchema = new azurerm.apimanagement.v20190101.ApiSchema("apiSchema", {
 *     apiId: "59d6bb8f1f7fab13dc67ec9b",
 *     contentType: "application/vnd.ms-azure-apim.xsd+xml",
 *     resourceGroupName: "rg1",
 *     schemaId: "ec12520d-9d48-4e7b-8f39-698ca2ac63f1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 */
export class ApiSchema extends pulumi.CustomResource {
    /**
     * Get an existing ApiSchema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiSchema {
        return new ApiSchema(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20190101:ApiSchema';

    /**
     * Returns true if the given object is an instance of ApiSchema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiSchema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiSchema.__pulumiType;
    }

    /**
     * Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml). </br> - `Swagger` Schema use `application/vnd.ms-azure-apim.swagger.definitions+json` </br> - `WSDL` Schema use `application/vnd.ms-azure-apim.xsd+xml` </br> - `OpenApi` Schema use `application/vnd.oai.openapi.components+json` </br> - `WADL Schema` use `application/vnd.ms-azure-apim.wadl.grammars+xml`. 
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * Properties of the Schema Document.
     */
    public /*out*/ readonly document!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiSchema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiSchemaArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.contentType === undefined) {
                throw new Error("Missing required property 'contentType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.schemaId === undefined) {
                throw new Error("Missing required property 'schemaId'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["schemaId"] = args ? args.schemaId : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["document"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["contentType"] = undefined /*out*/;
            inputs["document"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiSchema" }, { type: "azurerm:apimanagement/v20170301:ApiSchema" }, { type: "azurerm:apimanagement/v20180101:ApiSchema" }, { type: "azurerm:apimanagement/v20180601preview:ApiSchema" }, { type: "azurerm:apimanagement/v20191201:ApiSchema" }, { type: "azurerm:apimanagement/v20191201preview:ApiSchema" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiSchema.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiSchema resource.
 */
export interface ApiSchemaArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml). </br> - `Swagger` Schema use `application/vnd.ms-azure-apim.swagger.definitions+json` </br> - `WSDL` Schema use `application/vnd.ms-azure-apim.xsd+xml` </br> - `OpenApi` Schema use `application/vnd.oai.openapi.components+json` </br> - `WADL Schema` use `application/vnd.ms-azure-apim.wadl.grammars+xml`.
     */
    readonly contentType: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Schema identifier within an API. Must be unique in the current API Management service instance.
     */
    readonly schemaId: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Json escaped string defining the document representing the Schema.
     */
    readonly value?: pulumi.Input<string>;
}
