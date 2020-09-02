// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The top level Workspace resource container.
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
    public static readonly __pulumiType = 'azurerm:operationalinsights/v20151101preview:Workspace';

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
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * This is a legacy property and is not used anymore. Kept here for backward compatibility.
     */
    public /*out*/ readonly portalUrl!: pulumi.Output<string>;
    /**
     * The provisioning state of the workspace.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus. 
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * The SKU of the workspace.
     */
    public readonly sku!: pulumi.Output<outputs.operationalinsights.v20151101preview.SkuResponse | undefined>;
    /**
     * This is a read-only legacy property. It is always set to 'Azure' by the service. Kept here for backward compatibility.
     */
    public /*out*/ readonly source!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

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
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["customerId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["portalUrl"] = undefined /*out*/;
            inputs["source"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/latest:Workspace" }, { type: "azurerm:operationalinsights/v20200301preview:Workspace" }, { type: "azurerm:operationalinsights/v20200801:Workspace" }] };
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
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The provisioning state of the workspace.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The resource group name of the workspace.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus. 
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * The SKU of the workspace.
     */
    readonly sku?: pulumi.Input<inputs.operationalinsights.v20151101preview.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
