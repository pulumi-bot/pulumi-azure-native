// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The top level storage insight resource container.
 *
 * ## Example Usage
 * ### StorageInsightsCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const storageInsightConfig = new azurerm.operationalinsights.latest.StorageInsightConfig("storageInsightConfig", {
 *     containers: ["wad-iis-logfiles"],
 *     resourceGroupName: "OIAutoRest5123",
 *     storageAccount: {
 *         id: "/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945",
 *         key: "1234",
 *     },
 *     storageInsightName: "AzTestSI1110",
 *     tables: [
 *         "WADWindowsEventLogsTable",
 *         "LinuxSyslogVer2v0",
 *     ],
 *     workspaceName: "aztest5048",
 * });
 *
 * ```
 */
export class StorageInsightConfig extends pulumi.CustomResource {
    /**
     * Get an existing StorageInsightConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageInsightConfig {
        return new StorageInsightConfig(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/latest:StorageInsightConfig';

    /**
     * Returns true if the given object is an instance of StorageInsightConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageInsightConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageInsightConfig.__pulumiType;
    }

    /**
     * The names of the blob containers that the workspace should read
     */
    public readonly containers!: pulumi.Output<string[] | undefined>;
    /**
     * The ETag of the storage insight.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The status of the storage insight
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.operationalinsights.latest.StorageInsightStatusResponse>;
    /**
     * The storage account connection details
     */
    public readonly storageAccount!: pulumi.Output<outputs.operationalinsights.latest.StorageAccountResponse>;
    /**
     * The names of the Azure tables that the workspace should read
     */
    public readonly tables!: pulumi.Output<string[] | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StorageInsightConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageInsightConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageInsightConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as StorageInsightConfigArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageAccount === undefined) {
                throw new Error("Missing required property 'storageAccount'");
            }
            if (!args || args.storageInsightName === undefined) {
                throw new Error("Missing required property 'storageInsightName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["containers"] = args ? args.containers : undefined;
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccount"] = args ? args.storageAccount : undefined;
            inputs["storageInsightName"] = args ? args.storageInsightName : undefined;
            inputs["tables"] = args ? args.tables : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/v20150320:StorageInsightConfig" }, { type: "azurerm:operationalinsights/v20200301preview:StorageInsightConfig" }, { type: "azurerm:operationalinsights/v20200801:StorageInsightConfig" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(StorageInsightConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageInsightConfig resource.
 */
export interface StorageInsightConfigArgs {
    /**
     * The names of the blob containers that the workspace should read
     */
    readonly containers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ETag of the storage insight.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The storage account connection details
     */
    readonly storageAccount: pulumi.Input<inputs.operationalinsights.latest.StorageAccount>;
    /**
     * Name of the storageInsightsConfigs resource
     */
    readonly storageInsightName: pulumi.Input<string>;
    /**
     * The names of the Azure tables that the workspace should read
     */
    readonly tables?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
