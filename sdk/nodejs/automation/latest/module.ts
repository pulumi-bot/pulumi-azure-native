// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Definition of the module type.
 *
 * ## Create or update a module
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const module = new azurerm.automation.latest.Module("module", {
 *     automationAccountName: "myAutomationAccount33",
 *     contentLink: {
 *         contentHash: {
 *             algorithm: "sha265",
 *             value: "07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A",
 *         },
 *         uri: "https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip",
 *         version: "1.0.0.0",
 *     },
 *     moduleName: "OmsCompositeResources",
 *     resourceGroupName: "rg",
 * });
 *
 * ```
 */
export class Module extends pulumi.CustomResource {
    /**
     * Get an existing Module resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Module {
        return new Module(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:automation/latest:Module';

    /**
     * Returns true if the given object is an instance of Module.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Module {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Module.__pulumiType;
    }

    /**
     * Gets or sets the activity count of the module.
     */
    public /*out*/ readonly activityCount!: pulumi.Output<number | undefined>;
    /**
     * Gets or sets the contentLink of the module.
     */
    public readonly contentLink!: pulumi.Output<outputs.automation.latest.ContentLinkResponse | undefined>;
    /**
     * Gets or sets the creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the description.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the error info of the module.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.automation.latest.ModuleErrorInfoResponse | undefined>;
    /**
     * Gets or sets the etag of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets type of module, if its composite or not.
     */
    public /*out*/ readonly isComposite!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets the isGlobal flag of the module.
     */
    public /*out*/ readonly isGlobal!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * The Azure Region where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the provisioning state of the module.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the size in bytes of the module.
     */
    public /*out*/ readonly sizeInBytes!: pulumi.Output<number | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Gets or sets the version of the module.
     */
    public /*out*/ readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Module resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ModuleArgs | undefined;
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.contentLink === undefined) {
                throw new Error("Missing required property 'contentLink'");
            }
            if (!args || args.moduleName === undefined) {
                throw new Error("Missing required property 'moduleName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["contentLink"] = args ? args.contentLink : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["moduleName"] = args ? args.moduleName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["activityCount"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["isComposite"] = undefined /*out*/;
            inputs["isGlobal"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sizeInBytes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:automation/v20151031:Module" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Module.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Module resource.
 */
export interface ModuleArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the module content link.
     */
    readonly contentLink: pulumi.Input<inputs.automation.latest.ContentLink>;
    /**
     * Gets or sets the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of module.
     */
    readonly moduleName: pulumi.Input<string>;
    /**
     * Gets or sets name of the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the tags attached to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}