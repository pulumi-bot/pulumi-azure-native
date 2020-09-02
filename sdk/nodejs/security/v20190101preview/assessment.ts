// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Security assessment on a resource
 */
export class Assessment extends pulumi.CustomResource {
    /**
     * Get an existing Assessment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Assessment {
        return new Assessment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20190101preview:Assessment';

    /**
     * Returns true if the given object is an instance of Assessment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Assessment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Assessment.__pulumiType;
    }

    /**
     * Additional data regarding the assessment
     */
    public readonly additionalData!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * User friendly display name of the assessment
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * Links relevant to the assessment
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.security.v20190101preview.AssessmentLinksResponse | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Details of the resource that was assessed
     */
    public readonly resourceDetails!: pulumi.Output<outputs.security.v20190101preview.ResourceDetailsResponse>;
    /**
     * The result of the assessment
     */
    public readonly status!: pulumi.Output<outputs.security.v20190101preview.AssessmentStatusResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Assessment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssessmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssessmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AssessmentArgs | undefined;
            if (!args || args.assessmentName === undefined) {
                throw new Error("Missing required property 'assessmentName'");
            }
            if (!args || args.resourceDetails === undefined) {
                throw new Error("Missing required property 'resourceDetails'");
            }
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.status === undefined) {
                throw new Error("Missing required property 'status'");
            }
            inputs["additionalData"] = args ? args.additionalData : undefined;
            inputs["assessmentName"] = args ? args.assessmentName : undefined;
            inputs["resourceDetails"] = args ? args.resourceDetails : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["displayName"] = undefined /*out*/;
            inputs["links"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:security/latest:Assessment" }, { type: "azurerm:security/v20200101:Assessment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Assessment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Assessment resource.
 */
export interface AssessmentArgs {
    /**
     * Additional data regarding the assessment
     */
    readonly additionalData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Assessment Key - Unique key for the assessment type
     */
    readonly assessmentName: pulumi.Input<string>;
    /**
     * Details of the resource that was assessed
     */
    readonly resourceDetails: pulumi.Input<inputs.security.v20190101preview.ResourceDetails>;
    /**
     * The identifier of the resource.
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The result of the assessment
     */
    readonly status: pulumi.Input<inputs.security.v20190101preview.AssessmentStatus>;
}
