// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

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
    public static readonly __pulumiType = 'azurerm:eventgrid/latest:PrivateEndpointConnection';

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
     * GroupIds from the private link service resource.
     */
    public readonly groupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Private Endpoint resource for this Connection.
     */
    public readonly privateEndpoint!: pulumi.Output<outputs.eventgrid.latest.PrivateEndpointResponse | undefined>;
    /**
     * Details about the state of the connection.
     */
    public readonly privateLinkServiceConnectionState!: pulumi.Output<outputs.eventgrid.latest.ConnectionStateResponse | undefined>;
    /**
     * Provisioning state of the Private Endpoint Connection.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Type of the resource.
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
            if (!args || args.parentName === undefined) {
                throw new Error("Missing required property 'parentName'");
            }
            if (!args || args.parentType === undefined) {
                throw new Error("Missing required property 'parentType'");
            }
            if (!args || args.privateEndpointConnectionName === undefined) {
                throw new Error("Missing required property 'privateEndpointConnectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["groupIds"] = args ? args.groupIds : undefined;
            inputs["parentName"] = args ? args.parentName : undefined;
            inputs["parentType"] = args ? args.parentType : undefined;
            inputs["privateEndpoint"] = args ? args.privateEndpoint : undefined;
            inputs["privateEndpointConnectionName"] = args ? args.privateEndpointConnectionName : undefined;
            inputs["privateLinkServiceConnectionState"] = args ? args.privateLinkServiceConnectionState : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:eventgrid/v20200401preview:PrivateEndpointConnection" }, { type: "azurerm:eventgrid/v20200601:PrivateEndpointConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PrivateEndpointConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateEndpointConnection resource.
 */
export interface PrivateEndpointConnectionArgs {
    /**
     * GroupIds from the private link service resource.
     */
    readonly groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the parent resource (namely, either, the topic name or domain name).
     */
    readonly parentName: pulumi.Input<string>;
    /**
     * The type of the parent resource. This can be either \'topics\' or \'domains\'.
     */
    readonly parentType: pulumi.Input<string>;
    /**
     * The Private Endpoint resource for this Connection.
     */
    readonly privateEndpoint?: pulumi.Input<inputs.eventgrid.latest.PrivateEndpoint>;
    /**
     * The name of the private endpoint connection connection.
     */
    readonly privateEndpointConnectionName: pulumi.Input<string>;
    /**
     * Details about the state of the connection.
     */
    readonly privateLinkServiceConnectionState?: pulumi.Input<inputs.eventgrid.latest.ConnectionState>;
    /**
     * Provisioning state of the Private Endpoint Connection.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
