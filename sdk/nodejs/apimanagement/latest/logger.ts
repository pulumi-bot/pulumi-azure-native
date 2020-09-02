// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Logger details.
 *
 * ## Example Usage
 * ### ApiManagementCreateAILogger
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const logger = new azurerm.apimanagement.latest.Logger("logger", {
 *     credentials: {
 *         instrumentationKey: "11................a1",
 *     },
 *     description: "adding a new logger",
 *     loggerId: "loggerId",
 *     loggerType: "applicationInsights",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 * ### ApiManagementCreateEHLogger
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const logger = new azurerm.apimanagement.latest.Logger("logger", {
 *     credentials: {
 *         connectionString: "Endpoint=sb://hydraeventhub-ns.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=********=",
 *         name: "hydraeventhub",
 *     },
 *     description: "adding a new logger",
 *     loggerId: "loggerId",
 *     loggerType: "azureEventHub",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 */
export class Logger extends pulumi.CustomResource {
    /**
     * Get an existing Logger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Logger {
        return new Logger(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/latest:Logger';

    /**
     * Returns true if the given object is an instance of Logger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Logger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Logger.__pulumiType;
    }

    /**
     * The name and SendRule connection string of the event hub for azureEventHub logger.
     * Instrumentation key for applicationInsights logger.
     */
    public readonly credentials!: pulumi.Output<{[key: string]: string}>;
    /**
     * Logger description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether records are buffered in the logger before publishing. Default is assumed to be true.
     */
    public readonly isBuffered!: pulumi.Output<boolean | undefined>;
    /**
     * Logger type.
     */
    public readonly loggerType!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Logger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoggerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LoggerArgs | undefined;
            if (!args || args.credentials === undefined) {
                throw new Error("Missing required property 'credentials'");
            }
            if (!args || args.loggerId === undefined) {
                throw new Error("Missing required property 'loggerId'");
            }
            if (!args || args.loggerType === undefined) {
                throw new Error("Missing required property 'loggerType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["isBuffered"] = args ? args.isBuffered : undefined;
            inputs["loggerId"] = args ? args.loggerId : undefined;
            inputs["loggerType"] = args ? args.loggerType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/v20160707:Logger" }, { type: "azurerm:apimanagement/v20161010:Logger" }, { type: "azurerm:apimanagement/v20170301:Logger" }, { type: "azurerm:apimanagement/v20180101:Logger" }, { type: "azurerm:apimanagement/v20190101:Logger" }, { type: "azurerm:apimanagement/v20191201:Logger" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Logger.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Logger resource.
 */
export interface LoggerArgs {
    /**
     * The name and SendRule connection string of the event hub for azureEventHub logger.
     * Instrumentation key for applicationInsights logger.
     */
    readonly credentials: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Logger description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether records are buffered in the logger before publishing. Default is assumed to be true.
     */
    readonly isBuffered?: pulumi.Input<boolean>;
    /**
     * Logger identifier. Must be unique in the API Management service instance.
     */
    readonly loggerId: pulumi.Input<string>;
    /**
     * Logger type.
     */
    readonly loggerType: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
