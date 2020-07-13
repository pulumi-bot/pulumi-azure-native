// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Factory resource type.
 */
export class Factory extends pulumi.CustomResource {
    /**
     * Get an existing Factory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Factory {
        return new Factory(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datafactory:Factory';

    /**
     * Returns true if the given object is an instance of Factory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Factory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Factory.__pulumiType;
    }

    /**
     * Etag identifies change in the resource.
     */
    public /*out*/ readonly eTag!: pulumi.Output<string>;
    /**
     * Managed service identity of the factory.
     */
    public readonly identity!: pulumi.Output<outputs.datafactory.FactoryIdentityResponse | undefined>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the factory.
     */
    public readonly properties!: pulumi.Output<outputs.datafactory.FactoryPropertiesResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Factory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FactoryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["identity"] = args ? args.identity : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["eTag"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Factory.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Factory resource.
 */
export interface FactoryArgs {
    /**
     * Managed service identity of the factory.
     */
    readonly identity?: pulumi.Input<inputs.datafactory.FactoryIdentity>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The factory name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Properties of the factory.
     */
    readonly properties?: pulumi.Input<inputs.datafactory.FactoryProperties>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
