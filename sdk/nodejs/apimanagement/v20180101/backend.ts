// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Backend details.
 *
 * ## Example Usage
 * ### ApiManagementCreateBackendProxyBackend
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const backend = new azurerm.apimanagement.v20180101.Backend("backend", {
 *     backendid: "proxybackend",
 *     credentials: {
 *         authorization: {
 *             parameter: "opensesma",
 *             scheme: "Basic",
 *         },
 *         header: {
 *             "x-my-1": [
 *                 "val1",
 *                 "val2",
 *             ],
 *         },
 *         query: {
 *             sv: [
 *                 "xx",
 *                 "bb",
 *                 "cc",
 *             ],
 *         },
 *     },
 *     description: "description5308",
 *     protocol: "http",
 *     proxy: {
 *         password: "opensesame",
 *         url: "http://192.168.1.1:8080",
 *         username: "Contoso\\admin",
 *     },
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 *     tls: {
 *         validateCertificateChain: true,
 *         validateCertificateName: true,
 *     },
 *     url: "https://backendname2644/",
 * });
 *
 * ```
 * ### ApiManagementCreateBackendServiceFabric
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const backend = new azurerm.apimanagement.v20180101.Backend("backend", {
 *     backendid: "sfbackend",
 *     description: "Service Fabric Test App 1",
 *     properties: {
 *         serviceFabricCluster: {
 *             clientCertificatethumbprint: "EBA029198AA3E76EF0D70482626E5BCF148594A6",
 *             managementEndpoints: ["https://somecluster.com"],
 *             maxPartitionResolutionRetries: 5,
 *             serverX509Names: [{
 *                 issuerCertificateThumbprint: "IssuerCertificateThumbprint1",
 *                 name: "ServerCommonName1",
 *             }],
 *         },
 *     },
 *     protocol: "http",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 *     url: "fabric:/mytestapp/mytestservice",
 * });
 *
 * ```
 */
export class Backend extends pulumi.CustomResource {
    /**
     * Get an existing Backend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Backend {
        return new Backend(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20180101:Backend';

    /**
     * Returns true if the given object is an instance of Backend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backend.__pulumiType;
    }

    /**
     * Backend Credentials Contract Properties
     */
    public readonly credentials!: pulumi.Output<outputs.apimanagement.v20180101.BackendCredentialsContractResponse | undefined>;
    /**
     * Backend Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Backend Properties contract
     */
    public readonly properties!: pulumi.Output<outputs.apimanagement.v20180101.BackendPropertiesResponse>;
    /**
     * Backend communication protocol.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Backend Proxy Contract Properties
     */
    public readonly proxy!: pulumi.Output<outputs.apimanagement.v20180101.BackendProxyContractResponse | undefined>;
    /**
     * Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * Backend Title.
     */
    public readonly title!: pulumi.Output<string | undefined>;
    /**
     * Backend TLS Properties
     */
    public readonly tls!: pulumi.Output<outputs.apimanagement.v20180101.BackendTlsPropertiesResponse | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Runtime Url of the Backend.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Backend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.backendid === undefined) {
                throw new Error("Missing required property 'backendid'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["backendid"] = args ? args.backendid : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["proxy"] = args ? args.proxy : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["tls"] = args ? args.tls : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["credentials"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["protocol"] = undefined /*out*/;
            inputs["proxy"] = undefined /*out*/;
            inputs["resourceId"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["tls"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:Backend" }, { type: "azurerm:apimanagement/v20160707:Backend" }, { type: "azurerm:apimanagement/v20161010:Backend" }, { type: "azurerm:apimanagement/v20170301:Backend" }, { type: "azurerm:apimanagement/v20180601preview:Backend" }, { type: "azurerm:apimanagement/v20190101:Backend" }, { type: "azurerm:apimanagement/v20191201:Backend" }, { type: "azurerm:apimanagement/v20191201preview:Backend" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Backend.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Backend resource.
 */
export interface BackendArgs {
    /**
     * Identifier of the Backend entity. Must be unique in the current API Management service instance.
     */
    readonly backendid: pulumi.Input<string>;
    /**
     * Backend Credentials Contract Properties
     */
    readonly credentials?: pulumi.Input<inputs.apimanagement.v20180101.BackendCredentialsContract>;
    /**
     * Backend Description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Backend Properties contract
     */
    readonly properties?: pulumi.Input<inputs.apimanagement.v20180101.BackendProperties>;
    /**
     * Backend communication protocol.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * Backend Proxy Contract Properties
     */
    readonly proxy?: pulumi.Input<inputs.apimanagement.v20180101.BackendProxyContract>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Backend Title.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * Backend TLS Properties
     */
    readonly tls?: pulumi.Input<inputs.apimanagement.v20180101.BackendTlsProperties>;
    /**
     * Runtime Url of the Backend.
     */
    readonly url: pulumi.Input<string>;
}
