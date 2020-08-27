// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An server Active Directory Administrator.
 */
export class ServerAzureADAdministrator extends pulumi.CustomResource {
    /**
     * Get an existing ServerAzureADAdministrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerAzureADAdministrator {
        return new ServerAzureADAdministrator(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20140401:ServerAzureADAdministrator';

    /**
     * Returns true if the given object is an instance of ServerAzureADAdministrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerAzureADAdministrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerAzureADAdministrator.__pulumiType;
    }

    /**
     * The type of administrator.
     */
    public readonly administratorType!: pulumi.Output<AdministratorType>;
    /**
     * The server administrator login value.
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The server administrator Sid (Secure ID).
     */
    public readonly sid!: pulumi.Output<string>;
    /**
     * The server Active Directory Administrator tenant id.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ServerAzureADAdministrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerAzureADAdministratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerAzureADAdministratorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServerAzureADAdministratorArgs | undefined;
            if (!args || args.administratorName === undefined) {
                throw new Error("Missing required property 'administratorName'");
            }
            if (!args || args.administratorType === undefined) {
                throw new Error("Missing required property 'administratorType'");
            }
            if (!args || args.login === undefined) {
                throw new Error("Missing required property 'login'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.sid === undefined) {
                throw new Error("Missing required property 'sid'");
            }
            if (!args || args.tenantId === undefined) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["administratorName"] = args ? args.administratorName : undefined;
            inputs["administratorType"] = args ? args.administratorType : undefined;
            inputs["login"] = args ? args.login : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sid"] = args ? args.sid : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerAzureADAdministrator.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerAzureADAdministrator resource.
 */
export interface ServerAzureADAdministratorArgs {
    /**
     * Name of the server administrator resource.
     */
    readonly administratorName: pulumi.Input<string>;
    /**
     * The type of administrator.
     */
    readonly administratorType: pulumi.Input<AdministratorType>;
    /**
     * The server administrator login value.
     */
    readonly login: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The server administrator Sid (Secure ID).
     */
    readonly sid: pulumi.Input<string>;
    /**
     * The server Active Directory Administrator tenant id.
     */
    readonly tenantId: pulumi.Input<string>;
}
