// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Agent Pool.
 */
export class AgentPool extends pulumi.CustomResource {
    /**
     * Get an existing AgentPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentPool {
        return new AgentPool(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerservice/v20200201:AgentPool';

    /**
     * Returns true if the given object is an instance of AgentPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentPool.__pulumiType;
    }

    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of an agent pool.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.containerservice.v20200201.ManagedClusterAgentPoolProfilePropertiesResponse>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a AgentPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AgentPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AgentPoolArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["count"] = args ? args.count : undefined;
            inputs["enableAutoScaling"] = args ? args.enableAutoScaling : undefined;
            inputs["enableNodePublicIP"] = args ? args.enableNodePublicIP : undefined;
            inputs["maxCount"] = args ? args.maxCount : undefined;
            inputs["maxPods"] = args ? args.maxPods : undefined;
            inputs["minCount"] = args ? args.minCount : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeLabels"] = args ? args.nodeLabels : undefined;
            inputs["nodeTaints"] = args ? args.nodeTaints : undefined;
            inputs["orchestratorVersion"] = args ? args.orchestratorVersion : undefined;
            inputs["osDiskSizeGB"] = args ? args.osDiskSizeGB : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["scaleSetEvictionPolicy"] = args ? args.scaleSetEvictionPolicy : undefined;
            inputs["scaleSetPriority"] = args ? args.scaleSetPriority : undefined;
            inputs["spotMaxPrice"] = args ? args.spotMaxPrice : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
            inputs["vnetSubnetID"] = args ? args.vnetSubnetID : undefined;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AgentPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentPool resource.
 */
export interface AgentPoolArgs {
    /**
     * Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
     */
    readonly count?: pulumi.Input<number>;
    /**
     * Whether to enable auto-scaler
     */
    readonly enableAutoScaling?: pulumi.Input<boolean>;
    /**
     * Enable public IP for nodes
     */
    readonly enableNodePublicIP?: pulumi.Input<boolean>;
    /**
     * Maximum number of nodes for auto-scaling
     */
    readonly maxCount?: pulumi.Input<number>;
    /**
     * Maximum number of pods that can run on a node.
     */
    readonly maxPods?: pulumi.Input<number>;
    /**
     * Minimum number of nodes for auto-scaling
     */
    readonly minCount?: pulumi.Input<number>;
    /**
     * The name of the agent pool.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Agent pool node labels to be persisted across all nodes in agent pool.
     */
    readonly nodeLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
     */
    readonly nodeTaints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Version of orchestrator specified when creating the managed cluster.
     */
    readonly orchestratorVersion?: pulumi.Input<string>;
    /**
     * OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
     */
    readonly osDiskSizeGB?: pulumi.Input<number>;
    /**
     * OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the managed cluster resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * ScaleSetEvictionPolicy to be used to specify eviction policy for Spot or low priority virtual machine scale set. Default to Delete.
     */
    readonly scaleSetEvictionPolicy?: pulumi.Input<string>;
    /**
     * ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
     */
    readonly scaleSetPriority?: pulumi.Input<string>;
    /**
     * SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
     */
    readonly spotMaxPrice?: pulumi.Input<number>;
    /**
     * Agent pool tags to be persisted on the agent pool virtual machine scale set.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * AgentPoolType represents types of an agent pool
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Size of agent VMs.
     */
    readonly vmSize?: pulumi.Input<string>;
    /**
     * VNet SubnetID specifies the VNet's subnet identifier.
     */
    readonly vnetSubnetID?: pulumi.Input<string>;
}
