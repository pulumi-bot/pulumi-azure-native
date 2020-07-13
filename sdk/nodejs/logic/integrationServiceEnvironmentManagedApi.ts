// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The managed api definition.
 */
export class IntegrationServiceEnvironmentManagedApi extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationServiceEnvironmentManagedApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IntegrationServiceEnvironmentManagedApi {
        return new IntegrationServiceEnvironmentManagedApi(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:logic:IntegrationServiceEnvironmentManagedApi';

    /**
     * Returns true if the given object is an instance of IntegrationServiceEnvironmentManagedApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationServiceEnvironmentManagedApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationServiceEnvironmentManagedApi.__pulumiType;
    }

    /**
     * The resource location.
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The api resource properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.logic.ApiResourcePropertiesResponse>;
    /**
     * The resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IntegrationServiceEnvironmentManagedApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IntegrationServiceEnvironmentManagedApiArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.integrationServiceEnvironmentName === undefined) {
                throw new Error("Missing required property 'integrationServiceEnvironmentName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroup === undefined) {
                throw new Error("Missing required property 'resourceGroup'");
            }
        inputs["integrationServiceEnvironmentName"] = args ? args.integrationServiceEnvironmentName : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
        inputs["location"] = undefined /*out*/;
        inputs["properties"] = undefined /*out*/;
        inputs["tags"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IntegrationServiceEnvironmentManagedApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IntegrationServiceEnvironmentManagedApi resource.
 */
export interface IntegrationServiceEnvironmentManagedApiArgs {
    /**
     * The integration service environment name.
     */
    readonly integrationServiceEnvironmentName: pulumi.Input<string>;
    /**
     * The api name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroup: pulumi.Input<string>;
}
