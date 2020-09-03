// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Tag Contract details.
 *
 * ## Example Usage
 * ### ApiManagementCreateApiOperationTag
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const tagByOperation = new azurerm.apimanagement.v20191201.TagByOperation("tagByOperation", {
 *     apiId: "5931a75ae4bbd512a88c680b",
 *     operationId: "5931a75ae4bbd512a88c680a",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 *     tagId: "tagId1",
 * });
 *
 * ```
 */
export class TagByOperation extends pulumi.CustomResource {
    /**
     * Get an existing TagByOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TagByOperation {
        return new TagByOperation(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20191201:TagByOperation';

    /**
     * Returns true if the given object is an instance of TagByOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagByOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagByOperation.__pulumiType;
    }

    /**
     * Tag name.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a TagByOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagByOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagByOperationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as TagByOperationArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.operationId === undefined) {
                throw new Error("Missing required property 'operationId'");
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
            inputs["operationId"] = args ? args.operationId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["tagId"] = args ? args.tagId : undefined;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:TagByOperation" }, { type: "azurerm:apimanagement/v20170301:TagByOperation" }, { type: "azurerm:apimanagement/v20180101:TagByOperation" }, { type: "azurerm:apimanagement/v20180601preview:TagByOperation" }, { type: "azurerm:apimanagement/v20190101:TagByOperation" }, { type: "azurerm:apimanagement/v20191201preview:TagByOperation" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(TagByOperation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagByOperation resource.
 */
export interface TagByOperationArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Operation identifier within an API. Must be unique in the current API Management service instance.
     */
    readonly operationId: pulumi.Input<string>;
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
