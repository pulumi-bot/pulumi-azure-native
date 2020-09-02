// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents an instance of an Analysis Services resource.
 *
 * ## Example Usage
 * ### Create a server.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const serverDetails = new azurerm.analysisservices.v20170801.ServerDetails("serverDetails", {
 *     asAdministrators: {
 *         members: [
 *             "azsdktest@microsoft.com",
 *             "azsdktest2@microsoft.com",
 *         ],
 *     },
 *     location: "West US",
 *     resourceGroupName: "TestRG",
 *     serverName: "azsdktest",
 *     sku: {
 *         capacity: 1,
 *         name: "S1",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         testKey: "testValue",
 *     },
 * });
 *
 * ```
 */
export class ServerDetails extends pulumi.CustomResource {
    /**
     * Get an existing ServerDetails resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerDetails {
        return new ServerDetails(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:analysisservices/v20170801:ServerDetails';

    /**
     * Returns true if the given object is an instance of ServerDetails.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerDetails {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerDetails.__pulumiType;
    }

    /**
     * A collection of AS server administrators
     */
    public readonly asAdministrators!: pulumi.Output<outputs.analysisservices.v20170801.ServerAdministratorsResponse | undefined>;
    /**
     * The SAS container URI to the backup container.
     */
    public readonly backupBlobContainerUri!: pulumi.Output<string | undefined>;
    /**
     * The gateway details configured for the AS server.
     */
    public readonly gatewayDetails!: pulumi.Output<outputs.analysisservices.v20170801.GatewayDetailsResponse | undefined>;
    /**
     * The firewall settings for the AS server.
     */
    public readonly ipV4FirewallSettings!: pulumi.Output<outputs.analysisservices.v20170801.IPv4FirewallSettingsResponse | undefined>;
    /**
     * Location of the Analysis Services resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Analysis Services resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
     */
    public readonly querypoolConnectionMode!: pulumi.Output<string | undefined>;
    /**
     * The full name of the Analysis Services resource.
     */
    public /*out*/ readonly serverFullName!: pulumi.Output<string>;
    /**
     * The SKU of the Analysis Services resource.
     */
    public readonly sku!: pulumi.Output<outputs.analysisservices.v20170801.ResourceSkuResponse>;
    /**
     * The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the Analysis Services resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ServerDetails resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerDetailsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerDetailsArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServerDetailsArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["asAdministrators"] = args ? args.asAdministrators : undefined;
            inputs["backupBlobContainerUri"] = args ? args.backupBlobContainerUri : undefined;
            inputs["gatewayDetails"] = args ? args.gatewayDetails : undefined;
            inputs["ipV4FirewallSettings"] = args ? args.ipV4FirewallSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["querypoolConnectionMode"] = args ? args.querypoolConnectionMode : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serverFullName"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:analysisservices/latest:ServerDetails" }, { type: "azurerm:analysisservices/v20160516:ServerDetails" }, { type: "azurerm:analysisservices/v20170714:ServerDetails" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ServerDetails.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerDetails resource.
 */
export interface ServerDetailsArgs {
    /**
     * A collection of AS server administrators
     */
    readonly asAdministrators?: pulumi.Input<inputs.analysisservices.v20170801.ServerAdministrators>;
    /**
     * The SAS container URI to the backup container.
     */
    readonly backupBlobContainerUri?: pulumi.Input<string>;
    /**
     * The gateway details configured for the AS server.
     */
    readonly gatewayDetails?: pulumi.Input<inputs.analysisservices.v20170801.GatewayDetails>;
    /**
     * The firewall settings for the AS server.
     */
    readonly ipV4FirewallSettings?: pulumi.Input<inputs.analysisservices.v20170801.IPv4FirewallSettings>;
    /**
     * Location of the Analysis Services resource.
     */
    readonly location: pulumi.Input<string>;
    /**
     * How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
     */
    readonly querypoolConnectionMode?: pulumi.Input<string>;
    /**
     * The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The SKU of the Analysis Services resource.
     */
    readonly sku: pulumi.Input<inputs.analysisservices.v20170801.ResourceSku>;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
