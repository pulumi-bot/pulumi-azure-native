// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Api Operation details.
 */
export class ApiOperation extends pulumi.CustomResource {
    /**
     * Get an existing ApiOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiOperation {
        return new ApiOperation(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20170301:ApiOperation';

    /**
     * Returns true if the given object is an instance of ApiOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiOperation.__pulumiType;
    }

    /**
     * Description of the operation. May include HTML formatting tags.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Operation Name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
     */
    public readonly method!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Operation Policies
     */
    public readonly policies!: pulumi.Output<string | undefined>;
    /**
     * An entity containing request details.
     */
    public readonly request!: pulumi.Output<outputs.apimanagement.v20170301.RequestContractResponse | undefined>;
    /**
     * Array of Operation responses.
     */
    public readonly responses!: pulumi.Output<outputs.apimanagement.v20170301.ResponseContractResponse[] | undefined>;
    /**
     * Collection of URL template parameters.
     */
    public readonly templateParameters!: pulumi.Output<outputs.apimanagement.v20170301.ParameterContractResponse[] | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
     */
    public readonly urlTemplate!: pulumi.Output<string>;

    /**
     * Create a ApiOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiOperationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApiOperationArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.method === undefined) {
                throw new Error("Missing required property 'method'");
            }
            if (!args || args.operationId === undefined) {
                throw new Error("Missing required property 'operationId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.urlTemplate === undefined) {
                throw new Error("Missing required property 'urlTemplate'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["method"] = args ? args.method : undefined;
            inputs["operationId"] = args ? args.operationId : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["request"] = args ? args.request : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["responses"] = args ? args.responses : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["templateParameters"] = args ? args.templateParameters : undefined;
            inputs["urlTemplate"] = args ? args.urlTemplate : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiOperation" }, { type: "azurerm:apimanagement/v20160707:ApiOperation" }, { type: "azurerm:apimanagement/v20161010:ApiOperation" }, { type: "azurerm:apimanagement/v20180101:ApiOperation" }, { type: "azurerm:apimanagement/v20180601preview:ApiOperation" }, { type: "azurerm:apimanagement/v20190101:ApiOperation" }, { type: "azurerm:apimanagement/v20191201:ApiOperation" }, { type: "azurerm:apimanagement/v20191201preview:ApiOperation" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiOperation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiOperation resource.
 */
export interface ApiOperationArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Description of the operation. May include HTML formatting tags.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Operation Name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
     */
    readonly method: pulumi.Input<string>;
    /**
     * Operation identifier within an API. Must be unique in the current API Management service instance.
     */
    readonly operationId: pulumi.Input<string>;
    /**
     * Operation Policies
     */
    readonly policies?: pulumi.Input<string>;
    /**
     * An entity containing request details.
     */
    readonly request?: pulumi.Input<inputs.apimanagement.v20170301.RequestContract>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Array of Operation responses.
     */
    readonly responses?: pulumi.Input<pulumi.Input<inputs.apimanagement.v20170301.ResponseContract>[]>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Collection of URL template parameters.
     */
    readonly templateParameters?: pulumi.Input<pulumi.Input<inputs.apimanagement.v20170301.ParameterContract>[]>;
    /**
     * Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
     */
    readonly urlTemplate: pulumi.Input<string>;
}
