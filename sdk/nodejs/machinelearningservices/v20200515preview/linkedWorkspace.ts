// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Linked workspace.
 */
export class LinkedWorkspace extends pulumi.CustomResource {
    /**
     * Get an existing LinkedWorkspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LinkedWorkspace {
        return new LinkedWorkspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:machinelearningservices/v20200515preview:LinkedWorkspace';

    /**
     * Returns true if the given object is an instance of LinkedWorkspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedWorkspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedWorkspace.__pulumiType;
    }

    /**
     * Friendly name of the linked workspace.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * LinkedWorkspace specific properties.
     */
    public readonly properties!: pulumi.Output<outputs.machinelearningservices.v20200515preview.LinkedWorkspacePropsResponse>;
    /**
     * Resource type of linked workspace.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LinkedWorkspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedWorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.linkName === undefined) {
                throw new Error("Missing required property 'linkName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["linkName"] = args ? args.linkName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:machinelearningservices/v20200501preview:LinkedWorkspace" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LinkedWorkspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LinkedWorkspace resource.
 */
export interface LinkedWorkspaceArgs {
    /**
     * Friendly name of the linked workspace
     */
    readonly linkName: pulumi.Input<string>;
    /**
     * Friendly name of the linked workspace
     */
    readonly name?: pulumi.Input<string>;
    /**
     * LinkedWorkspace specific properties.
     */
    readonly properties?: pulumi.Input<inputs.machinelearningservices.v20200515preview.LinkedWorkspaceProps>;
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
