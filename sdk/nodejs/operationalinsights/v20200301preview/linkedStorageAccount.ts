// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Linked storage accounts top level resource container.
 */
export class LinkedStorageAccount extends pulumi.CustomResource {
    /**
     * Get an existing LinkedStorageAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LinkedStorageAccount {
        return new LinkedStorageAccount(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:operationalinsights/v20200301preview:LinkedStorageAccount';

    /**
     * Returns true if the given object is an instance of LinkedStorageAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedStorageAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedStorageAccount.__pulumiType;
    }

    /**
     * Linked storage accounts type.
     */
    public readonly dataSourceType!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Linked storage accounts resources ids.
     */
    public readonly storageAccountIds!: pulumi.Output<string[] | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LinkedStorageAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedStorageAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.dataSourceType === undefined) {
                throw new Error("Missing required property 'dataSourceType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["dataSourceType"] = args ? args.dataSourceType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountIds"] = args ? args.storageAccountIds : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["dataSourceType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["storageAccountIds"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:operationalinsights/latest:LinkedStorageAccount" }, { type: "azure-nextgen:operationalinsights/v20190801preview:LinkedStorageAccount" }, { type: "azure-nextgen:operationalinsights/v20200801:LinkedStorageAccount" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LinkedStorageAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LinkedStorageAccount resource.
 */
export interface LinkedStorageAccountArgs {
    /**
     * Linked storage accounts type.
     */
    readonly dataSourceType: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Linked storage accounts resources ids.
     */
    readonly storageAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
