// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The application resource.
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabric/v20200301:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Azure resource etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Describes the managed identities for an Azure resource.
     */
    public readonly identity!: pulumi.Output<outputs.servicefabric.v20200301.ManagedIdentityResponse | undefined>;
    /**
     * It will be deprecated in New API, resource location depends on the parent resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The application resource properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.servicefabric.v20200301.ApplicationResourcePropertiesResponse>;
    /**
     * Azure resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApplicationArgs | undefined;
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedIdentities"] = args ? args.managedIdentities : undefined;
            inputs["maximumNodes"] = args ? args.maximumNodes : undefined;
            inputs["metrics"] = args ? args.metrics : undefined;
            inputs["minimumNodes"] = args ? args.minimumNodes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["removeApplicationCapacity"] = args ? args.removeApplicationCapacity : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["typeName"] = args ? args.typeName : undefined;
            inputs["typeVersion"] = args ? args.typeVersion : undefined;
            inputs["upgradePolicy"] = args ? args.upgradePolicy : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * Describes the managed identities for an Azure resource.
     */
    readonly identity?: pulumi.Input<inputs.servicefabric.v20200301.ManagedIdentity>;
    /**
     * It will be deprecated in New API, resource location depends on the parent resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * List of user assigned identities for the application, each mapped to a friendly name.
     */
    readonly managedIdentities?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20200301.ApplicationUserAssignedIdentity>[]>;
    /**
     * The maximum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. By default, the value of this property is zero and it means that the services can be placed on any node.
     */
    readonly maximumNodes?: pulumi.Input<number>;
    /**
     * List of application capacity metric description.
     */
    readonly metrics?: pulumi.Input<inputs.servicefabric.v20200301.ApplicationMetricDescriptionList>;
    /**
     * The minimum number of nodes where Service Fabric will reserve capacity for this application. Note that this does not mean that the services of this application will be placed on all of those nodes. If this property is set to zero, no capacity will be reserved. The value of this property cannot be more than the value of the MaximumNodes property.
     */
    readonly minimumNodes?: pulumi.Input<number>;
    /**
     * The name of the application resource.
     */
    readonly name: pulumi.Input<string>;
    /**
     * List of application parameters with overridden values from their default values specified in the application manifest.
     */
    readonly parameters?: pulumi.Input<inputs.servicefabric.v20200301.ApplicationParameterList>;
    /**
     * Remove the current application capacity settings.
     */
    readonly removeApplicationCapacity?: pulumi.Input<boolean>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Azure resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The application type name as defined in the application manifest.
     */
    readonly typeName?: pulumi.Input<string>;
    /**
     * The version of the application type as defined in the application manifest.
     */
    readonly typeVersion?: pulumi.Input<string>;
    /**
     * Describes the policy for a monitored application upgrade.
     */
    readonly upgradePolicy?: pulumi.Input<inputs.servicefabric.v20200301.ApplicationUpgradePolicy>;
}
