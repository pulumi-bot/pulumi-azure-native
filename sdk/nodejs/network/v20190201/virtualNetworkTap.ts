// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Virtual Network Tap resource
 */
export class VirtualNetworkTap extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkTap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualNetworkTap {
        return new VirtualNetworkTap(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190201:VirtualNetworkTap';

    /**
     * Returns true if the given object is an instance of VirtualNetworkTap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkTap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkTap.__pulumiType;
    }

    /**
     * The reference to the private IP address on the internal Load Balancer that will receive the tap
     */
    public readonly destinationLoadBalancerFrontEndIPConfiguration!: pulumi.Output<outputs.network.v20190201.FrontendIPConfigurationResponse | undefined>;
    /**
     * The reference to the private IP Address of the collector nic that will receive the tap
     */
    public readonly destinationNetworkInterfaceIPConfiguration!: pulumi.Output<outputs.network.v20190201.NetworkInterfaceIPConfigurationResponse | undefined>;
    /**
     * The VXLAN destination port that will receive the tapped traffic.
     */
    public readonly destinationPort!: pulumi.Output<number | undefined>;
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
     * Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
     */
    public /*out*/ readonly networkInterfaceTapConfigurations!: pulumi.Output<outputs.network.v20190201.NetworkInterfaceTapConfigurationResponse[]>;
    /**
     * The provisioning state of the virtual network tap. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resourceGuid property of the virtual network tap.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VirtualNetworkTap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkTapArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.tapName === undefined) {
                throw new Error("Missing required property 'tapName'");
            }
            inputs["destinationLoadBalancerFrontEndIPConfiguration"] = args ? args.destinationLoadBalancerFrontEndIPConfiguration : undefined;
            inputs["destinationNetworkInterfaceIPConfiguration"] = args ? args.destinationNetworkInterfaceIPConfiguration : undefined;
            inputs["destinationPort"] = args ? args.destinationPort : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tapName"] = args ? args.tapName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaceTapConfigurations"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["destinationLoadBalancerFrontEndIPConfiguration"] = undefined /*out*/;
            inputs["destinationNetworkInterfaceIPConfiguration"] = undefined /*out*/;
            inputs["destinationPort"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaceTapConfigurations"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VirtualNetworkTap" }, { type: "azurerm:network/v20180801:VirtualNetworkTap" }, { type: "azurerm:network/v20181001:VirtualNetworkTap" }, { type: "azurerm:network/v20181101:VirtualNetworkTap" }, { type: "azurerm:network/v20181201:VirtualNetworkTap" }, { type: "azurerm:network/v20190401:VirtualNetworkTap" }, { type: "azurerm:network/v20190601:VirtualNetworkTap" }, { type: "azurerm:network/v20190701:VirtualNetworkTap" }, { type: "azurerm:network/v20190801:VirtualNetworkTap" }, { type: "azurerm:network/v20190901:VirtualNetworkTap" }, { type: "azurerm:network/v20191101:VirtualNetworkTap" }, { type: "azurerm:network/v20191201:VirtualNetworkTap" }, { type: "azurerm:network/v20200301:VirtualNetworkTap" }, { type: "azurerm:network/v20200401:VirtualNetworkTap" }, { type: "azurerm:network/v20200501:VirtualNetworkTap" }, { type: "azurerm:network/v20200601:VirtualNetworkTap" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualNetworkTap.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualNetworkTap resource.
 */
export interface VirtualNetworkTapArgs {
    /**
     * The reference to the private IP address on the internal Load Balancer that will receive the tap
     */
    readonly destinationLoadBalancerFrontEndIPConfiguration?: pulumi.Input<inputs.network.v20190201.FrontendIPConfiguration>;
    /**
     * The reference to the private IP Address of the collector nic that will receive the tap
     */
    readonly destinationNetworkInterfaceIPConfiguration?: pulumi.Input<inputs.network.v20190201.NetworkInterfaceIPConfiguration>;
    /**
     * The VXLAN destination port that will receive the tapped traffic.
     */
    readonly destinationPort?: pulumi.Input<number>;
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
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the virtual network tap.
     */
    readonly tapName: pulumi.Input<string>;
}
