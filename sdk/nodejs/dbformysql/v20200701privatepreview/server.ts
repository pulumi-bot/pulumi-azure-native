// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a server.
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:dbformysql/v20200701privatepreview:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
     */
    public readonly administratorLogin!: pulumi.Output<string | undefined>;
    /**
     * The password of the administrator login (required for server creation).
     */
    public readonly administratorLoginPassword!: pulumi.Output<string | undefined>;
    /**
     * The mode to create a new MySQL server.
     */
    public readonly createMode!: pulumi.Output<string | undefined>;
    /**
     * Earliest restore point creation time (ISO8601 format)
     */
    public readonly earliestRestoreDate!: pulumi.Output<string | undefined>;
    /**
     * The fully qualified domain name of a server.
     */
    public /*out*/ readonly fullyQualifiedDomainName!: pulumi.Output<string>;
    /**
     * The state of a HA server.
     */
    public /*out*/ readonly haState!: pulumi.Output<string>;
    /**
     * The Azure Active Directory identity of the server.
     */
    public readonly identity!: pulumi.Output<outputs.dbformysql.v20200701privatepreview.IdentityResponse | undefined>;
    /**
     * Status showing whether the server enabled infrastructure encryption.
     */
    public readonly infrastructureEncryption!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The primary server id of a replica server.
     */
    public /*out*/ readonly primaryServerId!: pulumi.Output<string>;
    /**
     * Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
     */
    public readonly publicNetworkAccess!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of replicas that a primary server can have.
     */
    public readonly replicaCapacity!: pulumi.Output<number | undefined>;
    /**
     * The replication role.
     */
    public readonly replicationRole!: pulumi.Output<string | undefined>;
    /**
     * Restore point creation time (ISO8601 format), specifying the time to restore from.
     */
    public readonly restorePointInTime!: pulumi.Output<string | undefined>;
    /**
     * The SKU (pricing tier) of the server.
     */
    public readonly sku!: pulumi.Output<outputs.dbformysql.v20200701privatepreview.SkuResponse | undefined>;
    /**
     * The source MySQL server name to restore from.
     */
    public readonly sourceServerId!: pulumi.Output<string | undefined>;
    /**
     * Enable ssl enforcement or not when connect to server.
     */
    public readonly sslEnforcement!: pulumi.Output<string | undefined>;
    /**
     * stand by count value can be either 0 or 1
     */
    public readonly standbyCount!: pulumi.Output<number | undefined>;
    /**
     * The state of a server.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Storage profile of a server.
     */
    public readonly storageProfile!: pulumi.Output<outputs.dbformysql.v20200701privatepreview.StorageProfileResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Server version.
     */
    public readonly version!: pulumi.Output<string | undefined>;
    /**
     * Vnet arguments.
     */
    public readonly vnetInjArgs!: pulumi.Output<outputs.dbformysql.v20200701privatepreview.VnetInjArgsResponse | undefined>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServerArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["administratorLogin"] = args ? args.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = args ? args.administratorLoginPassword : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["earliestRestoreDate"] = args ? args.earliestRestoreDate : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["infrastructureEncryption"] = args ? args.infrastructureEncryption : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            inputs["replicaCapacity"] = args ? args.replicaCapacity : undefined;
            inputs["replicationRole"] = args ? args.replicationRole : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["sourceServerId"] = args ? args.sourceServerId : undefined;
            inputs["sslEnforcement"] = args ? args.sslEnforcement : undefined;
            inputs["standbyCount"] = args ? args.standbyCount : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["vnetInjArgs"] = args ? args.vnetInjArgs : undefined;
            inputs["fullyQualifiedDomainName"] = undefined /*out*/;
            inputs["haState"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["primaryServerId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
     */
    readonly administratorLogin?: pulumi.Input<string>;
    /**
     * The password of the administrator login (required for server creation).
     */
    readonly administratorLoginPassword?: pulumi.Input<string>;
    /**
     * The mode to create a new MySQL server.
     */
    readonly createMode?: pulumi.Input<string>;
    /**
     * Earliest restore point creation time (ISO8601 format)
     */
    readonly earliestRestoreDate?: pulumi.Input<string>;
    /**
     * The Azure Active Directory identity of the server.
     */
    readonly identity?: pulumi.Input<inputs.dbformysql.v20200701privatepreview.Identity>;
    /**
     * Status showing whether the server enabled infrastructure encryption.
     */
    readonly infrastructureEncryption?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
     */
    readonly publicNetworkAccess?: pulumi.Input<string>;
    /**
     * The maximum number of replicas that a primary server can have.
     */
    readonly replicaCapacity?: pulumi.Input<number>;
    /**
     * The replication role.
     */
    readonly replicationRole?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Restore point creation time (ISO8601 format), specifying the time to restore from.
     */
    readonly restorePointInTime?: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The SKU (pricing tier) of the server.
     */
    readonly sku?: pulumi.Input<inputs.dbformysql.v20200701privatepreview.Sku>;
    /**
     * The source MySQL server name to restore from.
     */
    readonly sourceServerId?: pulumi.Input<string>;
    /**
     * Enable ssl enforcement or not when connect to server.
     */
    readonly sslEnforcement?: pulumi.Input<string>;
    /**
     * stand by count value can be either 0 or 1
     */
    readonly standbyCount?: pulumi.Input<number>;
    /**
     * Storage profile of a server.
     */
    readonly storageProfile?: pulumi.Input<inputs.dbformysql.v20200701privatepreview.StorageProfile>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Server version.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Vnet arguments.
     */
    readonly vnetInjArgs?: pulumi.Input<inputs.dbformysql.v20200701privatepreview.VnetInjArgs>;
}
