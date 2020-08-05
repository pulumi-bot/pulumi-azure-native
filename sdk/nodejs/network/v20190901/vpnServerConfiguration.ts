// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * VpnServerConfiguration Resource.
 */
export class VpnServerConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing VpnServerConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VpnServerConfiguration {
        return new VpnServerConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190901:VpnServerConfiguration';

    /**
     * Returns true if the given object is an instance of VpnServerConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnServerConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnServerConfiguration.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the P2SVpnServer configuration.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20190901.VpnServerConfigurationPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VpnServerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VpnServerConfigurationArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["aadAuthenticationParameters"] = args ? args.aadAuthenticationParameters : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["radiusClientRootCertificates"] = args ? args.radiusClientRootCertificates : undefined;
            inputs["radiusServerAddress"] = args ? args.radiusServerAddress : undefined;
            inputs["radiusServerRootCertificates"] = args ? args.radiusServerRootCertificates : undefined;
            inputs["radiusServerSecret"] = args ? args.radiusServerSecret : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpnAuthenticationTypes"] = args ? args.vpnAuthenticationTypes : undefined;
            inputs["vpnClientIpsecPolicies"] = args ? args.vpnClientIpsecPolicies : undefined;
            inputs["vpnClientRevokedCertificates"] = args ? args.vpnClientRevokedCertificates : undefined;
            inputs["vpnClientRootCertificates"] = args ? args.vpnClientRootCertificates : undefined;
            inputs["vpnProtocols"] = args ? args.vpnProtocols : undefined;
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
        super(VpnServerConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VpnServerConfiguration resource.
 */
export interface VpnServerConfigurationArgs {
    /**
     * The set of aad vpn authentication parameters.
     */
    readonly aadAuthenticationParameters?: pulumi.Input<inputs.network.v20190901.AadAuthenticationParameters>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the VpnServerConfiguration being created or updated.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Radius client root certificate of VpnServerConfiguration.
     */
    readonly radiusClientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20190901.VpnServerConfigRadiusClientRootCertificate>[]>;
    /**
     * The radius server address property of the VpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerAddress?: pulumi.Input<string>;
    /**
     * Radius Server root certificate of VpnServerConfiguration.
     */
    readonly radiusServerRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20190901.VpnServerConfigRadiusServerRootCertificate>[]>;
    /**
     * The radius secret property of the VpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerSecret?: pulumi.Input<string>;
    /**
     * The resource group name of the VpnServerConfiguration.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * VPN authentication types for the VpnServerConfiguration.
     */
    readonly vpnAuthenticationTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VpnClientIpsecPolicies for VpnServerConfiguration.
     */
    readonly vpnClientIpsecPolicies?: pulumi.Input<pulumi.Input<inputs.network.v20190901.IpsecPolicy>[]>;
    /**
     * VPN client revoked certificate of VpnServerConfiguration.
     */
    readonly vpnClientRevokedCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20190901.VpnServerConfigVpnClientRevokedCertificate>[]>;
    /**
     * VPN client root certificate of VpnServerConfiguration.
     */
    readonly vpnClientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20190901.VpnServerConfigVpnClientRootCertificate>[]>;
    /**
     * VPN protocols for the VpnServerConfiguration.
     */
    readonly vpnProtocols?: pulumi.Input<pulumi.Input<string>[]>;
}
