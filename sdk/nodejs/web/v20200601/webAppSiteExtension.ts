// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Site Extension Information.
 */
export class WebAppSiteExtension extends pulumi.CustomResource {
    /**
     * Get an existing WebAppSiteExtension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppSiteExtension {
        return new WebAppSiteExtension(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20200601:WebAppSiteExtension';

    /**
     * Returns true if the given object is an instance of WebAppSiteExtension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppSiteExtension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppSiteExtension.__pulumiType;
    }

    /**
     * List of authors.
     */
    public /*out*/ readonly authors!: pulumi.Output<string[] | undefined>;
    /**
     * Site Extension comment.
     */
    public /*out*/ readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Detailed description.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * Count of downloads.
     */
    public /*out*/ readonly downloadCount!: pulumi.Output<number | undefined>;
    /**
     * Site extension ID.
     */
    public /*out*/ readonly extensionId!: pulumi.Output<string | undefined>;
    /**
     * Site extension type.
     */
    public /*out*/ readonly extensionType!: pulumi.Output<SiteExtensionType | undefined>;
    /**
     * Extension URL.
     */
    public /*out*/ readonly extensionUrl!: pulumi.Output<string | undefined>;
    /**
     * Feed URL.
     */
    public /*out*/ readonly feedUrl!: pulumi.Output<string | undefined>;
    /**
     * Icon URL.
     */
    public /*out*/ readonly iconUrl!: pulumi.Output<string | undefined>;
    /**
     * Installed timestamp.
     */
    public /*out*/ readonly installedDateTime!: pulumi.Output<string | undefined>;
    /**
     * Installer command line parameters.
     */
    public /*out*/ readonly installerCommandLineParams!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public /*out*/ readonly kind!: pulumi.Output<string | undefined>;
    /**
     * License URL.
     */
    public /*out*/ readonly licenseUrl!: pulumi.Output<string | undefined>;
    /**
     * <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
     */
    public /*out*/ readonly localIsLatestVersion!: pulumi.Output<boolean | undefined>;
    /**
     * Local path.
     */
    public /*out*/ readonly localPath!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project URL.
     */
    public /*out*/ readonly projectUrl!: pulumi.Output<string | undefined>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Published timestamp.
     */
    public /*out*/ readonly publishedDateTime!: pulumi.Output<string | undefined>;
    /**
     * Summary description.
     */
    public /*out*/ readonly summary!: pulumi.Output<string | undefined>;
    public /*out*/ readonly title!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Version information.
     */
    public /*out*/ readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a WebAppSiteExtension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppSiteExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppSiteExtensionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppSiteExtensionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.siteExtensionId === undefined) {
                throw new Error("Missing required property 'siteExtensionId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["siteExtensionId"] = args ? args.siteExtensionId : undefined;
            inputs["authors"] = undefined /*out*/;
            inputs["comment"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["downloadCount"] = undefined /*out*/;
            inputs["extensionId"] = undefined /*out*/;
            inputs["extensionType"] = undefined /*out*/;
            inputs["extensionUrl"] = undefined /*out*/;
            inputs["feedUrl"] = undefined /*out*/;
            inputs["iconUrl"] = undefined /*out*/;
            inputs["installedDateTime"] = undefined /*out*/;
            inputs["installerCommandLineParams"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["licenseUrl"] = undefined /*out*/;
            inputs["localIsLatestVersion"] = undefined /*out*/;
            inputs["localPath"] = undefined /*out*/;
            inputs["projectUrl"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publishedDateTime"] = undefined /*out*/;
            inputs["summary"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/v20160801:WebAppSiteExtension" }, { type: "azurerm:web/v20180201:WebAppSiteExtension" }, { type: "azurerm:web/v20181101:WebAppSiteExtension" }, { type: "azurerm:web/v20190801:WebAppSiteExtension" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebAppSiteExtension.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppSiteExtension resource.
 */
export interface WebAppSiteExtensionArgs {
    /**
     * Site name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Site extension name.
     */
    readonly siteExtensionId: pulumi.Input<string>;
}
