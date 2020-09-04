// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Application gateway resource.
 */
export class ApplicationGateway extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationGateway {
        return new ApplicationGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:ApplicationGateway';

    /**
     * Returns true if the given object is an instance of ApplicationGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationGateway.__pulumiType;
    }

    /**
     * Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly authenticationCertificates!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayAuthenticationCertificateResponse[] | undefined>;
    /**
     * Autoscale Configuration.
     */
    public readonly autoscaleConfiguration!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayAutoscaleConfigurationResponse | undefined>;
    /**
     * Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly backendAddressPools!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayBackendAddressPoolResponse[] | undefined>;
    /**
     * Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly backendHttpSettingsCollection!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayBackendHttpSettingsResponse[] | undefined>;
    /**
     * Custom error configurations of the application gateway resource.
     */
    public readonly customErrorConfigurations!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayCustomErrorResponse[] | undefined>;
    /**
     * Whether FIPS is enabled on the application gateway resource.
     */
    public readonly enableFips!: pulumi.Output<boolean | undefined>;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    public readonly enableHttp2!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Reference to the FirewallPolicy resource.
     */
    public readonly firewallPolicy!: pulumi.Output<outputs.network.v20200601.SubResourceResponse | undefined>;
    /**
     * If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
     */
    public readonly forceFirewallPolicyAssociation!: pulumi.Output<boolean | undefined>;
    /**
     * Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly frontendIPConfigurations!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayFrontendIPConfigurationResponse[] | undefined>;
    /**
     * Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly frontendPorts!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayFrontendPortResponse[] | undefined>;
    /**
     * Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly gatewayIPConfigurations!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayIPConfigurationResponse[] | undefined>;
    /**
     * Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly httpListeners!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayHttpListenerResponse[] | undefined>;
    /**
     * The identity of the application gateway, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.network.v20200601.ManagedServiceIdentityResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Operational state of the application gateway resource.
     */
    public /*out*/ readonly operationalState!: pulumi.Output<string>;
    /**
     * Private Endpoint connections on application gateway.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayPrivateEndpointConnectionResponse[]>;
    /**
     * PrivateLink configurations on application gateway.
     */
    public readonly privateLinkConfigurations!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayPrivateLinkConfigurationResponse[] | undefined>;
    /**
     * Probes of the application gateway resource.
     */
    public readonly probes!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayProbeResponse[] | undefined>;
    /**
     * The provisioning state of the application gateway resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly redirectConfigurations!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayRedirectConfigurationResponse[] | undefined>;
    /**
     * Request routing rules of the application gateway resource.
     */
    public readonly requestRoutingRules!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayRequestRoutingRuleResponse[] | undefined>;
    /**
     * The resource GUID property of the application gateway resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Rewrite rules for the application gateway resource.
     */
    public readonly rewriteRuleSets!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayRewriteRuleSetResponse[] | undefined>;
    /**
     * SKU of the application gateway resource.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20200601.ApplicationGatewaySkuResponse | undefined>;
    /**
     * SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly sslCertificates!: pulumi.Output<outputs.network.v20200601.ApplicationGatewaySslCertificateResponse[] | undefined>;
    /**
     * SSL policy of the application gateway resource.
     */
    public readonly sslPolicy!: pulumi.Output<outputs.network.v20200601.ApplicationGatewaySslPolicyResponse | undefined>;
    /**
     * SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly sslProfiles!: pulumi.Output<outputs.network.v20200601.ApplicationGatewaySslProfileResponse[] | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly trustedClientCertificates!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayTrustedClientCertificateResponse[] | undefined>;
    /**
     * Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly trustedRootCertificates!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayTrustedRootCertificateResponse[] | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly urlPathMaps!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayUrlPathMapResponse[] | undefined>;
    /**
     * Web application firewall configuration.
     */
    public readonly webApplicationFirewallConfiguration!: pulumi.Output<outputs.network.v20200601.ApplicationGatewayWebApplicationFirewallConfigurationResponse | undefined>;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ApplicationGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.applicationGatewayName === undefined) {
                throw new Error("Missing required property 'applicationGatewayName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationGatewayName"] = args ? args.applicationGatewayName : undefined;
            inputs["authenticationCertificates"] = args ? args.authenticationCertificates : undefined;
            inputs["autoscaleConfiguration"] = args ? args.autoscaleConfiguration : undefined;
            inputs["backendAddressPools"] = args ? args.backendAddressPools : undefined;
            inputs["backendHttpSettingsCollection"] = args ? args.backendHttpSettingsCollection : undefined;
            inputs["customErrorConfigurations"] = args ? args.customErrorConfigurations : undefined;
            inputs["enableFips"] = args ? args.enableFips : undefined;
            inputs["enableHttp2"] = args ? args.enableHttp2 : undefined;
            inputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            inputs["forceFirewallPolicyAssociation"] = args ? args.forceFirewallPolicyAssociation : undefined;
            inputs["frontendIPConfigurations"] = args ? args.frontendIPConfigurations : undefined;
            inputs["frontendPorts"] = args ? args.frontendPorts : undefined;
            inputs["gatewayIPConfigurations"] = args ? args.gatewayIPConfigurations : undefined;
            inputs["httpListeners"] = args ? args.httpListeners : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["privateLinkConfigurations"] = args ? args.privateLinkConfigurations : undefined;
            inputs["probes"] = args ? args.probes : undefined;
            inputs["redirectConfigurations"] = args ? args.redirectConfigurations : undefined;
            inputs["requestRoutingRules"] = args ? args.requestRoutingRules : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rewriteRuleSets"] = args ? args.rewriteRuleSets : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["sslProfiles"] = args ? args.sslProfiles : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trustedClientCertificates"] = args ? args.trustedClientCertificates : undefined;
            inputs["trustedRootCertificates"] = args ? args.trustedRootCertificates : undefined;
            inputs["urlPathMaps"] = args ? args.urlPathMaps : undefined;
            inputs["webApplicationFirewallConfiguration"] = args ? args.webApplicationFirewallConfiguration : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationalState"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["authenticationCertificates"] = undefined /*out*/;
            inputs["autoscaleConfiguration"] = undefined /*out*/;
            inputs["backendAddressPools"] = undefined /*out*/;
            inputs["backendHttpSettingsCollection"] = undefined /*out*/;
            inputs["customErrorConfigurations"] = undefined /*out*/;
            inputs["enableFips"] = undefined /*out*/;
            inputs["enableHttp2"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["firewallPolicy"] = undefined /*out*/;
            inputs["forceFirewallPolicyAssociation"] = undefined /*out*/;
            inputs["frontendIPConfigurations"] = undefined /*out*/;
            inputs["frontendPorts"] = undefined /*out*/;
            inputs["gatewayIPConfigurations"] = undefined /*out*/;
            inputs["httpListeners"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationalState"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["privateLinkConfigurations"] = undefined /*out*/;
            inputs["probes"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["redirectConfigurations"] = undefined /*out*/;
            inputs["requestRoutingRules"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["rewriteRuleSets"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["sslCertificates"] = undefined /*out*/;
            inputs["sslPolicy"] = undefined /*out*/;
            inputs["sslProfiles"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["trustedClientCertificates"] = undefined /*out*/;
            inputs["trustedRootCertificates"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["urlPathMaps"] = undefined /*out*/;
            inputs["webApplicationFirewallConfiguration"] = undefined /*out*/;
            inputs["zones"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:ApplicationGateway" }, { type: "azurerm:network/v20150501preview:ApplicationGateway" }, { type: "azurerm:network/v20150615:ApplicationGateway" }, { type: "azurerm:network/v20160330:ApplicationGateway" }, { type: "azurerm:network/v20160601:ApplicationGateway" }, { type: "azurerm:network/v20160901:ApplicationGateway" }, { type: "azurerm:network/v20161201:ApplicationGateway" }, { type: "azurerm:network/v20170301:ApplicationGateway" }, { type: "azurerm:network/v20170601:ApplicationGateway" }, { type: "azurerm:network/v20170801:ApplicationGateway" }, { type: "azurerm:network/v20170901:ApplicationGateway" }, { type: "azurerm:network/v20171001:ApplicationGateway" }, { type: "azurerm:network/v20171101:ApplicationGateway" }, { type: "azurerm:network/v20180101:ApplicationGateway" }, { type: "azurerm:network/v20180201:ApplicationGateway" }, { type: "azurerm:network/v20180401:ApplicationGateway" }, { type: "azurerm:network/v20180601:ApplicationGateway" }, { type: "azurerm:network/v20180701:ApplicationGateway" }, { type: "azurerm:network/v20180801:ApplicationGateway" }, { type: "azurerm:network/v20181001:ApplicationGateway" }, { type: "azurerm:network/v20181101:ApplicationGateway" }, { type: "azurerm:network/v20181201:ApplicationGateway" }, { type: "azurerm:network/v20190201:ApplicationGateway" }, { type: "azurerm:network/v20190401:ApplicationGateway" }, { type: "azurerm:network/v20190601:ApplicationGateway" }, { type: "azurerm:network/v20190701:ApplicationGateway" }, { type: "azurerm:network/v20190801:ApplicationGateway" }, { type: "azurerm:network/v20190901:ApplicationGateway" }, { type: "azurerm:network/v20191101:ApplicationGateway" }, { type: "azurerm:network/v20191201:ApplicationGateway" }, { type: "azurerm:network/v20200301:ApplicationGateway" }, { type: "azurerm:network/v20200401:ApplicationGateway" }, { type: "azurerm:network/v20200501:ApplicationGateway" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApplicationGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationGateway resource.
 */
export interface ApplicationGatewayArgs {
    /**
     * The name of the application gateway.
     */
    readonly applicationGatewayName: pulumi.Input<string>;
    /**
     * Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly authenticationCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayAuthenticationCertificate>[]>;
    /**
     * Autoscale Configuration.
     */
    readonly autoscaleConfiguration?: pulumi.Input<inputs.network.v20200601.ApplicationGatewayAutoscaleConfiguration>;
    /**
     * Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly backendAddressPools?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayBackendAddressPool>[]>;
    /**
     * Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly backendHttpSettingsCollection?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayBackendHttpSettings>[]>;
    /**
     * Custom error configurations of the application gateway resource.
     */
    readonly customErrorConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayCustomError>[]>;
    /**
     * Whether FIPS is enabled on the application gateway resource.
     */
    readonly enableFips?: pulumi.Input<boolean>;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    readonly enableHttp2?: pulumi.Input<boolean>;
    /**
     * Reference to the FirewallPolicy resource.
     */
    readonly firewallPolicy?: pulumi.Input<inputs.network.v20200601.SubResource>;
    /**
     * If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
     */
    readonly forceFirewallPolicyAssociation?: pulumi.Input<boolean>;
    /**
     * Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly frontendIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayFrontendIPConfiguration>[]>;
    /**
     * Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly frontendPorts?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayFrontendPort>[]>;
    /**
     * Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly gatewayIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayIPConfiguration>[]>;
    /**
     * Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly httpListeners?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayHttpListener>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The identity of the application gateway, if configured.
     */
    readonly identity?: pulumi.Input<inputs.network.v20200601.ManagedServiceIdentity>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * PrivateLink configurations on application gateway.
     */
    readonly privateLinkConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayPrivateLinkConfiguration>[]>;
    /**
     * Probes of the application gateway resource.
     */
    readonly probes?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayProbe>[]>;
    /**
     * Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly redirectConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayRedirectConfiguration>[]>;
    /**
     * Request routing rules of the application gateway resource.
     */
    readonly requestRoutingRules?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayRequestRoutingRule>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Rewrite rules for the application gateway resource.
     */
    readonly rewriteRuleSets?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayRewriteRuleSet>[]>;
    /**
     * SKU of the application gateway resource.
     */
    readonly sku?: pulumi.Input<inputs.network.v20200601.ApplicationGatewaySku>;
    /**
     * SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly sslCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewaySslCertificate>[]>;
    /**
     * SSL policy of the application gateway resource.
     */
    readonly sslPolicy?: pulumi.Input<inputs.network.v20200601.ApplicationGatewaySslPolicy>;
    /**
     * SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly sslProfiles?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewaySslProfile>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly trustedClientCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayTrustedClientCertificate>[]>;
    /**
     * Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly trustedRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayTrustedRootCertificate>[]>;
    /**
     * URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly urlPathMaps?: pulumi.Input<pulumi.Input<inputs.network.v20200601.ApplicationGatewayUrlPathMap>[]>;
    /**
     * Web application firewall configuration.
     */
    readonly webApplicationFirewallConfiguration?: pulumi.Input<inputs.network.v20200601.ApplicationGatewayWebApplicationFirewallConfiguration>;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
