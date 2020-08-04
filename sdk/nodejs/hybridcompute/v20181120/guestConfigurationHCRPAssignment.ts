// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Guest configuration assignment is an association between a machine and guest configuration.
 */
export class GuestConfigurationHCRPAssignment extends pulumi.CustomResource {
    /**
     * Get an existing GuestConfigurationHCRPAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GuestConfigurationHCRPAssignment {
        return new GuestConfigurationHCRPAssignment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hybridcompute/v20181120:GuestConfigurationHCRPAssignment';

    /**
     * Returns true if the given object is an instance of GuestConfigurationHCRPAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GuestConfigurationHCRPAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GuestConfigurationHCRPAssignment.__pulumiType;
    }

    /**
     * Region where the VM is located.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of the guest configuration assignment.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Properties of the Guest configuration assignment.
     */
    public readonly properties!: pulumi.Output<outputs.hybridcompute.v20181120.GuestConfigurationAssignmentPropertiesResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GuestConfigurationHCRPAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GuestConfigurationHCRPAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GuestConfigurationHCRPAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as GuestConfigurationHCRPAssignmentArgs | undefined;
            if (!args || args.machineName === undefined) {
                throw new Error("Missing required property 'machineName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["machineName"] = args ? args.machineName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GuestConfigurationHCRPAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GuestConfigurationHCRPAssignment resource.
 */
export interface GuestConfigurationHCRPAssignmentArgs {
    /**
     * Region where the VM is located.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the ARC machine.
     */
    readonly machineName: pulumi.Input<string>;
    /**
     * Name of the guest configuration assignment.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Properties of the Guest configuration assignment.
     */
    readonly properties?: pulumi.Input<inputs.hybridcompute.v20181120.GuestConfigurationAssignmentProperties>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
