// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Machine Learning service object wrapped into ARM resource envelope.
 */
export class MachineLearningService extends pulumi.CustomResource {
    /**
     * Get an existing MachineLearningService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MachineLearningService {
        return new MachineLearningService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:machinelearningservices/v20200515preview:MachineLearningService';

    /**
     * Returns true if the given object is an instance of MachineLearningService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineLearningService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineLearningService.__pulumiType;
    }

    /**
     * The identity of the resource.
     */
    public /*out*/ readonly identity!: pulumi.Output<outputs.machinelearningservices.v20200515preview.IdentityResponse | undefined>;
    /**
     * Specifies the location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Service properties
     */
    public readonly properties!: pulumi.Output<outputs.machinelearningservices.v20200515preview.ACIServiceResponseResponse | outputs.machinelearningservices.v20200515preview.AKSVariantResponseResponse>;
    /**
     * The sku of the workspace.
     */
    public /*out*/ readonly sku!: pulumi.Output<outputs.machinelearningservices.v20200515preview.SkuResponse | undefined>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MachineLearningService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineLearningServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.computeType === undefined) {
                throw new Error("Missing required property 'computeType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["computeType"] = args ? args.computeType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environmentImageRequest"] = args ? args.environmentImageRequest : undefined;
            inputs["keys"] = args ? args.keys : undefined;
            inputs["kvTags"] = args ? args.kvTags : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["identity"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:machinelearningservices/v20200501preview:MachineLearningService" }, { type: "azurerm:machinelearningservices/v20200901preview:MachineLearningService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(MachineLearningService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a MachineLearningService resource.
 */
export interface MachineLearningServiceArgs {
    /**
     * The compute environment type for the service.
     */
    readonly computeType: pulumi.Input<string>;
    /**
     * The description of the service.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Environment, models and assets needed for inferencing.
     */
    readonly environmentImageRequest?: pulumi.Input<inputs.machinelearningservices.v20200515preview.CreateServiceRequestEnvironmentImageRequest>;
    /**
     * The authentication keys.
     */
    readonly keys?: pulumi.Input<inputs.machinelearningservices.v20200515preview.CreateServiceRequestKeys>;
    /**
     * The service tag dictionary. Tags are mutable.
     */
    readonly kvTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Azure location/region.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The service properties dictionary. Properties are immutable.
     */
    readonly properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the Azure Machine Learning service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
