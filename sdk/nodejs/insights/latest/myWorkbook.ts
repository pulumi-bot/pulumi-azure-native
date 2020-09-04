// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An Application Insights private workbook definition.
 *
 * ## Example Usage
 * ### WorkbookAdd
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const myWorkbook = new azurerm.insights.latest.MyWorkbook("myWorkbook", {
 *     category: "workbook",
 *     displayName: "Blah Blah Blah",
 *     id: "c0deea5e-3344-40f2-96f8-6f8e1c3b5722",
 *     kind: "user",
 *     location: "west us",
 *     name: "deadb33f-8bee-4d3b-a059-9be8dac93960",
 *     resourceGroupName: "my-resource-group",
 *     resourceName: "deadb33f-8bee-4d3b-a059-9be8dac93960",
 *     serializedData: "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}",
 *     sourceId: "/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens",
 *     tags: [
 *         "TagSample01",
 *         "TagSample02",
 *     ],
 * });
 *
 * ```
 */
export class MyWorkbook extends pulumi.CustomResource {
    /**
     * Get an existing MyWorkbook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MyWorkbook {
        return new MyWorkbook(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/latest:MyWorkbook';

    /**
     * Returns true if the given object is an instance of MyWorkbook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MyWorkbook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MyWorkbook.__pulumiType;
    }

    /**
     * Workbook category, as defined by the user at creation time.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * The user-defined name of the private workbook.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The kind of workbook. Choices are user and shared.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Configuration of this particular private workbook. Configuration data is a string containing valid JSON
     */
    public readonly serializedData!: pulumi.Output<string>;
    /**
     * Optional resourceId for a source resource.
     */
    public readonly sourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Date and time in UTC of the last modification that was made to this private workbook definition.
     */
    public /*out*/ readonly timeModified!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Unique user id of the specific user that owns this private workbook.
     */
    public /*out*/ readonly userId!: pulumi.Output<string>;
    /**
     * This instance's version of the data model. This can change as new features are added that can be marked private workbook.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a MyWorkbook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MyWorkbookArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.category === undefined) {
                throw new Error("Missing required property 'category'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.serializedData === undefined) {
                throw new Error("Missing required property 'serializedData'");
            }
            inputs["category"] = args ? args.category : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["serializedData"] = args ? args.serializedData : undefined;
            inputs["sourceId"] = args ? args.sourceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["timeModified"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
        } else {
            inputs["category"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serializedData"] = undefined /*out*/;
            inputs["sourceId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeModified"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:insights/v20150501:MyWorkbook" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(MyWorkbook.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a MyWorkbook resource.
 */
export interface MyWorkbookArgs {
    /**
     * Workbook category, as defined by the user at creation time.
     */
    readonly category: pulumi.Input<string>;
    /**
     * The user-defined name of the private workbook.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Azure resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The kind of workbook. Choices are user and shared.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Azure resource name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Application Insights component resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Configuration of this particular private workbook. Configuration data is a string containing valid JSON
     */
    readonly serializedData: pulumi.Input<string>;
    /**
     * Optional resourceId for a source resource.
     */
    readonly sourceId?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Azure resource type
     */
    readonly type?: pulumi.Input<string>;
    /**
     * This instance's version of the data model. This can change as new features are added that can be marked private workbook.
     */
    readonly version?: pulumi.Input<string>;
}
