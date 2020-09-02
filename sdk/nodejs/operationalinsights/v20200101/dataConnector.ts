// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Data connector.
 *
 * ## Example Usage
 * ### Creates or updates an Office365 data connector.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataConnector = new azurerm.operationalinsights.v20200101.DataConnector("dataConnector", {
 *     dataConnectorId: "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
 *     etag: "\"0300bf09-0000-0000-0000-5c37296e0000\"",
 *     kind: "Office365",
 *     resourceGroupName: "myRg",
 *     workspaceName: "myWorkspace",
 * });
 *
 * ```
 */
export class DataConnector extends pulumi.CustomResource {
    /**
     * Get an existing DataConnector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataConnector {
        return new DataConnector(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/v20200101:DataConnector';

    /**
     * Returns true if the given object is an instance of DataConnector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataConnector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataConnector.__pulumiType;
    }

    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The data connector kind
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataConnector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataConnectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DataConnectorArgs | undefined;
            if (!args || args.dataConnectorId === undefined) {
                throw new Error("Missing required property 'dataConnectorId'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["dataConnectorId"] = args ? args.dataConnectorId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/latest:DataConnector" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DataConnector.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataConnector resource.
 */
export interface DataConnectorArgs {
    /**
     * Connector ID
     */
    readonly dataConnectorId: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The data connector kind
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
