// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Deployment information.
 */
export class DeploymentAtManagementGroupScope extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentAtManagementGroupScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeploymentAtManagementGroupScope {
        return new DeploymentAtManagementGroupScope(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:management/v20191001:DeploymentAtManagementGroupScope';

    /**
     * Returns true if the given object is an instance of DeploymentAtManagementGroupScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentAtManagementGroupScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentAtManagementGroupScope.__pulumiType;
    }

    /**
     * the location of the deployment.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the deployment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Deployment properties.
     */
    public readonly properties!: pulumi.Output<outputs.management.v20191001.DeploymentPropertiesExtendedResponse>;
    /**
     * Deployment tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the deployment.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DeploymentAtManagementGroupScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentAtManagementGroupScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentAtManagementGroupScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DeploymentAtManagementGroupScopeArgs | undefined;
            if (!args || args.groupId === undefined) {
                throw new Error("Missing required property 'groupId'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DeploymentAtManagementGroupScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeploymentAtManagementGroupScope resource.
 */
export interface DeploymentAtManagementGroupScopeArgs {
    /**
     * The management group ID.
     */
    readonly groupId: pulumi.Input<string>;
    /**
     * The location to store the deployment data.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the deployment.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The deployment properties.
     */
    readonly properties: pulumi.Input<inputs.management.v20191001.DeploymentProperties>;
    /**
     * Deployment tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
