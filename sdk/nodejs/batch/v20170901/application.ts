// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Contains information about an application in a Batch account.
 *
 * ## Example Usage
 * ### ApplicationCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const application = new azurerm.batch.v20170901.Application("application", {
 *     accountName: "sampleacct",
 *     allowUpdates: false,
 *     applicationId: "app1",
 *     displayName: "myAppName",
 *     resourceGroupName: "default-azurebatch-japaneast",
 * });
 *
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:batch/v20170901:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * A value indicating whether packages within the application may be overwritten using the same version string.
     */
    public readonly allowUpdates!: pulumi.Output<boolean | undefined>;
    /**
     * The package to use if a client requests the application but does not specify a version.
     */
    public /*out*/ readonly defaultVersion!: pulumi.Output<string | undefined>;
    /**
     * The display name for the application.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The list of packages under this application.
     */
    public /*out*/ readonly packages!: pulumi.Output<outputs.batch.v20170901.ApplicationPackageResponse[] | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApplicationArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.applicationId === undefined) {
                throw new Error("Missing required property 'applicationId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["allowUpdates"] = args ? args.allowUpdates : undefined;
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["defaultVersion"] = undefined /*out*/;
            inputs["packages"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:batch/latest:Application" }, { type: "azurerm:batch/v20151201:Application" }, { type: "azurerm:batch/v20170101:Application" }, { type: "azurerm:batch/v20170501:Application" }, { type: "azurerm:batch/v20181201:Application" }, { type: "azurerm:batch/v20190401:Application" }, { type: "azurerm:batch/v20190801:Application" }, { type: "azurerm:batch/v20200301:Application" }, { type: "azurerm:batch/v20200501:Application" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The name of the Batch account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * A value indicating whether packages within the application may be overwritten using the same version string.
     */
    readonly allowUpdates?: pulumi.Input<boolean>;
    /**
     * The ID of the application.
     */
    readonly applicationId: pulumi.Input<string>;
    /**
     * The display name for the application.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
