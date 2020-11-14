// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getApplicationGateway(args: GetApplicationGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20150615:getApplicationGateway", {
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
 * Application gateway resource
 */
export interface GetApplicationGatewayResult {
    /**
     * Backend address pool of the application gateway resource.
     */
    readonly backendAddressPools?: outputs.network.v20150615.ApplicationGatewayBackendAddressPoolResponse[];
    /**
     * Backend http settings of the application gateway resource.
     */
    readonly backendHttpSettingsCollection?: outputs.network.v20150615.ApplicationGatewayBackendHttpSettingsResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * Frontend IP addresses of the application gateway resource.
     */
    readonly frontendIPConfigurations?: outputs.network.v20150615.ApplicationGatewayFrontendIPConfigurationResponse[];
    /**
     * Frontend ports of the application gateway resource.
     */
    readonly frontendPorts?: outputs.network.v20150615.ApplicationGatewayFrontendPortResponse[];
    /**
     * Gets or sets subnets of application gateway resource
     */
    readonly gatewayIPConfigurations?: outputs.network.v20150615.ApplicationGatewayIPConfigurationResponse[];
    /**
     * Http listeners of the application gateway resource.
     */
    readonly httpListeners?: outputs.network.v20150615.ApplicationGatewayHttpListenerResponse[];
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Operational state of the application gateway resource. Possible values are: 'Stopped', 'Started', 'Running', and 'Stopping'.
     */
    readonly operationalState: string;
    /**
     * Probes of the application gateway resource.
     */
    readonly probes?: outputs.network.v20150615.ApplicationGatewayProbeResponse[];
    /**
     * Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: string;
    /**
     * Request routing rules of the application gateway resource.
     */
    readonly requestRoutingRules?: outputs.network.v20150615.ApplicationGatewayRequestRoutingRuleResponse[];
    /**
     * Resource GUID property of the application gateway resource.
     */
    readonly resourceGuid?: string;
    /**
     * SKU of the application gateway resource.
     */
    readonly sku?: outputs.network.v20150615.ApplicationGatewaySkuResponse;
    /**
     * SSL certificates of the application gateway resource.
     */
    readonly sslCertificates?: outputs.network.v20150615.ApplicationGatewaySslCertificateResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * URL path map of the application gateway resource.
     */
    readonly urlPathMaps?: outputs.network.v20150615.ApplicationGatewayUrlPathMapResponse[];
}
