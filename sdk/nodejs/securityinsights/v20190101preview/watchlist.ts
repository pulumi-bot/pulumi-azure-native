// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a Watchlist in Azure Security Insights.
 */
export class Watchlist extends pulumi.CustomResource {
    /**
     * Get an existing Watchlist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Watchlist {
        return new Watchlist(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:securityinsights/v20190101preview:Watchlist';

    /**
     * Returns true if the given object is an instance of Watchlist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Watchlist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Watchlist.__pulumiType;
    }

    /**
     * Describes a user that created the watchlist
     */
    public readonly createdBy!: pulumi.Output<outputs.securityinsights.v20190101preview.UserInfoResponse | undefined>;
    /**
     * The time the watchlist was created
     */
    public readonly createdTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * The default duration of a watchlist (in ISO 8601 duration format)
     */
    public readonly defaultDuration!: pulumi.Output<string | undefined>;
    /**
     * A description of the watchlist
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the watchlist
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * List of labels relevant to this watchlist
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The last time the watchlist was updated
     */
    public readonly lastUpdatedTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The notes of the watchlist
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The provider of the watchlist
     */
    public readonly provider!: pulumi.Output<string>;
    /**
     * The source of the watchlist
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The tenantId where the watchlist belongs to.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Describes a user that updated the watchlist
     */
    public readonly updatedBy!: pulumi.Output<outputs.securityinsights.v20190101preview.UserInfoResponse | undefined>;
    /**
     * List of watchlist items.
     */
    public readonly watchlistItems!: pulumi.Output<outputs.securityinsights.v20190101preview.WatchlistItemResponse[] | undefined>;
    /**
     * The type of the watchlist
     */
    public readonly watchlistType!: pulumi.Output<string | undefined>;
    /**
     * The workspaceId where the watchlist belongs to.
     */
    public readonly workspaceId!: pulumi.Output<string | undefined>;

    /**
     * Create a Watchlist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WatchlistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WatchlistArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WatchlistArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.operationalInsightsResourceProvider === undefined) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if (!args || args.provider === undefined) {
                throw new Error("Missing required property 'provider'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            if (!args || args.watchlistAlias === undefined) {
                throw new Error("Missing required property 'watchlistAlias'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["createdBy"] = args ? args.createdBy : undefined;
            inputs["createdTimeUtc"] = args ? args.createdTimeUtc : undefined;
            inputs["defaultDuration"] = args ? args.defaultDuration : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["lastUpdatedTimeUtc"] = args ? args.lastUpdatedTimeUtc : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["provider"] = args ? args.provider : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["updatedBy"] = args ? args.updatedBy : undefined;
            inputs["watchlistAlias"] = args ? args.watchlistAlias : undefined;
            inputs["watchlistItems"] = args ? args.watchlistItems : undefined;
            inputs["watchlistType"] = args ? args.watchlistType : undefined;
            inputs["workspaceId"] = args ? args.workspaceId : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Watchlist.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Watchlist resource.
 */
export interface WatchlistArgs {
    /**
     * Describes a user that created the watchlist
     */
    readonly createdBy?: pulumi.Input<inputs.securityinsights.v20190101preview.UserInfo>;
    /**
     * The time the watchlist was created
     */
    readonly createdTimeUtc?: pulumi.Input<string>;
    /**
     * The default duration of a watchlist (in ISO 8601 duration format)
     */
    readonly defaultDuration?: pulumi.Input<string>;
    /**
     * A description of the watchlist
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the watchlist
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * List of labels relevant to this watchlist
     */
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The last time the watchlist was updated
     */
    readonly lastUpdatedTimeUtc?: pulumi.Input<string>;
    /**
     * The notes of the watchlist
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * The provider of the watchlist
     */
    readonly provider: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The source of the watchlist
     */
    readonly source: pulumi.Input<string>;
    /**
     * The tenantId where the watchlist belongs to.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Describes a user that updated the watchlist
     */
    readonly updatedBy?: pulumi.Input<inputs.securityinsights.v20190101preview.UserInfo>;
    /**
     * Watchlist Alias
     */
    readonly watchlistAlias: pulumi.Input<string>;
    /**
     * List of watchlist items.
     */
    readonly watchlistItems?: pulumi.Input<pulumi.Input<inputs.securityinsights.v20190101preview.WatchlistItem>[]>;
    /**
     * The type of the watchlist
     */
    readonly watchlistType?: pulumi.Input<string>;
    /**
     * The workspaceId where the watchlist belongs to.
     */
    readonly workspaceId?: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
