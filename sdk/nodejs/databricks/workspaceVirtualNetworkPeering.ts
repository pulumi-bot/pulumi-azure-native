// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Peerings in a VirtualNetwork resource
 */
export class WorkspaceVirtualNetworkPeering extends pulumi.CustomResource {
    /**
     * Get an existing WorkspaceVirtualNetworkPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkspaceVirtualNetworkPeering {
        return new WorkspaceVirtualNetworkPeering(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:databricks:WorkspaceVirtualNetworkPeering';

    /**
     * Returns true if the given object is an instance of WorkspaceVirtualNetworkPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkspaceVirtualNetworkPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkspaceVirtualNetworkPeering.__pulumiType;
    }

    /**
     * Name of the virtual network peering resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of properties for vNet Peering
     */
    public readonly properties!: pulumi.Output<outputs.databricks.VirtualNetworkPeeringPropertiesFormatResponse>;
    /**
     * type of the virtual network peering resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WorkspaceVirtualNetworkPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkspaceVirtualNetworkPeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["workspaceName"] = args ? args.workspaceName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WorkspaceVirtualNetworkPeering.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkspaceVirtualNetworkPeering resource.
 */
export interface WorkspaceVirtualNetworkPeeringArgs {
    /**
     * The name of the workspace vNet peering.
     */
    readonly name: pulumi.Input<string>;
    /**
     * List of properties for vNet Peering
     */
    readonly properties: pulumi.Input<inputs.databricks.VirtualNetworkPeeringPropertiesFormat>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
