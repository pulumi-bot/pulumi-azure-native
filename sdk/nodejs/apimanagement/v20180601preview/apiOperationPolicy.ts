// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Policy Contract details.
 */
export class ApiOperationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ApiOperationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiOperationPolicy {
        return new ApiOperationPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20180601preview:ApiOperationPolicy';

    /**
     * Returns true if the given object is an instance of ApiOperationPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiOperationPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiOperationPolicy.__pulumiType;
    }

    /**
     * Format of the policyContent.
     */
    public readonly contentFormat!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Json escaped Xml Encoded contents of the Policy.
     */
    public readonly policyContent!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiOperationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiOperationPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.operationId === undefined) {
                throw new Error("Missing required property 'operationId'");
            }
            if (!args || args.policyContent === undefined) {
                throw new Error("Missing required property 'policyContent'");
            }
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["contentFormat"] = args ? args.contentFormat : undefined;
            inputs["operationId"] = args ? args.operationId : undefined;
            inputs["policyContent"] = args ? args.policyContent : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["contentFormat"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["policyContent"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiOperationPolicy" }, { type: "azurerm:apimanagement/v20170301:ApiOperationPolicy" }, { type: "azurerm:apimanagement/v20180101:ApiOperationPolicy" }, { type: "azurerm:apimanagement/v20190101:ApiOperationPolicy" }, { type: "azurerm:apimanagement/v20191201:ApiOperationPolicy" }, { type: "azurerm:apimanagement/v20191201preview:ApiOperationPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiOperationPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiOperationPolicy resource.
 */
export interface ApiOperationPolicyArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Format of the policyContent.
     */
    readonly contentFormat?: pulumi.Input<string>;
    /**
     * Operation identifier within an API. Must be unique in the current API Management service instance.
     */
    readonly operationId: pulumi.Input<string>;
    /**
     * Json escaped Xml Encoded contents of the Policy.
     */
    readonly policyContent: pulumi.Input<string>;
    /**
     * The identifier of the Policy.
     */
    readonly policyId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
