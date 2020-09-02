// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A SQL virtual machine group.
 */
export class SqlVirtualMachineGroup extends pulumi.CustomResource {
    /**
     * Get an existing SqlVirtualMachineGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SqlVirtualMachineGroup {
        return new SqlVirtualMachineGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sqlvirtualmachine/v20170301preview:SqlVirtualMachineGroup';

    /**
     * Returns true if the given object is an instance of SqlVirtualMachineGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlVirtualMachineGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlVirtualMachineGroup.__pulumiType;
    }

    /**
     * Cluster type.
     */
    public /*out*/ readonly clusterConfiguration!: pulumi.Output<string>;
    /**
     * Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the group and the OS type.
     */
    public /*out*/ readonly clusterManagerType!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state to track the async operation status.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Scale type.
     */
    public /*out*/ readonly scaleType!: pulumi.Output<string>;
    /**
     * SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
     */
    public readonly sqlImageOffer!: pulumi.Output<string | undefined>;
    /**
     * SQL image sku.
     */
    public readonly sqlImageSku!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Cluster Active Directory domain profile.
     */
    public readonly wsfcDomainProfile!: pulumi.Output<outputs.sqlvirtualmachine.v20170301preview.WsfcDomainProfileResponse | undefined>;

    /**
     * Create a SqlVirtualMachineGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlVirtualMachineGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SqlVirtualMachineGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SqlVirtualMachineGroupArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sqlVirtualMachineGroupName === undefined) {
                throw new Error("Missing required property 'sqlVirtualMachineGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlImageOffer"] = args ? args.sqlImageOffer : undefined;
            inputs["sqlImageSku"] = args ? args.sqlImageSku : undefined;
            inputs["sqlVirtualMachineGroupName"] = args ? args.sqlVirtualMachineGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["wsfcDomainProfile"] = args ? args.wsfcDomainProfile : undefined;
            inputs["clusterConfiguration"] = undefined /*out*/;
            inputs["clusterManagerType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["scaleType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SqlVirtualMachineGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SqlVirtualMachineGroup resource.
 */
export interface SqlVirtualMachineGroupArgs {
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
     */
    readonly sqlImageOffer?: pulumi.Input<string>;
    /**
     * SQL image sku.
     */
    readonly sqlImageSku?: pulumi.Input<string>;
    /**
     * Name of the SQL virtual machine group.
     */
    readonly sqlVirtualMachineGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Cluster Active Directory domain profile.
     */
    readonly wsfcDomainProfile?: pulumi.Input<inputs.sqlvirtualmachine.v20170301preview.WsfcDomainProfile>;
}
