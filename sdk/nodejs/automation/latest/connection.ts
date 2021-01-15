// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Definition of the connection.
 * Latest API Version: 2019-06-01.
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:automation/latest:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * Gets or sets the connectionType of the connection.
     */
    public readonly connectionType!: pulumi.Output<outputs.automation.latest.ConnectionTypeAssociationPropertyResponse | undefined>;
    /**
     * Gets the creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Gets or sets the description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Gets the field definition values of the connection.
     */
    public readonly fieldDefinitionValues!: pulumi.Output<{[key: string]: string}>;
    /**
     * Gets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.automationAccountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.connectionName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.connectionType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'connectionType'");
            }
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["connectionName"] = args ? args.connectionName : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fieldDefinitionValues"] = args ? args.fieldDefinitionValues : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["connectionType"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["fieldDefinitionValues"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:automation/v20151031:Connection" }, { type: "azure-nextgen:automation/v20190601:Connection" }, { type: "azure-nextgen:automation/v20200113preview:Connection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Connection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The parameters supplied to the create or update connection operation.
     */
    readonly connectionName: pulumi.Input<string>;
    /**
     * Gets or sets the connectionType of the connection.
     */
    readonly connectionType: pulumi.Input<inputs.automation.latest.ConnectionTypeAssociationProperty>;
    /**
     * Gets or sets the description of the connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Gets or sets the field definition properties of the connection.
     */
    readonly fieldDefinitionValues?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Gets or sets the name of the connection.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
