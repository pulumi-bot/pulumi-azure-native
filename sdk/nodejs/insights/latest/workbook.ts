// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An Application Insights workbook definition.
 *
 * ## WorkbookAdd
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const workbook = new azurerm.insights.latest.Workbook("workbook", {
 *     category: "workbook",
 *     location: "west us",
 *     name: "deadb33f-8bee-4d3b-a059-9be8dac93960",
 *     resourceGroupName: "my-resource-group",
 *     resourceName: "deadb33f-8bee-4d3b-a059-9be8dac93960",
 *     serializedData: "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}",
 *     sharedTypeKind: "shared",
 *     sourceResourceId: "/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens",
 *     tags: [
 *         "TagSample01",
 *         "TagSample02",
 *     ],
 *     userId: "userId",
 *     workbookId: "deadb33f-8bee-4d3b-a059-9be8dac93960",
 * });
 *
 * ```
 */
export class Workbook extends pulumi.CustomResource {
    /**
     * Get an existing Workbook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workbook {
        return new Workbook(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/latest:Workbook';

    /**
     * Returns true if the given object is an instance of Workbook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workbook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workbook.__pulumiType;
    }

    /**
     * Workbook category, as defined by the user at creation time.
     */
    public readonly category!: pulumi.Output<string>;
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
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration of this particular workbook. Configuration data is a string containing valid JSON
     */
    public readonly serializedData!: pulumi.Output<string>;
    /**
     * Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
     */
    public readonly sharedTypeKind!: pulumi.Output<string>;
    /**
     * Optional resourceId for a source resource.
     */
    public readonly sourceResourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Date and time in UTC of the last modification that was made to this workbook definition.
     */
    public /*out*/ readonly timeModified!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Unique user id of the specific user that owns this workbook.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * This instance's version of the data model. This can change as new features are added that can be marked workbook.
     */
    public readonly version!: pulumi.Output<string | undefined>;
    /**
     * Internally assigned unique id of the workbook definition.
     */
    public readonly workbookId!: pulumi.Output<string>;

    /**
     * Create a Workbook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkbookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkbookArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WorkbookArgs | undefined;
            if (!args || args.category === undefined) {
                throw new Error("Missing required property 'category'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
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
            if (!args || args.sharedTypeKind === undefined) {
                throw new Error("Missing required property 'sharedTypeKind'");
            }
            if (!args || args.userId === undefined) {
                throw new Error("Missing required property 'userId'");
            }
            if (!args || args.workbookId === undefined) {
                throw new Error("Missing required property 'workbookId'");
            }
            inputs["category"] = args ? args.category : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["serializedData"] = args ? args.serializedData : undefined;
            inputs["sharedTypeKind"] = args ? args.sharedTypeKind : undefined;
            inputs["sourceResourceId"] = args ? args.sourceResourceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["workbookId"] = args ? args.workbookId : undefined;
            inputs["timeModified"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:insights/v20150501:Workbook" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Workbook.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workbook resource.
 */
export interface WorkbookArgs {
    /**
     * Workbook category, as defined by the user at creation time.
     */
    readonly category: pulumi.Input<string>;
    /**
     * The kind of workbook. Choices are user and shared.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The user-defined name of the workbook.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Application Insights component resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Configuration of this particular workbook. Configuration data is a string containing valid JSON
     */
    readonly serializedData: pulumi.Input<string>;
    /**
     * Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
     */
    readonly sharedTypeKind: pulumi.Input<string>;
    /**
     * Optional resourceId for a source resource.
     */
    readonly sourceResourceId?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Unique user id of the specific user that owns this workbook.
     */
    readonly userId: pulumi.Input<string>;
    /**
     * This instance's version of the data model. This can change as new features are added that can be marked workbook.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Internally assigned unique id of the workbook definition.
     */
    readonly workbookId: pulumi.Input<string>;
}