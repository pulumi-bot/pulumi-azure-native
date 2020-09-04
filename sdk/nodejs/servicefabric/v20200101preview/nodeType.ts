// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Describes a node type in the cluster, each node type represents sub set of nodes in the cluster.
 */
export class NodeType extends pulumi.CustomResource {
    /**
     * Get an existing NodeType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NodeType {
        return new NodeType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabric/v20200101preview:NodeType';

    /**
     * Returns true if the given object is an instance of NodeType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeType.__pulumiType;
    }

    /**
     * The range of ports from which cluster assigned port to Service Fabric applications.
     */
    public readonly applicationPorts!: pulumi.Output<outputs.servicefabric.v20200101preview.EndpointRangeDescriptionResponse | undefined>;
    /**
     * The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
     */
    public readonly capacities!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Disk size for each vm in the node type in GBs.
     */
    public readonly dataDiskSizeGB!: pulumi.Output<number>;
    /**
     * The range of ephemeral ports that nodes in this node type should be configured with.
     */
    public readonly ephemeralPorts!: pulumi.Output<outputs.servicefabric.v20200101preview.EndpointRangeDescriptionResponse | undefined>;
    /**
     * The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
     */
    public readonly isPrimary!: pulumi.Output<boolean>;
    /**
     * Azure resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
     */
    public readonly placementProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The provisioning state of the managed cluster resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Set of extensions that should be installed onto the virtual machines.
     */
    public readonly vmExtensions!: pulumi.Output<outputs.servicefabric.v20200101preview.VMSSExtensionResponse[] | undefined>;
    /**
     * The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
     */
    public readonly vmImageOffer!: pulumi.Output<string | undefined>;
    /**
     * The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
     */
    public readonly vmImagePublisher!: pulumi.Output<string | undefined>;
    /**
     * The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
     */
    public readonly vmImageSku!: pulumi.Output<string | undefined>;
    /**
     * The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
     */
    public readonly vmImageVersion!: pulumi.Output<string | undefined>;
    /**
     * The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
     */
    public readonly vmInstanceCount!: pulumi.Output<number>;
    /**
     * The secrets to install in the virtual machines.
     */
    public readonly vmSecrets!: pulumi.Output<outputs.servicefabric.v20200101preview.VaultSecretGroupResponse[] | undefined>;
    /**
     * The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
     */
    public readonly vmSize!: pulumi.Output<string | undefined>;

    /**
     * Create a NodeType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.dataDiskSizeGB === undefined) {
                throw new Error("Missing required property 'dataDiskSizeGB'");
            }
            if (!args || args.isPrimary === undefined) {
                throw new Error("Missing required property 'isPrimary'");
            }
            if (!args || args.nodeTypeName === undefined) {
                throw new Error("Missing required property 'nodeTypeName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vmInstanceCount === undefined) {
                throw new Error("Missing required property 'vmInstanceCount'");
            }
            inputs["applicationPorts"] = args ? args.applicationPorts : undefined;
            inputs["capacities"] = args ? args.capacities : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["dataDiskSizeGB"] = args ? args.dataDiskSizeGB : undefined;
            inputs["ephemeralPorts"] = args ? args.ephemeralPorts : undefined;
            inputs["isPrimary"] = args ? args.isPrimary : undefined;
            inputs["nodeTypeName"] = args ? args.nodeTypeName : undefined;
            inputs["placementProperties"] = args ? args.placementProperties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vmExtensions"] = args ? args.vmExtensions : undefined;
            inputs["vmImageOffer"] = args ? args.vmImageOffer : undefined;
            inputs["vmImagePublisher"] = args ? args.vmImagePublisher : undefined;
            inputs["vmImageSku"] = args ? args.vmImageSku : undefined;
            inputs["vmImageVersion"] = args ? args.vmImageVersion : undefined;
            inputs["vmInstanceCount"] = args ? args.vmInstanceCount : undefined;
            inputs["vmSecrets"] = args ? args.vmSecrets : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["applicationPorts"] = undefined /*out*/;
            inputs["capacities"] = undefined /*out*/;
            inputs["dataDiskSizeGB"] = undefined /*out*/;
            inputs["ephemeralPorts"] = undefined /*out*/;
            inputs["isPrimary"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["placementProperties"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmExtensions"] = undefined /*out*/;
            inputs["vmImageOffer"] = undefined /*out*/;
            inputs["vmImagePublisher"] = undefined /*out*/;
            inputs["vmImageSku"] = undefined /*out*/;
            inputs["vmImageVersion"] = undefined /*out*/;
            inputs["vmInstanceCount"] = undefined /*out*/;
            inputs["vmSecrets"] = undefined /*out*/;
            inputs["vmSize"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NodeType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NodeType resource.
 */
export interface NodeTypeArgs {
    /**
     * The range of ports from which cluster assigned port to Service Fabric applications.
     */
    readonly applicationPorts?: pulumi.Input<inputs.servicefabric.v20200101preview.EndpointRangeDescription>;
    /**
     * The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
     */
    readonly capacities?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * Disk size for each vm in the node type in GBs.
     */
    readonly dataDiskSizeGB: pulumi.Input<number>;
    /**
     * The range of ephemeral ports that nodes in this node type should be configured with.
     */
    readonly ephemeralPorts?: pulumi.Input<inputs.servicefabric.v20200101preview.EndpointRangeDescription>;
    /**
     * The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
     */
    readonly isPrimary: pulumi.Input<boolean>;
    /**
     * The name of the node type.
     */
    readonly nodeTypeName: pulumi.Input<string>;
    /**
     * The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
     */
    readonly placementProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Azure resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of extensions that should be installed onto the virtual machines.
     */
    readonly vmExtensions?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20200101preview.VMSSExtension>[]>;
    /**
     * The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
     */
    readonly vmImageOffer?: pulumi.Input<string>;
    /**
     * The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
     */
    readonly vmImagePublisher?: pulumi.Input<string>;
    /**
     * The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
     */
    readonly vmImageSku?: pulumi.Input<string>;
    /**
     * The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
     */
    readonly vmImageVersion?: pulumi.Input<string>;
    /**
     * The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
     */
    readonly vmInstanceCount: pulumi.Input<number>;
    /**
     * The secrets to install in the virtual machines.
     */
    readonly vmSecrets?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20200101preview.VaultSecretGroup>[]>;
    /**
     * The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
     */
    readonly vmSize?: pulumi.Input<string>;
}
