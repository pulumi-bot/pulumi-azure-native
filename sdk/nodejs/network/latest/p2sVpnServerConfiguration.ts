// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * P2SVpnServerConfiguration Resource.
 *
 * ## Example Usage
 * ### P2SVpnServerConfigurationPut
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const p2sVpnServerConfiguration = new azurerm.network.latest.P2sVpnServerConfiguration("p2sVpnServerConfiguration", {
 *     p2SVpnServerConfigRadiusClientRootCertificates: [{
 *         name: "p2sVpnServerConfigRadiusClientRootCert1",
 *     }],
 *     p2SVpnServerConfigRadiusServerRootCertificates: [{
 *         name: "p2sVpnServerConfigRadiusServerRootCert1",
 *     }],
 *     p2SVpnServerConfigVpnClientRevokedCertificates: [{
 *         name: "p2sVpnServerConfigVpnClientRevokedCert1",
 *     }],
 *     p2SVpnServerConfigVpnClientRootCertificates: [{
 *         name: "p2sVpnServerConfigVpnClientRootCert1",
 *     }],
 *     p2SVpnServerConfigurationName: "p2sVpnServerConfiguration1",
 *     radiusServerAddress: "8.9.9.9",
 *     radiusServerSecret: "123_abc",
 *     resourceGroupName: "rg1",
 *     virtualWanName: "virtualWan1",
 *     vpnClientIpsecPolicies: [{
 *         dhGroup: "DHGroup14",
 *         ikeEncryption: "AES256",
 *         ikeIntegrity: "SHA384",
 *         ipsecEncryption: "AES256",
 *         ipsecIntegrity: "SHA256",
 *         pfsGroup: "PFS14",
 *         saDataSizeKilobytes: 429497,
 *         saLifeTimeSeconds: 86472,
 *     }],
 *     vpnProtocols: ["IkeV2"],
 * });
 *
 * ```
 */
export class P2sVpnServerConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing P2sVpnServerConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): P2sVpnServerConfiguration {
        return new P2sVpnServerConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/latest:P2sVpnServerConfiguration';

    /**
     * Returns true if the given object is an instance of P2sVpnServerConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is P2sVpnServerConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === P2sVpnServerConfiguration.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * List of references to P2SVpnGateways.
     */
    public /*out*/ readonly p2SVpnGateways!: pulumi.Output<outputs.network.latest.SubResourceResponse[]>;
    /**
     * Radius client root certificate of P2SVpnServerConfiguration.
     */
    public readonly p2SVpnServerConfigRadiusClientRootCertificates!: pulumi.Output<outputs.network.latest.P2SVpnServerConfigRadiusClientRootCertificateResponse[] | undefined>;
    /**
     * Radius Server root certificate of P2SVpnServerConfiguration.
     */
    public readonly p2SVpnServerConfigRadiusServerRootCertificates!: pulumi.Output<outputs.network.latest.P2SVpnServerConfigRadiusServerRootCertificateResponse[] | undefined>;
    /**
     * VPN client revoked certificate of P2SVpnServerConfiguration.
     */
    public readonly p2SVpnServerConfigVpnClientRevokedCertificates!: pulumi.Output<outputs.network.latest.P2SVpnServerConfigVpnClientRevokedCertificateResponse[] | undefined>;
    /**
     * VPN client root certificate of P2SVpnServerConfiguration.
     */
    public readonly p2SVpnServerConfigVpnClientRootCertificates!: pulumi.Output<outputs.network.latest.P2SVpnServerConfigVpnClientRootCertificateResponse[] | undefined>;
    /**
     * The provisioning state of the P2S VPN server configuration resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
     */
    public readonly radiusServerAddress!: pulumi.Output<string | undefined>;
    /**
     * The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
     */
    public readonly radiusServerSecret!: pulumi.Output<string | undefined>;
    /**
     * VpnClientIpsecPolicies for P2SVpnServerConfiguration.
     */
    public readonly vpnClientIpsecPolicies!: pulumi.Output<outputs.network.latest.IpsecPolicyResponse[] | undefined>;
    /**
     * VPN protocols for the P2SVpnServerConfiguration.
     */
    public readonly vpnProtocols!: pulumi.Output<string[] | undefined>;

    /**
     * Create a P2sVpnServerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: P2sVpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: P2sVpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as P2sVpnServerConfigurationArgs | undefined;
            if (!args || args.p2SVpnServerConfigurationName === undefined) {
                throw new Error("Missing required property 'p2SVpnServerConfigurationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualWanName === undefined) {
                throw new Error("Missing required property 'virtualWanName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["p2SVpnServerConfigRadiusClientRootCertificates"] = args ? args.p2SVpnServerConfigRadiusClientRootCertificates : undefined;
            inputs["p2SVpnServerConfigRadiusServerRootCertificates"] = args ? args.p2SVpnServerConfigRadiusServerRootCertificates : undefined;
            inputs["p2SVpnServerConfigVpnClientRevokedCertificates"] = args ? args.p2SVpnServerConfigVpnClientRevokedCertificates : undefined;
            inputs["p2SVpnServerConfigVpnClientRootCertificates"] = args ? args.p2SVpnServerConfigVpnClientRootCertificates : undefined;
            inputs["p2SVpnServerConfigurationName"] = args ? args.p2SVpnServerConfigurationName : undefined;
            inputs["radiusServerAddress"] = args ? args.radiusServerAddress : undefined;
            inputs["radiusServerSecret"] = args ? args.radiusServerSecret : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["virtualWanName"] = args ? args.virtualWanName : undefined;
            inputs["vpnClientIpsecPolicies"] = args ? args.vpnClientIpsecPolicies : undefined;
            inputs["vpnProtocols"] = args ? args.vpnProtocols : undefined;
            inputs["p2SVpnGateways"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20180801:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20181001:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20181101:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20181201:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20190201:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20190401:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20190601:P2sVpnServerConfiguration" }, { type: "azurerm:network/v20190701:P2sVpnServerConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(P2sVpnServerConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a P2sVpnServerConfiguration resource.
 */
export interface P2sVpnServerConfigurationArgs {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Radius client root certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigRadiusClientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.latest.P2SVpnServerConfigRadiusClientRootCertificate>[]>;
    /**
     * Radius Server root certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigRadiusServerRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.latest.P2SVpnServerConfigRadiusServerRootCertificate>[]>;
    /**
     * VPN client revoked certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigVpnClientRevokedCertificates?: pulumi.Input<pulumi.Input<inputs.network.latest.P2SVpnServerConfigVpnClientRevokedCertificate>[]>;
    /**
     * VPN client root certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigVpnClientRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.latest.P2SVpnServerConfigVpnClientRootCertificate>[]>;
    /**
     * The name of the P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigurationName: pulumi.Input<string>;
    /**
     * The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerAddress?: pulumi.Input<string>;
    /**
     * The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerSecret?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualWan.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the VirtualWan.
     */
    readonly virtualWanName: pulumi.Input<string>;
    /**
     * VpnClientIpsecPolicies for P2SVpnServerConfiguration.
     */
    readonly vpnClientIpsecPolicies?: pulumi.Input<pulumi.Input<inputs.network.latest.IpsecPolicy>[]>;
    /**
     * VPN protocols for the P2SVpnServerConfiguration.
     */
    readonly vpnProtocols?: pulumi.Input<pulumi.Input<string>[]>;
}
