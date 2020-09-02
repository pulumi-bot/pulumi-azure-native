// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Private Endpoint Connection ARM resource.
 *
 * ## Example Usage
 * ### Approves or rejects a private endpoint connection for a site.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const webAppPrivateEndpointConnection = new azurerm.web.v20200601.WebAppPrivateEndpointConnection("webAppPrivateEndpointConnection", {
 *     name: "testSite",
 *     privateEndpointConnectionName: "connection",
 *     privateLinkServiceConnectionState: {
 *         actionsRequired: "",
 *         description: "Approved by admin.",
 *         status: "Approved",
 *     },
 *     resourceGroupName: "rg",
 * });
 *
 * ```
 */
export class WebAppPrivateEndpointConnection extends pulumi.CustomResource {
    /**
     * Get an existing WebAppPrivateEndpointConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppPrivateEndpointConnection {
        return new WebAppPrivateEndpointConnection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20200601:WebAppPrivateEndpointConnection';

    /**
     * Returns true if the given object is an instance of WebAppPrivateEndpointConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppPrivateEndpointConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppPrivateEndpointConnection.__pulumiType;
    }

    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * PrivateEndpoint of a remote private endpoint connection
     */
    public /*out*/ readonly privateEndpoint!: pulumi.Output<outputs.web.v20200601.ArmIdWrapperResponse | undefined>;
    /**
     * The state of a private link connection
     */
    public readonly privateLinkServiceConnectionState!: pulumi.Output<outputs.web.v20200601.PrivateLinkConnectionStateResponse | undefined>;
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebAppPrivateEndpointConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppPrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppPrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppPrivateEndpointConnectionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.privateEndpointConnectionName === undefined) {
                throw new Error("Missing required property 'privateEndpointConnectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateEndpointConnectionName"] = args ? args.privateEndpointConnectionName : undefined;
            inputs["privateLinkServiceConnectionState"] = args ? args.privateLinkServiceConnectionState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["privateEndpoint"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/latest:WebAppPrivateEndpointConnection" }, { type: "azurerm:web/v20190801:WebAppPrivateEndpointConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebAppPrivateEndpointConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppPrivateEndpointConnection resource.
 */
export interface WebAppPrivateEndpointConnectionArgs {
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the site.
     */
    readonly name: pulumi.Input<string>;
    readonly privateEndpointConnectionName: pulumi.Input<string>;
    /**
     * The state of a private link connection
     */
    readonly privateLinkServiceConnectionState?: pulumi.Input<inputs.web.v20200601.PrivateLinkConnectionState>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
