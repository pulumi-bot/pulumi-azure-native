// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * VpnConnection Resource.
 */
export class VpnGatewayVpnConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpnGatewayVpnConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VpnGatewayVpnConnection {
        return new VpnGatewayVpnConnection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network:VpnGatewayVpnConnection';

    /**
     * Returns true if the given object is an instance of VpnGatewayVpnConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnGatewayVpnConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnGatewayVpnConnection.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Properties of the VPN connection.
     */
    public readonly properties!: pulumi.Output<outputs.network.VpnConnectionPropertiesResponse>;

    /**
     * Create a VpnGatewayVpnConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpnGatewayVpnConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.connectionName === undefined) {
                throw new Error("Missing required property 'connectionName'");
            }
            if (!args || args.gatewayName === undefined) {
                throw new Error("Missing required property 'gatewayName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["connectionName"] = args ? args.connectionName : undefined;
        inputs["gatewayName"] = args ? args.gatewayName : undefined;
        inputs["id"] = args ? args.id : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["etag"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpnGatewayVpnConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VpnGatewayVpnConnection resource.
 */
export interface VpnGatewayVpnConnectionArgs {
    /**
     * The name of the connection.
     */
    readonly connectionName: pulumi.Input<string>;
    /**
     * The name of the gateway.
     */
    readonly gatewayName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Properties of the VPN connection.
     */
    readonly properties?: pulumi.Input<inputs.network.VpnConnectionProperties>;
    /**
     * The resource group name of the VpnGateway.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}