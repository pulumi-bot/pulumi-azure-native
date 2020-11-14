// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Application gateway resource
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
    public static readonly __pulumiType = 'azure-nextgen:network/v20171001:ApplicationGateway';

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
     * Authentication certificates of the application gateway resource.
     */
    public readonly authenticationCertificates!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayAuthenticationCertificateResponse[] | undefined>;
    /**
     * Backend address pool of the application gateway resource.
     */
    public readonly backendAddressPools!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayBackendAddressPoolResponse[] | undefined>;
    /**
     * Backend http settings of the application gateway resource.
     */
    public readonly backendHttpSettingsCollection!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayBackendHttpSettingsResponse[] | undefined>;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    public readonly enableHttp2!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Frontend IP addresses of the application gateway resource.
     */
    public readonly frontendIPConfigurations!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayFrontendIPConfigurationResponse[] | undefined>;
    /**
     * Frontend ports of the application gateway resource.
     */
    public readonly frontendPorts!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayFrontendPortResponse[] | undefined>;
    /**
     * Subnets of application the gateway resource.
     */
    public readonly gatewayIPConfigurations!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayIPConfigurationResponse[] | undefined>;
    /**
     * Http listeners of the application gateway resource.
     */
    public readonly httpListeners!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayHttpListenerResponse[] | undefined>;
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
     * Probes of the application gateway resource.
     */
    public readonly probes!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayProbeResponse[] | undefined>;
    /**
     * Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Redirect configurations of the application gateway resource.
     */
    public readonly redirectConfigurations!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayRedirectConfigurationResponse[] | undefined>;
    /**
     * Request routing rules of the application gateway resource.
     */
    public readonly requestRoutingRules!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayRequestRoutingRuleResponse[] | undefined>;
    /**
     * Resource GUID property of the application gateway resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * SKU of the application gateway resource.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20171001.ApplicationGatewaySkuResponse | undefined>;
    /**
     * SSL certificates of the application gateway resource.
     */
    public readonly sslCertificates!: pulumi.Output<outputs.network.v20171001.ApplicationGatewaySslCertificateResponse[] | undefined>;
    /**
     * SSL policy of the application gateway resource.
     */
    public readonly sslPolicy!: pulumi.Output<outputs.network.v20171001.ApplicationGatewaySslPolicyResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * URL path map of the application gateway resource.
     */
    public readonly urlPathMaps!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayUrlPathMapResponse[] | undefined>;
    /**
     * Web application firewall configuration.
     */
    public readonly webApplicationFirewallConfiguration!: pulumi.Output<outputs.network.v20171001.ApplicationGatewayWebApplicationFirewallConfigurationResponse | undefined>;

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
            inputs["backendAddressPools"] = args ? args.backendAddressPools : undefined;
            inputs["backendHttpSettingsCollection"] = args ? args.backendHttpSettingsCollection : undefined;
            inputs["enableHttp2"] = args ? args.enableHttp2 : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["frontendIPConfigurations"] = args ? args.frontendIPConfigurations : undefined;
            inputs["frontendPorts"] = args ? args.frontendPorts : undefined;
            inputs["gatewayIPConfigurations"] = args ? args.gatewayIPConfigurations : undefined;
            inputs["httpListeners"] = args ? args.httpListeners : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["probes"] = args ? args.probes : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["redirectConfigurations"] = args ? args.redirectConfigurations : undefined;
            inputs["requestRoutingRules"] = args ? args.requestRoutingRules : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["urlPathMaps"] = args ? args.urlPathMaps : undefined;
            inputs["webApplicationFirewallConfiguration"] = args ? args.webApplicationFirewallConfiguration : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["operationalState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["authenticationCertificates"] = undefined /*out*/;
            inputs["backendAddressPools"] = undefined /*out*/;
            inputs["backendHttpSettingsCollection"] = undefined /*out*/;
            inputs["enableHttp2"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["frontendIPConfigurations"] = undefined /*out*/;
            inputs["frontendPorts"] = undefined /*out*/;
            inputs["gatewayIPConfigurations"] = undefined /*out*/;
            inputs["httpListeners"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationalState"] = undefined /*out*/;
            inputs["probes"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["redirectConfigurations"] = undefined /*out*/;
            inputs["requestRoutingRules"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["sslCertificates"] = undefined /*out*/;
            inputs["sslPolicy"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["urlPathMaps"] = undefined /*out*/;
            inputs["webApplicationFirewallConfiguration"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:ApplicationGateway" }, { type: "azure-nextgen:network/v20150501preview:ApplicationGateway" }, { type: "azure-nextgen:network/v20150615:ApplicationGateway" }, { type: "azure-nextgen:network/v20160330:ApplicationGateway" }, { type: "azure-nextgen:network/v20160601:ApplicationGateway" }, { type: "azure-nextgen:network/v20160901:ApplicationGateway" }, { type: "azure-nextgen:network/v20161201:ApplicationGateway" }, { type: "azure-nextgen:network/v20170301:ApplicationGateway" }, { type: "azure-nextgen:network/v20170601:ApplicationGateway" }, { type: "azure-nextgen:network/v20170801:ApplicationGateway" }, { type: "azure-nextgen:network/v20170901:ApplicationGateway" }, { type: "azure-nextgen:network/v20171101:ApplicationGateway" }, { type: "azure-nextgen:network/v20180101:ApplicationGateway" }, { type: "azure-nextgen:network/v20180201:ApplicationGateway" }, { type: "azure-nextgen:network/v20180401:ApplicationGateway" }, { type: "azure-nextgen:network/v20180601:ApplicationGateway" }, { type: "azure-nextgen:network/v20180701:ApplicationGateway" }, { type: "azure-nextgen:network/v20180801:ApplicationGateway" }, { type: "azure-nextgen:network/v20181001:ApplicationGateway" }, { type: "azure-nextgen:network/v20181101:ApplicationGateway" }, { type: "azure-nextgen:network/v20181201:ApplicationGateway" }, { type: "azure-nextgen:network/v20190201:ApplicationGateway" }, { type: "azure-nextgen:network/v20190401:ApplicationGateway" }, { type: "azure-nextgen:network/v20190601:ApplicationGateway" }, { type: "azure-nextgen:network/v20190701:ApplicationGateway" }, { type: "azure-nextgen:network/v20190801:ApplicationGateway" }, { type: "azure-nextgen:network/v20190901:ApplicationGateway" }, { type: "azure-nextgen:network/v20191101:ApplicationGateway" }, { type: "azure-nextgen:network/v20191201:ApplicationGateway" }, { type: "azure-nextgen:network/v20200301:ApplicationGateway" }, { type: "azure-nextgen:network/v20200401:ApplicationGateway" }, { type: "azure-nextgen:network/v20200501:ApplicationGateway" }, { type: "azure-nextgen:network/v20200601:ApplicationGateway" }, { type: "azure-nextgen:network/v20200701:ApplicationGateway" }] };
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
     * Authentication certificates of the application gateway resource.
     */
    readonly authenticationCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayAuthenticationCertificate>[]>;
    /**
     * Backend address pool of the application gateway resource.
     */
    readonly backendAddressPools?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayBackendAddressPool>[]>;
    /**
     * Backend http settings of the application gateway resource.
     */
    readonly backendHttpSettingsCollection?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayBackendHttpSettings>[]>;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    readonly enableHttp2?: pulumi.Input<boolean>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Frontend IP addresses of the application gateway resource.
     */
    readonly frontendIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayFrontendIPConfiguration>[]>;
    /**
     * Frontend ports of the application gateway resource.
     */
    readonly frontendPorts?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayFrontendPort>[]>;
    /**
     * Subnets of application the gateway resource.
     */
    readonly gatewayIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayIPConfiguration>[]>;
    /**
     * Http listeners of the application gateway resource.
     */
    readonly httpListeners?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayHttpListener>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Probes of the application gateway resource.
     */
    readonly probes?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayProbe>[]>;
    /**
     * Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Redirect configurations of the application gateway resource.
     */
    readonly redirectConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayRedirectConfiguration>[]>;
    /**
     * Request routing rules of the application gateway resource.
     */
    readonly requestRoutingRules?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayRequestRoutingRule>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource GUID property of the application gateway resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * SKU of the application gateway resource.
     */
    readonly sku?: pulumi.Input<inputs.network.v20171001.ApplicationGatewaySku>;
    /**
     * SSL certificates of the application gateway resource.
     */
    readonly sslCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewaySslCertificate>[]>;
    /**
     * SSL policy of the application gateway resource.
     */
    readonly sslPolicy?: pulumi.Input<inputs.network.v20171001.ApplicationGatewaySslPolicy>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * URL path map of the application gateway resource.
     */
    readonly urlPathMaps?: pulumi.Input<pulumi.Input<inputs.network.v20171001.ApplicationGatewayUrlPathMap>[]>;
    /**
     * Web application firewall configuration.
     */
    readonly webApplicationFirewallConfiguration?: pulumi.Input<inputs.network.v20171001.ApplicationGatewayWebApplicationFirewallConfiguration>;
}
