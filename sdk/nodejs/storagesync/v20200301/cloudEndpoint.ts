// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Cloud Endpoint object.
 *
 * ## Example Usage
 * ### CloudEndpoints_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cloudEndpoint = new azurerm.storagesync.v20200301.CloudEndpoint("cloudEndpoint", {
 *     azureFileShareName: "cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4",
 *     cloudEndpointName: "SampleCloudEndpoint_1",
 *     friendlyName: "ankushbsubscriptionmgmtmab",
 *     resourceGroupName: "SampleResourceGroup_1",
 *     storageAccountResourceId: "/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage",
 *     storageAccountTenantId: "\"72f988bf-86f1-41af-91ab-2d7cd011db47\"",
 *     storageSyncServiceName: "SampleStorageSyncService_1",
 *     syncGroupName: "SampleSyncGroup_1",
 * });
 *
 * ```
 */
export class CloudEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing CloudEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CloudEndpoint {
        return new CloudEndpoint(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storagesync/v20200301:CloudEndpoint';

    /**
     * Returns true if the given object is an instance of CloudEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudEndpoint.__pulumiType;
    }

    /**
     * Azure file share name
     */
    public readonly azureFileShareName!: pulumi.Output<string | undefined>;
    /**
     * Backup Enabled
     */
    public /*out*/ readonly backupEnabled!: pulumi.Output<string>;
    /**
     * Friendly Name
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * Resource Last Operation Name
     */
    public /*out*/ readonly lastOperationName!: pulumi.Output<string | undefined>;
    /**
     * CloudEndpoint lastWorkflowId
     */
    public /*out*/ readonly lastWorkflowId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Partnership Id
     */
    public /*out*/ readonly partnershipId!: pulumi.Output<string | undefined>;
    /**
     * CloudEndpoint Provisioning State
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Storage Account Resource Id
     */
    public readonly storageAccountResourceId!: pulumi.Output<string | undefined>;
    /**
     * Storage Account Tenant Id
     */
    public readonly storageAccountTenantId!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CloudEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CloudEndpointArgs | undefined;
            if (!args || args.cloudEndpointName === undefined) {
                throw new Error("Missing required property 'cloudEndpointName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageSyncServiceName === undefined) {
                throw new Error("Missing required property 'storageSyncServiceName'");
            }
            if (!args || args.syncGroupName === undefined) {
                throw new Error("Missing required property 'syncGroupName'");
            }
            inputs["azureFileShareName"] = args ? args.azureFileShareName : undefined;
            inputs["cloudEndpointName"] = args ? args.cloudEndpointName : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountResourceId"] = args ? args.storageAccountResourceId : undefined;
            inputs["storageAccountTenantId"] = args ? args.storageAccountTenantId : undefined;
            inputs["storageSyncServiceName"] = args ? args.storageSyncServiceName : undefined;
            inputs["syncGroupName"] = args ? args.syncGroupName : undefined;
            inputs["backupEnabled"] = undefined /*out*/;
            inputs["lastOperationName"] = undefined /*out*/;
            inputs["lastWorkflowId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnershipId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storagesync/latest:CloudEndpoint" }, { type: "azurerm:storagesync/v20180402:CloudEndpoint" }, { type: "azurerm:storagesync/v20180701:CloudEndpoint" }, { type: "azurerm:storagesync/v20181001:CloudEndpoint" }, { type: "azurerm:storagesync/v20190201:CloudEndpoint" }, { type: "azurerm:storagesync/v20190301:CloudEndpoint" }, { type: "azurerm:storagesync/v20190601:CloudEndpoint" }, { type: "azurerm:storagesync/v20191001:CloudEndpoint" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(CloudEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloudEndpoint resource.
 */
export interface CloudEndpointArgs {
    /**
     * Azure file share name
     */
    readonly azureFileShareName?: pulumi.Input<string>;
    /**
     * Name of Cloud Endpoint object.
     */
    readonly cloudEndpointName: pulumi.Input<string>;
    /**
     * Friendly Name
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Storage Account Resource Id
     */
    readonly storageAccountResourceId?: pulumi.Input<string>;
    /**
     * Storage Account Tenant Id
     */
    readonly storageAccountTenantId?: pulumi.Input<string>;
    /**
     * Name of Storage Sync Service resource.
     */
    readonly storageSyncServiceName: pulumi.Input<string>;
    /**
     * Name of Sync Group resource.
     */
    readonly syncGroupName: pulumi.Input<string>;
}
