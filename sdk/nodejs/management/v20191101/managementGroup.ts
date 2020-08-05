// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The management group details.
 */
export class ManagementGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagementGroup {
        return new ManagementGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:management/v20191101:ManagementGroup';

    /**
     * Returns true if the given object is an instance of ManagementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementGroup.__pulumiType;
    }

    /**
     * The name of the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The generic properties of a management group.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.management.v20191101.ManagementGroupPropertiesResponse>;
    /**
     * The type of the resource.  For example, Microsoft.Management/managementGroups
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagementGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ManagementGroupArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["details"] = args ? args.details : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementGroup resource.
 */
export interface ManagementGroupArgs {
    /**
     * The details of a management group used during creation.
     */
    readonly details?: pulumi.Input<inputs.management.v20191101.CreateManagementGroupDetails>;
    /**
     * The friendly name of the management group. If no value is passed then this  field will be set to the groupId.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Management Group ID.
     */
    readonly name: pulumi.Input<string>;
}
