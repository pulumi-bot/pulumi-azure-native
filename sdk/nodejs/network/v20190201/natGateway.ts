// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Nat Gateway resource.
 *
 * ## Example Usage
 * ### Create nat gateway
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const natGateway = new azurerm.network.v20190201.NatGateway("natGateway", {
 *     natGatewayName: "test-natgateway",
 *     publicIpAddresses: [{
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1",
 *     }],
 *     publicIpPrefixes: [{
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1",
 *     }],
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class NatGateway extends pulumi.CustomResource {
    /**
     * Get an existing NatGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NatGateway {
        return new NatGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190201:NatGateway';

    /**
     * Returns true if the given object is an instance of NatGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatGateway.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The idle timeout of the nat gateway.
     */
    public readonly idleTimeoutInMinutes!: pulumi.Output<number | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * An array of public ip addresses associated with the nat gateway resource.
     */
    public readonly publicIpAddresses!: pulumi.Output<outputs.network.v20190201.SubResourceResponse[] | undefined>;
    /**
     * An array of public ip prefixes associated with the nat gateway resource.
     */
    public readonly publicIpPrefixes!: pulumi.Output<outputs.network.v20190201.SubResourceResponse[] | undefined>;
    /**
     * The resource GUID property of the nat gateway resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * The nat gateway SKU.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20190201.NatGatewaySkuResponse | undefined>;
    /**
     * An array of references to the subnets using this nat gateway resource.
     */
    public /*out*/ readonly subnets!: pulumi.Output<outputs.network.v20190201.SubResourceResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NatGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.natGatewayName === undefined) {
                throw new Error("Missing required property 'natGatewayName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["idleTimeoutInMinutes"] = args ? args.idleTimeoutInMinutes : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["natGatewayName"] = args ? args.natGatewayName : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["publicIpAddresses"] = args ? args.publicIpAddresses : undefined;
            inputs["publicIpPrefixes"] = args ? args.publicIpPrefixes : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["idleTimeoutInMinutes"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicIpAddresses"] = undefined /*out*/;
            inputs["publicIpPrefixes"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:NatGateway" }, { type: "azurerm:network/v20190401:NatGateway" }, { type: "azurerm:network/v20190601:NatGateway" }, { type: "azurerm:network/v20190701:NatGateway" }, { type: "azurerm:network/v20190801:NatGateway" }, { type: "azurerm:network/v20190901:NatGateway" }, { type: "azurerm:network/v20191101:NatGateway" }, { type: "azurerm:network/v20191201:NatGateway" }, { type: "azurerm:network/v20200301:NatGateway" }, { type: "azurerm:network/v20200401:NatGateway" }, { type: "azurerm:network/v20200501:NatGateway" }, { type: "azurerm:network/v20200601:NatGateway" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NatGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NatGateway resource.
 */
export interface NatGatewayArgs {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The idle timeout of the nat gateway.
     */
    readonly idleTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the nat gateway.
     */
    readonly natGatewayName: pulumi.Input<string>;
    /**
     * The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * An array of public ip addresses associated with the nat gateway resource.
     */
    readonly publicIpAddresses?: pulumi.Input<pulumi.Input<inputs.network.v20190201.SubResource>[]>;
    /**
     * An array of public ip prefixes associated with the nat gateway resource.
     */
    readonly publicIpPrefixes?: pulumi.Input<pulumi.Input<inputs.network.v20190201.SubResource>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the nat gateway resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * The nat gateway SKU.
     */
    readonly sku?: pulumi.Input<inputs.network.v20190201.NatGatewaySku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
