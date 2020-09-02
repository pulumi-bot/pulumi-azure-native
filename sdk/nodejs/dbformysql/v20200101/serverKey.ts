// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A MySQL Server key.
 *
 * ## Example Usage
 * ### Creates or updates a MySQL Server key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const serverKey = new azurerm.dbformysql.v20200101.ServerKey("serverKey", {
 *     keyName: "someVault_someKey_01234567890123456789012345678901",
 *     resourceGroupName: "testrg",
 *     serverKeyType: "AzureKeyVault",
 *     serverName: "testserver",
 *     uri: "https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901",
 * });
 *
 * ```
 */
export class ServerKey extends pulumi.CustomResource {
    /**
     * Get an existing ServerKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerKey {
        return new ServerKey(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:dbformysql/v20200101:ServerKey';

    /**
     * Returns true if the given object is an instance of ServerKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerKey.__pulumiType;
    }

    /**
     * The key creation date.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Kind of encryption protector used to protect the key.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The key type like 'AzureKeyVault'.
     */
    public readonly serverKeyType!: pulumi.Output<string>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The URI of the key.
     */
    public readonly uri!: pulumi.Output<string | undefined>;

    /**
     * Create a ServerKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServerKeyArgs | undefined;
            if (!args || args.keyName === undefined) {
                throw new Error("Missing required property 'keyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverKeyType === undefined) {
                throw new Error("Missing required property 'serverKeyType'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverKeyType"] = args ? args.serverKeyType : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["uri"] = args ? args.uri : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:dbformysql/latest:ServerKey" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ServerKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerKey resource.
 */
export interface ServerKeyArgs {
    /**
     * The name of the MySQL Server key to be operated on (updated or created).
     */
    readonly keyName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The key type like 'AzureKeyVault'.
     */
    readonly serverKeyType: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The URI of the key.
     */
    readonly uri?: pulumi.Input<string>;
}
