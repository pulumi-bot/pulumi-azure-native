// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
 */
export class BlobContainerImmutabilityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BlobContainerImmutabilityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BlobContainerImmutabilityPolicy {
        return new BlobContainerImmutabilityPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storage/v20190401:BlobContainerImmutabilityPolicy';

    /**
     * Returns true if the given object is an instance of BlobContainerImmutabilityPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlobContainerImmutabilityPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlobContainerImmutabilityPolicy.__pulumiType;
    }

    /**
     * Resource Etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The immutability period for the blobs in the container since the policy creation, in days.
     */
    public readonly immutabilityPeriodSinceCreationInDays!: pulumi.Output<number>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BlobContainerImmutabilityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlobContainerImmutabilityPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.containerName === undefined) {
                throw new Error("Missing required property 'containerName'");
            }
            if (!args || args.immutabilityPeriodSinceCreationInDays === undefined) {
                throw new Error("Missing required property 'immutabilityPeriodSinceCreationInDays'");
            }
            if (!args || args.immutabilityPolicyName === undefined) {
                throw new Error("Missing required property 'immutabilityPolicyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["immutabilityPeriodSinceCreationInDays"] = args ? args.immutabilityPeriodSinceCreationInDays : undefined;
            inputs["immutabilityPolicyName"] = args ? args.immutabilityPolicyName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["immutabilityPeriodSinceCreationInDays"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storage/latest:BlobContainerImmutabilityPolicy" }, { type: "azurerm:storage/v20180201:BlobContainerImmutabilityPolicy" }, { type: "azurerm:storage/v20180301preview:BlobContainerImmutabilityPolicy" }, { type: "azurerm:storage/v20180701:BlobContainerImmutabilityPolicy" }, { type: "azurerm:storage/v20181101:BlobContainerImmutabilityPolicy" }, { type: "azurerm:storage/v20190601:BlobContainerImmutabilityPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(BlobContainerImmutabilityPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BlobContainerImmutabilityPolicy resource.
 */
export interface BlobContainerImmutabilityPolicyArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
     */
    readonly containerName: pulumi.Input<string>;
    /**
     * The immutability period for the blobs in the container since the policy creation, in days.
     */
    readonly immutabilityPeriodSinceCreationInDays: pulumi.Input<number>;
    /**
     * The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
     */
    readonly immutabilityPolicyName: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
