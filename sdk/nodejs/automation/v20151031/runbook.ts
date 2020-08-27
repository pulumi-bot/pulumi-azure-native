// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Definition of the runbook type.
 */
export class Runbook extends pulumi.CustomResource {
    /**
     * Get an existing Runbook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Runbook {
        return new Runbook(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:automation/v20151031:Runbook';

    /**
     * Returns true if the given object is an instance of Runbook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Runbook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Runbook.__pulumiType;
    }

    /**
     * Gets or sets the creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the draft runbook properties.
     */
    public readonly draft!: pulumi.Output<outputs.automation.v20151031.RunbookDraftResponse | undefined>;
    /**
     * Gets or sets the etag of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the job count of the runbook.
     */
    public /*out*/ readonly jobCount!: pulumi.Output<number | undefined>;
    /**
     * Gets or sets the last modified by.
     */
    public /*out*/ readonly lastModifiedBy!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * The Azure Region where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the option to log activity trace of the runbook.
     */
    public readonly logActivityTrace!: pulumi.Output<number | undefined>;
    /**
     * Gets or sets progress log option.
     */
    public readonly logProgress!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets verbose log option.
     */
    public readonly logVerbose!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the runbook output types.
     */
    public /*out*/ readonly outputTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Gets or sets the runbook parameters.
     */
    public /*out*/ readonly parameters!: pulumi.Output<{[key: string]: outputs.automation.v20151031.RunbookParameterResponse} | undefined>;
    /**
     * Gets or sets the provisioning state of the runbook.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<RunbookProvisioningState | undefined>;
    /**
     * Gets or sets the published runbook content link.
     */
    public readonly publishContentLink!: pulumi.Output<outputs.automation.v20151031.ContentLinkResponse | undefined>;
    /**
     * Gets or sets the type of the runbook.
     */
    public readonly runbookType!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the state of the runbook.
     */
    public /*out*/ readonly state!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Runbook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RunbookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RunbookArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RunbookArgs | undefined;
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.runbookName === undefined) {
                throw new Error("Missing required property 'runbookName'");
            }
            if (!args || args.runbookType === undefined) {
                throw new Error("Missing required property 'runbookType'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["draft"] = args ? args.draft : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logActivityTrace"] = args ? args.logActivityTrace : undefined;
            inputs["logProgress"] = args ? args.logProgress : undefined;
            inputs["logVerbose"] = args ? args.logVerbose : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publishContentLink"] = args ? args.publishContentLink : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["runbookName"] = args ? args.runbookName : undefined;
            inputs["runbookType"] = args ? args.runbookType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["jobCount"] = undefined /*out*/;
            inputs["lastModifiedBy"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["outputTypes"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:automation/v20180630:Runbook" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Runbook.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Runbook resource.
 */
export interface RunbookArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the description of the runbook.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Gets or sets the draft runbook properties.
     */
    readonly draft?: pulumi.Input<inputs.automation.v20151031.RunbookDraft>;
    /**
     * Gets or sets the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Gets or sets the activity-level tracing options of the runbook.
     */
    readonly logActivityTrace?: pulumi.Input<number>;
    /**
     * Gets or sets progress log option.
     */
    readonly logProgress?: pulumi.Input<boolean>;
    /**
     * Gets or sets verbose log option.
     */
    readonly logVerbose?: pulumi.Input<boolean>;
    /**
     * Gets or sets the name of the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Gets or sets the published runbook content link.
     */
    readonly publishContentLink?: pulumi.Input<inputs.automation.v20151031.ContentLink>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The runbook name.
     */
    readonly runbookName: pulumi.Input<string>;
    /**
     * Gets or sets the type of the runbook.
     */
    readonly runbookType: pulumi.Input<string>;
    /**
     * Gets or sets the tags attached to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
