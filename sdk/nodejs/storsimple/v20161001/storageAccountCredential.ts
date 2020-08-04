// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The storage account credential
 */
export class StorageAccountCredential extends pulumi.CustomResource {
    /**
     * Get an existing StorageAccountCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageAccountCredential {
        return new StorageAccountCredential(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20161001:StorageAccountCredential';

    /**
     * Returns true if the given object is an instance of StorageAccountCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageAccountCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageAccountCredential.__pulumiType;
    }

    /**
     * The name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Credential properties
     */
    public readonly properties!: pulumi.Output<outputs.storsimple.v20161001.StorageAccountCredentialPropertiesResponse>;
    /**
     * The type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StorageAccountCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageAccountCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageAccountCredentialArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as StorageAccountCredentialArgs | undefined;
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StorageAccountCredential.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageAccountCredential resource.
 */
export interface StorageAccountCredentialArgs {
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The credential name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Credential properties
     */
    readonly properties: pulumi.Input<inputs.storsimple.v20161001.StorageAccountCredentialProperties>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
