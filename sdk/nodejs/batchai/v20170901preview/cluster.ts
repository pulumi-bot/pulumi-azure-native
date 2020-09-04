// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Contains information about a Cluster.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:batchai/v20170901preview:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Possible values are: steady and resizing. steady state indicates that the cluster is not resizing. There are no changes to the number of compute nodes in the cluster in progress. A cluster enters this state when it is created and when no operations are being performed on the cluster to change the number of compute nodes. resizing state indicates that the cluster is resizing; that is, compute nodes are being added to or removed from the cluster.
     */
    public /*out*/ readonly allocationState!: pulumi.Output<string>;
    public /*out*/ readonly allocationStateTransitionTime!: pulumi.Output<string>;
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public /*out*/ readonly currentNodeCount!: pulumi.Output<number>;
    /**
     * This element contains all the errors encountered by various compute nodes during node setup.
     */
    public /*out*/ readonly errors!: pulumi.Output<outputs.batchai.v20170901preview.BatchAIErrorResponse[] | undefined>;
    /**
     * The location of the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Use this to prepare the VM. NOTE: The volumes specified in mountVolumes are mounted first and then the setupTask is run. Therefore the setup task can use local mountPaths in its execution.
     */
    public readonly nodeSetup!: pulumi.Output<outputs.batchai.v20170901preview.NodeSetupResponse | undefined>;
    /**
     * Counts of various compute node states on the cluster.
     */
    public /*out*/ readonly nodeStateCounts!: pulumi.Output<outputs.batchai.v20170901preview.NodeStateCountsResponse>;
    /**
     * Possible value are: creating - Specifies that the cluster is being created. succeeded - Specifies that the cluster has been created successfully. failed - Specifies that the cluster creation has failed. deleting - Specifies that the cluster is being deleted.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    public /*out*/ readonly provisioningStateTransitionTime!: pulumi.Output<string>;
    /**
     * At least one of manual or autoScale settings must be specified. Only one of manual or autoScale settings can be specified. If autoScale settings are specified, the system automatically scales the cluster up and down (within the supplied limits) based on the pending jobs on the cluster.
     */
    public readonly scaleSettings!: pulumi.Output<outputs.batchai.v20170901preview.ScaleSettingsResponse | undefined>;
    /**
     * Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
     */
    public readonly subnet!: pulumi.Output<outputs.batchai.v20170901preview.ResourceIdResponse | undefined>;
    /**
     * The tags of the resource
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Settings for user account that gets created on each on the nodes of a cluster.
     */
    public readonly userAccountSettings!: pulumi.Output<outputs.batchai.v20170901preview.UserAccountSettingsResponse | undefined>;
    /**
     * Settings for OS image.
     */
    public readonly virtualMachineConfiguration!: pulumi.Output<outputs.batchai.v20170901preview.VirtualMachineConfigurationResponse | undefined>;
    /**
     * The default value is dedicated. The node can get preempted while the task is running if lowpriority is chosen. This is best suited if the workload is checkpointing and can be restarted.
     */
    public readonly vmPriority!: pulumi.Output<string | undefined>;
    /**
     * All virtual machines in a cluster are the same size. For information about available VM sizes for clusters using images from the Virtual Machines Marketplace (see Sizes for Virtual Machines (Linux) or Sizes for Virtual Machines (Windows). Batch AI service supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
     */
    public readonly vmSize!: pulumi.Output<string | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.userAccountSettings === undefined) {
                throw new Error("Missing required property 'userAccountSettings'");
            }
            if (!args || args.vmSize === undefined) {
                throw new Error("Missing required property 'vmSize'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["nodeSetup"] = args ? args.nodeSetup : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scaleSettings"] = args ? args.scaleSettings : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userAccountSettings"] = args ? args.userAccountSettings : undefined;
            inputs["virtualMachineConfiguration"] = args ? args.virtualMachineConfiguration : undefined;
            inputs["vmPriority"] = args ? args.vmPriority : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
            inputs["allocationState"] = undefined /*out*/;
            inputs["allocationStateTransitionTime"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["currentNodeCount"] = undefined /*out*/;
            inputs["errors"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeStateCounts"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningStateTransitionTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["allocationState"] = undefined /*out*/;
            inputs["allocationStateTransitionTime"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["currentNodeCount"] = undefined /*out*/;
            inputs["errors"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeSetup"] = undefined /*out*/;
            inputs["nodeStateCounts"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningStateTransitionTime"] = undefined /*out*/;
            inputs["scaleSettings"] = undefined /*out*/;
            inputs["subnet"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userAccountSettings"] = undefined /*out*/;
            inputs["virtualMachineConfiguration"] = undefined /*out*/;
            inputs["vmPriority"] = undefined /*out*/;
            inputs["vmSize"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:batchai/v20180301:Cluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The region in which to create the cluster.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Use this to prepare the VM. NOTE: The volumes specified in mountVolumes are mounted first and then the setupTask is run. Therefore the setup task can use local mountPaths in its execution.
     */
    readonly nodeSetup?: pulumi.Input<inputs.batchai.v20170901preview.NodeSetup>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * At least one of manual or autoScale settings must be specified. Only one of manual or autoScale settings can be specified. If autoScale settings are specified, the system automatically scales the cluster up and down (within the supplied limits) based on the pending jobs on the cluster.
     */
    readonly scaleSettings?: pulumi.Input<inputs.batchai.v20170901preview.ScaleSettings>;
    /**
     * Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
     */
    readonly subnet?: pulumi.Input<inputs.batchai.v20170901preview.ResourceId>;
    /**
     * The user specified tags associated with the Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Settings for user account that gets created on each on the nodes of a cluster.
     */
    readonly userAccountSettings: pulumi.Input<inputs.batchai.v20170901preview.UserAccountSettings>;
    /**
     * Settings for OS image.
     */
    readonly virtualMachineConfiguration?: pulumi.Input<inputs.batchai.v20170901preview.VirtualMachineConfiguration>;
    /**
     * Default is dedicated.
     */
    readonly vmPriority?: pulumi.Input<string>;
    /**
     * All virtual machines in a cluster are the same size. For information about available VM sizes for clusters using images from the Virtual Machines Marketplace (see Sizes for Virtual Machines (Linux) or Sizes for Virtual Machines (Windows). Batch AI service supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
     */
    readonly vmSize: pulumi.Input<string>;
}
