// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a disaster recovery configuration.
 */
export class ServerDisasterRecoveryConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing ServerDisasterRecoveryConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerDisasterRecoveryConfiguration {
        return new ServerDisasterRecoveryConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql:ServerDisasterRecoveryConfiguration';

    /**
     * Returns true if the given object is an instance of ServerDisasterRecoveryConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerDisasterRecoveryConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerDisasterRecoveryConfiguration.__pulumiType;
    }

    /**
     * Location of the server that contains this disaster recovery configuration.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties representing the resource.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.sql.DisasterRecoveryConfigurationPropertiesResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ServerDisasterRecoveryConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServerDisasterRecoveryConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
        inputs["name"] = args ? args.name : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["serverName"] = args ? args.serverName : undefined;
        inputs["location"] = undefined /*out*/;
        inputs["properties"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerDisasterRecoveryConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerDisasterRecoveryConfiguration resource.
 */
export interface ServerDisasterRecoveryConfigurationArgs {
    /**
     * The name of the disaster recovery configuration to be created/updated.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
}
