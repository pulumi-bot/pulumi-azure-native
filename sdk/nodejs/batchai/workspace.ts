// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Batch AI Workspace information.
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
    public static readonly __pulumiType = 'azurerm:batchai:Workspace';

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
     * The location of the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties associated with the workspace.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.batchai.WorkspacePropertiesResponse>;
    /**
     * The tags of the resource
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["properties"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * The region in which to create the Workspace.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The user specified tags associated with the Workspace.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}