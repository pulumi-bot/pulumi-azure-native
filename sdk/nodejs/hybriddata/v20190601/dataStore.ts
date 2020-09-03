// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Data store.
 *
 * ## Example Usage
 * ### DataStores_CreateOrUpdate_DataSinkPUT162
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataStore = new azurerm.hybriddata.v20190601.DataStore("dataStore", {
 *     customerSecrets: [
 *         {
 *             algorithm: "RSA1_5",
 *             keyIdentifier: "StorageAccountAccessKey",
 *             keyValue: "Of4H9eF03G8QuxvkZQEbFWv3YdN3U//WugzuqReQekbXXQyg+QSicVKrwSOOKVi1zWMYGbKg7d5/ES2gdz+O5ZEw89bvE4mJD/wQmkIsqhPnbN0gyVK6nZePXVUU1A+UzjLfvhSA6KyUQfzNAZ5/TLt6fo1JyQrKTtkvnkLFyfv1AqBZ+dw8JK3RZi/rEN8HD3R3qsBwUYfyEuGLGiujy2CGrr/1uPiUVMR6QuFWRsjm39eMSHa4maLg4tQ0IY/jIy8rMlx3KjF3CcCbPzAqEq5vXy37wkjZbus771te1gLSrzcpVKIMg4DrmgaoJ02jAu+izBjNgLXAFPSUneQ8yw==:ezMkh4PMhCnjJtYkpTaP0SdblP5VAeRe4glW2PgIzICHw3T8ZyGDoaTrCv4/m5wtcEhWdtxhta+j1MQWlK5MIA+hvf8QjIDIjQv696ov5y+pcFe/upd2ekGOei7FCwB2u7I8WnkAtIKTUkf6eDQBZXm26DjfG1Dlc+Mjjq+AerukEa6WpOyqrD7Qub26Pgmj4AsuBx19X1EAmTZacubkoiNagXM8V+IDanHOhLMvfgQ7rw8oZhWfofxi4m+eJqjOXXaqSyorNK8UEcqP6P9pDP8AN8ulXEx6rZy2B5Oi0vSV+wlRLbUuQslga4ItOGxctW/ZX8uWozt+5A3k4URt6A==",
 *         },
 *         {
 *             algorithm: "RSA1_5",
 *             keyIdentifier: "StorageAccountAccessKeyForQueue",
 *             keyValue: "Of4H9eF03G8QuxvkZQEbFWv3YdN3U//WugzuqReQekbXXQyg+QSicVKrwSOOKVi1zWMYGbKg7d5/ES2gdz+O5ZEw89bvE4mJD/wQmkIsqhPnbN0gyVK6nZePXVUU1A+UzjLfvhSA6KyUQfzNAZ5/TLt6fo1JyQrKTtkvnkLFyfv1AqBZ+dw8JK3RZi/rEN8HD3R3qsBwUYfyEuGLGiujy2CGrr/1uPiUVMR6QuFWRsjm39eMSHa4maLg4tQ0IY/jIy8rMlx3KjF3CcCbPzAqEq5vXy37wkjZbus771te1gLSrzcpVKIMg4DrmgaoJ02jAu+izBjNgLXAFPSUneQ8yw==:ezMkh4PMhCnjJtYkpTaP0SdblP5VAeRe4glW2PgIzICHw3T8ZyGDoaTrCv4/m5wtcEhWdtxhta+j1MQWlK5MIA+hvf8QjIDIjQv696ov5y+pcFe/upd2ekGOei7FCwB2u7I8WnkAtIKTUkf6eDQBZXm26DjfG1Dlc+Mjjq+AerukEa6WpOyqrD7Qub26Pgmj4AsuBx19X1EAmTZacubkoiNagXM8V+IDanHOhLMvfgQ7rw8oZhWfofxi4m+eJqjOXXaqSyorNK8UEcqP6P9pDP8AN8ulXEx6rZy2B5Oi0vSV+wlRLbUuQslga4ItOGxctW/ZX8uWozt+5A3k4URt6A==",
 *         },
 *     ],
 *     dataManagerName: "TestAzureSDKOperations",
 *     dataStoreName: "TestAzureStorage1",
 *     dataStoreTypeId: "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount",
 *     extendedProperties: {
 *         extendedSaKey: undefined,
 *         extendedSaName: "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
 *         storageAccountNameForQueue: "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
 *     },
 *     repositoryId: "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
 *     resourceGroupName: "ResourceGroupForSDKTest",
 *     state: "Enabled",
 * });
 *
 * ```
 * ### DataStores_CreateOrUpdate_DataSourcePUT162
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataStore = new azurerm.hybriddata.v20190601.DataStore("dataStore", {
 *     customerSecrets: [{
 *         algorithm: "RSA1_5",
 *         keyIdentifier: "ServiceEncryptionKey",
 *         keyValue: "EVuEBV40qv23xDRL4NZBuMms4e3So6ikHjrQYRvG9NloqxdgPOg+ZYzpho5lytI4fmv0ANmRIvDiDboRXcUVSjbB9T2gm19fMIuwZa4FK2+LYEgMqKK1GaLkk7xC8f5IeFUXLo6KyBBpaAiayTnWDcHuYEpMiGrV7trDDcbhMRefO3CHecmH3Z7ye8L0RQ/e7WW8GlCKZj3m0BaG7OrJgjai8gyDfMfGAG5rTqEhDVh2hLQ+TjvUjcOFwHvJusqKTENtbJTNQYmL9wZXsnwBvUwxqrGieILNB7V3GD1Ow9OiV0UCDW1e9LnMueukg+l7YJCU9FUhIPh/nSif6p32zw==:jCfio+pDtY3BSPZDpDJ0L6QdXLYMeOmxaFWtYTOZkNqNTgT8Loc/KSQRjtWS5K4N4btbznuSJ/dzg0aZEzlXgKDSuZgMfd4Ch92ZwAC/BkeCmVrTjiKJsoQXO1IICCUf7GHGBbYnnpsNJcEn4vyc9NXyKwOBjeU+I9AyK7PtYiC03RLpTS6xttFCICteBV0uoBHAiV0chZ5VIIUUMjO9u8EhHqRY7NNcGbWdVJeAb6J3vH4E/DHkQj+DXlpjcLvmK/uqBwxfNju30RJhR04Nmz6zcv/zTcvS0uN5hEPQoHLyv84hjnc4omg/gmNjo2cDW64QxA3RTJ5Sl///4xTBkg==",
 *     }],
 *     dataManagerName: "TestAzureSDKOperations",
 *     dataStoreName: "TestStorSimpleSource1",
 *     dataStoreTypeId: "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series",
 *     extendedProperties: {
 *         extendedSaKey: undefined,
 *         resourceId: "/subscriptions/c5fc377d-0085-41b9-86b7-cc96dc56d1e9/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600",
 *     },
 *     repositoryId: "/subscriptions/c5fc377d-0085-41b9-86b7-cc96dc56d1e9/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600",
 *     resourceGroupName: "ResourceGroupForSDKTest",
 *     state: "Enabled",
 * });
 *
 * ```
 */
