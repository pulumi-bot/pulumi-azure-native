// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Class representing a database principal assignment.
 */
export class DatabasePrincipalAssignment extends pulumi.CustomResource {
    /**
     * Get an existing DatabasePrincipalAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabasePrincipalAssignment {
        return new DatabasePrincipalAssignment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:kusto/v20200614:DatabasePrincipalAssignment';

    /**
     * Returns true if the given object is an instance of DatabasePrincipalAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabasePrincipalAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabasePrincipalAssignment.__pulumiType;
    }

    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The database principal.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.kusto.v20200614.DatabasePrincipalPropertiesResponse>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabasePrincipalAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabasePrincipalAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabasePrincipalAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DatabasePrincipalAssignmentArgs | undefined;
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.principalId === undefined) {
                throw new Error("Missing required property 'principalId'");
            }
            if (!args || args.principalType === undefined) {
                throw new Error("Missing required property 'principalType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["principalId"] = args ? args.principalId : undefined;
            inputs["principalType"] = args ? args.principalType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabasePrincipalAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabasePrincipalAssignment resource.
 */
export interface DatabasePrincipalAssignmentArgs {
    /**
     * The name of the Kusto cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The name of the database in the Kusto cluster.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The name of the Kusto principalAssignment.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
     */
    readonly principalId: pulumi.Input<string>;
    /**
     * Principal type.
     */
    readonly principalType: pulumi.Input<string>;
    /**
     * The name of the resource group containing the Kusto cluster.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Database principal role.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The tenant id of the principal
     */
    readonly tenantId?: pulumi.Input<string>;
}
