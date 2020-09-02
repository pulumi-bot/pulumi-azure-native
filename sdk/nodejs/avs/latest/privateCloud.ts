// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A private cloud resource
 *
 * ## Example Usage
 * ### PrivateClouds_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const privateCloud = new azurerm.avs.latest.PrivateCloud("privateCloud", {
 *     location: "eastus2",
 *     managementCluster: {
 *         clusterSize: 4,
 *     },
 *     networkBlock: "192.168.48.0/22",
 *     privateCloudName: "cloud1",
 *     resourceGroupName: "group1",
 *     sku: {
 *         name: "AV36",
 *     },
 *     tags: {},
 * });
 *
 * ```
 */
export class PrivateCloud extends pulumi.CustomResource {
    /**
     * Get an existing PrivateCloud resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateCloud {
        return new PrivateCloud(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:avs/latest:PrivateCloud';

    /**
     * Returns true if the given object is an instance of PrivateCloud.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateCloud {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateCloud.__pulumiType;
    }

    /**
     * An ExpressRoute Circuit
     */
    public /*out*/ readonly circuit!: pulumi.Output<outputs.avs.latest.CircuitResponse | undefined>;
    /**
     * The endpoints
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.avs.latest.EndpointsResponse>;
    /**
     * vCenter Single Sign On Identity Sources
     */
    public readonly identitySources!: pulumi.Output<outputs.avs.latest.IdentitySourceResponse[] | undefined>;
    /**
     * Connectivity to internet is enabled or disabled
     */
    public readonly internet!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The default cluster used for management
     */
    public readonly managementCluster!: pulumi.Output<outputs.avs.latest.ManagementClusterResponse>;
    /**
     * Network used to access vCenter Server and NSX-T Manager
     */
    public /*out*/ readonly managementNetwork!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
     */
    public readonly networkBlock!: pulumi.Output<string>;
    /**
     * Thumbprint of the NSX-T Manager SSL certificate
     */
    public /*out*/ readonly nsxtCertificateThumbprint!: pulumi.Output<string>;
    /**
     * Optionally, set the NSX-T Manager password when the private cloud is created
     */
    public readonly nsxtPassword!: pulumi.Output<string | undefined>;
    /**
     * Used for virtual machine cold migration, cloning, and snapshot migration
     */
    public /*out*/ readonly provisioningNetwork!: pulumi.Output<string>;
    /**
     * The provisioning state
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The private cloud SKU
     */
    public readonly sku!: pulumi.Output<outputs.avs.latest.SkuResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Thumbprint of the vCenter Server SSL certificate
     */
    public /*out*/ readonly vcenterCertificateThumbprint!: pulumi.Output<string>;
    /**
     * Optionally, set the vCenter admin password when the private cloud is created
     */
    public readonly vcenterPassword!: pulumi.Output<string | undefined>;
    /**
     * Used for live migration of virtual machines
     */
    public /*out*/ readonly vmotionNetwork!: pulumi.Output<string>;

    /**
     * Create a PrivateCloud resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateCloudArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateCloudArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PrivateCloudArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.managementCluster === undefined) {
                throw new Error("Missing required property 'managementCluster'");
            }
            if (!args || args.networkBlock === undefined) {
                throw new Error("Missing required property 'networkBlock'");
            }
            if (!args || args.privateCloudName === undefined) {
                throw new Error("Missing required property 'privateCloudName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["identitySources"] = args ? args.identitySources : undefined;
            inputs["internet"] = args ? args.internet : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementCluster"] = args ? args.managementCluster : undefined;
            inputs["networkBlock"] = args ? args.networkBlock : undefined;
            inputs["nsxtPassword"] = args ? args.nsxtPassword : undefined;
            inputs["privateCloudName"] = args ? args.privateCloudName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vcenterPassword"] = args ? args.vcenterPassword : undefined;
            inputs["circuit"] = undefined /*out*/;
            inputs["endpoints"] = undefined /*out*/;
            inputs["managementNetwork"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nsxtCertificateThumbprint"] = undefined /*out*/;
            inputs["provisioningNetwork"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vcenterCertificateThumbprint"] = undefined /*out*/;
            inputs["vmotionNetwork"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:avs/v20200320:PrivateCloud" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PrivateCloud.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateCloud resource.
 */
export interface PrivateCloudArgs {
    /**
     * vCenter Single Sign On Identity Sources
     */
    readonly identitySources?: pulumi.Input<pulumi.Input<inputs.avs.latest.IdentitySource>[]>;
    /**
     * Connectivity to internet is enabled or disabled
     */
    readonly internet?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The default cluster used for management
     */
    readonly managementCluster: pulumi.Input<inputs.avs.latest.ManagementCluster>;
    /**
     * The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
     */
    readonly networkBlock: pulumi.Input<string>;
    /**
     * Optionally, set the NSX-T Manager password when the private cloud is created
     */
    readonly nsxtPassword?: pulumi.Input<string>;
    /**
     * Name of the private cloud
     */
    readonly privateCloudName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The private cloud SKU
     */
    readonly sku: pulumi.Input<inputs.avs.latest.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optionally, set the vCenter admin password when the private cloud is created
     */
    readonly vcenterPassword?: pulumi.Input<string>;
}
