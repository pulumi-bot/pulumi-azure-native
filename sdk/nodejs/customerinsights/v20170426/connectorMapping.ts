// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The connector mapping resource format.
 */
export class ConnectorMapping extends pulumi.CustomResource {
    /**
     * Get an existing ConnectorMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectorMapping {
        return new ConnectorMapping(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/v20170426:ConnectorMapping';

    /**
     * Returns true if the given object is an instance of ConnectorMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectorMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectorMapping.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The connector mapping definition.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.customerinsights.v20170426.ConnectorMappingResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ConnectorMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectorMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ConnectorMappingArgs | undefined;
            if (!args || args.connectorName === undefined) {
                throw new Error("Missing required property 'connectorName'");
            }
            if (!args || args.entityType === undefined) {
                throw new Error("Missing required property 'entityType'");
            }
            if (!args || args.entityTypeName === undefined) {
                throw new Error("Missing required property 'entityTypeName'");
            }
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.mappingProperties === undefined) {
                throw new Error("Missing required property 'mappingProperties'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["connectorName"] = args ? args.connectorName : undefined;
            inputs["connectorType"] = args ? args.connectorType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["entityType"] = args ? args.entityType : undefined;
            inputs["entityTypeName"] = args ? args.entityTypeName : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["mappingProperties"] = args ? args.mappingProperties : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConnectorMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectorMapping resource.
 */
export interface ConnectorMappingArgs {
    /**
     * The name of the connector.
     */
    readonly connectorName: pulumi.Input<string>;
    /**
     * Type of connector.
     */
    readonly connectorType?: pulumi.Input<string>;
    /**
     * The description of the connector mapping.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display name for the connector mapping.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Defines which entity type the file should map to.
     */
    readonly entityType: pulumi.Input<string>;
    /**
     * The mapping entity name.
     */
    readonly entityTypeName: pulumi.Input<string>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * The properties of the mapping.
     */
    readonly mappingProperties: pulumi.Input<inputs.customerinsights.v20170426.ConnectorMappingProperties>;
    /**
     * The name of the connector mapping.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
