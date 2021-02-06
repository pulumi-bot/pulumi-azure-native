// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Definition of ARM tracked top level resource.
 */
export class DataCollectionRule extends pulumi.CustomResource {
    /**
     * Get an existing DataCollectionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataCollectionRule {
        return new DataCollectionRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:insights/v20191101preview:DataCollectionRule';

    /**
     * Returns true if the given object is an instance of DataCollectionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataCollectionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataCollectionRule.__pulumiType;
    }

    /**
     * The specification of data flows.
     */
    public readonly dataFlows!: pulumi.Output<outputs.insights.v20191101preview.DataFlowResponse[]>;
    /**
     * The specification of data sources. 
     * This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
     */
    public readonly dataSources!: pulumi.Output<outputs.insights.v20191101preview.DataCollectionRuleResponseDataSources | undefined>;
    /**
     * Description of the data collection rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The specification of destinations.
     */
    public readonly destinations!: pulumi.Output<outputs.insights.v20191101preview.DataCollectionRuleResponseDestinations>;
    /**
     * Resource entity tag (ETag).
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataCollectionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataCollectionRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.dataCollectionRuleName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dataCollectionRuleName'");
            }
            if ((!args || args.dataFlows === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dataFlows'");
            }
            if ((!args || args.destinations === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'destinations'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dataCollectionRuleName"] = args ? args.dataCollectionRuleName : undefined;
            inputs["dataFlows"] = args ? args.dataFlows : undefined;
            inputs["dataSources"] = args ? args.dataSources : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinations"] = args ? args.destinations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["dataFlows"] = undefined /*out*/;
            inputs["dataSources"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["destinations"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataCollectionRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataCollectionRule resource.
 */
export interface DataCollectionRuleArgs {
    /**
     * The name of the data collection rule. The name is case insensitive.
     */
    readonly dataCollectionRuleName: pulumi.Input<string>;
    /**
     * The specification of data flows.
     */
    readonly dataFlows: pulumi.Input<pulumi.Input<inputs.insights.v20191101preview.DataFlow>[]>;
    /**
     * The specification of data sources. 
     * This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
     */
    readonly dataSources?: pulumi.Input<inputs.insights.v20191101preview.DataCollectionRuleDataSources>;
    /**
     * Description of the data collection rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The specification of destinations.
     */
    readonly destinations: pulumi.Input<inputs.insights.v20191101preview.DataCollectionRuleDestinations>;
    /**
     * The geo-location where the resource lives.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
