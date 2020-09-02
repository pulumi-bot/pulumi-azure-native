// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Subnet in a VirtualNetwork resource
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20150501preview:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    /**
     * Gets or sets Address prefix for the subnet.
     */
    public readonly addressPrefix!: pulumi.Output<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Gets array of references to the network interface IP configurations using subnet
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.v20150501preview.SubResourceResponse[] | undefined>;
    /**
     * Gets name of the resource that is unique within a resource group. This name can be used to access the resource
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the reference of the NetworkSecurityGroup resource
     */
    public readonly networkSecurityGroup!: pulumi.Output<outputs.network.v20150501preview.SubResourceResponse | undefined>;
    /**
     * Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the reference of the RouteTable resource
     */
    public readonly routeTable!: pulumi.Output<outputs.network.v20150501preview.SubResourceResponse | undefined>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SubnetArgs | undefined;
            if (!args || args.addressPrefix === undefined) {
                throw new Error("Missing required property 'addressPrefix'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subnetName === undefined) {
                throw new Error("Missing required property 'subnetName'");
            }
            if (!args || args.virtualNetworkName === undefined) {
                throw new Error("Missing required property 'virtualNetworkName'");
            }
            inputs["addressPrefix"] = args ? args.addressPrefix : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkSecurityGroup"] = args ? args.networkSecurityGroup : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeTable"] = args ? args.routeTable : undefined;
            inputs["subnetName"] = args ? args.subnetName : undefined;
            inputs["virtualNetworkName"] = args ? args.virtualNetworkName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:Subnet" }, { type: "azurerm:network/v20150615:Subnet" }, { type: "azurerm:network/v20160330:Subnet" }, { type: "azurerm:network/v20160601:Subnet" }, { type: "azurerm:network/v20160901:Subnet" }, { type: "azurerm:network/v20161201:Subnet" }, { type: "azurerm:network/v20170301:Subnet" }, { type: "azurerm:network/v20170601:Subnet" }, { type: "azurerm:network/v20170801:Subnet" }, { type: "azurerm:network/v20170901:Subnet" }, { type: "azurerm:network/v20171001:Subnet" }, { type: "azurerm:network/v20171101:Subnet" }, { type: "azurerm:network/v20180101:Subnet" }, { type: "azurerm:network/v20180201:Subnet" }, { type: "azurerm:network/v20180401:Subnet" }, { type: "azurerm:network/v20180601:Subnet" }, { type: "azurerm:network/v20180701:Subnet" }, { type: "azurerm:network/v20180801:Subnet" }, { type: "azurerm:network/v20181001:Subnet" }, { type: "azurerm:network/v20181101:Subnet" }, { type: "azurerm:network/v20181201:Subnet" }, { type: "azurerm:network/v20190201:Subnet" }, { type: "azurerm:network/v20190401:Subnet" }, { type: "azurerm:network/v20190601:Subnet" }, { type: "azurerm:network/v20190701:Subnet" }, { type: "azurerm:network/v20190801:Subnet" }, { type: "azurerm:network/v20190901:Subnet" }, { type: "azurerm:network/v20191101:Subnet" }, { type: "azurerm:network/v20191201:Subnet" }, { type: "azurerm:network/v20200301:Subnet" }, { type: "azurerm:network/v20200401:Subnet" }, { type: "azurerm:network/v20200501:Subnet" }, { type: "azurerm:network/v20200601:Subnet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Subnet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * Gets or sets Address prefix for the subnet.
     */
    readonly addressPrefix: pulumi.Input<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Gets array of references to the network interface IP configurations using subnet
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.SubResource>[]>;
    /**
     * Gets name of the resource that is unique within a resource group. This name can be used to access the resource
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Gets or sets the reference of the NetworkSecurityGroup resource
     */
    readonly networkSecurityGroup?: pulumi.Input<inputs.network.v20150501preview.SubResource>;
    /**
     * Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the reference of the RouteTable resource
     */
    readonly routeTable?: pulumi.Input<inputs.network.v20150501preview.SubResource>;
    /**
     * The name of the subnet.
     */
    readonly subnetName: pulumi.Input<string>;
    /**
     * The name of the virtual network.
     */
    readonly virtualNetworkName: pulumi.Input<string>;
}
