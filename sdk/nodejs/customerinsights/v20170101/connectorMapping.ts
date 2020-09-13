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
        return new ConnectorMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/v20170101:ConnectorMapping';

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
     * The connector mapping name
     */
    public /*out*/ readonly connectorMappingName!: pulumi.Output<string>;
    /**
     * The connector name.
     */
    public readonly connectorName!: pulumi.Output<string>;
    /**
     * Type of connector.
     */
    public readonly connectorType!: pulumi.Output<string | undefined>;
    /**
     * The created time.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The DataFormat ID.
     */
    public /*out*/ readonly dataFormatId!: pulumi.Output<string>;
    /**
     * The description of the connector mapping.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Display name for the connector mapping.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Defines which entity type the file should map to.
     */
    public readonly entityType!: pulumi.Output<string>;
    /**
     * The mapping entity name.
     */
    public readonly entityTypeName!: pulumi.Output<string>;
    /**
     * The last modified time.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The properties of the mapping.
     */
    public readonly mappingProperties!: pulumi.Output<outputs.customerinsights.v20170101.ConnectorMappingPropertiesResponse>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The next run time based on customer's settings.
     */
    public /*out*/ readonly nextRunTime!: pulumi.Output<string>;
    /**
     * The RunId.
     */
    public /*out*/ readonly runId!: pulumi.Output<string>;
    /**
     * State of connector mapping.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
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
    constructor(name: string, args: ConnectorMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
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
            if (!args || args.mappingName === undefined) {
                throw new Error("Missing required property 'mappingName'");
            }
            if (!args || args.mappingProperties === undefined) {
                throw new Error("Missing required property 'mappingProperties'");
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
            inputs["mappingName"] = args ? args.mappingName : undefined;
            inputs["mappingProperties"] = args ? args.mappingProperties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["connectorMappingName"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["dataFormatId"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nextRunTime"] = undefined /*out*/;
            inputs["runId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["connectorMappingName"] = undefined /*out*/;
            inputs["connectorName"] = undefined /*out*/;
            inputs["connectorType"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["dataFormatId"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["entityType"] = undefined /*out*/;
            inputs["entityTypeName"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["mappingProperties"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nextRunTime"] = undefined /*out*/;
            inputs["runId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/latest:ConnectorMapping" }, { type: "azurerm:customerinsights/v20170426:ConnectorMapping" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
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
     * The name of the connector mapping.
     */
    readonly mappingName: pulumi.Input<string>;
    /**
     * The properties of the mapping.
     */
    readonly mappingProperties: pulumi.Input<inputs.customerinsights.v20170101.ConnectorMappingProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
