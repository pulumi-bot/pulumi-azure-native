// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Properties of the blob container, including Id, resource name, resource type, Etag.
 *
 * ## Example Usage
 * ### PutContainers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const blobContainer = new azurerm.storage.v20180701.BlobContainer("blobContainer", {
 *     accountName: "sto328",
 *     containerName: "container6185",
 *     resourceGroupName: "res3376",
 * });
 *
 * ```
 */
export class BlobContainer extends pulumi.CustomResource {
    /**
     * Get an existing BlobContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BlobContainer {
        return new BlobContainer(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storage/v20180701:BlobContainer';

    /**
     * Returns true if the given object is an instance of BlobContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlobContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlobContainer.__pulumiType;
    }

    /**
     * Resource Etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
     */
    public /*out*/ readonly hasImmutabilityPolicy!: pulumi.Output<boolean>;
    /**
     * The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
     */
    public /*out*/ readonly hasLegalHold!: pulumi.Output<boolean>;
    /**
     * The ImmutabilityPolicy property of the container.
     */
    public /*out*/ readonly immutabilityPolicy!: pulumi.Output<outputs.storage.v20180701.ImmutabilityPolicyPropertiesResponse>;
    /**
     * Returns the date and time the container was last modified.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
     */
    public /*out*/ readonly leaseDuration!: pulumi.Output<string>;
    /**
     * Lease state of the container.
     */
    public /*out*/ readonly leaseState!: pulumi.Output<string>;
    /**
     * The lease status of the container.
     */
    public /*out*/ readonly leaseStatus!: pulumi.Output<string>;
    /**
     * The LegalHold property of the container.
     */
    public /*out*/ readonly legalHold!: pulumi.Output<outputs.storage.v20180701.LegalHoldPropertiesResponse>;
    /**
     * A name-value pair to associate with the container as metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether data in the container may be accessed publicly and the level of access.
     */
    public readonly publicAccess!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BlobContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlobContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlobContainerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as BlobContainerArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.containerName === undefined) {
                throw new Error("Missing required property 'containerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["publicAccess"] = args ? args.publicAccess : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["hasImmutabilityPolicy"] = undefined /*out*/;
            inputs["hasLegalHold"] = undefined /*out*/;
            inputs["immutabilityPolicy"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["leaseDuration"] = undefined /*out*/;
            inputs["leaseState"] = undefined /*out*/;
            inputs["leaseStatus"] = undefined /*out*/;
            inputs["legalHold"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storage/latest:BlobContainer" }, { type: "azurerm:storage/v20180201:BlobContainer" }, { type: "azurerm:storage/v20181101:BlobContainer" }, { type: "azurerm:storage/v20190401:BlobContainer" }, { type: "azurerm:storage/v20190601:BlobContainer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(BlobContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BlobContainer resource.
 */
export interface BlobContainerArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
     */
    readonly containerName: pulumi.Input<string>;
    /**
     * A name-value pair to associate with the container as metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether data in the container may be accessed publicly and the level of access.
     */
    readonly publicAccess?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
