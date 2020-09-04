// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An application type version resource for the specified application type name resource.
 */
export class ApplicationTypeVersion extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationTypeVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationTypeVersion {
        return new ApplicationTypeVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabric/v20190601preview:ApplicationTypeVersion';

    /**
     * Returns true if the given object is an instance of ApplicationTypeVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationTypeVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationTypeVersion.__pulumiType;
    }

    /**
     * The URL to the application package
     */
    public readonly appPackageUrl!: pulumi.Output<string>;
    /**
     * List of application type parameters that can be overridden when creating or updating the application.
     */
    public /*out*/ readonly defaultParameterList!: pulumi.Output<{[key: string]: string}>;
    /**
     * Azure resource etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * It will be deprecated in New API, resource location depends on the parent resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current deployment or provisioning state, which only appears in the response
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApplicationTypeVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationTypeVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.appPackageUrl === undefined) {
                throw new Error("Missing required property 'appPackageUrl'");
            }
            if (!args || args.applicationTypeName === undefined) {
                throw new Error("Missing required property 'applicationTypeName'");
            }
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["appPackageUrl"] = args ? args.appPackageUrl : undefined;
            inputs["applicationTypeName"] = args ? args.applicationTypeName : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["defaultParameterList"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["appPackageUrl"] = undefined /*out*/;
            inputs["defaultParameterList"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:servicefabric/latest:ApplicationTypeVersion" }, { type: "azurerm:servicefabric/v20170701preview:ApplicationTypeVersion" }, { type: "azurerm:servicefabric/v20190301:ApplicationTypeVersion" }, { type: "azurerm:servicefabric/v20190301preview:ApplicationTypeVersion" }, { type: "azurerm:servicefabric/v20191101preview:ApplicationTypeVersion" }, { type: "azurerm:servicefabric/v20200301:ApplicationTypeVersion" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApplicationTypeVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationTypeVersion resource.
 */
export interface ApplicationTypeVersionArgs {
    /**
     * The URL to the application package
     */
    readonly appPackageUrl: pulumi.Input<string>;
    /**
     * The name of the application type name resource.
     */
    readonly applicationTypeName: pulumi.Input<string>;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * It will be deprecated in New API, resource location depends on the parent resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Azure resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The application type version.
     */
    readonly version: pulumi.Input<string>;
}
