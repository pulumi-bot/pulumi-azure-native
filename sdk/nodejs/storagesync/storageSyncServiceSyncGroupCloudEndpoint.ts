// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Cloud Endpoint object.
 */
export class StorageSyncServiceSyncGroupCloudEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing StorageSyncServiceSyncGroupCloudEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageSyncServiceSyncGroupCloudEndpoint {
        return new StorageSyncServiceSyncGroupCloudEndpoint(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storagesync:StorageSyncServiceSyncGroupCloudEndpoint';

    /**
     * Returns true if the given object is an instance of StorageSyncServiceSyncGroupCloudEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageSyncServiceSyncGroupCloudEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageSyncServiceSyncGroupCloudEndpoint.__pulumiType;
    }

    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Cloud Endpoint properties.
     */
    public readonly properties!: pulumi.Output<outputs.storagesync.CloudEndpointPropertiesResponse>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StorageSyncServiceSyncGroupCloudEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StorageSyncServiceSyncGroupCloudEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
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
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["storageSyncServiceName"] = args ? args.storageSyncServiceName : undefined;
        inputs["syncGroupName"] = args ? args.syncGroupName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StorageSyncServiceSyncGroupCloudEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageSyncServiceSyncGroupCloudEndpoint resource.
 */
export interface StorageSyncServiceSyncGroupCloudEndpointArgs {
    /**
     * Name of Cloud Endpoint object.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The parameters used to create the cloud endpoint.
     */
    readonly properties?: pulumi.Input<inputs.storagesync.CloudEndpointCreateParametersProperties>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Storage Sync Service resource.
     */
    readonly storageSyncServiceName: pulumi.Input<string>;
    /**
     * Name of Sync Group resource.
     */
    readonly syncGroupName: pulumi.Input<string>;
}