export class DataStore extends pulumi.CustomResource {
    /**
     * Get an existing DataStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataStore {
        return new DataStore(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hybriddata/v20190601:DataStore';

    /**
     * Returns true if the given object is an instance of DataStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataStore.__pulumiType;
    }

    /**
     * List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
     */
    public readonly customerSecrets!: pulumi.Output<outputs.hybriddata.v20190601.CustomerSecretResponse[] | undefined>;
    /**
     * The arm id of the data store type.
     */
    public readonly dataStoreTypeId!: pulumi.Output<string>;
    /**
     * A generic json used differently by each data source type.
     */
    public readonly extendedProperties!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Arm Id for the manager resource to which the data source is associated. This is optional.
     */
    public readonly repositoryId!: pulumi.Output<string | undefined>;
    /**
     * State of the data source.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DataStoreArgs | undefined;
            if (!args || args.dataManagerName === undefined) {
                throw new Error("Missing required property 'dataManagerName'");
            }
            if (!args || args.dataStoreName === undefined) {
                throw new Error("Missing required property 'dataStoreName'");
            }
            if (!args || args.dataStoreTypeId === undefined) {
                throw new Error("Missing required property 'dataStoreTypeId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.state === undefined) {
                throw new Error("Missing required property 'state'");
            }
            inputs["customerSecrets"] = args ? args.customerSecrets : undefined;
            inputs["dataManagerName"] = args ? args.dataManagerName : undefined;
            inputs["dataStoreName"] = args ? args.dataStoreName : undefined;
            inputs["dataStoreTypeId"] = args ? args.dataStoreTypeId : undefined;
            inputs["extendedProperties"] = args ? args.extendedProperties : undefined;
            inputs["repositoryId"] = args ? args.repositoryId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:hybriddata/latest:DataStore" }, { type: "azurerm:hybriddata/v20160601:DataStore" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DataStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataStore resource.
 */
export interface DataStoreArgs {
    /**
     * List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
     */
    readonly customerSecrets?: pulumi.Input<pulumi.Input<inputs.hybriddata.v20190601.CustomerSecret>[]>;
    /**
     * The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
     */
    readonly dataManagerName: pulumi.Input<string>;
    /**
     * The data store/repository name to be created or updated.
     */
    readonly dataStoreName: pulumi.Input<string>;
    /**
     * The arm id of the data store type.
     */
    readonly dataStoreTypeId: pulumi.Input<string>;
    /**
     * A generic json used differently by each data source type.
     */
    readonly extendedProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Arm Id for the manager resource to which the data source is associated. This is optional.
     */
    readonly repositoryId?: pulumi.Input<string>;
    /**
     * The Resource Group Name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * State of the data source.
     */
    readonly state: pulumi.Input<string>;
}
