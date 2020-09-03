// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Azure Firewall resource.
 *
 * ## Example Usage
 * ### Create Azure Firewall
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const azureFirewall = new azurerm.network.v20200301.AzureFirewall("azureFirewall", {
 *     applicationRuleCollections: [{
 *         name: "apprulecoll",
 *     }],
 *     azureFirewallName: "azurefirewall",
 *     ipConfigurations: [{
 *         name: "azureFirewallIpConfiguration",
 *     }],
 *     location: "West US",
 *     natRuleCollections: [{
 *         name: "natrulecoll",
 *     }],
 *     networkRuleCollections: [{
 *         name: "netrulecoll",
 *     }],
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "AZFW_VNet",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 *     zones: [],
 * });
 *
 * ```
 * ### Create Azure Firewall With Additional Properties
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const azureFirewall = new azurerm.network.v20200301.AzureFirewall("azureFirewall", {
 *     additionalProperties: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 *     applicationRuleCollections: [{
 *         name: "apprulecoll",
 *     }],
 *     azureFirewallName: "azurefirewall",
 *     ipConfigurations: [{
 *         name: "azureFirewallIpConfiguration",
 *     }],
 *     location: "West US",
 *     natRuleCollections: [{
 *         name: "natrulecoll",
 *     }],
 *     networkRuleCollections: [{
 *         name: "netrulecoll",
 *     }],
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "AZFW_VNet",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 *     zones: [],
 * });
 *
 * ```
 * ### Create Azure Firewall With IpGroups
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const azureFirewall = new azurerm.network.v20200301.AzureFirewall("azureFirewall", {
 *     applicationRuleCollections: [{
 *         name: "apprulecoll",
 *     }],
 *     azureFirewallName: "azurefirewall",
 *     ipConfigurations: [{
 *         name: "azureFirewallIpConfiguration",
 *     }],
 *     location: "West US",
 *     natRuleCollections: [{
 *         name: "natrulecoll",
 *     }],
 *     networkRuleCollections: [{
 *         name: "netrulecoll",
 *     }],
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "AZFW_VNet",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 *     zones: [],
 * });
 *
 * ```
 * ### Create Azure Firewall With Zones
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const azureFirewall = new azurerm.network.v20200301.AzureFirewall("azureFirewall", {
 *     applicationRuleCollections: [{
 *         name: "apprulecoll",
 *     }],
 *     azureFirewallName: "azurefirewall",
 *     ipConfigurations: [{
 *         name: "azureFirewallIpConfiguration",
 *     }],
 *     location: "West US 2",
 *     natRuleCollections: [{
 *         name: "natrulecoll",
 *     }],
 *     networkRuleCollections: [{
 *         name: "netrulecoll",
 *     }],
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "AZFW_VNet",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 *     zones: [
 *         "1",
 *         "2",
 *         "3",
 *     ],
 * });
 *
 * ```
 * ### Create Azure Firewall With management subnet
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const azureFirewall = new azurerm.network.v20200301.AzureFirewall("azureFirewall", {
 *     applicationRuleCollections: [{
 *         name: "apprulecoll",
 *     }],
 *     azureFirewallName: "azurefirewall",
 *     ipConfigurations: [{
 *         name: "azureFirewallIpConfiguration",
 *     }],
 *     location: "West US",
 *     managementIpConfiguration: {
 *         name: "azureFirewallMgmtIpConfiguration",
 *     },
 *     natRuleCollections: [{
 *         name: "natrulecoll",
 *     }],
 *     networkRuleCollections: [{
 *         name: "netrulecoll",
 *     }],
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "AZFW_VNet",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 *     zones: [],
 * });
 *
 * ```
 * ### Create Azure Firewall in virtual Hub
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const azureFirewall = new azurerm.network.v20200301.AzureFirewall("azureFirewall", {
 *     azureFirewallName: "azurefirewall",
 *     firewallPolicy: {
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1",
 *     },
 *     location: "West US",
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "AZFW_Hub",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 *     virtualHub: {
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
 *     },
 *     zones: [],
 * });
 *
 * ```
 */
