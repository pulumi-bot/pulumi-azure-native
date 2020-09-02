// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Application Insights workbook template definition.
 */
export class WorkbookTemplate extends pulumi.CustomResource {
    /**
     * Get an existing WorkbookTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkbookTemplate {
        return new WorkbookTemplate(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/v20191017preview:WorkbookTemplate';

    /**
     * Returns true if the given object is an instance of WorkbookTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkbookTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkbookTemplate.__pulumiType;
    }

    /**
     * Information about the author of the workbook template.
     */
    public readonly author!: pulumi.Output<string | undefined>;
    /**
     * Workbook galleries supported by the template.
     */
    public readonly galleries!: pulumi.Output<outputs.insights.v20191017preview.WorkbookTemplateGalleryResponse[]>;
    /**
     * Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
     */
    public readonly localized!: pulumi.Output<{[key: string]: outputs.insights.v20191017preview.WorkbookTemplateLocalizedGalleryResponse[]} | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Valid JSON object containing workbook template payload.
     */
    public readonly templateData!: pulumi.Output<{[key: string]: any}>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WorkbookTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkbookTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkbookTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WorkbookTemplateArgs | undefined;
            if (!args || args.galleries === undefined) {
                throw new Error("Missing required property 'galleries'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.templateData === undefined) {
                throw new Error("Missing required property 'templateData'");
            }
            inputs["author"] = args ? args.author : undefined;
            inputs["galleries"] = args ? args.galleries : undefined;
            inputs["localized"] = args ? args.localized : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateData"] = args ? args.templateData : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WorkbookTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkbookTemplate resource.
 */
export interface WorkbookTemplateArgs {
    /**
     * Information about the author of the workbook template.
     */
    readonly author?: pulumi.Input<string>;
    /**
     * Workbook galleries supported by the template.
     */
    readonly galleries: pulumi.Input<pulumi.Input<inputs.insights.v20191017preview.WorkbookTemplateGallery>[]>;
    /**
     * Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
     */
    readonly localized?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.Input<inputs.insights.v20191017preview.WorkbookTemplateLocalizedGallery>[]>}>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Application Insights component resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Valid JSON object containing workbook template payload.
     */
    readonly templateData: pulumi.Input<{[key: string]: any}>;
}
