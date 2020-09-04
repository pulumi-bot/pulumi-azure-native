// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A report resource.
 */
export class Report extends pulumi.CustomResource {
    /**
     * Get an existing Report resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Report {
        return new Report(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:costmanagement/v20180801preview:Report';

    /**
     * Returns true if the given object is an instance of Report.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Report {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Report.__pulumiType;
    }

    /**
     * Has definition for the report.
     */
    public readonly definition!: pulumi.Output<outputs.costmanagement.v20180801preview.ReportDefinitionResponse>;
    /**
     * Has delivery information for the report.
     */
    public readonly deliveryInfo!: pulumi.Output<outputs.costmanagement.v20180801preview.ReportDeliveryInfoResponse>;
    /**
     * The format of the report being delivered.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Has schedule information for the report.
     */
    public readonly schedule!: pulumi.Output<outputs.costmanagement.v20180801preview.ReportScheduleResponse | undefined>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Report resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.definition === undefined) {
                throw new Error("Missing required property 'definition'");
            }
            if (!args || args.deliveryInfo === undefined) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if (!args || args.reportName === undefined) {
                throw new Error("Missing required property 'reportName'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["reportName"] = args ? args.reportName : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["definition"] = undefined /*out*/;
            inputs["deliveryInfo"] = undefined /*out*/;
            inputs["format"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Report.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Report resource.
 */
export interface ReportArgs {
    /**
     * Has definition for the report.
     */
    readonly definition: pulumi.Input<inputs.costmanagement.v20180801preview.ReportDefinition>;
    /**
     * Has delivery information for the report.
     */
    readonly deliveryInfo: pulumi.Input<inputs.costmanagement.v20180801preview.ReportDeliveryInfo>;
    /**
     * The format of the report being delivered.
     */
    readonly format?: pulumi.Input<string>;
    /**
     * Report Name.
     */
    readonly reportName: pulumi.Input<string>;
    /**
     * Has schedule information for the report.
     */
    readonly schedule?: pulumi.Input<inputs.costmanagement.v20180801preview.ReportSchedule>;
}
