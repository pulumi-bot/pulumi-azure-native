// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
 */
export class PrivateEndpointConnection extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateEndpointConnection {
        return new PrivateEndpointConnection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:search/latest:PrivateEndpointConnection';

    /**
     * Returns true if the given object is an instance of PrivateEndpointConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointConnection.__pulumiType;
    }

    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
     */
    public readonly properties!: pulumi.Output<outputs.search.latest.PrivateEndpointConnectionPropertiesResponse>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PrivateEndpointConnectionArgs | undefined;
            if (!args || args.privateEndpointConnectionName === undefined) {
                throw new Error("Missing required property 'privateEndpointConnectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.searchServiceName === undefined) {
                throw new Error("Missing required property 'searchServiceName'");
            }
            inputs["privateEndpointConnectionName"] = args ? args.privateEndpointConnectionName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["searchServiceName"] = args ? args.searchServiceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:search/v20191001preview:PrivateEndpointConnection" }, { type: "azurerm:search/v20200313:PrivateEndpointConnection" }, { type: "azurerm:search/v20200801:PrivateEndpointConnection" }, { type: "azurerm:search/v20200801preview:PrivateEndpointConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PrivateEndpointConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateEndpointConnection resource.
 */
export interface PrivateEndpointConnectionArgs {
    /**
     * The name of the private endpoint connection to the Azure Cognitive Search service with the specified resource group.
     */
    readonly privateEndpointConnectionName: pulumi.Input<string>;
    /**
     * Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
     */
    readonly properties?: pulumi.Input<inputs.search.latest.PrivateEndpointConnectionProperties>;
    /**
     * The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Azure Cognitive Search service associated with the specified resource group.
     */
    readonly searchServiceName: pulumi.Input<string>;
}