export class AzureFirewall extends pulumi.CustomResource {
    /**
     * Get an existing AzureFirewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AzureFirewall {
        return new AzureFirewall(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200301:AzureFirewall';

    /**
     * Returns true if the given object is an instance of AzureFirewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureFirewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureFirewall.__pulumiType;
    }

    /**
     * The additional properties used to further config this azure firewall.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Collection of application rule collections used by Azure Firewall.
     */
    public readonly applicationRuleCollections!: pulumi.Output<outputs.network.v20200301.AzureFirewallApplicationRuleCollectionResponse[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The firewallPolicy associated with this azure firewall.
     */
    public readonly firewallPolicy!: pulumi.Output<outputs.network.v20200301.SubResourceResponse | undefined>;
    /**
     * IP addresses associated with AzureFirewall.
     */
    public /*out*/ readonly hubIpAddresses!: pulumi.Output<outputs.network.v20200301.HubIPAddressesResponse>;
    /**
     * IP configuration of the Azure Firewall resource.
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.v20200301.AzureFirewallIPConfigurationResponse[] | undefined>;
    /**
     * IpGroups associated with AzureFirewall.
     */
    public /*out*/ readonly ipGroups!: pulumi.Output<outputs.network.v20200301.AzureFirewallIpGroupsResponse[]>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * IP configuration of the Azure Firewall used for management traffic.
     */
    public readonly managementIpConfiguration!: pulumi.Output<outputs.network.v20200301.AzureFirewallIPConfigurationResponse | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Collection of NAT rule collections used by Azure Firewall.
     */
    public readonly natRuleCollections!: pulumi.Output<outputs.network.v20200301.AzureFirewallNatRuleCollectionResponse[] | undefined>;
    /**
     * Collection of network rule collections used by Azure Firewall.
     */
    public readonly networkRuleCollections!: pulumi.Output<outputs.network.v20200301.AzureFirewallNetworkRuleCollectionResponse[] | undefined>;
    /**
     * The provisioning state of the Azure firewall resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The Azure Firewall Resource SKU.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20200301.AzureFirewallSkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The operation mode for Threat Intelligence.
     */
    public readonly threatIntelMode!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The virtualHub to which the firewall belongs.
     */
    public readonly virtualHub!: pulumi.Output<outputs.network.v20200301.SubResourceResponse | undefined>;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AzureFirewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureFirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureFirewallArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AzureFirewallArgs | undefined;
            if (!args || args.azureFirewallName === undefined) {
                throw new Error("Missing required property 'azureFirewallName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["applicationRuleCollections"] = args ? args.applicationRuleCollections : undefined;
            inputs["azureFirewallName"] = args ? args.azureFirewallName : undefined;
            inputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementIpConfiguration"] = args ? args.managementIpConfiguration : undefined;
            inputs["natRuleCollections"] = args ? args.natRuleCollections : undefined;
            inputs["networkRuleCollections"] = args ? args.networkRuleCollections : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["threatIntelMode"] = args ? args.threatIntelMode : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["hubIpAddresses"] = undefined /*out*/;
            inputs["ipGroups"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:AzureFirewall" }, { type: "azurerm:network/v20180401:AzureFirewall" }, { type: "azurerm:network/v20180601:AzureFirewall" }, { type: "azurerm:network/v20180701:AzureFirewall" }, { type: "azurerm:network/v20180801:AzureFirewall" }, { type: "azurerm:network/v20181001:AzureFirewall" }, { type: "azurerm:network/v20181101:AzureFirewall" }, { type: "azurerm:network/v20181201:AzureFirewall" }, { type: "azurerm:network/v20190201:AzureFirewall" }, { type: "azurerm:network/v20190401:AzureFirewall" }, { type: "azurerm:network/v20190601:AzureFirewall" }, { type: "azurerm:network/v20190701:AzureFirewall" }, { type: "azurerm:network/v20190801:AzureFirewall" }, { type: "azurerm:network/v20190901:AzureFirewall" }, { type: "azurerm:network/v20191101:AzureFirewall" }, { type: "azurerm:network/v20191201:AzureFirewall" }, { type: "azurerm:network/v20200401:AzureFirewall" }, { type: "azurerm:network/v20200501:AzureFirewall" }, { type: "azurerm:network/v20200601:AzureFirewall" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AzureFirewall.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AzureFirewall resource.
 */
export interface AzureFirewallArgs {
    /**
     * The additional properties used to further config this azure firewall.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Collection of application rule collections used by Azure Firewall.
     */
    readonly applicationRuleCollections?: pulumi.Input<pulumi.Input<inputs.network.v20200301.AzureFirewallApplicationRuleCollection>[]>;
    /**
     * The name of the Azure Firewall.
     */
    readonly azureFirewallName: pulumi.Input<string>;
    /**
     * The firewallPolicy associated with this azure firewall.
     */
    readonly firewallPolicy?: pulumi.Input<inputs.network.v20200301.SubResource>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * IP configuration of the Azure Firewall resource.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200301.AzureFirewallIPConfiguration>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * IP configuration of the Azure Firewall used for management traffic.
     */
    readonly managementIpConfiguration?: pulumi.Input<inputs.network.v20200301.AzureFirewallIPConfiguration>;
    /**
     * Collection of NAT rule collections used by Azure Firewall.
     */
    readonly natRuleCollections?: pulumi.Input<pulumi.Input<inputs.network.v20200301.AzureFirewallNatRuleCollection>[]>;
    /**
     * Collection of network rule collections used by Azure Firewall.
     */
    readonly networkRuleCollections?: pulumi.Input<pulumi.Input<inputs.network.v20200301.AzureFirewallNetworkRuleCollection>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Azure Firewall Resource SKU.
     */
    readonly sku?: pulumi.Input<inputs.network.v20200301.AzureFirewallSku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The operation mode for Threat Intelligence.
     */
    readonly threatIntelMode?: pulumi.Input<string>;
    /**
     * The virtualHub to which the firewall belongs.
     */
    readonly virtualHub?: pulumi.Input<inputs.network.v20200301.SubResource>;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
