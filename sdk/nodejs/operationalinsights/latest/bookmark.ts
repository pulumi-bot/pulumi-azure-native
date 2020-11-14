// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents a bookmark in Azure Security Insights.
 */
export class Bookmark extends pulumi.CustomResource {
    /**
     * Get an existing Bookmark resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Bookmark {
        return new Bookmark(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:operationalinsights/latest:Bookmark';

    /**
     * Returns true if the given object is an instance of Bookmark.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bookmark {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bookmark.__pulumiType;
    }

    /**
     * The time the bookmark was created
     */
    public readonly created!: pulumi.Output<string | undefined>;
    /**
     * Describes a user that created the bookmark
     */
    public readonly createdBy!: pulumi.Output<outputs.operationalinsights.latest.UserInfoResponse | undefined>;
    /**
     * The display name of the bookmark
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Describes an incident that relates to bookmark
     */
    public readonly incidentInfo!: pulumi.Output<outputs.operationalinsights.latest.IncidentInfoResponse | undefined>;
    /**
     * List of labels relevant to this bookmark
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The notes of the bookmark
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The query of the bookmark.
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * The query result of the bookmark.
     */
    public readonly queryResult!: pulumi.Output<string | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The last time the bookmark was updated
     */
    public readonly updated!: pulumi.Output<string | undefined>;
    /**
     * Describes a user that updated the bookmark
     */
    public readonly updatedBy!: pulumi.Output<outputs.operationalinsights.latest.UserInfoResponse | undefined>;

    /**
     * Create a Bookmark resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BookmarkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.bookmarkId === undefined) {
                throw new Error("Missing required property 'bookmarkId'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.query === undefined) {
                throw new Error("Missing required property 'query'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["bookmarkId"] = args ? args.bookmarkId : undefined;
            inputs["created"] = args ? args.created : undefined;
            inputs["createdBy"] = args ? args.createdBy : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["incidentInfo"] = args ? args.incidentInfo : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["queryResult"] = args ? args.queryResult : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["updated"] = args ? args.updated : undefined;
            inputs["updatedBy"] = args ? args.updatedBy : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["created"] = undefined /*out*/;
            inputs["createdBy"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["incidentInfo"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notes"] = undefined /*out*/;
            inputs["query"] = undefined /*out*/;
            inputs["queryResult"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updated"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:operationalinsights/v20200101:Bookmark" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Bookmark.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Bookmark resource.
 */
export interface BookmarkArgs {
    /**
     * Bookmark ID
     */
    readonly bookmarkId: pulumi.Input<string>;
    /**
     * The time the bookmark was created
     */
    readonly created?: pulumi.Input<string>;
    /**
     * Describes a user that created the bookmark
     */
    readonly createdBy?: pulumi.Input<inputs.operationalinsights.latest.UserInfo>;
    /**
     * The display name of the bookmark
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Describes an incident that relates to bookmark
     */
    readonly incidentInfo?: pulumi.Input<inputs.operationalinsights.latest.IncidentInfo>;
    /**
     * List of labels relevant to this bookmark
     */
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The notes of the bookmark
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * The query of the bookmark.
     */
    readonly query: pulumi.Input<string>;
    /**
     * The query result of the bookmark.
     */
    readonly queryResult?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The last time the bookmark was updated
     */
    readonly updated?: pulumi.Input<string>;
    /**
     * Describes a user that updated the bookmark
     */
    readonly updatedBy?: pulumi.Input<inputs.operationalinsights.latest.UserInfo>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
