// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Registration definition.
 */
export class RegistrationDefinition extends pulumi.CustomResource {
    /**
     * Get an existing RegistrationDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegistrationDefinition {
        return new RegistrationDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:managedservices/v20180601preview:RegistrationDefinition';

    /**
     * Returns true if the given object is an instance of RegistrationDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistrationDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistrationDefinition.__pulumiType;
    }

    /**
     * Name of the registration definition.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Plan details for the managed services.
     */
    public readonly plan!: pulumi.Output<outputs.managedservices.v20180601preview.PlanResponse | undefined>;
    /**
     * Properties of a registration definition.
     */
    public readonly properties!: pulumi.Output<outputs.managedservices.v20180601preview.RegistrationDefinitionPropertiesResponse>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RegistrationDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.registrationDefinitionId === undefined) {
                throw new Error("Missing required property 'registrationDefinitionId'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["plan"] = args ? args.plan : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["registrationDefinitionId"] = args ? args.registrationDefinitionId : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["plan"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:managedservices/latest:RegistrationDefinition" }, { type: "azurerm:managedservices/v20190401preview:RegistrationDefinition" }, { type: "azurerm:managedservices/v20190601:RegistrationDefinition" }, { type: "azurerm:managedservices/v20190901:RegistrationDefinition" }, { type: "azurerm:managedservices/v20200201preview:RegistrationDefinition" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RegistrationDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegistrationDefinition resource.
 */
export interface RegistrationDefinitionArgs {
    /**
     * Plan details for the managed services.
     */
    readonly plan?: pulumi.Input<inputs.managedservices.v20180601preview.Plan>;
    /**
     * Properties of a registration definition.
     */
    readonly properties?: pulumi.Input<inputs.managedservices.v20180601preview.RegistrationDefinitionProperties>;
    /**
     * Guid of the registration definition.
     */
    readonly registrationDefinitionId: pulumi.Input<string>;
    /**
     * Scope of the resource.
     */
    readonly scope: pulumi.Input<string>;
}
