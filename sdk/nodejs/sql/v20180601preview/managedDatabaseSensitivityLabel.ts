// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A sensitivity label.
 */
export class ManagedDatabaseSensitivityLabel extends pulumi.CustomResource {
    /**
     * Get an existing ManagedDatabaseSensitivityLabel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedDatabaseSensitivityLabel {
        return new ManagedDatabaseSensitivityLabel(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20180601preview:ManagedDatabaseSensitivityLabel';

    /**
     * Returns true if the given object is an instance of ManagedDatabaseSensitivityLabel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedDatabaseSensitivityLabel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedDatabaseSensitivityLabel.__pulumiType;
    }

    /**
     * The information type.
     */
    public readonly informationType!: pulumi.Output<string | undefined>;
    /**
     * The information type ID.
     */
    public readonly informationTypeId!: pulumi.Output<string | undefined>;
    /**
     * Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
     */
    public /*out*/ readonly isDisabled!: pulumi.Output<boolean>;
    /**
     * The label ID.
     */
    public readonly labelId!: pulumi.Output<string | undefined>;
    /**
     * The label name.
     */
    public readonly labelName!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly rank!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagedDatabaseSensitivityLabel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedDatabaseSensitivityLabelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedDatabaseSensitivityLabelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ManagedDatabaseSensitivityLabelArgs | undefined;
            if (!args || args.columnName === undefined) {
                throw new Error("Missing required property 'columnName'");
            }
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.managedInstanceName === undefined) {
                throw new Error("Missing required property 'managedInstanceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.schemaName === undefined) {
                throw new Error("Missing required property 'schemaName'");
            }
            if (!args || args.sensitivityLabelSource === undefined) {
                throw new Error("Missing required property 'sensitivityLabelSource'");
            }
            if (!args || args.tableName === undefined) {
                throw new Error("Missing required property 'tableName'");
            }
            inputs["columnName"] = args ? args.columnName : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["informationType"] = args ? args.informationType : undefined;
            inputs["informationTypeId"] = args ? args.informationTypeId : undefined;
            inputs["labelId"] = args ? args.labelId : undefined;
            inputs["labelName"] = args ? args.labelName : undefined;
            inputs["managedInstanceName"] = args ? args.managedInstanceName : undefined;
            inputs["rank"] = args ? args.rank : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["schemaName"] = args ? args.schemaName : undefined;
            inputs["sensitivityLabelSource"] = args ? args.sensitivityLabelSource : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
            inputs["isDisabled"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagedDatabaseSensitivityLabel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedDatabaseSensitivityLabel resource.
 */
export interface ManagedDatabaseSensitivityLabelArgs {
    /**
     * The name of the column.
     */
    readonly columnName: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The information type.
     */
    readonly informationType?: pulumi.Input<string>;
    /**
     * The information type ID.
     */
    readonly informationTypeId?: pulumi.Input<string>;
    /**
     * The label ID.
     */
    readonly labelId?: pulumi.Input<string>;
    /**
     * The label name.
     */
    readonly labelName?: pulumi.Input<string>;
    /**
     * The name of the managed instance.
     */
    readonly managedInstanceName: pulumi.Input<string>;
    readonly rank?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the schema.
     */
    readonly schemaName: pulumi.Input<string>;
    /**
     * The source of the sensitivity label.
     */
    readonly sensitivityLabelSource: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    readonly tableName: pulumi.Input<string>;
}
