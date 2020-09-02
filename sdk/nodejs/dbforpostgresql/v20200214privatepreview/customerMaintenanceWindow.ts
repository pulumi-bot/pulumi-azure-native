// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Represents a server firewall rule.
 */
export class CustomerMaintenanceWindow extends pulumi.CustomResource {
    /**
     * Get an existing CustomerMaintenanceWindow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomerMaintenanceWindow {
        return new CustomerMaintenanceWindow(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:dbforpostgresql/v20200214privatepreview:CustomerMaintenanceWindow';

    /**
     * Returns true if the given object is an instance of CustomerMaintenanceWindow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerMaintenanceWindow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerMaintenanceWindow.__pulumiType;
    }

    /**
     * The day of week of the customer maintenance window to start
     */
    public readonly dayOfWeek!: pulumi.Output<number>;
    /**
     * The duration of the customer maintenance window to run.
     */
    public readonly durationInMinutes!: pulumi.Output<number | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The starting hour of the customer maintenance window.
     */
    public readonly startHour!: pulumi.Output<number>;
    /**
     * The starting minutes of the customer maintenance window.
     */
    public readonly startMinute!: pulumi.Output<number>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CustomerMaintenanceWindow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerMaintenanceWindowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomerMaintenanceWindowArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CustomerMaintenanceWindowArgs | undefined;
            if (!args || args.dayOfWeek === undefined) {
                throw new Error("Missing required property 'dayOfWeek'");
            }
            if (!args || args.maintenanceWindowName === undefined) {
                throw new Error("Missing required property 'maintenanceWindowName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.startHour === undefined) {
                throw new Error("Missing required property 'startHour'");
            }
            if (!args || args.startMinute === undefined) {
                throw new Error("Missing required property 'startMinute'");
            }
            inputs["dayOfWeek"] = args ? args.dayOfWeek : undefined;
            inputs["durationInMinutes"] = args ? args.durationInMinutes : undefined;
            inputs["maintenanceWindowName"] = args ? args.maintenanceWindowName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["startHour"] = args ? args.startHour : undefined;
            inputs["startMinute"] = args ? args.startMinute : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CustomerMaintenanceWindow.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomerMaintenanceWindow resource.
 */
export interface CustomerMaintenanceWindowArgs {
    /**
     * The day of week of the customer maintenance window to start
     */
    readonly dayOfWeek: pulumi.Input<number>;
    /**
     * The duration of the customer maintenance window to run.
     */
    readonly durationInMinutes?: pulumi.Input<number>;
    /**
     * The name of the maintenance window.
     */
    readonly maintenanceWindowName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The starting hour of the customer maintenance window.
     */
    readonly startHour: pulumi.Input<number>;
    /**
     * The starting minutes of the customer maintenance window.
     */
    readonly startMinute: pulumi.Input<number>;
}
