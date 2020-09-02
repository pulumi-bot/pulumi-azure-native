// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Deployment script object.
 */
export class DeploymentScript extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentScript resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeploymentScript {
        return new DeploymentScript(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:resources/v20191001preview:DeploymentScript';

    /**
     * Returns true if the given object is an instance of DeploymentScript.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentScript {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentScript.__pulumiType;
    }

    /**
     * Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
     */
    public readonly identity!: pulumi.Output<outputs.resources.v20191001preview.ManagedServiceIdentityResponse>;
    /**
     * Type of the script.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The location of the ACI and the storage account for the deployment script.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of this resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The system metadata related to this resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.resources.v20191001preview.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of this resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DeploymentScript resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentScriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentScriptArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DeploymentScriptArgs | undefined;
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.scriptName === undefined) {
                throw new Error("Missing required property 'scriptName'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scriptName"] = args ? args.scriptName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DeploymentScript.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeploymentScript resource.
 */
export interface DeploymentScriptArgs {
    /**
     * Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
     */
    readonly identity: pulumi.Input<inputs.resources.v20191001preview.ManagedServiceIdentity>;
    /**
     * Type of the script.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The location of the ACI and the storage account for the deployment script.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the deployment script.
     */
    readonly scriptName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
