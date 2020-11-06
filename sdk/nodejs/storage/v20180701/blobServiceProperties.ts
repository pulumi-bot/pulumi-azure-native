// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The properties of a storage account’s Blob service.
 */
export class BlobServiceProperties extends pulumi.CustomResource {
    /**
     * Get an existing BlobServiceProperties resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BlobServiceProperties {
        return new BlobServiceProperties(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:storage/v20180701:BlobServiceProperties';

    /**
     * Returns true if the given object is an instance of BlobServiceProperties.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlobServiceProperties {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlobServiceProperties.__pulumiType;
    }

    /**
     * Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
     */
    public readonly cors!: pulumi.Output<outputs.storage.v20180701.CorsRulesResponse | undefined>;
    /**
     * DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
     */
    public readonly defaultServiceVersion!: pulumi.Output<string | undefined>;
    /**
     * The blob service properties for soft delete.
     */
    public readonly deleteRetentionPolicy!: pulumi.Output<outputs.storage.v20180701.DeleteRetentionPolicyResponse | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BlobServiceProperties resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlobServicePropertiesArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.blobServicesName === undefined) {
                throw new Error("Missing required property 'blobServicesName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["blobServicesName"] = args ? args.blobServicesName : undefined;
            inputs["cors"] = args ? args.cors : undefined;
            inputs["defaultServiceVersion"] = args ? args.defaultServiceVersion : undefined;
            inputs["deleteRetentionPolicy"] = args ? args.deleteRetentionPolicy : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["cors"] = undefined /*out*/;
            inputs["defaultServiceVersion"] = undefined /*out*/;
            inputs["deleteRetentionPolicy"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storage/latest:BlobServiceProperties" }, { type: "azure-nextgen:storage/v20181101:BlobServiceProperties" }, { type: "azure-nextgen:storage/v20190401:BlobServiceProperties" }, { type: "azure-nextgen:storage/v20190601:BlobServiceProperties" }, { type: "azure-nextgen:storage/v20200801preview:BlobServiceProperties" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(BlobServiceProperties.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BlobServiceProperties resource.
 */
export interface BlobServicePropertiesArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
     */
    readonly blobServicesName: pulumi.Input<string>;
    /**
     * Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Blob service.
     */
    readonly cors?: pulumi.Input<inputs.storage.v20180701.CorsRules>;
    /**
     * DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
     */
    readonly defaultServiceVersion?: pulumi.Input<string>;
    /**
     * The blob service properties for soft delete.
     */
    readonly deleteRetentionPolicy?: pulumi.Input<inputs.storage.v20180701.DeleteRetentionPolicy>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
