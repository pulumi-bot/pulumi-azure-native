// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Description of an API Management service resource.
 */
export class ApiManagementService extends pulumi.CustomResource {
    /**
     * Get an existing ApiManagementService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiManagementService {
        return new ApiManagementService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20160707:ApiManagementService';

    /**
     * Returns true if the given object is an instance of ApiManagementService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiManagementService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiManagementService.__pulumiType;
    }

    /**
     * ETag of the resource.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Datacenter location of the API Management service.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the API Management service.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the API Management service.
     */
    public readonly properties!: pulumi.Output<outputs.apimanagement.v20160707.ApiManagementServicePropertiesResponse>;
    /**
     * SKU properties of the API Management service.
     */
    public readonly sku!: pulumi.Output<outputs.apimanagement.v20160707.ApiManagementServiceSkuPropertiesResponse>;
    /**
     * API Management service tags. A maximum of 10 tags can be provided for a resource, and each tag must have a key no greater than 128 characters (and a value no greater than 256 characters).
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type of the API Management service.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiManagementService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApiManagementServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiManagementServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApiManagementServiceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApiManagementService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiManagementService resource.
 */
export interface ApiManagementServiceArgs {
    /**
     * ETag of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Datacenter location of the API Management service.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Properties of the API Management service.
     */
    readonly properties: pulumi.Input<inputs.apimanagement.v20160707.ApiManagementServiceProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SKU properties of the API Management service.
     */
    readonly sku: pulumi.Input<inputs.apimanagement.v20160707.ApiManagementServiceSkuProperties>;
    /**
     * API Management service tags. A maximum of 10 tags can be provided for a resource, and each tag must have a key no greater than 128 characters (and a value no greater than 256 characters).
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}