// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The top level Workspace resource container.
 *
 * ## Example Usage
 * ### WorkspacesCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const workspace = new azurerm.operationalinsights.v20200801.Workspace("workspace", {
 *     location: "australiasoutheast",
 *     resourceGroupName: "oiautorest6685",
 *     retentionInDays: 30,
 *     sku: {
 *         name: "PerGB2018",
 *     },
 *     tags: {
 *         tag1: "val1",
 *     },
 *     workspaceName: "oiautorest6685",
 * });
 *
 * ```
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/v20200801:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * This is a read-only property. Represents the ID associated with the workspace.
     */
    public /*out*/ readonly customerId!: pulumi.Output<string>;
    /**
     * The ETag of the workspace.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of linked private link scope resources.
     */
    public /*out*/ readonly privateLinkScopedResources!: pulumi.Output<outputs.operationalinsights.v20200801.PrivateLinkScopedResourceResponse[]>;
    /**
     * The provisioning state of the workspace.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The network access type for accessing Log Analytics ingestion.
     */
    public readonly publicNetworkAccessForIngestion!: pulumi.Output<string | undefined>;
    /**
     * The network access type for accessing Log Analytics query.
     */
    public readonly publicNetworkAccessForQuery!: pulumi.Output<string | undefined>;
    /**
     * The workspace data retention in days, between 30 and 730.
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * The SKU of the workspace.
     */
    public readonly sku!: pulumi.Output<outputs.operationalinsights.v20200801.WorkspaceSkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The daily volume cap for ingestion.
     */
    public readonly workspaceCapping!: pulumi.Output<outputs.operationalinsights.v20200801.WorkspaceCappingResponse | undefined>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WorkspaceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["publicNetworkAccessForIngestion"] = args ? args.publicNetworkAccessForIngestion : undefined;
            inputs["publicNetworkAccessForQuery"] = args ? args.publicNetworkAccessForQuery : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceCapping"] = args ? args.workspaceCapping : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["customerId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateLinkScopedResources"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/latest:Workspace" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * The ETag of the workspace.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The provisioning state of the workspace.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The network access type for accessing Log Analytics ingestion.
     */
    readonly publicNetworkAccessForIngestion?: pulumi.Input<string>;
    /**
     * The network access type for accessing Log Analytics query.
     */
    readonly publicNetworkAccessForQuery?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The workspace data retention in days, between 30 and 730.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * The SKU of the workspace.
     */
    readonly sku?: pulumi.Input<inputs.operationalinsights.v20200801.WorkspaceSku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The daily volume cap for ingestion.
     */
    readonly workspaceCapping?: pulumi.Input<inputs.operationalinsights.v20200801.WorkspaceCapping>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
