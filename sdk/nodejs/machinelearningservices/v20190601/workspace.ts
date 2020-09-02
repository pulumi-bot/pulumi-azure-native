// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An object that represents a machine learning workspace.
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
    public static readonly __pulumiType = 'azurerm:machinelearningservices/v20190601:Workspace';

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
     * ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
     */
    public readonly applicationInsights!: pulumi.Output<string | undefined>;
    /**
     * ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
     */
    public readonly containerRegistry!: pulumi.Output<string | undefined>;
    /**
     * The creation time of the machine learning workspace in ISO8601 format.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of this workspace.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Url for the discovery service to identify regional endpoints for machine learning experimentation services
     */
    public readonly discoveryUrl!: pulumi.Output<string | undefined>;
    /**
     * The friendly name for this workspace. This name in mutable
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.machinelearningservices.v20190601.IdentityResponse | undefined>;
    /**
     * ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
     */
    public readonly keyVault!: pulumi.Output<string | undefined>;
    /**
     * Specifies the location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
     */
    public readonly storageAccount!: pulumi.Output<string | undefined>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The immutable id associated with this workspace.
     */
    public /*out*/ readonly workspaceId!: pulumi.Output<string>;

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
            inputs["applicationInsights"] = args ? args.applicationInsights : undefined;
            inputs["containerRegistry"] = args ? args.containerRegistry : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["discoveryUrl"] = args ? args.discoveryUrl : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["keyVault"] = args ? args.keyVault : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccount"] = args ? args.storageAccount : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:machinelearningservices/latest:Workspace" }, { type: "azurerm:machinelearningservices/v20180301preview:Workspace" }, { type: "azurerm:machinelearningservices/v20181119:Workspace" }, { type: "azurerm:machinelearningservices/v20190501:Workspace" }, { type: "azurerm:machinelearningservices/v20191101:Workspace" }, { type: "azurerm:machinelearningservices/v20200101:Workspace" }, { type: "azurerm:machinelearningservices/v20200218preview:Workspace" }, { type: "azurerm:machinelearningservices/v20200301:Workspace" }, { type: "azurerm:machinelearningservices/v20200401:Workspace" }, { type: "azurerm:machinelearningservices/v20200501preview:Workspace" }, { type: "azurerm:machinelearningservices/v20200515preview:Workspace" }, { type: "azurerm:machinelearningservices/v20200601:Workspace" }, { type: "azurerm:machinelearningservices/v20200901preview:Workspace" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly applicationInsights?: pulumi.Input<string>;
    /**
     * ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly containerRegistry?: pulumi.Input<string>;
    /**
     * The description of this workspace.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Url for the discovery service to identify regional endpoints for machine learning experimentation services
     */
    readonly discoveryUrl?: pulumi.Input<string>;
    /**
     * The friendly name for this workspace. This name in mutable
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.machinelearningservices.v20190601.Identity>;
    /**
     * ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly keyVault?: pulumi.Input<string>;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly storageAccount?: pulumi.Input<string>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
