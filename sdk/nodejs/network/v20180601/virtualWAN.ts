// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * VirtualWAN Resource.
 */
export class VirtualWAN extends pulumi.CustomResource {
    /**
     * Get an existing VirtualWAN resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualWAN {
        return new VirtualWAN(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20180601:VirtualWAN';

    /**
     * Returns true if the given object is an instance of VirtualWAN.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualWAN {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualWAN.__pulumiType;
    }

    /**
     * Vpn encryption to be disabled or not.
     */
    public readonly disableVpnEncryption!: pulumi.Output<boolean | undefined>;
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
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * List of VirtualHubs in the VirtualWAN.
     */
    public /*out*/ readonly virtualHubs!: pulumi.Output<outputs.network.v20180601.SubResourceResponse[]>;
    public /*out*/ readonly vpnSites!: pulumi.Output<outputs.network.v20180601.SubResourceResponse[]>;

    /**
     * Create a VirtualWAN resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualWANArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.virtualWANName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'virtualWANName'");
            }
            inputs["disableVpnEncryption"] = args ? args.disableVpnEncryption : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualWANName"] = args ? args.virtualWANName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualHubs"] = undefined /*out*/;
            inputs["vpnSites"] = undefined /*out*/;
        } else {
            inputs["disableVpnEncryption"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualHubs"] = undefined /*out*/;
            inputs["vpnSites"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:VirtualWAN" }, { type: "azure-nextgen:network/v20180401:VirtualWAN" }, { type: "azure-nextgen:network/v20180701:VirtualWAN" }, { type: "azure-nextgen:network/v20180801:VirtualWAN" }, { type: "azure-nextgen:network/v20181001:VirtualWAN" }, { type: "azure-nextgen:network/v20181101:VirtualWAN" }, { type: "azure-nextgen:network/v20181201:VirtualWAN" }, { type: "azure-nextgen:network/v20190201:VirtualWAN" }, { type: "azure-nextgen:network/v20190401:VirtualWAN" }, { type: "azure-nextgen:network/v20190601:VirtualWAN" }, { type: "azure-nextgen:network/v20190701:VirtualWAN" }, { type: "azure-nextgen:network/v20190801:VirtualWAN" }, { type: "azure-nextgen:network/v20190901:VirtualWAN" }, { type: "azure-nextgen:network/v20191101:VirtualWAN" }, { type: "azure-nextgen:network/v20191201:VirtualWAN" }, { type: "azure-nextgen:network/v20200301:VirtualWAN" }, { type: "azure-nextgen:network/v20200401:VirtualWAN" }, { type: "azure-nextgen:network/v20200501:VirtualWAN" }, { type: "azure-nextgen:network/v20200601:VirtualWAN" }, { type: "azure-nextgen:network/v20200701:VirtualWAN" }, { type: "azure-nextgen:network/v20200801:VirtualWAN" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualWAN.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualWAN resource.
 */
export interface VirtualWANArgs {
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
    readonly location?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualWan.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the VirtualWAN being created or updated.
     */
    readonly virtualWANName: pulumi.Input<string>;
}
