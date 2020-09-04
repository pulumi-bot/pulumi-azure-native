// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Storage Sync Service object.
 *
 * ## Example Usage
 * ### StorageSyncServices_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const storageSyncService = new azurerm.storagesync.v20180402.StorageSyncService("storageSyncService", {
 *     location: "WestUS",
 *     resourceGroupName: "SampleResourceGroup_1",
 *     storageSyncServiceName: "SampleStorageSyncService_1",
 *     tags: {},
 * });
 *
 * ```
 */
export class StorageSyncService extends pulumi.CustomResource {
    /**
     * Get an existing StorageSyncService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageSyncService {
        return new StorageSyncService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storagesync/v20180402:StorageSyncService';

    /**
     * Returns true if the given object is an instance of StorageSyncService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageSyncService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageSyncService.__pulumiType;
    }

    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Storage Sync service status.
     */
    public /*out*/ readonly storageSyncServiceStatus!: pulumi.Output<number>;
    /**
     * Storage Sync service Uid
     */
    public /*out*/ readonly storageSyncServiceUid!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StorageSyncService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageSyncServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageSyncServiceName === undefined) {
                throw new Error("Missing required property 'storageSyncServiceName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageSyncServiceName"] = args ? args.storageSyncServiceName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["storageSyncServiceStatus"] = undefined /*out*/;
            inputs["storageSyncServiceUid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["storageSyncServiceStatus"] = undefined /*out*/;
            inputs["storageSyncServiceUid"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storagesync/latest:StorageSyncService" }, { type: "azurerm:storagesync/v20170605preview:StorageSyncService" }, { type: "azurerm:storagesync/v20180701:StorageSyncService" }, { type: "azurerm:storagesync/v20181001:StorageSyncService" }, { type: "azurerm:storagesync/v20190201:StorageSyncService" }, { type: "azurerm:storagesync/v20190301:StorageSyncService" }, { type: "azurerm:storagesync/v20190601:StorageSyncService" }, { type: "azurerm:storagesync/v20191001:StorageSyncService" }, { type: "azurerm:storagesync/v20200301:StorageSyncService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(StorageSyncService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageSyncService resource.
 */
export interface StorageSyncServiceArgs {
    /**
     * Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Storage Sync Service resource.
     */
    readonly storageSyncServiceName: pulumi.Input<string>;
    /**
     * Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
