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
        return new VpnServerConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:VpnServerConfiguration';

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
     * The set of aad vpn authentication parameters.
     */
    public readonly aadAuthenticationParameters!: pulumi.Output<outputs.network.v20200601.AadAuthenticationParametersResponse | undefined>;
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
     * List of references to P2SVpnGateways.
     */
    public /*out*/ readonly p2SVpnGateways!: pulumi.Output<outputs.network.v20200601.P2SVpnGatewayResponse[]>;
    /**
     * The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Radius client root certificate of VpnServerConfiguration.
     */
    public readonly radiusClientRootCertificates!: pulumi.Output<outputs.network.v20200601.VpnServerConfigRadiusClientRootCertificateResponse[] | undefined>;
    /**
     * The radius server address property of the VpnServerConfiguration resource for point to site client connection.
     */
    public readonly radiusServerAddress!: pulumi.Output<string | undefined>;
    /**
     * Radius Server root certificate of VpnServerConfiguration.
     */
    public readonly radiusServerRootCertificates!: pulumi.Output<outputs.network.v20200601.VpnServerConfigRadiusServerRootCertificateResponse[] | undefined>;
    /**
     * The radius secret property of the VpnServerConfiguration resource for point to site client connection.
     */
    public readonly radiusServerSecret!: pulumi.Output<string | undefined>;
    /**
     * Multiple Radius Server configuration for VpnServerConfiguration.
     */
    public readonly radiusServers!: pulumi.Output<outputs.network.v20200601.RadiusServerResponse[] | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * VPN authentication types for the VpnServerConfiguration.
     */
    public readonly vpnAuthenticationTypes!: pulumi.Output<string[] | undefined>;
    /**
     * VpnClientIpsecPolicies for VpnServerConfiguration.
     */
    public readonly vpnClientIpsecPolicies!: pulumi.Output<outputs.network.v20200601.IpsecPolicyResponse[] | undefined>;
    /**
     * VPN client revoked certificate of VpnServerConfiguration.
     */
    public readonly vpnClientRevokedCertificates!: pulumi.Output<outputs.network.v20200601.VpnServerConfigVpnClientRevokedCertificateResponse[] | undefined>;
    /**
     * VPN client root certificate of VpnServerConfiguration.
     */
    public readonly vpnClientRootCertificates!: pulumi.Output<outputs.network.v20200601.VpnServerConfigVpnClientRootCertificateResponse[] | undefined>;
    /**
     * VPN protocols for the VpnServerConfiguration.
     */
    public readonly vpnProtocols!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VpnServerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vpnServerConfigurationName === undefined) {
                throw new Error("Missing required property 'vpnServerConfigurationName'");
            }
            inputs["aadAuthenticationParameters"] = args ? args.aadAuthenticationParameters : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["radiusClientRootCertificates"] = args ? args.radiusClientRootCertificates : undefined;
            inputs["radiusServerAddress"] = args ? args.radiusServerAddress : undefined;
            inputs["radiusServerRootCertificates"] = args ? args.radiusServerRootCertificates : undefined;
            inputs["radiusServerSecret"] = args ? args.radiusServerSecret : undefined;
            inputs["radiusServers"] = args ? args.radiusServers : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpnAuthenticationTypes"] = args ? args.vpnAuthenticationTypes : undefined;
            inputs["vpnClientIpsecPolicies"] = args ? args.vpnClientIpsecPolicies : undefined;
            inputs["vpnClientRevokedCertificates"] = args ? args.vpnClientRevokedCertificates : undefined;
            inputs["vpnClientRootCertificates"] = args ? args.vpnClientRootCertificates : undefined;
            inputs["vpnProtocols"] = args ? args.vpnProtocols : undefined;
            inputs["vpnServerConfigurationName"] = args ? args.vpnServerConfigurationName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["p2SVpnGateways"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["aadAuthenticationParameters"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["p2SVpnGateways"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["radiusClientRootCertificates"] = undefined /*out*/;
            inputs["radiusServerAddress"] = undefined /*out*/;
            inputs["radiusServerRootCertificates"] = undefined /*out*/;
            inputs["radiusServerSecret"] = undefined /*out*/;
            inputs["radiusServers"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vpnAuthenticationTypes"] = undefined /*out*/;
            inputs["vpnClientIpsecPolicies"] = undefined /*out*/;
            inputs["vpnClientRevokedCertificates"] = undefined /*out*/;
            inputs["vpnClientRootCertificates"] = undefined /*out*/;
            inputs["vpnProtocols"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VpnServerConfiguration" }, { type: "azurerm:network/v20190801:VpnServerConfiguration" }, { type: "azurerm:network/v20190901:VpnServerConfiguration" }, { type: "azurerm:network/v20191101:VpnServerConfiguration" }, { type: "azurerm:network/v20191201:VpnServerConfiguration" }, { type: "azurerm:network/v20200301:VpnServerConfiguration" }, { type: "azurerm:network/v20200401:VpnServerConfiguration" }, { type: "azurerm:network/v20200501:VpnServerConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
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
    readonly aadAuthenticationParameters?: pulumi.Input<inputs.network.v20200601.AadAuthenticationParameters>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the VpnServerConfiguration that is unique within a resource group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Radius client root certificate of VpnServerConfiguration.
     */
    readonly radiusClientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.VpnServerConfigRadiusClientRootCertificate>[]>;
    /**
     * The radius server address property of the VpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerAddress?: pulumi.Input<string>;
    /**
     * Radius Server root certificate of VpnServerConfiguration.
     */
    readonly radiusServerRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.VpnServerConfigRadiusServerRootCertificate>[]>;
    /**
     * The radius secret property of the VpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerSecret?: pulumi.Input<string>;
    /**
     * Multiple Radius Server configuration for VpnServerConfiguration.
     */
    readonly radiusServers?: pulumi.Input<pulumi.Input<inputs.network.v20200601.RadiusServer>[]>;
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
    readonly vpnClientIpsecPolicies?: pulumi.Input<pulumi.Input<inputs.network.v20200601.IpsecPolicy>[]>;
    /**
     * VPN client revoked certificate of VpnServerConfiguration.
     */
    readonly vpnClientRevokedCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.VpnServerConfigVpnClientRevokedCertificate>[]>;
    /**
     * VPN client root certificate of VpnServerConfiguration.
     */
    readonly vpnClientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.VpnServerConfigVpnClientRootCertificate>[]>;
    /**
     * VPN protocols for the VpnServerConfiguration.
     */
    readonly vpnProtocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the VpnServerConfiguration being created or updated.
     */
    readonly vpnServerConfigurationName: pulumi.Input<string>;
}
