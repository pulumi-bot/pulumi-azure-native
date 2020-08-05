// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Peering in an ExpressRouteCircuit resource.
 */
export class ExpressRouteCircuitPeering extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExpressRouteCircuitPeering {
        return new ExpressRouteCircuitPeering(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190201:ExpressRouteCircuitPeering';

    /**
     * Returns true if the given object is an instance of ExpressRouteCircuitPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRouteCircuitPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRouteCircuitPeering.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Properties of the express route circuit peering.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20190201.ExpressRouteCircuitPeeringPropertiesFormatResponse>;

    /**
     * Create a ExpressRouteCircuitPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRouteCircuitPeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExpressRouteCircuitPeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ExpressRouteCircuitPeeringArgs | undefined;
            if (!args || args.circuitName === undefined) {
                throw new Error("Missing required property 'circuitName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["azureASN"] = args ? args.azureASN : undefined;
            inputs["circuitName"] = args ? args.circuitName : undefined;
            inputs["connections"] = args ? args.connections : undefined;
            inputs["expressRouteConnection"] = args ? args.expressRouteConnection : undefined;
            inputs["gatewayManagerEtag"] = args ? args.gatewayManagerEtag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipv6PeeringConfig"] = args ? args.ipv6PeeringConfig : undefined;
            inputs["lastModifiedBy"] = args ? args.lastModifiedBy : undefined;
            inputs["microsoftPeeringConfig"] = args ? args.microsoftPeeringConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peerASN"] = args ? args.peerASN : undefined;
            inputs["peeringType"] = args ? args.peeringType : undefined;
            inputs["primaryAzurePort"] = args ? args.primaryAzurePort : undefined;
            inputs["primaryPeerAddressPrefix"] = args ? args.primaryPeerAddressPrefix : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeFilter"] = args ? args.routeFilter : undefined;
            inputs["secondaryAzurePort"] = args ? args.secondaryAzurePort : undefined;
            inputs["secondaryPeerAddressPrefix"] = args ? args.secondaryPeerAddressPrefix : undefined;
            inputs["sharedKey"] = args ? args.sharedKey : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["stats"] = args ? args.stats : undefined;
            inputs["vlanId"] = args ? args.vlanId : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ExpressRouteCircuitPeering.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExpressRouteCircuitPeering resource.
 */
export interface ExpressRouteCircuitPeeringArgs {
    /**
     * The Azure ASN.
     */
    readonly azureASN?: pulumi.Input<number>;
    /**
     * The name of the express route circuit.
     */
    readonly circuitName: pulumi.Input<string>;
    /**
     * The list of circuit connections associated with Azure Private Peering for this circuit.
     */
    readonly connections?: pulumi.Input<pulumi.Input<inputs.network.v20190201.ExpressRouteCircuitConnection>[]>;
    /**
     * The ExpressRoute connection.
     */
    readonly expressRouteConnection?: pulumi.Input<inputs.network.v20190201.ExpressRouteConnectionId>;
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
    readonly ipv6PeeringConfig?: pulumi.Input<inputs.network.v20190201.Ipv6ExpressRouteCircuitPeeringConfig>;
    /**
     * Gets whether the provider or the customer last modified the peering.
     */
    readonly lastModifiedBy?: pulumi.Input<string>;
    /**
     * The Microsoft peering configuration.
     */
    readonly microsoftPeeringConfig?: pulumi.Input<inputs.network.v20190201.ExpressRouteCircuitPeeringConfig>;
    /**
     * The name of the peering.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The peer ASN.
     */
    readonly peerASN?: pulumi.Input<number>;
    /**
     * The peering type.
     */
    readonly peeringType?: pulumi.Input<string>;
    /**
     * The primary port.
     */
    readonly primaryAzurePort?: pulumi.Input<string>;
    /**
     * The primary address prefix.
     */
    readonly primaryPeerAddressPrefix?: pulumi.Input<string>;
    /**
     * Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The reference of the RouteFilter resource.
     */
    readonly routeFilter?: pulumi.Input<inputs.network.v20190201.RouteFilter>;
    /**
     * The secondary port.
     */
    readonly secondaryAzurePort?: pulumi.Input<string>;
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
     * Gets peering stats.
     */
    readonly stats?: pulumi.Input<inputs.network.v20190201.ExpressRouteCircuitStats>;
    /**
     * The VLAN ID.
     */
    readonly vlanId?: pulumi.Input<number>;
}
