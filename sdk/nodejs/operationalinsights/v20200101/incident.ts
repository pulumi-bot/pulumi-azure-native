// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents an incident in Azure Security Insights.
 *
 * ## Example Usage
 * ### Creates or updates an incident.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const incident = new azurerm.operationalinsights.v20200101.Incident("incident", {
 *     classification: "FalsePositive",
 *     classificationComment: "Not a malicious activity",
 *     classificationReason: "IncorrectAlertLogic",
 *     description: "This is a demo incident",
 *     etag: "\"0300bf09-0000-0000-0000-5c37296e0000\"",
 *     firstActivityTimeUtc: "2019-01-01T13:00:30Z",
 *     incidentId: "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
 *     lastActivityTimeUtc: "2019-01-01T13:05:30Z",
 *     owner: {
 *         objectId: "2046feea-040d-4a46-9e2b-91c2941bfa70",
 *     },
 *     resourceGroupName: "myRg",
 *     severity: "High",
 *     status: "Closed",
 *     title: "My incident",
 *     workspaceName: "myWorkspace",
 * });
 *
 * ```
 */
export class Incident extends pulumi.CustomResource {
    /**
     * Get an existing Incident resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Incident {
        return new Incident(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/v20200101:Incident';

    /**
     * Returns true if the given object is an instance of Incident.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Incident {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Incident.__pulumiType;
    }

    /**
     * Additional data on the incident
     */
    public /*out*/ readonly additionalData!: pulumi.Output<outputs.operationalinsights.v20200101.IncidentAdditionalDataResponse>;
    /**
     * The reason the incident was closed
     */
    public readonly classification!: pulumi.Output<string | undefined>;
    /**
     * Describes the reason the incident was closed
     */
    public readonly classificationComment!: pulumi.Output<string | undefined>;
    /**
     * The classification reason the incident was closed with
     */
    public readonly classificationReason!: pulumi.Output<string | undefined>;
    /**
     * The time the incident was created
     */
    public /*out*/ readonly createdTimeUtc!: pulumi.Output<string>;
    /**
     * The description of the incident
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The time of the first activity in the incident
     */
    public readonly firstActivityTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * A sequential number
     */
    public /*out*/ readonly incidentNumber!: pulumi.Output<number>;
    /**
     * The deep-link url to the incident in Azure portal
     */
    public /*out*/ readonly incidentUrl!: pulumi.Output<string>;
    /**
     * List of labels relevant to this incident
     */
    public readonly labels!: pulumi.Output<outputs.operationalinsights.v20200101.IncidentLabelResponse[] | undefined>;
    /**
     * The time of the last activity in the incident
     */
    public readonly lastActivityTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * The last time the incident was updated
     */
    public /*out*/ readonly lastModifiedTimeUtc!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Describes a user that the incident is assigned to
     */
    public readonly owner!: pulumi.Output<outputs.operationalinsights.v20200101.IncidentOwnerInfoResponse | undefined>;
    /**
     * List of resource ids of Analytic rules related to the incident
     */
    public /*out*/ readonly relatedAnalyticRuleIds!: pulumi.Output<string[]>;
    /**
     * The severity of the incident
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * The status of the incident
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The title of the incident
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Incident resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IncidentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IncidentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IncidentArgs | undefined;
            if (!args || args.incidentId === undefined) {
                throw new Error("Missing required property 'incidentId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.severity === undefined) {
                throw new Error("Missing required property 'severity'");
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
            inputs["classification"] = args ? args.classification : undefined;
            inputs["classificationComment"] = args ? args.classificationComment : undefined;
            inputs["classificationReason"] = args ? args.classificationReason : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["firstActivityTimeUtc"] = args ? args.firstActivityTimeUtc : undefined;
            inputs["incidentId"] = args ? args.incidentId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["lastActivityTimeUtc"] = args ? args.lastActivityTimeUtc : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["severity"] = args ? args.severity : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["additionalData"] = undefined /*out*/;
            inputs["createdTimeUtc"] = undefined /*out*/;
            inputs["incidentNumber"] = undefined /*out*/;
            inputs["incidentUrl"] = undefined /*out*/;
            inputs["lastModifiedTimeUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["relatedAnalyticRuleIds"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/latest:Incident" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Incident.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Incident resource.
 */
export interface IncidentArgs {
    /**
     * The reason the incident was closed
     */
    readonly classification?: pulumi.Input<string>;
    /**
     * Describes the reason the incident was closed
     */
    readonly classificationComment?: pulumi.Input<string>;
    /**
     * The classification reason the incident was closed with
     */
    readonly classificationReason?: pulumi.Input<string>;
    /**
     * The description of the incident
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The time of the first activity in the incident
     */
    readonly firstActivityTimeUtc?: pulumi.Input<string>;
    /**
     * Incident ID
     */
    readonly incidentId: pulumi.Input<string>;
    /**
     * List of labels relevant to this incident
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.operationalinsights.v20200101.IncidentLabel>[]>;
    /**
     * The time of the last activity in the incident
     */
    readonly lastActivityTimeUtc?: pulumi.Input<string>;
    /**
     * Describes a user that the incident is assigned to
     */
    readonly owner?: pulumi.Input<inputs.operationalinsights.v20200101.IncidentOwnerInfo>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The severity of the incident
     */
    readonly severity: pulumi.Input<string>;
    /**
     * The status of the incident
     */
    readonly status: pulumi.Input<string>;
    /**
     * The title of the incident
     */
    readonly title: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
