// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
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
        return new ApplicationGateway(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20161201:ApplicationGateway';

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
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the application gateway.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20161201.ApplicationGatewayPropertiesFormatResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApplicationGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApplicationGatewayArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["authenticationCertificates"] = args ? args.authenticationCertificates : undefined;
            inputs["backendAddressPools"] = args ? args.backendAddressPools : undefined;
            inputs["backendHttpSettingsCollection"] = args ? args.backendHttpSettingsCollection : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["frontendIPConfigurations"] = args ? args.frontendIPConfigurations : undefined;
            inputs["frontendPorts"] = args ? args.frontendPorts : undefined;
            inputs["gatewayIPConfigurations"] = args ? args.gatewayIPConfigurations : undefined;
            inputs["httpListeners"] = args ? args.httpListeners : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["probes"] = args ? args.probes : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["requestRoutingRules"] = args ? args.requestRoutingRules : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["urlPathMaps"] = args ? args.urlPathMaps : undefined;
            inputs["webApplicationFirewallConfiguration"] = args ? args.webApplicationFirewallConfiguration : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApplicationGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationGateway resource.
 */
export interface ApplicationGatewayArgs {
    /**
     * Authentication certificates of the application gateway resource.
     */
    readonly authenticationCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayAuthenticationCertificate>[]>;
    /**
     * Backend address pool of the application gateway resource.
     */
    readonly backendAddressPools?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayBackendAddressPool>[]>;
    /**
     * Backend http settings of the application gateway resource.
     */
    readonly backendHttpSettingsCollection?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayBackendHttpSettings>[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Frontend IP addresses of the application gateway resource.
     */
    readonly frontendIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayFrontendIPConfiguration>[]>;
    /**
     * Frontend ports of the application gateway resource.
     */
    readonly frontendPorts?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayFrontendPort>[]>;
    /**
     * Subnets of application the gateway resource.
     */
    readonly gatewayIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayIPConfiguration>[]>;
    /**
     * Http listeners of the application gateway resource.
     */
    readonly httpListeners?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayHttpListener>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the application gateway.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Probes of the application gateway resource.
     */
    readonly probes?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayProbe>[]>;
    /**
     * Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Request routing rules of the application gateway resource.
     */
    readonly requestRoutingRules?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayRequestRoutingRule>[]>;
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
    readonly sku?: pulumi.Input<inputs.network.v20161201.ApplicationGatewaySku>;
    /**
     * SSL certificates of the application gateway resource.
     */
    readonly sslCertificates?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewaySslCertificate>[]>;
    /**
     * SSL policy of the application gateway resource.
     */
    readonly sslPolicy?: pulumi.Input<inputs.network.v20161201.ApplicationGatewaySslPolicy>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * URL path map of the application gateway resource.
     */
    readonly urlPathMaps?: pulumi.Input<pulumi.Input<inputs.network.v20161201.ApplicationGatewayUrlPathMap>[]>;
    /**
     * Web application firewall configuration.
     */
    readonly webApplicationFirewallConfiguration?: pulumi.Input<inputs.network.v20161201.ApplicationGatewayWebApplicationFirewallConfiguration>;
}
