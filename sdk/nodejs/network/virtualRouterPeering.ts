// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Virtual Router Peering resource.
 */
export class VirtualRouterPeering extends pulumi.CustomResource {
    /**
     * Get an existing VirtualRouterPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualRouterPeering {
        return new VirtualRouterPeering(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network:VirtualRouterPeering';

    /**
     * Returns true if the given object is an instance of VirtualRouterPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualRouterPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualRouterPeering.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Name of the virtual router peering that is unique within a virtual router.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The properties of the Virtual Router Peering.
     */
    public readonly properties!: pulumi.Output<outputs.network.VirtualRouterPeeringPropertiesResponse>;
    /**
     * Peering type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VirtualRouterPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VirtualRouterPeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.peeringName === undefined) {
                throw new Error("Missing required property 'peeringName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualRouterName === undefined) {
                throw new Error("Missing required property 'virtualRouterName'");
            }
        inputs["id"] = args ? args.id : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["peeringName"] = args ? args.peeringName : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["virtualRouterName"] = args ? args.virtualRouterName : undefined;
        inputs["etag"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualRouterPeering.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualRouterPeering resource.
 */
export interface VirtualRouterPeeringArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Name of the virtual router peering that is unique within a virtual router.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Virtual Router Peering.
     */
    readonly peeringName: pulumi.Input<string>;
    /**
     * The properties of the Virtual Router Peering.
     */
    readonly properties?: pulumi.Input<inputs.network.VirtualRouterPeeringProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Virtual Router.
     */
    readonly virtualRouterName: pulumi.Input<string>;
}
