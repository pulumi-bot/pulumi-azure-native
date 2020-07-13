// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Definition of the runbook type.
 */
export class AutomationAccountRunbook extends pulumi.CustomResource {
    /**
     * Get an existing AutomationAccountRunbook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AutomationAccountRunbook {
        return new AutomationAccountRunbook(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:automation:AutomationAccountRunbook';

    /**
     * Returns true if the given object is an instance of AutomationAccountRunbook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutomationAccountRunbook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutomationAccountRunbook.__pulumiType;
    }

    /**
     * Gets or sets the etag of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The Azure Region where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the runbook properties.
     */
    public readonly properties!: pulumi.Output<outputs.automation.RunbookPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AutomationAccountRunbook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutomationAccountRunbookArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.runbookName === undefined) {
                throw new Error("Missing required property 'runbookName'");
            }
        inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["runbookName"] = args ? args.runbookName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["etag"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AutomationAccountRunbook.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AutomationAccountRunbook resource.
 */
export interface AutomationAccountRunbookArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Gets or sets the name of the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Gets or sets runbook create or update properties.
     */
    readonly properties: pulumi.Input<inputs.automation.RunbookCreateOrUpdateProperties>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The runbook name.
     */
    readonly runbookName: pulumi.Input<string>;
    /**
     * Gets or sets the tags attached to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
