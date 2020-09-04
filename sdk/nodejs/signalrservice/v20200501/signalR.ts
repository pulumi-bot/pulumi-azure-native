// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A class represent a SignalR service resource.
 *
 * ## Example Usage
 * ### SignalR_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const signalR = new azurerm.signalrservice.v20200501.SignalR("signalR", {
 *     cors: {
 *         allowedOrigins: [
 *             "https://foo.com",
 *             "https://bar.com",
 *         ],
 *     },
 *     features: [
 *         {
 *             flag: "ServiceMode",
 *             properties: {},
 *             value: "Serverless",
 *         },
 *         {
 *             flag: "EnableConnectivityLogs",
 *             properties: {},
 *             value: "True",
 *         },
 *         {
 *             flag: "EnableMessagingLogs",
 *             properties: {},
 *             value: "False",
 *         },
 *     ],
 *     kind: "SignalR",
 *     location: "eastus",
 *     networkACLs: {
 *         defaultAction: "Deny",
 *         privateEndpoints: [{
 *             allow: ["ServerConnection"],
 *             name: "mySignalRService.1fa229cd-bf3f-47f0-8c49-afb36723997e",
 *         }],
 *         publicNetwork: {
 *             allow: ["ClientConnection"],
 *         },
 *     },
 *     resourceGroupName: "myResourceGroup",
 *     resourceName: "mySignalRService",
 *     sku: {
 *         capacity: 1,
 *         name: "Standard_S1",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         key1: "value1",
 *     },
 *     upstream: {
 *         templates: [{
 *             categoryPattern: "*",
 *             eventPattern: "connect,disconnect",
 *             hubPattern: "*",
 *             urlTemplate: "https://example.com/chat/api/connect",
 *         }],
 *     },
 * });
 *
 * ```
 */
export class SignalR extends pulumi.CustomResource {
    /**
     * Get an existing SignalR resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SignalR {
        return new SignalR(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:signalrservice/v20200501:SignalR';

    /**
     * Returns true if the given object is an instance of SignalR.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SignalR {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SignalR.__pulumiType;
    }

    /**
     * Cross-Origin Resource Sharing (CORS) settings.
     */
    public readonly cors!: pulumi.Output<outputs.signalrservice.v20200501.SignalRCorsSettingsResponse | undefined>;
    /**
     * The publicly accessible IP of the SignalR service.
     */
    public /*out*/ readonly externalIP!: pulumi.Output<string>;
    /**
     * List of SignalR featureFlags. e.g. ServiceMode.
     * 
     * FeatureFlags that are not included in the parameters for the update operation will not be modified.
     * And the response will only include featureFlags that are explicitly set. 
     * When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
     * But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
     */
    public readonly features!: pulumi.Output<outputs.signalrservice.v20200501.SignalRFeatureResponse[] | undefined>;
    /**
     * FQDN of the SignalR service instance. Format: xxx.service.signalr.net
     */
    public /*out*/ readonly hostName!: pulumi.Output<string>;
    /**
     * Prefix for the hostName of the SignalR service. Retained for future use.
     * The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
     */
    public readonly hostNamePrefix!: pulumi.Output<string | undefined>;
    /**
     * The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network ACLs
     */
    public readonly networkACLs!: pulumi.Output<outputs.signalrservice.v20200501.SignalRNetworkACLsResponse | undefined>;
    /**
     * Private endpoint connections to the SignalR resource.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.signalrservice.v20200501.PrivateEndpointConnectionResponse[]>;
    /**
     * Provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for browser/client side usage.
     */
    public /*out*/ readonly publicPort!: pulumi.Output<number>;
    /**
     * The publicly accessible port of the SignalR service which is designed for customer server side usage.
     */
    public /*out*/ readonly serverPort!: pulumi.Output<number>;
    /**
     * The billing information of the resource.(e.g. Free, Standard)
     */
    public readonly sku!: pulumi.Output<outputs.signalrservice.v20200501.ResourceSkuResponse | undefined>;
    /**
     * Tags of the service which is a list of key value pairs that describe the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Upstream settings when the Azure SignalR is in server-less mode.
     */
    public readonly upstream!: pulumi.Output<outputs.signalrservice.v20200501.ServerlessUpstreamSettingsResponse | undefined>;
    /**
     * Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a SignalR resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SignalRArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["cors"] = args ? args.cors : undefined;
            inputs["features"] = args ? args.features : undefined;
            inputs["hostNamePrefix"] = args ? args.hostNamePrefix : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkACLs"] = args ? args.networkACLs : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upstream"] = args ? args.upstream : undefined;
            inputs["externalIP"] = undefined /*out*/;
            inputs["hostName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicPort"] = undefined /*out*/;
            inputs["serverPort"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        } else {
            inputs["cors"] = undefined /*out*/;
            inputs["externalIP"] = undefined /*out*/;
            inputs["features"] = undefined /*out*/;
            inputs["hostName"] = undefined /*out*/;
            inputs["hostNamePrefix"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkACLs"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicPort"] = undefined /*out*/;
            inputs["serverPort"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upstream"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:signalrservice/latest:SignalR" }, { type: "azurerm:signalrservice/v20180301preview:SignalR" }, { type: "azurerm:signalrservice/v20181001:SignalR" }, { type: "azurerm:signalrservice/v20200701preview:SignalR" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SignalR.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SignalR resource.
 */
export interface SignalRArgs {
    /**
     * Cross-Origin Resource Sharing (CORS) settings.
     */
    readonly cors?: pulumi.Input<inputs.signalrservice.v20200501.SignalRCorsSettings>;
    /**
     * List of SignalR featureFlags. e.g. ServiceMode.
     * 
     * FeatureFlags that are not included in the parameters for the update operation will not be modified.
     * And the response will only include featureFlags that are explicitly set. 
     * When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
     * But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
     */
    readonly features?: pulumi.Input<pulumi.Input<inputs.signalrservice.v20200501.SignalRFeature>[]>;
    /**
     * Prefix for the hostName of the SignalR service. Retained for future use.
     * The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
     */
    readonly hostNamePrefix?: pulumi.Input<string>;
    /**
     * The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Network ACLs
     */
    readonly networkACLs?: pulumi.Input<inputs.signalrservice.v20200501.SignalRNetworkACLs>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the SignalR resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * The billing information of the resource.(e.g. Free, Standard)
     */
    readonly sku?: pulumi.Input<inputs.signalrservice.v20200501.ResourceSku>;
    /**
     * Tags of the service which is a list of key value pairs that describe the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Upstream settings when the Azure SignalR is in server-less mode.
     */
    readonly upstream?: pulumi.Input<inputs.signalrservice.v20200501.ServerlessUpstreamSettings>;
}
