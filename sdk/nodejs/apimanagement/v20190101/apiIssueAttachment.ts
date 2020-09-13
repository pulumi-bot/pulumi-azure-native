// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Issue Attachment Contract details.
 */
export class ApiIssueAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ApiIssueAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiIssueAttachment {
        return new ApiIssueAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20190101:ApiIssueAttachment';

    /**
     * Returns true if the given object is an instance of ApiIssueAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiIssueAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiIssueAttachment.__pulumiType;
    }

    /**
     * An HTTP link or Base64-encoded binary data.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
     */
    public readonly contentFormat!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Filename by which the binary data will be saved.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiIssueAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiIssueAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.attachmentId === undefined) {
                throw new Error("Missing required property 'attachmentId'");
            }
            if (!args || args.content === undefined) {
                throw new Error("Missing required property 'content'");
            }
            if (!args || args.contentFormat === undefined) {
                throw new Error("Missing required property 'contentFormat'");
            }
            if (!args || args.issueId === undefined) {
                throw new Error("Missing required property 'issueId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["attachmentId"] = args ? args.attachmentId : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["contentFormat"] = args ? args.contentFormat : undefined;
            inputs["issueId"] = args ? args.issueId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["content"] = undefined /*out*/;
            inputs["contentFormat"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiIssueAttachment" }, { type: "azurerm:apimanagement/v20170301:ApiIssueAttachment" }, { type: "azurerm:apimanagement/v20180101:ApiIssueAttachment" }, { type: "azurerm:apimanagement/v20180601preview:ApiIssueAttachment" }, { type: "azurerm:apimanagement/v20191201:ApiIssueAttachment" }, { type: "azurerm:apimanagement/v20191201preview:ApiIssueAttachment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiIssueAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiIssueAttachment resource.
 */
export interface ApiIssueAttachmentArgs {
    /**
     * API identifier. Must be unique in the current API Management service instance.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Attachment identifier within an Issue. Must be unique in the current Issue.
     */
    readonly attachmentId: pulumi.Input<string>;
    /**
     * An HTTP link or Base64-encoded binary data.
     */
    readonly content: pulumi.Input<string>;
    /**
     * Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
     */
    readonly contentFormat: pulumi.Input<string>;
    /**
     * Issue identifier. Must be unique in the current API Management service instance.
     */
    readonly issueId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Filename by which the binary data will be saved.
     */
    readonly title: pulumi.Input<string>;
}
