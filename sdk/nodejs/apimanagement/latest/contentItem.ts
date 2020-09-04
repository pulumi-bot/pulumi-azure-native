// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Content type contract details.
 *
 * ## Example Usage
 * ### ApiManagementCreateContentItem
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const contentItem = new azurerm.apimanagement.latest.ContentItem("contentItem", {
 *     contentItemId: "4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8",
 *     contentTypeId: "page",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 */
export class ContentItem extends pulumi.CustomResource {
    /**
     * Get an existing ContentItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContentItem {
        return new ContentItem(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/latest:ContentItem';

    /**
     * Returns true if the given object is an instance of ContentItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentItem.__pulumiType;
    }

    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ContentItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContentItemArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.contentItemId === undefined) {
                throw new Error("Missing required property 'contentItemId'");
            }
            if (!args || args.contentTypeId === undefined) {
                throw new Error("Missing required property 'contentTypeId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["contentItemId"] = args ? args.contentItemId : undefined;
            inputs["contentTypeId"] = args ? args.contentTypeId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/v20191201:ContentItem" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ContentItem.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContentItem resource.
 */
export interface ContentItemArgs {
    /**
     * Content item identifier.
     */
    readonly contentItemId: pulumi.Input<string>;
    /**
     * Content type identifier.
     */
    readonly contentTypeId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
