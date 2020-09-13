// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Represents a relation between two resources
 */
export class BookmarkRelation extends pulumi.CustomResource {
    /**
     * Get an existing BookmarkRelation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BookmarkRelation {
        return new BookmarkRelation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:securityinsights/v20190101preview:BookmarkRelation';

    /**
     * Returns true if the given object is an instance of BookmarkRelation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BookmarkRelation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BookmarkRelation.__pulumiType;
    }

    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource ID of the related resource
     */
    public readonly relatedResourceId!: pulumi.Output<string>;
    /**
     * The resource kind of the related resource
     */
    public /*out*/ readonly relatedResourceKind!: pulumi.Output<string>;
    /**
     * The name of the related resource
     */
    public /*out*/ readonly relatedResourceName!: pulumi.Output<string>;
    /**
     * The resource type of the related resource
     */
    public /*out*/ readonly relatedResourceType!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BookmarkRelation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BookmarkRelationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.bookmarkId === undefined) {
                throw new Error("Missing required property 'bookmarkId'");
            }
            if (!args || args.operationalInsightsResourceProvider === undefined) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if (!args || args.relatedResourceId === undefined) {
                throw new Error("Missing required property 'relatedResourceId'");
            }
            if (!args || args.relationName === undefined) {
                throw new Error("Missing required property 'relationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["bookmarkId"] = args ? args.bookmarkId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["relatedResourceId"] = args ? args.relatedResourceId : undefined;
            inputs["relationName"] = args ? args.relationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["relatedResourceKind"] = undefined /*out*/;
            inputs["relatedResourceName"] = undefined /*out*/;
            inputs["relatedResourceType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["relatedResourceId"] = undefined /*out*/;
            inputs["relatedResourceKind"] = undefined /*out*/;
            inputs["relatedResourceName"] = undefined /*out*/;
            inputs["relatedResourceType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BookmarkRelation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BookmarkRelation resource.
 */
export interface BookmarkRelationArgs {
    /**
     * Bookmark ID
     */
    readonly bookmarkId: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * The resource ID of the related resource
     */
    readonly relatedResourceId: pulumi.Input<string>;
    /**
     * Relation Name
     */
    readonly relationName: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
