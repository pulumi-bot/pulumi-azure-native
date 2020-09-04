// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a case in Azure Security Insights.
 */
export class Case extends pulumi.CustomResource {
    /**
     * Get an existing Case resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Case {
        return new Case(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:securityinsights/v20190101preview:Case';

    /**
     * Returns true if the given object is an instance of Case.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Case {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Case.__pulumiType;
    }

    /**
     * a sequential number
     */
    public /*out*/ readonly caseNumber!: pulumi.Output<number>;
    /**
     * The reason the case was closed
     */
    public readonly closeReason!: pulumi.Output<string | undefined>;
    /**
     * the case close reason details
     */
    public readonly closedReasonText!: pulumi.Output<string | undefined>;
    /**
     * The time the case was created
     */
    public /*out*/ readonly createdTimeUtc!: pulumi.Output<string>;
    /**
     * The description of the case
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The end time of the case
     */
    public readonly endTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * List of labels relevant to this case
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * the last comment in the case
     */
    public /*out*/ readonly lastComment!: pulumi.Output<string>;
    /**
     * The last time the case was updated
     */
    public /*out*/ readonly lastUpdatedTimeUtc!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Describes a user that the case is assigned to
     */
    public readonly owner!: pulumi.Output<outputs.securityinsights.v20190101preview.UserInfoResponse | undefined>;
    /**
     * List of related alert identifiers
     */
    public /*out*/ readonly relatedAlertIds!: pulumi.Output<string[]>;
    /**
     * The severity of the case
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * The start time of the case
     */
    public readonly startTimeUtc!: pulumi.Output<string>;
    /**
     * The status of the case
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The tactics associated with case
     */
    public /*out*/ readonly tactics!: pulumi.Output<string[]>;
    /**
     * The title of the case
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * the number of total comments in the case
     */
    public /*out*/ readonly totalComments!: pulumi.Output<number>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Case resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaseArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.caseId === undefined) {
                throw new Error("Missing required property 'caseId'");
            }
            if (!args || args.operationalInsightsResourceProvider === undefined) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.severity === undefined) {
                throw new Error("Missing required property 'severity'");
            }
            if (!args || args.startTimeUtc === undefined) {
                throw new Error("Missing required property 'startTimeUtc'");
            }
            if (!args || args.status === undefined) {
                throw new Error("Missing required property 'status'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["caseId"] = args ? args.caseId : undefined;
            inputs["closeReason"] = args ? args.closeReason : undefined;
            inputs["closedReasonText"] = args ? args.closedReasonText : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["endTimeUtc"] = args ? args.endTimeUtc : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["severity"] = args ? args.severity : undefined;
            inputs["startTimeUtc"] = args ? args.startTimeUtc : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["caseNumber"] = undefined /*out*/;
            inputs["createdTimeUtc"] = undefined /*out*/;
            inputs["lastComment"] = undefined /*out*/;
            inputs["lastUpdatedTimeUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["relatedAlertIds"] = undefined /*out*/;
            inputs["tactics"] = undefined /*out*/;
            inputs["totalComments"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["caseNumber"] = undefined /*out*/;
            inputs["closeReason"] = undefined /*out*/;
            inputs["closedReasonText"] = undefined /*out*/;
            inputs["createdTimeUtc"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["endTimeUtc"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["lastComment"] = undefined /*out*/;
            inputs["lastUpdatedTimeUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["owner"] = undefined /*out*/;
            inputs["relatedAlertIds"] = undefined /*out*/;
            inputs["severity"] = undefined /*out*/;
            inputs["startTimeUtc"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tactics"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["totalComments"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Case.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Case resource.
 */
export interface CaseArgs {
    /**
     * Case ID
     */
    readonly caseId: pulumi.Input<string>;
    /**
     * The reason the case was closed
     */
    readonly closeReason?: pulumi.Input<string>;
    /**
     * the case close reason details
     */
    readonly closedReasonText?: pulumi.Input<string>;
    /**
     * The description of the case
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The end time of the case
     */
    readonly endTimeUtc?: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * List of labels relevant to this case
     */
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * Describes a user that the case is assigned to
     */
    readonly owner?: pulumi.Input<inputs.securityinsights.v20190101preview.UserInfo>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The severity of the case
     */
    readonly severity: pulumi.Input<string>;
    /**
     * The start time of the case
     */
    readonly startTimeUtc: pulumi.Input<string>;
    /**
     * The status of the case
     */
    readonly status: pulumi.Input<string>;
    /**
     * The title of the case
     */
    readonly title: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
