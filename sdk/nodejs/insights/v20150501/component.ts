// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Application Insights component definition.
 */
export class Component extends pulumi.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Component {
        return new Component(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/v20150501:Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    /**
     * The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties that define an Application Insights component resource.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.insights.v20150501.ApplicationInsightsComponentPropertiesResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComponentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ComponentArgs | undefined;
            if (!args || args.Application_Type === undefined) {
                throw new Error("Missing required property 'Application_Type'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["Application_Type"] = args ? args.Application_Type : undefined;
            inputs["DisableIpMasking"] = args ? args.DisableIpMasking : undefined;
            inputs["Flow_Type"] = args ? args.Flow_Type : undefined;
            inputs["HockeyAppId"] = args ? args.HockeyAppId : undefined;
            inputs["ImmediatePurgeDataOn30Days"] = args ? args.ImmediatePurgeDataOn30Days : undefined;
            inputs["IngestionMode"] = args ? args.IngestionMode : undefined;
            inputs["Request_Source"] = args ? args.Request_Source : undefined;
            inputs["RetentionInDays"] = args ? args.RetentionInDays : undefined;
            inputs["SamplingPercentage"] = args ? args.SamplingPercentage : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Component.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    /**
     * Type of application being monitored.
     */
    readonly Application_Type: pulumi.Input<string>;
    /**
     * Disable IP masking.
     */
    readonly DisableIpMasking?: pulumi.Input<boolean>;
    /**
     * Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
     */
    readonly Flow_Type?: pulumi.Input<string>;
    /**
     * The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
     */
    readonly HockeyAppId?: pulumi.Input<string>;
    /**
     * Purge data immediately after 30 days.
     */
    readonly ImmediatePurgeDataOn30Days?: pulumi.Input<boolean>;
    /**
     * Indicates the flow of the ingestion.
     */
    readonly IngestionMode?: pulumi.Input<string>;
    /**
     * Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
     */
    readonly Request_Source?: pulumi.Input<string>;
    /**
     * Retention period in days.
     */
    readonly RetentionInDays?: pulumi.Input<number>;
    /**
     * Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
     */
    readonly SamplingPercentage?: pulumi.Input<number>;
    /**
     * The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the Application Insights component resource.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
