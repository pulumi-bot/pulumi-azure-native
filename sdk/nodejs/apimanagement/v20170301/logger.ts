// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Logger details.
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
        return new Logger(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20170301:Logger';

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
     * Sampling settings for an ApplicationInsights logger.
     */
    public readonly sampling!: pulumi.Output<outputs.apimanagement.v20170301.LoggerSamplingContractResponse | undefined>;
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
    constructor(name: string, args: LoggerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.credentials === undefined) {
                throw new Error("Missing required property 'credentials'");
            }
            if (!args || args.loggerType === undefined) {
                throw new Error("Missing required property 'loggerType'");
            }
            if (!args || args.loggerid === undefined) {
                throw new Error("Missing required property 'loggerid'");
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
            inputs["loggerType"] = args ? args.loggerType : undefined;
            inputs["loggerid"] = args ? args.loggerid : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sampling"] = args ? args.sampling : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["credentials"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["isBuffered"] = undefined /*out*/;
            inputs["loggerType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sampling"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:Logger" }, { type: "azurerm:apimanagement/v20160707:Logger" }, { type: "azurerm:apimanagement/v20161010:Logger" }, { type: "azurerm:apimanagement/v20180101:Logger" }, { type: "azurerm:apimanagement/v20180601preview:Logger" }, { type: "azurerm:apimanagement/v20190101:Logger" }, { type: "azurerm:apimanagement/v20191201:Logger" }, { type: "azurerm:apimanagement/v20191201preview:Logger" }] };
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
     * Logger type.
     */
    readonly loggerType: pulumi.Input<string>;
    /**
     * Logger identifier. Must be unique in the API Management service instance.
     */
    readonly loggerid: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sampling settings for an ApplicationInsights logger.
     */
    readonly sampling?: pulumi.Input<inputs.apimanagement.v20170301.LoggerSamplingContract>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
