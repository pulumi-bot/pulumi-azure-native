// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Source control configuration for an app.
 */
export class WebAppSourceControl extends pulumi.CustomResource {
    /**
     * Get an existing WebAppSourceControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppSourceControl {
        return new WebAppSourceControl(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20180201:WebAppSourceControl';

    /**
     * Returns true if the given object is an instance of WebAppSourceControl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppSourceControl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppSourceControl.__pulumiType;
    }

    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * SiteSourceControl resource specific properties
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.web.v20180201.SiteSourceControlResponseProperties>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebAppSourceControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppSourceControlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppSourceControlArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppSourceControlArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["branch"] = args ? args.branch : undefined;
            inputs["deploymentRollbackEnabled"] = args ? args.deploymentRollbackEnabled : undefined;
            inputs["isManualIntegration"] = args ? args.isManualIntegration : undefined;
            inputs["isMercurial"] = args ? args.isMercurial : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["repoUrl"] = args ? args.repoUrl : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WebAppSourceControl.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppSourceControl resource.
 */
export interface WebAppSourceControlArgs {
    /**
     * Name of branch to use for deployment.
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
     */
    readonly deploymentRollbackEnabled?: pulumi.Input<boolean>;
    /**
     * <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
     */
    readonly isManualIntegration?: pulumi.Input<boolean>;
    /**
     * <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
     */
    readonly isMercurial?: pulumi.Input<boolean>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the app.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Repository or source control URL.
     */
    readonly repoUrl?: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
