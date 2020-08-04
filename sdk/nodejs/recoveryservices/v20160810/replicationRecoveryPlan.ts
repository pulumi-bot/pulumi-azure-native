// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Recovery plan details.
 */
export class ReplicationRecoveryPlan extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationRecoveryPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicationRecoveryPlan {
        return new ReplicationRecoveryPlan(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:recoveryservices/v20160810:ReplicationRecoveryPlan';

    /**
     * Returns true if the given object is an instance of ReplicationRecoveryPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationRecoveryPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationRecoveryPlan.__pulumiType;
    }

    /**
     * Resource Location
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The custom details.
     */
    public readonly properties!: pulumi.Output<outputs.recoveryservices.v20160810.RecoveryPlanPropertiesResponse>;
    /**
     * Resource Type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReplicationRecoveryPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationRecoveryPlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationRecoveryPlanArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ReplicationRecoveryPlanArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReplicationRecoveryPlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReplicationRecoveryPlan resource.
 */
export interface ReplicationRecoveryPlanArgs {
    /**
     * Recovery plan name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Recovery plan creation properties.
     */
    readonly properties: pulumi.Input<inputs.recoveryservices.v20160810.CreateRecoveryPlanInputProperties>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the recovery services vault.
     */
    readonly resourceName: pulumi.Input<string>;
}
