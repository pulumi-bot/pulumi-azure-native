// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Virtual Appliance Site resource.
 */
export class VirtualHubBgpConnection extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHubBgpConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualHubBgpConnection {
        return new VirtualHubBgpConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20200701:VirtualHubBgpConnection';

    /**
     * Returns true if the given object is an instance of VirtualHubBgpConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHubBgpConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHubBgpConnection.__pulumiType;
    }

    /**
     * The current state of the VirtualHub to Peer.
     */
    public /*out*/ readonly connectionState!: pulumi.Output<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Name of the connection.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Peer ASN.
     */
    public readonly peerAsn!: pulumi.Output<number | undefined>;
    /**
     * Peer IP.
     */
    public readonly peerIp!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Connection type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VirtualHubBgpConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubBgpConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.connectionName === undefined) {
                throw new Error("Missing required property 'connectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualHubName === undefined) {
                throw new Error("Missing required property 'virtualHubName'");
            }
            inputs["connectionName"] = args ? args.connectionName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peerAsn"] = args ? args.peerAsn : undefined;
            inputs["peerIp"] = args ? args.peerIp : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["virtualHubName"] = args ? args.virtualHubName : undefined;
            inputs["connectionState"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["connectionState"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peerAsn"] = undefined /*out*/;
            inputs["peerIp"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:VirtualHubBgpConnection" }, { type: "azure-nextgen:network/v20200501:VirtualHubBgpConnection" }, { type: "azure-nextgen:network/v20200601:VirtualHubBgpConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualHubBgpConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualHubBgpConnection resource.
 */
export interface VirtualHubBgpConnectionArgs {
    /**
     * The name of the connection.
     */
    readonly connectionName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Name of the connection.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Peer ASN.
     */
    readonly peerAsn?: pulumi.Input<number>;
    /**
     * Peer IP.
     */
    readonly peerIp?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the VirtualHub.
     */
    readonly virtualHubName: pulumi.Input<string>;
}
