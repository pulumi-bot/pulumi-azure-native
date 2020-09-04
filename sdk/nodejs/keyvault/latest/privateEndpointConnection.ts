// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Private endpoint connection resource.
 *
 * ## Example Usage
 * ### KeyVaultPutPrivateEndpointConnection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const privateEndpointConnection = new azurerm.keyvault.latest.PrivateEndpointConnection("privateEndpointConnection", {
 *     privateEndpointConnectionName: "sample-pec",
 *     privateLinkServiceConnectionState: {
 *         description: "My name is Joe and I'm approving this.",
 *         status: "Approved",
 *     },
 *     resourceGroupName: "sample-group",
 *     vaultName: "sample-vault",
 * });
 *
 * ```
 */
export class PrivateEndpointConnection extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateEndpointConnection {
        return new PrivateEndpointConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:keyvault/latest:PrivateEndpointConnection';

    /**
     * Returns true if the given object is an instance of PrivateEndpointConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointConnection.__pulumiType;
    }

    /**
     * Azure location of the key vault resource.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Name of the key vault resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of the private endpoint object.
     */
    public /*out*/ readonly privateEndpoint!: pulumi.Output<outputs.keyvault.latest.PrivateEndpointResponse | undefined>;
    /**
     * Approval state of the private link connection.
     */
    public readonly privateLinkServiceConnectionState!: pulumi.Output<outputs.keyvault.latest.PrivateLinkServiceConnectionStateResponse | undefined>;
    /**
     * Provisioning state of the private endpoint connection.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Tags assigned to the key vault resource.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type of the key vault resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.privateEndpointConnectionName === undefined) {
                throw new Error("Missing required property 'privateEndpointConnectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vaultName === undefined) {
                throw new Error("Missing required property 'vaultName'");
            }
            inputs["privateEndpointConnectionName"] = args ? args.privateEndpointConnectionName : undefined;
            inputs["privateLinkServiceConnectionState"] = args ? args.privateLinkServiceConnectionState : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["vaultName"] = args ? args.vaultName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpoint"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpoint"] = undefined /*out*/;
            inputs["privateLinkServiceConnectionState"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:keyvault/v20180214:PrivateEndpointConnection" }, { type: "azurerm:keyvault/v20190901:PrivateEndpointConnection" }, { type: "azurerm:keyvault/v20200401preview:PrivateEndpointConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PrivateEndpointConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateEndpointConnection resource.
 */
export interface PrivateEndpointConnectionArgs {
    /**
     * Name of the private endpoint connection associated with the key vault.
     */
    readonly privateEndpointConnectionName: pulumi.Input<string>;
    /**
     * Approval state of the private link connection.
     */
    readonly privateLinkServiceConnectionState?: pulumi.Input<inputs.keyvault.latest.PrivateLinkServiceConnectionState>;
    /**
     * Provisioning state of the private endpoint connection.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Name of the resource group that contains the key vault.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the key vault.
     */
    readonly vaultName: pulumi.Input<string>;
}
