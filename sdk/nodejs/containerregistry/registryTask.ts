// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The task that has the ARM resource and task properties.
 * The task will have all information to schedule a run against it.
 */
export class RegistryTask extends pulumi.CustomResource {
    /**
     * Get an existing RegistryTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegistryTask {
        return new RegistryTask(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerregistry:RegistryTask';

    /**
     * Returns true if the given object is an instance of RegistryTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryTask.__pulumiType;
    }

    /**
     * Identity for the resource.
     */
    public readonly identity!: pulumi.Output<outputs.containerregistry.IdentityPropertiesResponse | undefined>;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of a task.
     */
    public readonly properties!: pulumi.Output<outputs.containerregistry.TaskPropertiesResponse>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RegistryTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegistryTaskArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["identity"] = args ? args.identity : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["registryName"] = args ? args.registryName : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegistryTask.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegistryTask resource.
 */
export interface RegistryTaskArgs {
    /**
     * Identity for the resource.
     */
    readonly identity?: pulumi.Input<inputs.containerregistry.IdentityProperties>;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the container registry task.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The properties of a task.
     */
    readonly properties?: pulumi.Input<inputs.containerregistry.TaskProperties>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
