// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Contains information about an Azure Batch account.
 */
export class BatchAccount extends pulumi.CustomResource {
    /**
     * Get an existing BatchAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BatchAccount {
        return new BatchAccount(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:batch/v20170501:BatchAccount';

    /**
     * Returns true if the given object is an instance of BatchAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BatchAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BatchAccount.__pulumiType;
    }

    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties associated with the account.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.batch.v20170501.BatchAccountPropertiesResponse>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BatchAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BatchAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BatchAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as BatchAccountArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoStorage"] = args ? args.autoStorage : undefined;
            inputs["keyVaultReference"] = args ? args.keyVaultReference : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["poolAllocationMode"] = args ? args.poolAllocationMode : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BatchAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BatchAccount resource.
 */
export interface BatchAccountArgs {
    /**
     * The properties related to the auto-storage account.
     */
    readonly autoStorage?: pulumi.Input<inputs.batch.v20170501.AutoStorageBaseProperties>;
    /**
     * A reference to the Azure key vault associated with the Batch account.
     */
    readonly keyVaultReference?: pulumi.Input<inputs.batch.v20170501.KeyVaultReference>;
    /**
     * The region in which to create the account.
     */
    readonly location: pulumi.Input<string>;
    /**
     * A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24 characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is used to access the Batch service in the region in which the account is created. For example: http://accountname.region.batch.azure.com/.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService.
     */
    readonly poolAllocationMode?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The user-specified tags associated with the account.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
