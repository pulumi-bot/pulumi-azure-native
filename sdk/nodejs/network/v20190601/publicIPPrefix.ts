// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Public IP prefix resource.
 *
 * ## Example Usage
 * ### Create public IP prefix allocation method
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const publicIPPrefix = new azurerm.network.v20190601.PublicIPPrefix("publicIPPrefix", {
 *     location: "westus",
 *     prefixLength: 30,
 *     publicIPAddressVersion: "IPv4",
 *     publicIpPrefixName: "test-ipprefix",
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "Standard",
 *     },
 * });
 *
 * ```
 * ### Create public IP prefix defaults
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const publicIPPrefix = new azurerm.network.v20190601.PublicIPPrefix("publicIPPrefix", {
 *     location: "westus",
 *     prefixLength: 30,
 *     publicIpPrefixName: "test-ipprefix",
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "Standard",
 *     },
 * });
 *
 * ```
 */
export class PublicIPPrefix extends pulumi.CustomResource {
    /**
     * Get an existing PublicIPPrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PublicIPPrefix {
        return new PublicIPPrefix(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190601:PublicIPPrefix';

    /**
     * Returns true if the given object is an instance of PublicIPPrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicIPPrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicIPPrefix.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The allocated Prefix.
     */
    public readonly ipPrefix!: pulumi.Output<string | undefined>;
    /**
     * The list of tags associated with the public IP prefix.
     */
    public readonly ipTags!: pulumi.Output<outputs.network.v20190601.IpTagResponse[] | undefined>;
    /**
     * The reference to load balancer frontend IP configuration associated with the public IP prefix.
     */
    public /*out*/ readonly loadBalancerFrontendIpConfiguration!: pulumi.Output<outputs.network.v20190601.SubResourceResponse>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Length of the Public IP Prefix.
     */
    public readonly prefixLength!: pulumi.Output<number | undefined>;
    /**
     * The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The public IP address version.
     */
    public readonly publicIPAddressVersion!: pulumi.Output<string | undefined>;
    /**
     * The list of all referenced PublicIPAddresses.
     */
    public readonly publicIPAddresses!: pulumi.Output<outputs.network.v20190601.ReferencedPublicIpAddressResponse[] | undefined>;
    /**
     * The resource GUID property of the public IP prefix resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * The public IP prefix SKU.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20190601.PublicIPPrefixSkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a PublicIPPrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicIPPrefixArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicIPPrefixArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PublicIPPrefixArgs | undefined;
            if (!args || args.publicIpPrefixName === undefined) {
                throw new Error("Missing required property 'publicIpPrefixName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipPrefix"] = args ? args.ipPrefix : undefined;
            inputs["ipTags"] = args ? args.ipTags : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["prefixLength"] = args ? args.prefixLength : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["publicIPAddressVersion"] = args ? args.publicIPAddressVersion : undefined;
            inputs["publicIPAddresses"] = args ? args.publicIPAddresses : undefined;
            inputs["publicIpPrefixName"] = args ? args.publicIpPrefixName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["loadBalancerFrontendIpConfiguration"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:PublicIPPrefix" }, { type: "azurerm:network/v20180701:PublicIPPrefix" }, { type: "azurerm:network/v20180801:PublicIPPrefix" }, { type: "azurerm:network/v20181001:PublicIPPrefix" }, { type: "azurerm:network/v20181101:PublicIPPrefix" }, { type: "azurerm:network/v20181201:PublicIPPrefix" }, { type: "azurerm:network/v20190201:PublicIPPrefix" }, { type: "azurerm:network/v20190401:PublicIPPrefix" }, { type: "azurerm:network/v20190701:PublicIPPrefix" }, { type: "azurerm:network/v20190801:PublicIPPrefix" }, { type: "azurerm:network/v20190901:PublicIPPrefix" }, { type: "azurerm:network/v20191101:PublicIPPrefix" }, { type: "azurerm:network/v20191201:PublicIPPrefix" }, { type: "azurerm:network/v20200301:PublicIPPrefix" }, { type: "azurerm:network/v20200401:PublicIPPrefix" }, { type: "azurerm:network/v20200501:PublicIPPrefix" }, { type: "azurerm:network/v20200601:PublicIPPrefix" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PublicIPPrefix.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PublicIPPrefix resource.
 */
export interface PublicIPPrefixArgs {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The allocated Prefix.
     */
    readonly ipPrefix?: pulumi.Input<string>;
    /**
     * The list of tags associated with the public IP prefix.
     */
    readonly ipTags?: pulumi.Input<pulumi.Input<inputs.network.v20190601.IpTag>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Length of the Public IP Prefix.
     */
    readonly prefixLength?: pulumi.Input<number>;
    /**
     * The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The public IP address version.
     */
    readonly publicIPAddressVersion?: pulumi.Input<string>;
    /**
     * The list of all referenced PublicIPAddresses.
     */
    readonly publicIPAddresses?: pulumi.Input<pulumi.Input<inputs.network.v20190601.ReferencedPublicIpAddress>[]>;
    /**
     * The name of the public IP prefix.
     */
    readonly publicIpPrefixName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the public IP prefix resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * The public IP prefix SKU.
     */
    readonly sku?: pulumi.Input<inputs.network.v20190601.PublicIPPrefixSku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
