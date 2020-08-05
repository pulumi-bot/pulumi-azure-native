// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A report config resource.
 */
export class ReportConfigByResourceGroupName extends pulumi.CustomResource {
    /**
     * Get an existing ReportConfigByResourceGroupName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReportConfigByResourceGroupName {
        return new ReportConfigByResourceGroupName(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:costmanagement/v20180531:ReportConfigByResourceGroupName';

    /**
     * Returns true if the given object is an instance of ReportConfigByResourceGroupName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportConfigByResourceGroupName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportConfigByResourceGroupName.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the report config.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.costmanagement.v20180531.ReportConfigPropertiesResponse>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReportConfigByResourceGroupName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportConfigByResourceGroupNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReportConfigByResourceGroupNameArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ReportConfigByResourceGroupNameArgs | undefined;
            if (!args || args.definition === undefined) {
                throw new Error("Missing required property 'definition'");
            }
            if (!args || args.deliveryInfo === undefined) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReportConfigByResourceGroupName.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReportConfigByResourceGroupName resource.
 */
export interface ReportConfigByResourceGroupNameArgs {
    /**
     * Has definition for the report config.
     */
    readonly definition: pulumi.Input<inputs.costmanagement.v20180531.ReportConfigDefinition>;
    /**
     * Has delivery information for the report config.
     */
    readonly deliveryInfo: pulumi.Input<inputs.costmanagement.v20180531.ReportConfigDeliveryInfo>;
    /**
     * The format of the report being delivered.
     */
    readonly format?: pulumi.Input<string>;
    /**
     * Report Config Name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Azure Resource Group Name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Has schedule information for the report config.
     */
    readonly schedule?: pulumi.Input<inputs.costmanagement.v20180531.ReportConfigSchedule>;
}
