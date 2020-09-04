// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Data Lake Analytics compute policy information.
 *
 * ## Example Usage
 * ### Creates or updates the specified compute policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const computePolicy = new azurerm.datalakeanalytics.v20161101.ComputePolicy("computePolicy", {
 *     accountName: "contosoadla",
 *     computePolicyName: "test_policy",
 *     maxDegreeOfParallelismPerJob: 10,
 *     minPriorityPerJob: 30,
 *     objectId: "776b9091-8916-4638-87f7-9c989a38da98",
 *     objectType: "User",
 *     resourceGroupName: "contosorg",
 * });
 *
 * ```
 */
export class ComputePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ComputePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ComputePolicy {
        return new ComputePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datalakeanalytics/v20161101:ComputePolicy';

    /**
     * Returns true if the given object is an instance of ComputePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputePolicy.__pulumiType;
    }

    /**
     * The maximum degree of parallelism per job this user can use to submit jobs.
     */
    public readonly maxDegreeOfParallelismPerJob!: pulumi.Output<number>;
    /**
     * The minimum priority per job this user can use to submit jobs.
     */
    public readonly minPriorityPerJob!: pulumi.Output<number>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The AAD object identifier for the entity to create a policy for.
     */
    public readonly objectId!: pulumi.Output<string>;
    /**
     * The type of AAD object the object identifier refers to.
     */
    public readonly objectType!: pulumi.Output<string>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ComputePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.computePolicyName === undefined) {
                throw new Error("Missing required property 'computePolicyName'");
            }
            if (!args || args.objectId === undefined) {
                throw new Error("Missing required property 'objectId'");
            }
            if (!args || args.objectType === undefined) {
                throw new Error("Missing required property 'objectType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["computePolicyName"] = args ? args.computePolicyName : undefined;
            inputs["maxDegreeOfParallelismPerJob"] = args ? args.maxDegreeOfParallelismPerJob : undefined;
            inputs["minPriorityPerJob"] = args ? args.minPriorityPerJob : undefined;
            inputs["objectId"] = args ? args.objectId : undefined;
            inputs["objectType"] = args ? args.objectType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["maxDegreeOfParallelismPerJob"] = undefined /*out*/;
            inputs["minPriorityPerJob"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["objectId"] = undefined /*out*/;
            inputs["objectType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:datalakeanalytics/latest:ComputePolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ComputePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ComputePolicy resource.
 */
export interface ComputePolicyArgs {
    /**
     * The name of the Data Lake Analytics account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the compute policy to create or update.
     */
    readonly computePolicyName: pulumi.Input<string>;
    /**
     * The maximum degree of parallelism per job this user can use to submit jobs. This property, the min priority per job property, or both must be passed.
     */
    readonly maxDegreeOfParallelismPerJob?: pulumi.Input<number>;
    /**
     * The minimum priority per job this user can use to submit jobs. This property, the max degree of parallelism per job property, or both must be passed.
     */
    readonly minPriorityPerJob?: pulumi.Input<number>;
    /**
     * The AAD object identifier for the entity to create a policy for.
     */
    readonly objectId: pulumi.Input<string>;
    /**
     * The type of AAD object the object identifier refers to.
     */
    readonly objectType: pulumi.Input<string>;
    /**
     * The name of the Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
