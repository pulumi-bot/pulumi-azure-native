// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Azure Active Directory only authentication.
 */
export class ServerAzureADOnlyAuthentication extends pulumi.CustomResource {
    /**
     * Get an existing ServerAzureADOnlyAuthentication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerAzureADOnlyAuthentication {
        return new ServerAzureADOnlyAuthentication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20200202preview:ServerAzureADOnlyAuthentication';

    /**
     * Returns true if the given object is an instance of ServerAzureADOnlyAuthentication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerAzureADOnlyAuthentication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerAzureADOnlyAuthentication.__pulumiType;
    }

    /**
     * Azure Active Directory only Authentication enabled.
     */
    public readonly azureADOnlyAuthentication!: pulumi.Output<boolean>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ServerAzureADOnlyAuthentication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerAzureADOnlyAuthenticationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.authenticationName === undefined) {
                throw new Error("Missing required property 'authenticationName'");
            }
            if (!args || args.azureADOnlyAuthentication === undefined) {
                throw new Error("Missing required property 'azureADOnlyAuthentication'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["authenticationName"] = args ? args.authenticationName : undefined;
            inputs["azureADOnlyAuthentication"] = args ? args.azureADOnlyAuthentication : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["azureADOnlyAuthentication"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerAzureADOnlyAuthentication.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerAzureADOnlyAuthentication resource.
 */
export interface ServerAzureADOnlyAuthenticationArgs {
    /**
     * The name of server azure active directory only authentication.
     */
    readonly authenticationName: pulumi.Input<string>;
    /**
     * Azure Active Directory only Authentication enabled.
     */
    readonly azureADOnlyAuthentication: pulumi.Input<boolean>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
}
