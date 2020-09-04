// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Tap configuration in a Network Interface
 */
export class NetworkInterfaceTapConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceTapConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkInterfaceTapConfiguration {
        return new NetworkInterfaceTapConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20180801:NetworkInterfaceTapConfiguration';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceTapConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceTapConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceTapConfiguration.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the network interface tap configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Sub Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The reference of the Virtual Network Tap resource.
     */
    public readonly virtualNetworkTap!: pulumi.Output<outputs.network.v20180801.VirtualNetworkTapResponse | undefined>;

    /**
     * Create a NetworkInterfaceTapConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceTapConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.networkInterfaceName === undefined) {
                throw new Error("Missing required property 'networkInterfaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.tapConfigurationName === undefined) {
                throw new Error("Missing required property 'tapConfigurationName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaceName"] = args ? args.networkInterfaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tapConfigurationName"] = args ? args.tapConfigurationName : undefined;
            inputs["virtualNetworkTap"] = args ? args.virtualNetworkTap : undefined;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkTap"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20181001:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20181101:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20181201:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20190201:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20190401:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20190601:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20190701:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20190801:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20190901:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20191101:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20191201:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20200301:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20200401:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20200501:NetworkInterfaceTapConfiguration" }, { type: "azurerm:network/v20200601:NetworkInterfaceTapConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NetworkInterfaceTapConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkInterfaceTapConfiguration resource.
 */
export interface NetworkInterfaceTapConfigurationArgs {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the network interface.
     */
    readonly networkInterfaceName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the tap configuration.
     */
    readonly tapConfigurationName: pulumi.Input<string>;
    /**
     * The reference of the Virtual Network Tap resource.
     */
    readonly virtualNetworkTap?: pulumi.Input<inputs.network.v20180801.VirtualNetworkTap>;
}
