// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Virtual Network resource.
 */
export class VirtualNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualNetwork {
        return new VirtualNetwork(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190201:VirtualNetwork';

    /**
     * Returns true if the given object is an instance of VirtualNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetwork.__pulumiType;
    }

    /**
     * The AddressSpace that contains an array of IP address ranges that can be used by subnets.
     */
    public readonly addressSpace!: pulumi.Output<outputs.network.v20190201.AddressSpaceResponse | undefined>;
    /**
     * The DDoS protection plan associated with the virtual network.
     */
    public readonly ddosProtectionPlan!: pulumi.Output<outputs.network.v20190201.SubResourceResponse | undefined>;
    /**
     * The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
     */
    public readonly dhcpOptions!: pulumi.Output<outputs.network.v20190201.DhcpOptionsResponse | undefined>;
    /**
     * Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
     */
    public readonly enableDdosProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if VM protection is enabled for all the subnets in the virtual network.
     */
    public readonly enableVmProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The resourceGuid property of the Virtual Network resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * A list of subnets in a Virtual Network.
     */
    public readonly subnets!: pulumi.Output<outputs.network.v20190201.SubnetResponse[] | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of peerings in a Virtual Network.
     */
    public readonly virtualNetworkPeerings!: pulumi.Output<outputs.network.v20190201.VirtualNetworkPeeringResponse[] | undefined>;

    /**
     * Create a VirtualNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualNetworkArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualNetworkName === undefined) {
                throw new Error("Missing required property 'virtualNetworkName'");
            }
            inputs["addressSpace"] = args ? args.addressSpace : undefined;
            inputs["ddosProtectionPlan"] = args ? args.ddosProtectionPlan : undefined;
            inputs["dhcpOptions"] = args ? args.dhcpOptions : undefined;
            inputs["enableDdosProtection"] = args ? args.enableDdosProtection : undefined;
            inputs["enableVmProtection"] = args ? args.enableVmProtection : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkName"] = args ? args.virtualNetworkName : undefined;
            inputs["virtualNetworkPeerings"] = args ? args.virtualNetworkPeerings : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VirtualNetwork" }, { type: "azurerm:network/v20150501preview:VirtualNetwork" }, { type: "azurerm:network/v20150615:VirtualNetwork" }, { type: "azurerm:network/v20160330:VirtualNetwork" }, { type: "azurerm:network/v20160601:VirtualNetwork" }, { type: "azurerm:network/v20160901:VirtualNetwork" }, { type: "azurerm:network/v20161201:VirtualNetwork" }, { type: "azurerm:network/v20170301:VirtualNetwork" }, { type: "azurerm:network/v20170601:VirtualNetwork" }, { type: "azurerm:network/v20170801:VirtualNetwork" }, { type: "azurerm:network/v20170901:VirtualNetwork" }, { type: "azurerm:network/v20171001:VirtualNetwork" }, { type: "azurerm:network/v20171101:VirtualNetwork" }, { type: "azurerm:network/v20180101:VirtualNetwork" }, { type: "azurerm:network/v20180201:VirtualNetwork" }, { type: "azurerm:network/v20180401:VirtualNetwork" }, { type: "azurerm:network/v20180601:VirtualNetwork" }, { type: "azurerm:network/v20180701:VirtualNetwork" }, { type: "azurerm:network/v20180801:VirtualNetwork" }, { type: "azurerm:network/v20181001:VirtualNetwork" }, { type: "azurerm:network/v20181101:VirtualNetwork" }, { type: "azurerm:network/v20181201:VirtualNetwork" }, { type: "azurerm:network/v20190401:VirtualNetwork" }, { type: "azurerm:network/v20190601:VirtualNetwork" }, { type: "azurerm:network/v20190701:VirtualNetwork" }, { type: "azurerm:network/v20190801:VirtualNetwork" }, { type: "azurerm:network/v20190901:VirtualNetwork" }, { type: "azurerm:network/v20191101:VirtualNetwork" }, { type: "azurerm:network/v20191201:VirtualNetwork" }, { type: "azurerm:network/v20200301:VirtualNetwork" }, { type: "azurerm:network/v20200401:VirtualNetwork" }, { type: "azurerm:network/v20200501:VirtualNetwork" }, { type: "azurerm:network/v20200601:VirtualNetwork" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualNetwork.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualNetwork resource.
 */
export interface VirtualNetworkArgs {
    /**
     * The AddressSpace that contains an array of IP address ranges that can be used by subnets.
     */
    readonly addressSpace?: pulumi.Input<inputs.network.v20190201.AddressSpace>;
    /**
     * The DDoS protection plan associated with the virtual network.
     */
    readonly ddosProtectionPlan?: pulumi.Input<inputs.network.v20190201.SubResource>;
    /**
     * The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
     */
    readonly dhcpOptions?: pulumi.Input<inputs.network.v20190201.DhcpOptions>;
    /**
     * Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
     */
    readonly enableDdosProtection?: pulumi.Input<boolean>;
    /**
     * Indicates if VM protection is enabled for all the subnets in the virtual network.
     */
    readonly enableVmProtection?: pulumi.Input<boolean>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resourceGuid property of the Virtual Network resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * A list of subnets in a Virtual Network.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<inputs.network.v20190201.Subnet>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the virtual network.
     */
    readonly virtualNetworkName: pulumi.Input<string>;
    /**
     * A list of peerings in a Virtual Network.
     */
    readonly virtualNetworkPeerings?: pulumi.Input<pulumi.Input<inputs.network.v20190201.VirtualNetworkPeering>[]>;
}
