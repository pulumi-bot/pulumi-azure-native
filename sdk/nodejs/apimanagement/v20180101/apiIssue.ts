// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Issue Contract details.
 */
export class ApiIssue extends pulumi.CustomResource {
    /**
     * Get an existing ApiIssue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiIssue {
        return new ApiIssue(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20180101:ApiIssue';

    /**
     * Returns true if the given object is an instance of ApiIssue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiIssue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiIssue.__pulumiType;
    }

    /**
     * A resource identifier for the API the issue was created for.
     */
    public readonly apiId!: pulumi.Output<string | undefined>;
    /**
     * Date and time when the issue was created.
     */
    public readonly createdDate!: pulumi.Output<string | undefined>;
    /**
     * Text describing the issue.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Status of the issue.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * The issue title.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A resource identifier for the user created the issue.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a ApiIssue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiIssueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiIssueArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApiIssueArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.issueId === undefined) {
                throw new Error("Missing required property 'issueId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            if (!args || args.userId === undefined) {
                throw new Error("Missing required property 'userId'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["createdDate"] = args ? args.createdDate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["issueId"] = args ? args.issueId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:ApiIssue" }, { type: "azurerm:apimanagement/v20170301:ApiIssue" }, { type: "azurerm:apimanagement/v20180601preview:ApiIssue" }, { type: "azurerm:apimanagement/v20190101:ApiIssue" }, { type: "azurerm:apimanagement/v20191201:ApiIssue" }, { type: "azurerm:apimanagement/v20191201preview:ApiIssue" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApiIssue.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiIssue resource.
 */
export interface ApiIssueArgs {
    /**
     * A resource identifier for the API the issue was created for.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * Date and time when the issue was created.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * Text describing the issue.
     */
    readonly description: pulumi.Input<string>;
    /**
     * Issue identifier. Must be unique in the current API Management service instance.
     */
    readonly issueId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Status of the issue.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The issue title.
     */
    readonly title: pulumi.Input<string>;
    /**
     * A resource identifier for the user created the issue.
     */
    readonly userId: pulumi.Input<string>;
}
