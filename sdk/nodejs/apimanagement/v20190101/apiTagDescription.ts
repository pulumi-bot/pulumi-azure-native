// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Contract details.
 */
export class ApiTagDescription extends pulumi.CustomResource {
    /**
     * Get an existing ApiTagDescription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiTagDescription {
        return new ApiTagDescription(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20190101:ApiTagDescription';

    /**
     * Returns true if the given object is an instance of ApiTagDescription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiTagDescription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiTagDescription.__pulumiType;
    }

    /**
     * Description of the Tag.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Tag name.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Description of the external resources describing the tag.
     */
    public readonly externalDocsDescription!: pulumi.Output<string | undefined>;
    /**
     * Absolute URL of external resources describing the tag.
     */
    public readonly externalDocsUrl!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiTagDescription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiTagDescriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.tagId === undefined) {
                throw new Error("Missing required property 'tagId'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["externalDocsDescription"] = args ? args.externalDocsDescription : undefined;
            inputs["externalDocsUrl"] = args ? args.externalDocsUrl : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["tagId"] = args ? args.tagId : undefined;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["externalDocsDescription"] = undefined /*out*/;
            inputs["externalDocsUrl"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiTagDescription" }, { type: "azurerm:apimanagement/v20170301:ApiTagDescription" }, { type: "azurerm:apimanagement/v20180101:ApiTagDescription" }, { type: "azurerm:apimanagement/v20180601preview:ApiTagDescription" }, { type: "azurerm:apimanagement/v20191201:ApiTagDescription" }, { type: "azurerm:apimanagement/v20191201preview:ApiTagDescription" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiTagDescription.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiTagDescription resource.
 */
export interface ApiTagDescriptionArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Description of the Tag.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Description of the external resources describing the tag.
     */
    readonly externalDocsDescription?: pulumi.Input<string>;
    /**
     * Absolute URL of external resources describing the tag.
     */
    readonly externalDocsUrl?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Tag identifier. Must be unique in the current API Management service instance.
     */
    readonly tagId: pulumi.Input<string>;
}
