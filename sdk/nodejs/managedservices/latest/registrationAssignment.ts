// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Registration assignment.
 *
 * ## Put Registration Assignment
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const registrationAssignment = new azurerm.managedservices.latest.RegistrationAssignment("registrationAssignment", {
 *     registrationAssignmentId: "26c128c2-fefa-4340-9bb1-6e081c90ada2",
 *     scope: "subscription/0afefe50-734e-4610-8a82-a144ahf49dea",
 * });
 *
 * ```
 */
export class RegistrationAssignment extends pulumi.CustomResource {
    /**
     * Get an existing RegistrationAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegistrationAssignment {
        return new RegistrationAssignment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:managedservices/latest:RegistrationAssignment';

    /**
     * Returns true if the given object is an instance of RegistrationAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistrationAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistrationAssignment.__pulumiType;
    }

    /**
     * Name of the registration assignment.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of a registration assignment.
     */
    public readonly properties!: pulumi.Output<outputs.managedservices.latest.RegistrationAssignmentPropertiesResponse>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RegistrationAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistrationAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RegistrationAssignmentArgs | undefined;
            if (!args || args.registrationAssignmentId === undefined) {
                throw new Error("Missing required property 'registrationAssignmentId'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["properties"] = args ? args.properties : undefined;
            inputs["registrationAssignmentId"] = args ? args.registrationAssignmentId : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:managedservices/v20190601:RegistrationAssignment" }, { type: "azurerm:managedservices/v20190901:RegistrationAssignment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RegistrationAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegistrationAssignment resource.
 */
export interface RegistrationAssignmentArgs {
    /**
     * Properties of a registration assignment.
     */
    readonly properties?: pulumi.Input<inputs.managedservices.latest.RegistrationAssignmentProperties>;
    /**
     * Guid of the registration assignment.
     */
    readonly registrationAssignmentId: pulumi.Input<string>;
    /**
     * Scope of the resource.
     */
    readonly scope: pulumi.Input<string>;
}