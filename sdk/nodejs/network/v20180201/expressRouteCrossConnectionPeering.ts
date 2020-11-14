// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Peering in an ExpressRoute Cross Connection resource.
 */
export class ExpressRouteCrossConnectionPeering extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRouteCrossConnectionPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExpressRouteCrossConnectionPeering {
        return new ExpressRouteCrossConnectionPeering(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20180201:ExpressRouteCrossConnectionPeering';

    /**
     * Returns true if the given object is an instance of ExpressRouteCrossConnectionPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRouteCrossConnectionPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRouteCrossConnectionPeering.__pulumiType;
    }

    /**
     * The Azure ASN.
     */
    public /*out*/ readonly azureASN!: pulumi.Output<number>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The GatewayManager Etag.
     */
    public readonly gatewayManagerEtag!: pulumi.Output<string | undefined>;
    /**
     * The IPv6 peering configuration.
     */
    public readonly ipv6PeeringConfig!: pulumi.Output<outputs.network.v20180201.Ipv6ExpressRouteCircuitPeeringConfigResponse | undefined>;
    /**
     * Gets whether the provider or the customer last modified the peering.
     */
    public readonly lastModifiedBy!: pulumi.Output<string | undefined>;
    /**
     * The Microsoft peering configuration.
     */
    public readonly microsoftPeeringConfig!: pulumi.Output<outputs.network.v20180201.ExpressRouteCircuitPeeringConfigResponse | undefined>;
    /**
     * Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The peer ASN.
     */
    public readonly peerASN!: pulumi.Output<number | undefined>;
    /**
     * The peering type.
     */
    public readonly peeringType!: pulumi.Output<string | undefined>;
    /**
     * The primary port.
     */
    public /*out*/ readonly primaryAzurePort!: pulumi.Output<string>;
    /**
     * The primary address prefix.
     */
    public readonly primaryPeerAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The secondary port.
     */
    public /*out*/ readonly secondaryAzurePort!: pulumi.Output<string>;
    /**
     * The secondary address prefix.
     */
    public readonly secondaryPeerAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * The shared key.
     */
    public readonly sharedKey!: pulumi.Output<string | undefined>;
    /**
     * The peering state.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * The VLAN ID.
     */
    public readonly vlanId!: pulumi.Output<number | undefined>;

    /**
     * Create a ExpressRouteCrossConnectionPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRouteCrossConnectionPeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.crossConnectionName === undefined) {
                throw new Error("Missing required property 'crossConnectionName'");
            }
            if (!args || args.peeringName === undefined) {
                throw new Error("Missing required property 'peeringName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["crossConnectionName"] = args ? args.crossConnectionName : undefined;
            inputs["gatewayManagerEtag"] = args ? args.gatewayManagerEtag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipv6PeeringConfig"] = args ? args.ipv6PeeringConfig : undefined;
            inputs["lastModifiedBy"] = args ? args.lastModifiedBy : undefined;
            inputs["microsoftPeeringConfig"] = args ? args.microsoftPeeringConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peerASN"] = args ? args.peerASN : undefined;
            inputs["peeringName"] = args ? args.peeringName : undefined;
            inputs["peeringType"] = args ? args.peeringType : undefined;
            inputs["primaryPeerAddressPrefix"] = args ? args.primaryPeerAddressPrefix : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["secondaryPeerAddressPrefix"] = args ? args.secondaryPeerAddressPrefix : undefined;
            inputs["sharedKey"] = args ? args.sharedKey : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["vlanId"] = args ? args.vlanId : undefined;
            inputs["azureASN"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["primaryAzurePort"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["secondaryAzurePort"] = undefined /*out*/;
        } else {
            inputs["azureASN"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["gatewayManagerEtag"] = undefined /*out*/;
            inputs["ipv6PeeringConfig"] = undefined /*out*/;
            inputs["lastModifiedBy"] = undefined /*out*/;
            inputs["microsoftPeeringConfig"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peerASN"] = undefined /*out*/;
            inputs["peeringType"] = undefined /*out*/;
            inputs["primaryAzurePort"] = undefined /*out*/;
            inputs["primaryPeerAddressPrefix"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["secondaryAzurePort"] = undefined /*out*/;
            inputs["secondaryPeerAddressPrefix"] = undefined /*out*/;
            inputs["sharedKey"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["vlanId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20180401:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20180601:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20180701:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20180801:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20181001:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20181101:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20181201:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20190201:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20190401:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20190601:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20190701:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20190801:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20190901:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20191101:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20191201:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20200301:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20200401:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20200501:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20200601:ExpressRouteCrossConnectionPeering" }, { type: "azure-nextgen:network/v20200701:ExpressRouteCrossConnectionPeering" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ExpressRouteCrossConnectionPeering.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExpressRouteCrossConnectionPeering resource.
 */
export interface ExpressRouteCrossConnectionPeeringArgs {
    /**
     * The name of the ExpressRouteCrossConnection.
     */
    readonly crossConnectionName: pulumi.Input<string>;
    /**
     * The GatewayManager Etag.
     */
    readonly gatewayManagerEtag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The IPv6 peering configuration.
     */
    readonly ipv6PeeringConfig?: pulumi.Input<inputs.network.v20180201.Ipv6ExpressRouteCircuitPeeringConfig>;
    /**
     * Gets whether the provider or the customer last modified the peering.
     */
    readonly lastModifiedBy?: pulumi.Input<string>;
    /**
     * The Microsoft peering configuration.
     */
    readonly microsoftPeeringConfig?: pulumi.Input<inputs.network.v20180201.ExpressRouteCircuitPeeringConfig>;
    /**
     * Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The peer ASN.
     */
    readonly peerASN?: pulumi.Input<number>;
    /**
     * The name of the peering.
     */
    readonly peeringName: pulumi.Input<string>;
    /**
     * The peering type.
     */
    readonly peeringType?: pulumi.Input<string>;
    /**
     * The primary address prefix.
     */
    readonly primaryPeerAddressPrefix?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The secondary address prefix.
     */
    readonly secondaryPeerAddressPrefix?: pulumi.Input<string>;
    /**
     * The shared key.
     */
    readonly sharedKey?: pulumi.Input<string>;
    /**
     * The peering state.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The VLAN ID.
     */
    readonly vlanId?: pulumi.Input<number>;
}
