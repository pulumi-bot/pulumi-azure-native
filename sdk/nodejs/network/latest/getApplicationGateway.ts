// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApplicationGateway(args: GetApplicationGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/latest:getApplicationGateway", {
        "applicationGatewayName": args.applicationGatewayName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationGatewayArgs {
    /**
     * The name of the application gateway.
     */
    readonly applicationGatewayName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Application gateway resource.
 */
export interface GetApplicationGatewayResult {
    /**
     * Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly authenticationCertificates?: outputs.network.latest.ApplicationGatewayAuthenticationCertificateResponse[];
    /**
     * Autoscale Configuration.
     */
    readonly autoscaleConfiguration?: outputs.network.latest.ApplicationGatewayAutoscaleConfigurationResponse;
    /**
     * Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly backendAddressPools?: outputs.network.latest.ApplicationGatewayBackendAddressPoolResponse[];
    /**
     * Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly backendHttpSettingsCollection?: outputs.network.latest.ApplicationGatewayBackendHttpSettingsResponse[];
    /**
     * Custom error configurations of the application gateway resource.
     */
    readonly customErrorConfigurations?: outputs.network.latest.ApplicationGatewayCustomErrorResponse[];
    /**
     * Whether FIPS is enabled on the application gateway resource.
     */
    readonly enableFips?: boolean;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    readonly enableHttp2?: boolean;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Reference to the FirewallPolicy resource.
     */
    readonly firewallPolicy?: outputs.network.latest.SubResourceResponse;
    /**
     * If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
     */
    readonly forceFirewallPolicyAssociation?: boolean;
    /**
     * Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly frontendIPConfigurations?: outputs.network.latest.ApplicationGatewayFrontendIPConfigurationResponse[];
    /**
     * Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly frontendPorts?: outputs.network.latest.ApplicationGatewayFrontendPortResponse[];
    /**
     * Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly gatewayIPConfigurations?: outputs.network.latest.ApplicationGatewayIPConfigurationResponse[];
    /**
     * Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly httpListeners?: outputs.network.latest.ApplicationGatewayHttpListenerResponse[];
    /**
     * The identity of the application gateway, if configured.
     */
    readonly identity?: outputs.network.latest.ManagedServiceIdentityResponse;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Operational state of the application gateway resource.
     */
    readonly operationalState: string;
    /**
     * Private Endpoint connections on application gateway.
     */
    readonly privateEndpointConnections: outputs.network.latest.ApplicationGatewayPrivateEndpointConnectionResponse[];
    /**
     * PrivateLink configurations on application gateway.
     */
    readonly privateLinkConfigurations?: outputs.network.latest.ApplicationGatewayPrivateLinkConfigurationResponse[];
    /**
     * Probes of the application gateway resource.
     */
    readonly probes?: outputs.network.latest.ApplicationGatewayProbeResponse[];
    /**
     * The provisioning state of the application gateway resource.
     */
    readonly provisioningState: string;
    /**
     * Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly redirectConfigurations?: outputs.network.latest.ApplicationGatewayRedirectConfigurationResponse[];
    /**
     * Request routing rules of the application gateway resource.
     */
    readonly requestRoutingRules?: outputs.network.latest.ApplicationGatewayRequestRoutingRuleResponse[];
    /**
     * The resource GUID property of the application gateway resource.
     */
    readonly resourceGuid: string;
    /**
     * Rewrite rules for the application gateway resource.
     */
    readonly rewriteRuleSets?: outputs.network.latest.ApplicationGatewayRewriteRuleSetResponse[];
    /**
     * SKU of the application gateway resource.
     */
    readonly sku?: outputs.network.latest.ApplicationGatewaySkuResponse;
    /**
     * SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly sslCertificates?: outputs.network.latest.ApplicationGatewaySslCertificateResponse[];
    /**
     * SSL policy of the application gateway resource.
     */
    readonly sslPolicy?: outputs.network.latest.ApplicationGatewaySslPolicyResponse;
    /**
     * SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly sslProfiles?: outputs.network.latest.ApplicationGatewaySslProfileResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly trustedClientCertificates?: outputs.network.latest.ApplicationGatewayTrustedClientCertificateResponse[];
    /**
     * Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly trustedRootCertificates?: outputs.network.latest.ApplicationGatewayTrustedRootCertificateResponse[];
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    readonly urlPathMaps?: outputs.network.latest.ApplicationGatewayUrlPathMapResponse[];
    /**
     * Web application firewall configuration.
     */
    readonly webApplicationFirewallConfiguration?: outputs.network.latest.ApplicationGatewayWebApplicationFirewallConfigurationResponse;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    readonly zones?: string[];
}