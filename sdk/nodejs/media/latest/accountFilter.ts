// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Account Filter.
 *
 * ## Example Usage
 * ### Create an Account Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const accountFilter = new azurerm.media.latest.AccountFilter("accountFilter", {
 *     accountName: "contosomedia",
 *     filterName: "newAccountFilter",
 *     firstQuality: {
 *         bitrate: 128000,
 *     },
 *     presentationTimeRange: {
 *         endTimestamp: 170000000,
 *         forceEndTimestamp: false,
 *         liveBackoffDuration: 0,
 *         presentationWindowDuration: 9.223372036854776e+18,
 *         startTimestamp: 0,
 *         timescale: 10000000,
 *     },
 *     resourceGroupName: "contoso",
 *     tracks: [
 *         {
 *             trackSelections: [
 *                 {
 *                     operation: "Equal",
 *                     property: "Type",
 *                     value: "Audio",
 *                 },
 *                 {
 *                     operation: "NotEqual",
 *                     property: "Language",
 *                     value: "en",
 *                 },
 *                 {
 *                     operation: "NotEqual",
 *                     property: "FourCC",
 *                     value: "EC-3",
 *                 },
 *             ],
 *         },
 *         {
 *             trackSelections: [
 *                 {
 *                     operation: "Equal",
 *                     property: "Type",
 *                     value: "Video",
 *                 },
 *                 {
 *                     operation: "Equal",
 *                     property: "Bitrate",
 *                     value: "3000000-5000000",
 *                 },
 *             ],
 *         },
 *     ],
 * });
 *
 * ```
 */
export class AccountFilter extends pulumi.CustomResource {
    /**
     * Get an existing AccountFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccountFilter {
        return new AccountFilter(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/latest:AccountFilter';

    /**
     * Returns true if the given object is an instance of AccountFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountFilter.__pulumiType;
    }

    /**
     * The first quality.
     */
    public readonly firstQuality!: pulumi.Output<outputs.media.latest.FirstQualityResponse | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The presentation time range.
     */
    public readonly presentationTimeRange!: pulumi.Output<outputs.media.latest.PresentationTimeRangeResponse | undefined>;
    /**
     * The tracks selection conditions.
     */
    public readonly tracks!: pulumi.Output<outputs.media.latest.FilterTrackSelectionResponse[] | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AccountFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountFilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountFilterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AccountFilterArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.filterName === undefined) {
                throw new Error("Missing required property 'filterName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["filterName"] = args ? args.filterName : undefined;
            inputs["firstQuality"] = args ? args.firstQuality : undefined;
            inputs["presentationTimeRange"] = args ? args.presentationTimeRange : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tracks"] = args ? args.tracks : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:media/v20180701:AccountFilter" }, { type: "azurerm:media/v20200501:AccountFilter" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AccountFilter.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccountFilter resource.
 */
export interface AccountFilterArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The Account Filter name
     */
    readonly filterName: pulumi.Input<string>;
    /**
     * The first quality.
     */
    readonly firstQuality?: pulumi.Input<inputs.media.latest.FirstQuality>;
    /**
     * The presentation time range.
     */
    readonly presentationTimeRange?: pulumi.Input<inputs.media.latest.PresentationTimeRange>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tracks selection conditions.
     */
    readonly tracks?: pulumi.Input<pulumi.Input<inputs.media.latest.FilterTrackSelection>[]>;
}
