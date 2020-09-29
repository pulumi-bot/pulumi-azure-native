// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Cloud Endpoint object.
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
        return new CloudEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:storagesync/v20180701:CloudEndpoint';

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
     * Backup Enabled
     */
    public /*out*/ readonly backupEnabled!: pulumi.Output<boolean>;
    /**
     * Friendly Name
     */
    public /*out*/ readonly friendlyName!: pulumi.Output<string | undefined>;
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
     * Storage Account Share name
     */
    public readonly storageAccountShareName!: pulumi.Output<string | undefined>;
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
    constructor(name: string, args: CloudEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
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
            inputs["cloudEndpointName"] = args ? args.cloudEndpointName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountResourceId"] = args ? args.storageAccountResourceId : undefined;
            inputs["storageAccountShareName"] = args ? args.storageAccountShareName : undefined;
            inputs["storageAccountTenantId"] = args ? args.storageAccountTenantId : undefined;
            inputs["storageSyncServiceName"] = args ? args.storageSyncServiceName : undefined;
            inputs["syncGroupName"] = args ? args.syncGroupName : undefined;
            inputs["backupEnabled"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["lastOperationName"] = undefined /*out*/;
            inputs["lastWorkflowId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnershipId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backupEnabled"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["lastOperationName"] = undefined /*out*/;
            inputs["lastWorkflowId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnershipId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["storageAccountResourceId"] = undefined /*out*/;
            inputs["storageAccountShareName"] = undefined /*out*/;
            inputs["storageAccountTenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storagesync/latest:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20170605preview:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20180402:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20181001:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20190201:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20190301:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20190601:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20191001:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20200301:CloudEndpoint" }, { type: "azure-nextgen:storagesync/v20200901:CloudEndpoint" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(CloudEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloudEndpoint resource.
 */
export interface CloudEndpointArgs {
    /**
     * Name of Cloud Endpoint object.
     */
    readonly cloudEndpointName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Storage Account Resource Id
     */
    readonly storageAccountResourceId?: pulumi.Input<string>;
    /**
     * Storage Account Share name
     */
    readonly storageAccountShareName?: pulumi.Input<string>;
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
