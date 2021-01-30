// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The storage account blob inventory policy.
 * Latest API Version: 2019-06-01.
 */
export class BlobInventoryPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BlobInventoryPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BlobInventoryPolicy {
        return new BlobInventoryPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:storage/latest:BlobInventoryPolicy';

    /**
     * Returns true if the given object is an instance of BlobInventoryPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlobInventoryPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlobInventoryPolicy.__pulumiType;
    }

    /**
     * Returns the last modified date and time of the blob inventory policy.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The storage account blob inventory policy object. It is composed of policy rules.
     */
    public readonly policy!: pulumi.Output<outputs.storage.latest.BlobInventoryPolicySchemaResponse>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.storage.latest.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BlobInventoryPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlobInventoryPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.accountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.blobInventoryPolicyName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'blobInventoryPolicyName'");
            }
            if ((!args || args.policy === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["blobInventoryPolicyName"] = args ? args.blobInventoryPolicyName : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["policy"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storage/v20190601:BlobInventoryPolicy" }, { type: "azure-nextgen:storage/v20200801preview:BlobInventoryPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(BlobInventoryPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BlobInventoryPolicy resource.
 */
export interface BlobInventoryPolicyArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the storage account blob inventory policy. It should always be 'default'
     */
    readonly blobInventoryPolicyName: pulumi.Input<string>;
    /**
     * The storage account blob inventory policy object. It is composed of policy rules.
     */
    readonly policy: pulumi.Input<inputs.storage.latest.BlobInventoryPolicySchema>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
