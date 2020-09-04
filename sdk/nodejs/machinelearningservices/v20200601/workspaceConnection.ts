// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Workspace connection.
 *
 * ## Example Usage
 * ### CreateWorkspaceConnection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const workspaceConnection = new azurerm.machinelearningservices.v20200601.WorkspaceConnection("workspaceConnection", {
 *     authType: "PAT",
 *     category: "ACR",
 *     connectionName: "connection-1",
 *     name: "connection-1",
 *     resourceGroupName: "resourceGroup-1",
 *     target: "www.facebook.com",
 *     value: "secrets",
 *     workspaceName: "workspace-1",
 * });
 *
 * ```
 */
export class WorkspaceConnection extends pulumi.CustomResource {
    /**
     * Get an existing WorkspaceConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkspaceConnection {
        return new WorkspaceConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:machinelearningservices/v20200601:WorkspaceConnection';

    /**
     * Returns true if the given object is an instance of WorkspaceConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkspaceConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkspaceConnection.__pulumiType;
    }

    /**
     * Authorization type of the workspace connection.
     */
    public readonly authType!: pulumi.Output<string | undefined>;
    /**
     * Category of the workspace connection.
     */
    public readonly category!: pulumi.Output<string | undefined>;
    /**
     * Friendly name of the workspace connection.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Target of the workspace connection.
     */
    public readonly target!: pulumi.Output<string | undefined>;
    /**
     * Resource type of workspace connection.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Value details of the workspace connection.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a WorkspaceConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.connectionName === undefined) {
                throw new Error("Missing required property 'connectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["authType"] = args ? args.authType : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["connectionName"] = args ? args.connectionName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["authType"] = undefined /*out*/;
            inputs["category"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["target"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:machinelearningservices/latest:WorkspaceConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WorkspaceConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkspaceConnection resource.
 */
export interface WorkspaceConnectionArgs {
    /**
     * Authorization type of the workspace connection.
     */
    readonly authType?: pulumi.Input<string>;
    /**
     * Category of the workspace connection.
     */
    readonly category?: pulumi.Input<string>;
    /**
     * Friendly name of the workspace connection
     */
    readonly connectionName: pulumi.Input<string>;
    /**
     * Friendly name of the workspace connection
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Target of the workspace connection.
     */
    readonly target?: pulumi.Input<string>;
    /**
     * Value details of the workspace connection.
     */
    readonly value?: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
