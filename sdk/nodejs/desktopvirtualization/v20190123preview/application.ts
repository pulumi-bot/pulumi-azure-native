// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Schema for Application properties.
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
        return new Application(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:desktopvirtualization/v20190123preview:Application';

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
     * Command Line Arguments for Application.
     */
    public readonly commandLineArguments!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
     */
    public readonly commandLineSetting!: pulumi.Output<string>;
    /**
     * Description of Application.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies a path for the executable file for the application.
     */
    public readonly filePath!: pulumi.Output<string | undefined>;
    /**
     * Friendly name of Application.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * the icon a 64 bit string as a byte array.
     */
    public /*out*/ readonly iconContent!: pulumi.Output<string>;
    /**
     * Hash of the icon.
     */
    public /*out*/ readonly iconHash!: pulumi.Output<string>;
    /**
     * Index of the icon.
     */
    public readonly iconIndex!: pulumi.Output<number | undefined>;
    /**
     * Path to icon.
     */
    public readonly iconPath!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether to show the RemoteApp program in the RD Web Access server.
     */
    public readonly showInPortal!: pulumi.Output<boolean | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.applicationGroupName === undefined) {
                throw new Error("Missing required property 'applicationGroupName'");
            }
            if (!args || args.applicationName === undefined) {
                throw new Error("Missing required property 'applicationName'");
            }
            if (!args || args.commandLineSetting === undefined) {
                throw new Error("Missing required property 'commandLineSetting'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationGroupName"] = args ? args.applicationGroupName : undefined;
            inputs["applicationName"] = args ? args.applicationName : undefined;
            inputs["commandLineArguments"] = args ? args.commandLineArguments : undefined;
            inputs["commandLineSetting"] = args ? args.commandLineSetting : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["filePath"] = args ? args.filePath : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["iconIndex"] = args ? args.iconIndex : undefined;
            inputs["iconPath"] = args ? args.iconPath : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["showInPortal"] = args ? args.showInPortal : undefined;
            inputs["iconContent"] = undefined /*out*/;
            inputs["iconHash"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["commandLineArguments"] = undefined /*out*/;
            inputs["commandLineSetting"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["filePath"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["iconContent"] = undefined /*out*/;
            inputs["iconHash"] = undefined /*out*/;
            inputs["iconIndex"] = undefined /*out*/;
            inputs["iconPath"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["showInPortal"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:desktopvirtualization/v20190924preview:Application" }, { type: "azure-nextgen:desktopvirtualization/v20191210preview:Application" }, { type: "azure-nextgen:desktopvirtualization/v20200921preview:Application" }, { type: "azure-nextgen:desktopvirtualization/v20201019preview:Application" }, { type: "azure-nextgen:desktopvirtualization/v20201102preview:Application" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The name of the application group
     */
    readonly applicationGroupName: pulumi.Input<string>;
    /**
     * The name of the application within the specified application group
     */
    readonly applicationName: pulumi.Input<string>;
    /**
     * Command Line Arguments for Application.
     */
    readonly commandLineArguments?: pulumi.Input<string>;
    /**
     * Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
     */
    readonly commandLineSetting: pulumi.Input<string>;
    /**
     * Description of Application.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies a path for the executable file for the application.
     */
    readonly filePath?: pulumi.Input<string>;
    /**
     * Friendly name of Application.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * Index of the icon.
     */
    readonly iconIndex?: pulumi.Input<number>;
    /**
     * Path to icon.
     */
    readonly iconPath?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies whether to show the RemoteApp program in the RD Web Access server.
     */
    readonly showInPortal?: pulumi.Input<boolean>;
}
