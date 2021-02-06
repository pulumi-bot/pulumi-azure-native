// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Information about workspace.
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
        return new Workspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:databricks/v20180401:Workspace';

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
     * The workspace provider authorizations.
     */
    public readonly authorizations!: pulumi.Output<outputs.databricks.v20180401.WorkspaceProviderAuthorizationResponse[] | undefined>;
    /**
     * Indicates the Object ID, PUID and Application ID of entity that created the workspace.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<outputs.databricks.v20180401.CreatedByResponse | undefined>;
    /**
     * Specifies the date and time when the workspace is created.
     */
    public /*out*/ readonly createdDateTime!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The managed resource group Id.
     */
    public readonly managedResourceGroupId!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The workspace's custom parameters.
     */
    public readonly parameters!: pulumi.Output<outputs.databricks.v20180401.WorkspaceCustomParametersResponse | undefined>;
    /**
     * The workspace provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The SKU of the resource.
     */
    public readonly sku!: pulumi.Output<outputs.databricks.v20180401.SkuResponse | undefined>;
    /**
     * The details of Managed Identity of Storage Account
     */
    public /*out*/ readonly storageAccountIdentity!: pulumi.Output<outputs.databricks.v20180401.ManagedIdentityConfigurationResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The blob URI where the UI definition file is located.
     */
    public readonly uiDefinitionUri!: pulumi.Output<string | undefined>;
    /**
     * Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.
     */
    public /*out*/ readonly updatedBy!: pulumi.Output<outputs.databricks.v20180401.CreatedByResponse | undefined>;
    /**
     * The unique identifier of the databricks workspace in databricks control plane.
     */
    public /*out*/ readonly workspaceId!: pulumi.Output<string>;
    /**
     * The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
     */
    public /*out*/ readonly workspaceUrl!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.managedResourceGroupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'managedResourceGroupId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["authorizations"] = args ? args.authorizations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedResourceGroupId"] = args ? args.managedResourceGroupId : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["uiDefinitionUri"] = args ? args.uiDefinitionUri : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["createdBy"] = undefined /*out*/;
            inputs["createdDateTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["storageAccountIdentity"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
            inputs["workspaceUrl"] = undefined /*out*/;
        } else {
            inputs["authorizations"] = undefined /*out*/;
            inputs["createdBy"] = undefined /*out*/;
            inputs["createdDateTime"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedResourceGroupId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["storageAccountIdentity"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uiDefinitionUri"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
            inputs["workspaceUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:databricks/latest:Workspace" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * The workspace provider authorizations.
     */
    readonly authorizations?: pulumi.Input<pulumi.Input<inputs.databricks.v20180401.WorkspaceProviderAuthorization>[]>;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The managed resource group Id.
     */
    readonly managedResourceGroupId: pulumi.Input<string>;
    /**
     * The workspace's custom parameters.
     */
    readonly parameters?: pulumi.Input<inputs.databricks.v20180401.WorkspaceCustomParameters>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the resource.
     */
    readonly sku?: pulumi.Input<inputs.databricks.v20180401.Sku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The blob URI where the UI definition file is located.
     */
    readonly uiDefinitionUri?: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
