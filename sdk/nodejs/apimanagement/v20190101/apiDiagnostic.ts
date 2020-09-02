// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Diagnostic details.
 *
 * ## Example Usage
 * ### ApiManagementCreateApiDiagnostic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const apiDiagnostic = new azurerm.apimanagement.v20190101.ApiDiagnostic("apiDiagnostic", {
 *     alwaysLog: "allErrors",
 *     apiId: "57d1f7558aa04f15146d9d8a",
 *     backend: {
 *         request: {
 *             body: {
 *                 bytes: 512,
 *             },
 *             headers: ["Content-type"],
 *         },
 *         response: {
 *             body: {
 *                 bytes: 512,
 *             },
 *             headers: ["Content-type"],
 *         },
 *     },
 *     diagnosticId: "applicationinsights",
 *     frontend: {
 *         request: {
 *             body: {
 *                 bytes: 512,
 *             },
 *             headers: ["Content-type"],
 *         },
 *         response: {
 *             body: {
 *                 bytes: 512,
 *             },
 *             headers: ["Content-type"],
 *         },
 *     },
 *     loggerId: "/loggers/applicationinsights",
 *     resourceGroupName: "rg1",
 *     sampling: {
 *         percentage: 50,
 *         samplingType: "fixed",
 *     },
 *     serviceName: "apimService1",
 * });
 *
 * ```
 */
export class ApiDiagnostic extends pulumi.CustomResource {
    /**
     * Get an existing ApiDiagnostic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiDiagnostic {
        return new ApiDiagnostic(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20190101:ApiDiagnostic';

    /**
     * Returns true if the given object is an instance of ApiDiagnostic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiDiagnostic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiDiagnostic.__pulumiType;
    }

    /**
     * Specifies for what type of messages sampling settings should not apply.
     */
    public readonly alwaysLog!: pulumi.Output<string | undefined>;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Backend
     */
    public readonly backend!: pulumi.Output<outputs.apimanagement.v20190101.PipelineDiagnosticSettingsResponse | undefined>;
    /**
     * Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
     */
    public readonly enableHttpCorrelationHeaders!: pulumi.Output<boolean | undefined>;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
     */
    public readonly frontend!: pulumi.Output<outputs.apimanagement.v20190101.PipelineDiagnosticSettingsResponse | undefined>;
    /**
     * Sets correlation protocol to use for Application Insights diagnostics.
     */
    public readonly httpCorrelationProtocol!: pulumi.Output<string | undefined>;
    /**
     * Resource Id of a target logger.
     */
    public readonly loggerId!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Sampling settings for Diagnostic.
     */
    public readonly sampling!: pulumi.Output<outputs.apimanagement.v20190101.SamplingSettingsResponse | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The verbosity level applied to traces emitted by trace policies.
     */
    public readonly verbosity!: pulumi.Output<string | undefined>;

    /**
     * Create a ApiDiagnostic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiDiagnosticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiDiagnosticArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApiDiagnosticArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.diagnosticId === undefined) {
                throw new Error("Missing required property 'diagnosticId'");
            }
            if (!args || args.loggerId === undefined) {
                throw new Error("Missing required property 'loggerId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["alwaysLog"] = args ? args.alwaysLog : undefined;
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["diagnosticId"] = args ? args.diagnosticId : undefined;
            inputs["enableHttpCorrelationHeaders"] = args ? args.enableHttpCorrelationHeaders : undefined;
            inputs["frontend"] = args ? args.frontend : undefined;
            inputs["httpCorrelationProtocol"] = args ? args.httpCorrelationProtocol : undefined;
            inputs["loggerId"] = args ? args.loggerId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sampling"] = args ? args.sampling : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["verbosity"] = args ? args.verbosity : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiDiagnostic" }, { type: "azurerm:apimanagement/v20170301:ApiDiagnostic" }, { type: "azurerm:apimanagement/v20180101:ApiDiagnostic" }, { type: "azurerm:apimanagement/v20191201:ApiDiagnostic" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiDiagnostic.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiDiagnostic resource.
 */
export interface ApiDiagnosticArgs {
    /**
     * Specifies for what type of messages sampling settings should not apply.
     */
    readonly alwaysLog?: pulumi.Input<string>;
    /**
     * API identifier. Must be unique in the current API Management service instance.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Backend
     */
    readonly backend?: pulumi.Input<inputs.apimanagement.v20190101.PipelineDiagnosticSettings>;
    /**
     * Diagnostic identifier. Must be unique in the current API Management service instance.
     */
    readonly diagnosticId: pulumi.Input<string>;
    /**
     * Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
     */
    readonly enableHttpCorrelationHeaders?: pulumi.Input<boolean>;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
     */
    readonly frontend?: pulumi.Input<inputs.apimanagement.v20190101.PipelineDiagnosticSettings>;
    /**
     * Sets correlation protocol to use for Application Insights diagnostics.
     */
    readonly httpCorrelationProtocol?: pulumi.Input<string>;
    /**
     * Resource Id of a target logger.
     */
    readonly loggerId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sampling settings for Diagnostic.
     */
    readonly sampling?: pulumi.Input<inputs.apimanagement.v20190101.SamplingSettings>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * The verbosity level applied to traces emitted by trace policies.
     */
    readonly verbosity?: pulumi.Input<string>;
}
