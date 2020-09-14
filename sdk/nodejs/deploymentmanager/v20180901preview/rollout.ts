// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Defines the PUT rollout request body.
 */
export class Rollout extends pulumi.CustomResource {
    /**
     * Get an existing Rollout resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Rollout {
        return new Rollout(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:deploymentmanager/v20180901preview:Rollout';

    /**
     * Returns true if the given object is an instance of Rollout.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rollout {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rollout.__pulumiType;
    }

    /**
     * The reference to the artifact source resource Id where the payload is located.
     */
    public readonly artifactSourceId!: pulumi.Output<string | undefined>;
    /**
     * The version of the build being deployed.
     */
    public readonly buildVersion!: pulumi.Output<string>;
    /**
     * Identity for the resource.
     */
    public readonly identity!: pulumi.Output<outputs.deploymentmanager.v20180901preview.IdentityResponse>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of step groups that define the orchestration.
     */
    public readonly stepGroups!: pulumi.Output<outputs.deploymentmanager.v20180901preview.StepResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
     */
    public readonly targetServiceTopologyId!: pulumi.Output<string>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Rollout resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RolloutArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.buildVersion === undefined) {
                throw new Error("Missing required property 'buildVersion'");
            }
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.rolloutName === undefined) {
                throw new Error("Missing required property 'rolloutName'");
            }
            if (!args || args.stepGroups === undefined) {
                throw new Error("Missing required property 'stepGroups'");
            }
            if (!args || args.targetServiceTopologyId === undefined) {
                throw new Error("Missing required property 'targetServiceTopologyId'");
            }
            inputs["artifactSourceId"] = args ? args.artifactSourceId : undefined;
            inputs["buildVersion"] = args ? args.buildVersion : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rolloutName"] = args ? args.rolloutName : undefined;
            inputs["stepGroups"] = args ? args.stepGroups : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetServiceTopologyId"] = args ? args.targetServiceTopologyId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["artifactSourceId"] = undefined /*out*/;
            inputs["buildVersion"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["stepGroups"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["targetServiceTopologyId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:deploymentmanager/latest:Rollout" }, { type: "azurerm:deploymentmanager/v20191101preview:Rollout" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Rollout.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Rollout resource.
 */
export interface RolloutArgs {
    /**
     * The reference to the artifact source resource Id where the payload is located.
     */
    readonly artifactSourceId?: pulumi.Input<string>;
    /**
     * The version of the build being deployed.
     */
    readonly buildVersion: pulumi.Input<string>;
    /**
     * Identity for the resource.
     */
    readonly identity: pulumi.Input<inputs.deploymentmanager.v20180901preview.Identity>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The rollout name.
     */
    readonly rolloutName: pulumi.Input<string>;
    /**
     * The list of step groups that define the orchestration.
     */
    readonly stepGroups: pulumi.Input<pulumi.Input<inputs.deploymentmanager.v20180901preview.Step>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
     */
    readonly targetServiceTopologyId: pulumi.Input<string>;
}
