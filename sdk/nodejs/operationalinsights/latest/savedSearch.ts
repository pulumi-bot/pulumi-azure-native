// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Value object for saved search results.
 */
export class SavedSearch extends pulumi.CustomResource {
    /**
     * Get an existing SavedSearch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SavedSearch {
        return new SavedSearch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:operationalinsights/latest:SavedSearch';

    /**
     * Returns true if the given object is an instance of SavedSearch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SavedSearch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SavedSearch.__pulumiType;
    }

    /**
     * The category of the saved search. This helps the user to find a saved search faster. 
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * Saved search display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The ETag of the saved search.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The function alias if query serves as a function.
     */
    public readonly functionAlias!: pulumi.Output<string | undefined>;
    /**
     * The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
     */
    public readonly functionParameters!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The query expression for the saved search.
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * The tags attached to the saved search.
     */
    public readonly tags!: pulumi.Output<outputs.operationalinsights.latest.TagResponse[] | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The version number of the query language. The current version is 2 and is the default.
     */
    public readonly version!: pulumi.Output<number | undefined>;

    /**
     * Create a SavedSearch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SavedSearchArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.category === undefined) {
                throw new Error("Missing required property 'category'");
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
            if (!args || args.savedSearchId === undefined) {
                throw new Error("Missing required property 'savedSearchId'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["category"] = args ? args.category : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["functionAlias"] = args ? args.functionAlias : undefined;
            inputs["functionParameters"] = args ? args.functionParameters : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["savedSearchId"] = args ? args.savedSearchId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["category"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["functionAlias"] = undefined /*out*/;
            inputs["functionParameters"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["query"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:operationalinsights/v20150320:SavedSearch" }, { type: "azure-nextgen:operationalinsights/v20200301preview:SavedSearch" }, { type: "azure-nextgen:operationalinsights/v20200801:SavedSearch" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SavedSearch.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SavedSearch resource.
 */
export interface SavedSearchArgs {
    /**
     * The category of the saved search. This helps the user to find a saved search faster. 
     */
    readonly category: pulumi.Input<string>;
    /**
     * Saved search display name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The ETag of the saved search.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The function alias if query serves as a function.
     */
    readonly functionAlias?: pulumi.Input<string>;
    /**
     * The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
     */
    readonly functionParameters?: pulumi.Input<string>;
    /**
     * The query expression for the saved search.
     */
    readonly query: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The id of the saved search.
     */
    readonly savedSearchId: pulumi.Input<string>;
    /**
     * The tags attached to the saved search.
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.operationalinsights.latest.Tag>[]>;
    /**
     * The version number of the query language. The current version is 2 and is the default.
     */
    readonly version?: pulumi.Input<number>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
