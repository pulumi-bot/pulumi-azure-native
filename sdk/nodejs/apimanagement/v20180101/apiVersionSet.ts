// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Api Version Set Contract details.
 */
export class ApiVersionSet extends pulumi.CustomResource {
    /**
     * Get an existing ApiVersionSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiVersionSet {
        return new ApiVersionSet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20180101:ApiVersionSet';

    /**
     * Returns true if the given object is an instance of ApiVersionSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiVersionSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiVersionSet.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Api VersionSet contract properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.apimanagement.v20180101.ApiVersionSetContractPropertiesResponse>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiVersionSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiVersionSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiVersionSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApiVersionSetArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.versioningScheme === undefined) {
                throw new Error("Missing required property 'versioningScheme'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["versionHeaderName"] = args ? args.versionHeaderName : undefined;
            inputs["versionQueryName"] = args ? args.versionQueryName : undefined;
            inputs["versioningScheme"] = args ? args.versioningScheme : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApiVersionSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiVersionSet resource.
 */
export interface ApiVersionSetArgs {
    /**
     * Description of API Version Set.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of API Version Set
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Api Version Set identifier. Must be unique in the current API Management service instance.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
     */
    readonly versionHeaderName?: pulumi.Input<string>;
    /**
     * Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
     */
    readonly versionQueryName?: pulumi.Input<string>;
    /**
     * An value that determines where the API Version identifer will be located in a HTTP request.
     */
    readonly versioningScheme: pulumi.Input<string>;
}
