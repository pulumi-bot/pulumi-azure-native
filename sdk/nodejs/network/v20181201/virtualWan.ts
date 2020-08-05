// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * VirtualWAN Resource.
 */
export class VirtualWan extends pulumi.CustomResource {
    /**
     * Get an existing VirtualWan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualWan {
        return new VirtualWan(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20181201:VirtualWan';

    /**
     * Returns true if the given object is an instance of VirtualWan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualWan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualWan.__pulumiType;
    }

    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Parameters for VirtualWAN
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20181201.VirtualWanPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VirtualWan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualWanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualWanArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualWanArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["allowBranchToBranchTraffic"] = args ? args.allowBranchToBranchTraffic : undefined;
            inputs["allowVnetToVnetTraffic"] = args ? args.allowVnetToVnetTraffic : undefined;
            inputs["disableVpnEncryption"] = args ? args.disableVpnEncryption : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["office365LocalBreakoutCategory"] = args ? args.office365LocalBreakoutCategory : undefined;
            inputs["p2SVpnServerConfigurations"] = args ? args.p2SVpnServerConfigurations : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityProviderName"] = args ? args.securityProviderName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualWan.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualWan resource.
 */
export interface VirtualWanArgs {
    /**
     * True if branch to branch traffic is allowed.
     */
    readonly allowBranchToBranchTraffic?: pulumi.Input<boolean>;
    /**
     * True if Vnet to Vnet traffic is allowed.
     */
    readonly allowVnetToVnetTraffic?: pulumi.Input<boolean>;
    /**
     * Vpn encryption to be disabled or not.
     */
    readonly disableVpnEncryption?: pulumi.Input<boolean>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the VirtualWAN being created or updated.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The office local breakout category.
     */
    readonly office365LocalBreakoutCategory?: pulumi.Input<string>;
    /**
     * List of all P2SVpnServerConfigurations associated with the virtual wan.
     */
    readonly p2SVpnServerConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20181201.P2SVpnServerConfiguration>[]>;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualWan.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Security Provider name.
     */
    readonly securityProviderName?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
