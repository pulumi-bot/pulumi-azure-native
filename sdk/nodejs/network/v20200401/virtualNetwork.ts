// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
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
        return new VirtualNetwork(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20200401:VirtualNetwork';

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
    public readonly addressSpace!: pulumi.Output<outputs.network.v20200401.AddressSpaceResponse | undefined>;
    /**
     * Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
     */
    public readonly bgpCommunities!: pulumi.Output<outputs.network.v20200401.VirtualNetworkBgpCommunitiesResponse | undefined>;
    /**
     * The DDoS protection plan associated with the virtual network.
     */
    public readonly ddosProtectionPlan!: pulumi.Output<outputs.network.v20200401.SubResourceResponse | undefined>;
    /**
     * The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
     */
    public readonly dhcpOptions!: pulumi.Output<outputs.network.v20200401.DhcpOptionsResponse | undefined>;
    /**
     * Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
     */
    public readonly enableDdosProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if VM protection is enabled for all the subnets in the virtual network.
     */
    public readonly enableVmProtection!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Array of IpAllocation which reference this VNET.
     */
    public readonly ipAllocations!: pulumi.Output<outputs.network.v20200401.SubResourceResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the virtual network resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resourceGuid property of the Virtual Network resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * A list of subnets in a Virtual Network.
     */
    public readonly subnets!: pulumi.Output<outputs.network.v20200401.SubnetResponse[] | undefined>;
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
    public readonly virtualNetworkPeerings!: pulumi.Output<outputs.network.v20200401.VirtualNetworkPeeringResponse[] | undefined>;

    /**
     * Create a VirtualNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.virtualNetworkName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'virtualNetworkName'");
            }
            inputs["addressSpace"] = args ? args.addressSpace : undefined;
            inputs["bgpCommunities"] = args ? args.bgpCommunities : undefined;
            inputs["ddosProtectionPlan"] = args ? args.ddosProtectionPlan : undefined;
            inputs["dhcpOptions"] = args ? args.dhcpOptions : undefined;
            inputs["enableDdosProtection"] = args ? args.enableDdosProtection : undefined;
            inputs["enableVmProtection"] = args ? args.enableVmProtection : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipAllocations"] = args ? args.ipAllocations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkName"] = args ? args.virtualNetworkName : undefined;
            inputs["virtualNetworkPeerings"] = args ? args.virtualNetworkPeerings : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["addressSpace"] = undefined /*out*/;
            inputs["bgpCommunities"] = undefined /*out*/;
            inputs["ddosProtectionPlan"] = undefined /*out*/;
            inputs["dhcpOptions"] = undefined /*out*/;
            inputs["enableDdosProtection"] = undefined /*out*/;
            inputs["enableVmProtection"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["ipAllocations"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkPeerings"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:VirtualNetwork" }, { type: "azure-nextgen:network/v20150501preview:VirtualNetwork" }, { type: "azure-nextgen:network/v20150615:VirtualNetwork" }, { type: "azure-nextgen:network/v20160330:VirtualNetwork" }, { type: "azure-nextgen:network/v20160601:VirtualNetwork" }, { type: "azure-nextgen:network/v20160901:VirtualNetwork" }, { type: "azure-nextgen:network/v20161201:VirtualNetwork" }, { type: "azure-nextgen:network/v20170301:VirtualNetwork" }, { type: "azure-nextgen:network/v20170601:VirtualNetwork" }, { type: "azure-nextgen:network/v20170801:VirtualNetwork" }, { type: "azure-nextgen:network/v20170901:VirtualNetwork" }, { type: "azure-nextgen:network/v20171001:VirtualNetwork" }, { type: "azure-nextgen:network/v20171101:VirtualNetwork" }, { type: "azure-nextgen:network/v20180101:VirtualNetwork" }, { type: "azure-nextgen:network/v20180201:VirtualNetwork" }, { type: "azure-nextgen:network/v20180401:VirtualNetwork" }, { type: "azure-nextgen:network/v20180601:VirtualNetwork" }, { type: "azure-nextgen:network/v20180701:VirtualNetwork" }, { type: "azure-nextgen:network/v20180801:VirtualNetwork" }, { type: "azure-nextgen:network/v20181001:VirtualNetwork" }, { type: "azure-nextgen:network/v20181101:VirtualNetwork" }, { type: "azure-nextgen:network/v20181201:VirtualNetwork" }, { type: "azure-nextgen:network/v20190201:VirtualNetwork" }, { type: "azure-nextgen:network/v20190401:VirtualNetwork" }, { type: "azure-nextgen:network/v20190601:VirtualNetwork" }, { type: "azure-nextgen:network/v20190701:VirtualNetwork" }, { type: "azure-nextgen:network/v20190801:VirtualNetwork" }, { type: "azure-nextgen:network/v20190901:VirtualNetwork" }, { type: "azure-nextgen:network/v20191101:VirtualNetwork" }, { type: "azure-nextgen:network/v20191201:VirtualNetwork" }, { type: "azure-nextgen:network/v20200301:VirtualNetwork" }, { type: "azure-nextgen:network/v20200501:VirtualNetwork" }, { type: "azure-nextgen:network/v20200601:VirtualNetwork" }, { type: "azure-nextgen:network/v20200701:VirtualNetwork" }, { type: "azure-nextgen:network/v20200801:VirtualNetwork" }] };
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
    readonly addressSpace?: pulumi.Input<inputs.network.v20200401.AddressSpace>;
    /**
     * Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
     */
    readonly bgpCommunities?: pulumi.Input<inputs.network.v20200401.VirtualNetworkBgpCommunities>;
    /**
     * The DDoS protection plan associated with the virtual network.
     */
    readonly ddosProtectionPlan?: pulumi.Input<inputs.network.v20200401.SubResource>;
    /**
     * The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
     */
    readonly dhcpOptions?: pulumi.Input<inputs.network.v20200401.DhcpOptions>;
    /**
     * Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
     */
    readonly enableDdosProtection?: pulumi.Input<boolean>;
    /**
     * Indicates if VM protection is enabled for all the subnets in the virtual network.
     */
    readonly enableVmProtection?: pulumi.Input<boolean>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Array of IpAllocation which reference this VNET.
     */
    readonly ipAllocations?: pulumi.Input<pulumi.Input<inputs.network.v20200401.SubResource>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A list of subnets in a Virtual Network.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<inputs.network.v20200401.Subnet>[]>;
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
    readonly virtualNetworkPeerings?: pulumi.Input<pulumi.Input<inputs.network.v20200401.VirtualNetworkPeering>[]>;
}
