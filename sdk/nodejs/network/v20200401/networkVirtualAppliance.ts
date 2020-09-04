// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * NetworkVirtualAppliance Resource.
 *
 * ## Example Usage
 * ### Create NetworkVirtualAppliance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const networkVirtualAppliance = new azurerm.network.v20200401.NetworkVirtualAppliance("networkVirtualAppliance", {
 *     bootStrapConfigurationBlob: ["https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"],
 *     cloudInitConfigurationBlob: ["https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"],
 *     identity: {
 *         type: "UserAssigned",
 *     },
 *     location: "West US",
 *     networkVirtualApplianceName: "nva",
 *     resourceGroupName: "rg1",
 *     sku: {
 *         bundledScaleUnit: "1",
 *         marketPlaceVersion: "12.1",
 *         vendor: "Cisco SDWAN",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     virtualApplianceAsn: 10000,
 *     virtualHub: {
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
 *     },
 * });
 *
 * ```
 */
export class NetworkVirtualAppliance extends pulumi.CustomResource {
    /**
     * Get an existing NetworkVirtualAppliance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkVirtualAppliance {
        return new NetworkVirtualAppliance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200401:NetworkVirtualAppliance';

    /**
     * Returns true if the given object is an instance of NetworkVirtualAppliance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkVirtualAppliance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkVirtualAppliance.__pulumiType;
    }

    /**
     * BootStrapConfigurationBlob storage URLs.
     */
    public readonly bootStrapConfigurationBlob!: pulumi.Output<string[] | undefined>;
    /**
     * CloudInitConfigurationBlob storage URLs.
     */
    public readonly cloudInitConfigurationBlob!: pulumi.Output<string[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The service principal that has read access to cloud-init and config blob.
     */
    public readonly identity!: pulumi.Output<outputs.network.v20200401.ManagedServiceIdentityResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Network Virtual Appliance SKU.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20200401.VirtualApplianceSkuPropertiesResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * VirtualAppliance ASN.
     */
    public readonly virtualApplianceAsn!: pulumi.Output<number | undefined>;
    /**
     * List of Virtual Appliance Network Interfaces.
     */
    public /*out*/ readonly virtualApplianceNics!: pulumi.Output<outputs.network.v20200401.VirtualApplianceNicPropertiesResponse[]>;
    /**
     * The Virtual Hub where Network Virtual Appliance is being deployed.
     */
    public readonly virtualHub!: pulumi.Output<outputs.network.v20200401.SubResourceResponse | undefined>;

    /**
     * Create a NetworkVirtualAppliance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkVirtualApplianceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.networkVirtualApplianceName === undefined) {
                throw new Error("Missing required property 'networkVirtualApplianceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["bootStrapConfigurationBlob"] = args ? args.bootStrapConfigurationBlob : undefined;
            inputs["cloudInitConfigurationBlob"] = args ? args.cloudInitConfigurationBlob : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkVirtualApplianceName"] = args ? args.networkVirtualApplianceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualApplianceAsn"] = args ? args.virtualApplianceAsn : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualApplianceNics"] = undefined /*out*/;
        } else {
            inputs["bootStrapConfigurationBlob"] = undefined /*out*/;
            inputs["cloudInitConfigurationBlob"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualApplianceAsn"] = undefined /*out*/;
            inputs["virtualApplianceNics"] = undefined /*out*/;
            inputs["virtualHub"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:NetworkVirtualAppliance" }, { type: "azurerm:network/v20191201:NetworkVirtualAppliance" }, { type: "azurerm:network/v20200301:NetworkVirtualAppliance" }, { type: "azurerm:network/v20200501:NetworkVirtualAppliance" }, { type: "azurerm:network/v20200601:NetworkVirtualAppliance" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NetworkVirtualAppliance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkVirtualAppliance resource.
 */
export interface NetworkVirtualApplianceArgs {
    /**
     * BootStrapConfigurationBlob storage URLs.
     */
    readonly bootStrapConfigurationBlob?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * CloudInitConfigurationBlob storage URLs.
     */
    readonly cloudInitConfigurationBlob?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The service principal that has read access to cloud-init and config blob.
     */
    readonly identity?: pulumi.Input<inputs.network.v20200401.ManagedServiceIdentity>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of Network Virtual Appliance.
     */
    readonly networkVirtualApplianceName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Network Virtual Appliance SKU.
     */
    readonly sku?: pulumi.Input<inputs.network.v20200401.VirtualApplianceSkuProperties>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * VirtualAppliance ASN.
     */
    readonly virtualApplianceAsn?: pulumi.Input<number>;
    /**
     * The Virtual Hub where Network Virtual Appliance is being deployed.
     */
    readonly virtualHub?: pulumi.Input<inputs.network.v20200401.SubResource>;
}
