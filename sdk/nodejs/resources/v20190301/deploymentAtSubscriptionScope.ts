// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Deployment information.
 */
export class DeploymentAtSubscriptionScope extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentAtSubscriptionScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeploymentAtSubscriptionScope {
        return new DeploymentAtSubscriptionScope(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:resources/v20190301:DeploymentAtSubscriptionScope';

    /**
     * Returns true if the given object is an instance of DeploymentAtSubscriptionScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentAtSubscriptionScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentAtSubscriptionScope.__pulumiType;
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
    public readonly properties!: pulumi.Output<outputs.resources.v20190301.DeploymentPropertiesExtendedResponse>;
    /**
     * The type of the deployment.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DeploymentAtSubscriptionScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentAtSubscriptionScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentAtSubscriptionScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DeploymentAtSubscriptionScopeArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DeploymentAtSubscriptionScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeploymentAtSubscriptionScope resource.
 */
export interface DeploymentAtSubscriptionScopeArgs {
    /**
     * The location to store the deployment data.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the deployment.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The deployment properties.
     */
    readonly properties: pulumi.Input<inputs.resources.v20190301.DeploymentProperties>;
}
