// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Diagnostic details.
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
        return new ApiDiagnostic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:apimanagement/v20200601preview:ApiDiagnostic';

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
    public readonly backend!: pulumi.Output<outputs.apimanagement.v20200601preview.PipelineDiagnosticSettingsResponse | undefined>;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
     */
    public readonly frontend!: pulumi.Output<outputs.apimanagement.v20200601preview.PipelineDiagnosticSettingsResponse | undefined>;
    /**
     * Sets correlation protocol to use for Application Insights diagnostics.
     */
    public readonly httpCorrelationProtocol!: pulumi.Output<string | undefined>;
    /**
     * Log the ClientIP. Default is false.
     */
    public readonly logClientIp!: pulumi.Output<boolean | undefined>;
    /**
     * Resource Id of a target logger.
     */
    public readonly loggerId!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The format of the Operation Name for Application Insights telemetries. Default is Name.
     */
    public readonly operationNameFormat!: pulumi.Output<string | undefined>;
    /**
     * Sampling settings for Diagnostic.
     */
    public readonly sampling!: pulumi.Output<outputs.apimanagement.v20200601preview.SamplingSettingsResponse | undefined>;
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
    constructor(name: string, args: ApiDiagnosticArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
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
            inputs["frontend"] = args ? args.frontend : undefined;
            inputs["httpCorrelationProtocol"] = args ? args.httpCorrelationProtocol : undefined;
            inputs["logClientIp"] = args ? args.logClientIp : undefined;
            inputs["loggerId"] = args ? args.loggerId : undefined;
            inputs["operationNameFormat"] = args ? args.operationNameFormat : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sampling"] = args ? args.sampling : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["verbosity"] = args ? args.verbosity : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["alwaysLog"] = undefined /*out*/;
            inputs["backend"] = undefined /*out*/;
            inputs["frontend"] = undefined /*out*/;
            inputs["httpCorrelationProtocol"] = undefined /*out*/;
            inputs["logClientIp"] = undefined /*out*/;
            inputs["loggerId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationNameFormat"] = undefined /*out*/;
            inputs["sampling"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["verbosity"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:apimanagement/latest:ApiDiagnostic" }, { type: "azure-nextgen:apimanagement/v20170301:ApiDiagnostic" }, { type: "azure-nextgen:apimanagement/v20180101:ApiDiagnostic" }, { type: "azure-nextgen:apimanagement/v20180601preview:ApiDiagnostic" }, { type: "azure-nextgen:apimanagement/v20190101:ApiDiagnostic" }, { type: "azure-nextgen:apimanagement/v20191201:ApiDiagnostic" }, { type: "azure-nextgen:apimanagement/v20191201preview:ApiDiagnostic" }] };
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
    readonly backend?: pulumi.Input<inputs.apimanagement.v20200601preview.PipelineDiagnosticSettings>;
    /**
     * Diagnostic identifier. Must be unique in the current API Management service instance.
     */
    readonly diagnosticId: pulumi.Input<string>;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
     */
    readonly frontend?: pulumi.Input<inputs.apimanagement.v20200601preview.PipelineDiagnosticSettings>;
    /**
     * Sets correlation protocol to use for Application Insights diagnostics.
     */
    readonly httpCorrelationProtocol?: pulumi.Input<string>;
    /**
     * Log the ClientIP. Default is false.
     */
    readonly logClientIp?: pulumi.Input<boolean>;
    /**
     * Resource Id of a target logger.
     */
    readonly loggerId: pulumi.Input<string>;
    /**
     * The format of the Operation Name for Application Insights telemetries. Default is Name.
     */
    readonly operationNameFormat?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sampling settings for Diagnostic.
     */
    readonly sampling?: pulumi.Input<inputs.apimanagement.v20200601preview.SamplingSettings>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * The verbosity level applied to traces emitted by trace policies.
     */
    readonly verbosity?: pulumi.Input<string>;
}
