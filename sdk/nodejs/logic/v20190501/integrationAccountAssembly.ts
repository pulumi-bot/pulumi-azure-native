// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * The assembly definition.
 */
export class IntegrationAccountAssembly extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationAccountAssembly resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IntegrationAccountAssembly {
        return new IntegrationAccountAssembly(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:logic/v20190501:IntegrationAccountAssembly';

    /**
     * Returns true if the given object is an instance of IntegrationAccountAssembly.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationAccountAssembly {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationAccountAssembly.__pulumiType;
    }

    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets the resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The assembly properties.
     */
    public readonly properties!: pulumi.Output<outputs.logic.v20190501.AssemblyPropertiesResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IntegrationAccountAssembly resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationAccountAssemblyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.assemblyArtifactName === undefined) {
                throw new Error("Missing required property 'assemblyArtifactName'");
            }
            if (!args || args.integrationAccountName === undefined) {
                throw new Error("Missing required property 'integrationAccountName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["assemblyArtifactName"] = args ? args.assemblyArtifactName : undefined;
            inputs["integrationAccountName"] = args ? args.integrationAccountName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:logic/latest:IntegrationAccountAssembly" }, { type: "azure-nextgen:logic/v20160601:IntegrationAccountAssembly" }, { type: "azure-nextgen:logic/v20180701preview:IntegrationAccountAssembly" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IntegrationAccountAssembly.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IntegrationAccountAssembly resource.
 */
export interface IntegrationAccountAssemblyArgs {
    /**
     * The assembly artifact name.
     */
    readonly assemblyArtifactName: pulumi.Input<string>;
    /**
     * The integration account name.
     */
    readonly integrationAccountName: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The assembly properties.
     */
    readonly properties: pulumi.Input<inputs.logic.v20190501.AssemblyProperties>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
