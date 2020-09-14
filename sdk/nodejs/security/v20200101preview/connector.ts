// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The connector setting
 */
export class Connector extends pulumi.CustomResource {
    /**
     * Get an existing Connector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connector {
        return new Connector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20200101preview:Connector';

    /**
     * Returns true if the given object is an instance of Connector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connector.__pulumiType;
    }

    /**
     * Settings for authentication management, these settings are relevant only for the cloud connector.
     */
    public readonly authenticationDetails!: pulumi.Output<outputs.security.v20200101preview.AwAssumeRoleAuthenticationDetailsPropertiesResponse | outputs.security.v20200101preview.AwsCredsAuthenticationDetailsPropertiesResponse | outputs.security.v20200101preview.GcpCredentialsDetailsPropertiesResponse | undefined>;
    /**
     * Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
     */
    public readonly hybridComputeSettings!: pulumi.Output<outputs.security.v20200101preview.HybridComputeSettingsPropertiesResponse | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Connector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.connectorName === undefined) {
                throw new Error("Missing required property 'connectorName'");
            }
            inputs["authenticationDetails"] = args ? args.authenticationDetails : undefined;
            inputs["connectorName"] = args ? args.connectorName : undefined;
            inputs["hybridComputeSettings"] = args ? args.hybridComputeSettings : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["authenticationDetails"] = undefined /*out*/;
            inputs["hybridComputeSettings"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:security/latest:Connector" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Connector.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connector resource.
 */
export interface ConnectorArgs {
    /**
     * Settings for authentication management, these settings are relevant only for the cloud connector.
     */
    readonly authenticationDetails?: pulumi.Input<inputs.security.v20200101preview.AwAssumeRoleAuthenticationDetailsProperties | inputs.security.v20200101preview.AwsCredsAuthenticationDetailsProperties | inputs.security.v20200101preview.GcpCredentialsDetailsProperties>;
    /**
     * Name of the cloud account connector
     */
    readonly connectorName: pulumi.Input<string>;
    /**
     * Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
     */
    readonly hybridComputeSettings?: pulumi.Input<inputs.security.v20200101preview.HybridComputeSettingsProperties>;
}
